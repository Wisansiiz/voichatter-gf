package group

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "voichatter/api/group/v1"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/service"
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

func (s sGroup) GroupList(ctx context.Context, serverId uint64) (res *v1.GroupListRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	count, err := dao.Member.Ctx(ctx).Where("server_id = ? AND user_id = ?", serverId, userId).Count()
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New("权限不足")
	}
	var groupList []model.GroupList
	err = dao.Group.Ctx(ctx).
		Where("server_id = ?", serverId).
		Scan(&groupList)
	if err != nil {
		return nil, err
	}
	for i, group := range groupList {
		err = dao.Channel.Ctx(ctx).
			Where("group_id", group.GroupId).
			Scan(&groupList[i].ChannelList)
		if err != nil {
			return nil, err
		}
	}
	var channelList []model.ChannelInfo
	err = dao.Channel.Ctx(ctx).
		Where("server_id = ? AND group_id IS NULL", serverId).
		Scan(&channelList)
	if err != nil {
		return nil, err
	}
	return &v1.GroupListRes{
		ChannelList: &channelList,
		GroupList:   &groupList,
	}, nil
}
