package server

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "voichatter/api/server/v1"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/model/entity"
	"voichatter/internal/service"
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
	var servers []model.Server
	err = g.Model("server s").
		Fields("s.*").
		LeftJoin("member m", "s.server_id = m.server_id").
		Where("m.user_id = ?", userId).
		Scan(&servers)
	if err != nil {
		return nil, errResponse.DbOperationError("查询服务器列表出错了")
	}
	return &v1.ServerListRes{
		ServerList: &servers,
	}, err
}

func (s *sServer) ServerCreate(ctx context.Context, in model.ServerCreateInput) (res *v1.ServerCreateRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	server := entity.Server{
		ServerName:    in.ServerName,
		CreatorUserId: userId,
		CreateDate:    gtime.Now(),
		ServerType:    in.ServerType,
		ServerImgUrl:  in.ServerImgUrl,
	}
	lastInsertId, err := dao.Server.Ctx(ctx).InsertAndGetId(&server)
	if err != nil {
		return nil, errResponse.DbOperationError("查询出错")
	}
	member := entity.Member{
		UserId:       server.CreatorUserId,
		ServerId:     uint64(lastInsertId),
		JoinDate:     gtime.Now(),
		SPermissions: "owner",
	}
	_, err = dao.Member.Ctx(ctx).Insert(&member)
	if err != nil {
		return nil, errResponse.DbOperationError("查询出错")
	}
	return &v1.ServerCreateRes{
		Server: &model.Server{
			ServerId:     uint64(lastInsertId),
			ServerName:   in.ServerName,
			ServerType:   in.ServerType,
			ServerImgUrl: in.ServerImgUrl,
		},
	}, nil
}

func (s *sServer) ServerJoin(ctx context.Context, serverId uint64) (res *v1.ServerJoinRes, err error) {
	var server *model.Server
	// 服务器是否公开
	err = dao.Server.Ctx(ctx).
		Where("server_id = ? AND server_type = ?", serverId, "public").
		Scan(&server)
	if err != nil {
		return nil, errResponse.DbOperationError("查询出错")
	}
	if server == nil {
		return nil, errResponse.OperationFailed("服务器不存在或不是公开的")
	}
	userId := gconv.Uint64(ctx.Value("userId"))
	// 判断是否已经加入过服务器
	count, err := g.DB().Model("member").
		Where("server_id = ? AND user_id = ?", serverId, userId).
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
	return &v1.ServerJoinRes{
		Server: server,
	}, nil
}

func (s *sServer) ServerDel(ctx context.Context, serverId uint64) (res *v1.ServerDelRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	result, err := dao.Server.Ctx(ctx).
		Where("server_id = ? AND creator_user_id = ?", serverId, userId).
		Delete()
	row, _ := result.RowsAffected()
	if err != nil || row == 0 {
		return nil, errResponse.OperationFailed("查询失败, 权限不足")
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
	return &v1.ServerModifyNameRes{
		ServerInfo: server,
	}, nil
}
