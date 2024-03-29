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
	"voichatter/internal/model/entity"
	"voichatter/internal/service"
	"voichatter/utility/auth"
	"voichatter/utility/cache"
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

	// 查询群列表
	var groupList []*model.GroupList
	getGroupList, err := g.Redis().Get(ctx, serverVoGroup)
	if err != nil {
		return nil, errResponse.DbOperationError("查询失败")
	}
	if err = getGroupList.Struct(&groupList); err != nil {
		return nil, errResponse.OperationFailed("失败")
	}
	// 查询频道列表
	var channelList []*model.ChannelInfo
	getChannelList, err := g.Redis().Get(ctx, serverVoChannel)
	if err = getChannelList.Struct(&channelList); err != nil {
		return nil, errResponse.OperationFailed("失败")
	}
	if groupList != nil || channelList != nil {
		res = &v1.GroupListRes{
			ChannelList: channelList,
			GroupList:   groupList,
		}
	}
	// 有数据返回
	if res != nil {
		return
	}

	count, err := dao.Member.Ctx(ctx).Where("server_id = ? AND user_id = ?", serverId, userId).Count()
	if err != nil {
		return nil, errResponse.DbOperationError("查询失败")
	}
	if count == 0 {
		return nil, errResponse.OperationFailed("权限不足")
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
		ChannelList: channelList,
		GroupList:   groupList,
	}, nil
}

func (s *sGroup) GroupCreate(ctx context.Context, in model.GroupCreateInput) (res *v1.GroupCreateRes, err error) {
	if err = auth.IsServerCreator(ctx, in.ServerId); err != nil {
		return nil, err
	}
	id, err := dao.Group.Ctx(ctx).InsertAndGetId(entity.Group{
		ServerId:  in.ServerId,
		GroupName: in.GroupName,
	})
	if err != nil {
		return nil, errResponse.DbOperationError("新增失败")
	}
	groupInfo := model.Group{
		GroupId:   gconv.Uint64(id),
		ServerId:  in.ServerId,
		GroupName: in.GroupName,
	}
	if err = cache.DelGroupVoCache(ctx, in.ServerId); err != nil {
		return nil, err
	}
	return &v1.GroupCreateRes{
		Group: &groupInfo,
	}, nil
}

func (s *sGroup) GroupModify(ctx context.Context, in model.Group) (res *v1.GroupModifyRes, err error) {
	if err = auth.IsServerCreator(ctx, in.ServerId); err != nil {
		return nil, err
	}
	update, err := dao.Group.Ctx(ctx).
		Fields("group_name").
		Data(entity.Group{GroupName: in.GroupName}).
		Where("server_id = ? AND group_id = ?", in.ServerId, in.GroupId).
		Update()
	if err != nil || update == nil {
		return nil, errResponse.OperationFailed("无该分组")
	}
	// 删除缓存
	if err = cache.DelGroupVoCache(ctx, in.ServerId); err != nil {
		return nil, err
	}
	var group *model.Group
	if err = dao.Group.Ctx(ctx).Where("group_id = ?", in.GroupId).Scan(&group); err != nil {
		return nil, errResponse.DbOperationError("查询失败")
	}
	return &v1.GroupModifyRes{
		Group: group,
	}, nil
}

func (s *sGroup) GroupRemove(ctx context.Context, in model.GroupRemoveInput) (res *v1.GroupRemoveRes, err error) {
	if err = auth.IsServerCreator(ctx, in.ServerId); err != nil {
		return nil, err
	}
	_, err = dao.Group.Ctx(ctx).Where("group_id = ?", in.GroupId).Delete()
	if err != nil {
		return nil, errResponse.DbOperationError("删除失败")
	}
	_, err = dao.Channel.Ctx(ctx).Where("group_id = ?", in.GroupId).Delete()
	if err != nil {
		return nil, errResponse.DbOperationError("删除分组下频道失败")
	}
	if err = cache.DelGroupVoCache(ctx, in.ServerId); err != nil {
		return nil, err
	}
	return nil, nil
}
