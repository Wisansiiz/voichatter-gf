package server

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "voichatter/api/server/v1"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/model/entity"
	"voichatter/internal/service"
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
	}
	lastInsertId, err := dao.Server.Ctx(ctx).InsertAndGetId(&server)
	if err != nil {
		return nil, err
	}
	member := entity.Member{
		UserId:       server.CreatorUserId,
		ServerId:     uint64(lastInsertId),
		JoinDate:     gtime.Now(),
		SPermissions: "admin",
	}
	_, err = dao.Member.Ctx(ctx).Insert(&member)
	if err != nil {
		return nil, err
	}
	return
}

func (s *sServer) ServerJoin(ctx context.Context, serverId uint64) (res *v1.ServerJoinRes, err error) {
	// 使用gDB ORM进行查询操作
	count, err := dao.Server.Ctx(ctx).
		Where("server_id = ? AND server_type = ?", serverId, "public").
		Count()
	if err != nil {
		return nil, gerror.New("服务器不存在")
	}
	if count == 0 {
		return nil, gerror.New("服务器不存在或不是公开的")
	}
	//user := service.BizCtx().Get(ctx).User
	userId := gconv.Uint64(ctx.Value("userId"))
	// 判断是否已经加入过服务器
	count, err = g.DB().Model("member").
		Where("server_id = ? AND user_id = ?", serverId, userId).
		Count()
	if err != nil {
		return nil, errors.New("查询是否加入过服务器时出错")
	}
	if count > 0 {
		return nil, gerror.New("已经加入过服务器")
	}
	// 添加成员
	m := entity.Member{
		ServerId:     serverId,
		UserId:       userId,
		JoinDate:     gtime.Now(),
		SPermissions: "member",
	}
	_, err = g.DB().Model("member").Insert(m)
	if err != nil {
		return nil, gerror.New("添加成员失败")
	}
	return
}

func (s *sServer) ServerDel(ctx context.Context, serverId uint64) (res *v1.ServerDelRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	result, err := dao.Server.Ctx(ctx).
		Where("server_id = ? AND creator_user_id = ?", serverId, userId).
		Delete()
	row, _ := result.RowsAffected()
	if err != nil || row == 0 {
		return nil, gerror.New("权限不足")
	}
	return nil, nil
}
