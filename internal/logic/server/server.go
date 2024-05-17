package server

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "voichatter/api/server/v1"
	"voichatter/internal/consts"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/model/entity"
	"voichatter/internal/service"
	"voichatter/utility/auth"
	"voichatter/utility/cache"
	"voichatter/utility/errResponse"
)

type (
	sServer struct{}
)

func init() {
	service.RegisterServer(New())
}

func New() service.IServer {
	return &sServer{}
}
func (s *sServer) ServerList(ctx context.Context, _ *v1.ServerListReq) (res *v1.ServerListRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))

	serverList := fmt.Sprintf("%s-%d", consts.ServerList, userId)

	var servers []*model.Server
	get, err := g.Redis().Get(ctx, serverList)
	if err != nil {
		return nil, errResponse.DbOperationError("查询服务器列表出错了")
	}
	if err = get.Struct(&servers); err != nil {
		return nil, errResponse.OperationFailed("服务器列表转换出错")
	}
	if servers != nil {
		return &v1.ServerListRes{
			ServerList: servers,
		}, nil
	}
	err = g.Model("server s").
		Fields("s.*").
		LeftJoin("member m", "s.server_id = m.server_id").
		Where("m.user_id = ?", userId).
		Scan(&servers)
	if err != nil {
		return nil, errResponse.DbOperationError("查询服务器列表出错了")
	}
	err = g.Redis().SetEX(ctx, serverList, servers, consts.OneDaySec)
	if err != nil {
		return nil, errResponse.DbOperationError("设置服务器列表缓存出错了")
	}
	return &v1.ServerListRes{
		ServerList: servers,
	}, nil
}

func (s *sServer) ServerCreate(ctx context.Context, in model.ServerCreateInput) (res *v1.ServerCreateRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	createDate := gtime.Now()
	server := entity.Server{
		ServerName:    in.ServerName,
		CreatorUserId: userId,
		CreateDate:    createDate,
		ServerType:    in.ServerType,
		ServerImgUrl:  in.ServerImgUrl,
	}
	var lastInsertId int64
	err = dao.Server.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		lastInsertId, err = dao.Server.Ctx(ctx).InsertAndGetId(&server)
		if err != nil {
			return err
		}
		member := entity.Member{
			UserId:       server.CreatorUserId,
			ServerId:     uint64(lastInsertId),
			JoinDate:     gtime.Now(),
			SPermissions: "owner",
		}
		_, err = dao.Member.Ctx(ctx).Insert(&member)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, errResponse.DbOperationError("创建服务器出错了")
	}

	if err = cache.DelServerListsCache(ctx, uint64(lastInsertId)); err != nil {
		return nil, err
	}

	return &v1.ServerCreateRes{
		Server: &model.Server{
			ServerId:      uint64(lastInsertId),
			ServerName:    in.ServerName,
			CreatorUserId: userId,
			CreateDate:    createDate,
			ServerType:    in.ServerType,
			ServerImgUrl:  in.ServerImgUrl,
		},
	}, nil
}

func (s *sServer) ServerJoin(ctx context.Context, serverId uint64, link string) (res *v1.ServerJoinRes, err error) {
	var server *model.Server
	if link == "" {
		// 服务器是否公开
		err = dao.Server.Ctx(ctx).
			Where("server_id = ? AND server_type = ?", serverId, "public").
			Scan(&server)
	} else {
		linkTar := fmt.Sprintf("%s-%s", consts.InviteLink, link)
		get, err := g.Redis().Get(ctx, linkTar)
		if err != nil {
			return nil, errResponse.DbOperationError("查询邀请链接出错")
		}
		err = dao.Server.Ctx(ctx).
			Where("server_id = ?", get.Uint64()).
			Scan(&server)
	}
	if err != nil {
		return nil, errResponse.DbOperationError("查询出错")
	}
	if server == nil {
		return nil, errResponse.OperationFailed("服务器不存在或不是公开的")
	}
	userId := gconv.Uint64(ctx.Value("userId"))
	// 判断是否已经加入过服务器
	count, err := dao.Member.Ctx(ctx).
		Where("server_id = ? AND user_id = ?", server.ServerId, userId).
		Count()
	if err != nil {
		return nil, errResponse.DbOperationError("查询是否加入过服务器时出错")
	}
	if count > 0 {
		return nil, errResponse.OperationFailed("已经加入过该服务器")
	}
	// 添加成员
	m := entity.Member{
		ServerId:     serverId,
		UserId:       userId,
		JoinDate:     gtime.Now(),
		SPermissions: "member",
	}
	_, err = dao.Member.Ctx(ctx).Insert(m)
	if err != nil {
		return nil, errResponse.DbOperationError("添加成员失败")
	}
	if err = cache.DelServerUsersCache(ctx, serverId); err != nil {
		return nil, err
	}
	if err = cache.DelServerListCache(ctx, userId); err != nil {
		return nil, err
	}
	return &v1.ServerJoinRes{
		Server: server,
	}, nil
}

// ServerInviteLink 生成邀请链接
func (s *sServer) ServerInviteLink(ctx context.Context, serverId uint64) (res string, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	err = auth.IsServerCreator(ctx, serverId)
	if err != nil {
		return "", err
	}
	count, err := dao.Server.Ctx(ctx).
		Where("server_id = ? AND creator_user_id = ?", serverId, userId).
		Count()
	if err != nil {
		return "", errResponse.DbOperationError("查询拥有者时失败")
	}
	if count == 0 {
		return "", errResponse.OperationFailed("权限不足")
	}
	token := make([]byte, 8)
	_, err = rand.Read(token)
	if err != nil {
		return "", err
	}
	link := hex.EncodeToString(token)
	g.Dump(link)
	linkTar := fmt.Sprintf("%s-%s", consts.InviteLink, link)
	err = g.Redis().SetEX(ctx, linkTar, serverId, consts.OneDaySec)
	if err != nil {
		return "", errResponse.DbOperationError("生成邀请链接失败")
	}
	return link, nil
}

func (s *sServer) ServerDel(ctx context.Context, serverId uint64) (res *v1.ServerDelRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	err = dao.Server.Ctx(ctx).Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		result, err := dao.Server.Ctx(ctx).
			Where("server_id = ? AND creator_user_id = ?", serverId, userId).
			Delete()
		row, _ := result.RowsAffected()
		if err != nil || row == 0 {
			//return nil, errResponse.OperationFailed("权限不足")
			return errResponse.OperationFailed("权限不足")
		}
		result, err = dao.Member.Ctx(ctx).Delete("server_id = ?", serverId)
		if err != nil {
			//return nil, errResponse.DbOperationError("删除成员失败")
			return errResponse.DbOperationError("删除成员失败")
		}
		return nil
	})
	if err != nil {
		return nil, err
	}

	if err = cache.DelServerListsCache(ctx, serverId); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *sServer) ServerModifyName(ctx context.Context, serverId uint64, serverName string) (res *v1.ServerModifyNameRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	count, err := dao.Server.Ctx(ctx).Where("server_id = ? AND creator_user_id = ?", serverId, userId).Count()
	if err != nil {
		return nil, errResponse.DbOperationError("查询拥有者时失败")
	}
	if count == 0 {
		return nil, errResponse.OperationFailed("权限不足")
	}
	result, err := dao.Server.Ctx(ctx).Update(g.Map{
		"server_name": serverName,
	}, "server_id = ?", serverId)
	if err != nil {
		return nil, errResponse.DbOperationError("修改失败")
	}
	row, _ := result.RowsAffected()
	if row == 0 {
		return nil, errResponse.OperationFailed("名字相同")
	}
	var server *model.Server
	err = dao.Server.Ctx(ctx).Where("server_id = ?", serverId).Scan(&server)
	if err != nil {
		return nil, errResponse.DbOperationError("查询失败")
	}

	if err = cache.DelServerListsCache(ctx, serverId); err != nil {
		return nil, err
	}

	return &v1.ServerModifyNameRes{
		ServerInfo: server,
	}, nil
}

func (s *sServer) ServerSearch(ctx context.Context, serverName string) (res *v1.ServerSearchRes, err error) {
	var servers []*model.Server
	err = dao.Server.Ctx(ctx).
		Where("server_name LIKE ? AND server_type = ?", "%"+serverName+"%", "public").
		Scan(&servers)
	if err != nil {
		return nil, errResponse.DbOperationErrorDefault()
	}
	return &v1.ServerSearchRes{
		Servers: servers,
	}, nil
}

func (s *sServer) ServerPages(ctx context.Context, in model.ServerPagesInput) (res []*model.ServerPages, total int, err error) {
	err = dao.Server.Ctx(ctx).
		Where("server_name LIKE ? AND server_type = ?", "%"+in.ServerName+"%", "public").
		Page(in.Page, in.PageSize).
		Scan(&res)
	if err != nil {
		return nil, 0, errResponse.DbOperationErrorDefault()
	}
	count, err := dao.Server.Ctx(ctx).
		Where("server_name LIKE ? AND server_type = ?", "%"+in.ServerName+"%", "public").
		Count()
	return res, count, nil
}

func (s *sServer) ServerInfo(ctx context.Context, serverId uint64) (res *v1.ServerInfoRes, err error) {
	if err = auth.IsServerCreator(ctx, serverId); err != nil {
		return nil, errResponse.NotAuthorized("你不是服务器创建者不能进行服务器设置")
	}
	var server *model.Server
	if err = dao.Server.Ctx(ctx).Where("server_id = ?", serverId).Scan(&server); err != nil {
		return nil, errResponse.DbOperationErrorDefault()
	}
	return &v1.ServerInfoRes{
		ServerInfo: server,
	}, nil
}

func (s *sServer) ServerInfoUpd(ctx context.Context, in model.ServerInfoUpdInput) (res *v1.ServerInfoUpdRes, err error) {
	if err = auth.IsServerCreator(ctx, in.ServerId); err != nil {
		return nil, errResponse.NotAuthorized("你不是服务器创建者不能进行服务器设置")
	}
	_, err = dao.Server.Ctx(ctx).
		Fields("server_description", "server_name", "server_type").
		Data(&in).
		Where("server_id = ?", in.ServerId).
		Update()
	if err != nil {
		return nil, errResponse.DbOperationError("更新失败")
	}
	if err = cache.DelServerListsCache(ctx, in.ServerId); err != nil {
		return nil, err
	}
	var server *model.Server
	err = dao.Server.Ctx(ctx).Where("server_id = ?", in.ServerId).Scan(&server)
	if err != nil {
		return nil, errResponse.DbOperationErrorDefault()
	}
	return &v1.ServerInfoUpdRes{
		ServerInfo: server,
	}, nil
}

func (s *sServer) ServerCount(ctx context.Context, _ *v1.ServerCountReq) (res *v1.ServerCountRes, err error) {
	count, err := dao.Server.Ctx(ctx).Count()
	if err != nil {
		return nil, errResponse.DbOperationErrorDefault()
	}
	return &v1.ServerCountRes{
		Count: uint64(count),
	}, nil
}

func (s *sServer) ServerImg(ctx context.Context, serverId uint64, file *ghttp.UploadFile) (res *v1.ServerImgRes, err error) {
	if err = auth.IsServerCreator(ctx, serverId); err != nil {
		return nil, errResponse.NotAuthorized("你不是服务器创建者不能进行服务器设置")
	}
	url, err := service.Qiniu().UploadFile(ctx, file, "server")
	if err != nil {
		return nil, errResponse.CodeInvalidParameter("上传失败")
	}
	update, err := dao.Server.Ctx(ctx).Fields("server_img_url").
		Data(g.Map{
			"server_img_url": url,
		}).
		Where("server_id = ?", serverId).
		Update()
	if err != nil || update == nil {
		return nil, errResponse.DbOperationError("操作失败")
	}
	if err = cache.DelServerListsCache(ctx, serverId); err != nil {
		return nil, err
	}
	var server *model.Server
	err = dao.Server.Ctx(ctx).
		Where("server_id = ?", serverId).
		Scan(&server)
	if err != nil {
		return nil, errResponse.DbOperationErrorDefault()
	}
	return &v1.ServerImgRes{
		ServerInfo: server,
	}, nil
}
