package group

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "voichatter/api/group/v1"
	"voichatter/internal/consts"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/service"
	"voichatter/utility/errResponse"
)

type (
	sGroup struct{}
)

func init() {
	service.RegisterGroup(New())
}
func New() service.IGroup {
	return &sGroup{}
}

func (s *sGroup) GroupList(ctx context.Context, serverId uint64) (res *v1.GroupListRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))

	serverVoGroup := fmt.Sprintf("%s-%d-%s-%d", consts.ServerId, serverId, consts.GroupList, userId)
	serverVoChannel := fmt.Sprintf("%s-%d-%s-%d", consts.ServerId, serverId, consts.ChannelList, userId)

	count, err := dao.Member.Ctx(ctx).Where("server_id = ? AND user_id = ?", serverId, userId).Count()
	if err != nil {
		return nil, errResponse.DbOperationError("查询失败")
	}
	if count == 0 {
		return nil, errResponse.OperationFailed("权限不足")
	}
	// 查询群列表
	var groupList []model.GroupList
	getGroupList, err := g.Redis().Get(ctx, serverVoGroup)
	if err != nil {
		return nil, errResponse.DbOperationError("查询失败")
	}
	if err = gconv.Struct(getGroupList, &groupList); err != nil {
		return nil, errResponse.OperationFailed("失败")
	}
	// 查询频道列表
	var channelList []model.ChannelInfo
	getChannelList, err := g.Redis().Get(ctx, serverVoChannel)
	if err = gconv.Struct(getChannelList, &channelList); err != nil {
		return nil, errResponse.OperationFailed("失败")
	}
	if groupList != nil || channelList != nil {
		res = &v1.GroupListRes{
			ChannelList: &channelList,
			GroupList:   &groupList,
		}
	}
	// 有数据返回
	if res != nil {
		return
	}
	// 没数据从数据库中查询
	err = dao.Group.Ctx(ctx).
		Where("server_id = ?", serverId).
		Scan(&groupList)
	if err != nil {
		return nil, errResponse.DbOperationError("查询失败")
	}
	for i, group := range groupList {
		err = dao.Channel.Ctx(ctx).
			Where("group_id = ?", group.GroupId).
			Scan(&groupList[i].ChannelList)
		if err != nil {
			return nil, errResponse.DbOperationError("查询失败")
		}
	}
	err = dao.Channel.Ctx(ctx).
		Where("server_id = ? AND group_id IS NULL", serverId).
		Scan(&channelList)
	if err != nil {
		return nil, errResponse.DbOperationError("查询失败")
	}
	// 再将数据存入redis，缓存时间为一天
	if err = g.Redis().SetEX(ctx, serverVoGroup, groupList, int64(gtime.D)); err != nil {
		return nil, errResponse.DbOperationError("设置失败")
	}
	if err = g.Redis().SetEX(ctx, serverVoChannel, channelList, int64(gtime.D)); err != nil {
		return nil, errResponse.DbOperationError("设置失败")
	}
	return &v1.GroupListRes{
		ChannelList: &channelList,
		GroupList:   &groupList,
	}, nil
}

func (s *sGroup) GroupCreate(ctx context.Context, req *v1.GroupCreateReq) (res *v1.GroupCreateRes, err error) {

	panic("implement me")
}

func (s *sGroup) GroupModify(ctx context.Context, req *v1.GroupModifyReq) (res *v1.GroupModifyRes, err error) {

	panic("implement me")
}

func (s *sGroup) GroupRemove(ctx context.Context, req *v1.GroupRemoveReq) (res *v1.GroupRemoveRes, err error) {

	panic("implement me")
}
