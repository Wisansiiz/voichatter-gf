package channel

import (
	"context"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "voichatter/api/channel/v1"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/model/entity"
	"voichatter/internal/service"
	"voichatter/utility/cache"
	"voichatter/utility/errResponse"
)

type (
	sChannel struct{}
)

func init() {
	service.RegisterChannel(New())
}
func New() service.IChannel {
	return &sChannel{}
}

func (s *sChannel) ChannelCreate(ctx context.Context, in model.ChannelCreateInput) (res *v1.ChannelCreateRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	count, err := dao.Member.Ctx(ctx).
		Where("server_id = ? AND user_id = ?", in.ServerId, userId).
		Where("s_permissions = ?", "admin").
		WhereOr("s_permissions = ?", "owner").
		Count()
	if err != nil {
		return nil, errResponse.DbOperationErrorDefault()
	}
	if count == 0 {
		return nil, errResponse.OperationFailed("权限不足")
	}
	channelId, err := dao.Channel.Ctx(ctx).
		InsertAndGetId(&entity.Channel{
			ServerId:     in.ServerId,
			ChannelName:  in.ChannelName,
			Type:         in.Type,
			CreateUserId: userId,
			CreationDate: gtime.Now(),
		})
	if err != nil {
		return nil, errResponse.DbOperationError("新增失败")
	}
	// 删除缓存
	err = cache.DelGroupVoCache(ctx, in.ServerId)
	if err != nil {
		return nil, err
	}
	return &v1.ChannelCreateRes{
		Channel: &model.ChannelInfo{
			ChannelId:    gconv.Uint64(channelId),
			ChannelName:  in.ChannelName,
			Type:         in.Type,
			ServerId:     in.ServerId,
			CreateUserId: userId,
		},
	}, nil
}

func (s *sChannel) ChannelModify(ctx context.Context, in model.ChannelModifyInput) (res *v1.ChannelModifyRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	count, err := dao.Server.Ctx(ctx).Where("server_id = ? AND creator_user_id = ?", in.ServerId, userId).Count()
	if err != nil || count == 0 {
		return nil, errResponse.DbOperationError("权限不足")
	}
	// 删除缓存
	err = cache.DelGroupVoCache(ctx, in.ServerId)
	if err != nil {
		return nil, err
	}
	update, err := dao.Channel.Ctx(ctx).
		Fields("channel_name").
		Data(&entity.Channel{ChannelName: in.ChannelName}).
		Where("channel_id = ?", in.ChannelId).
		Update()
	if err != nil || update == nil {
		return nil, errResponse.DbOperationError("修改失败")
	}
	var channelInfo model.ChannelInfo
	err = dao.Channel.Ctx(ctx).Where("channel_id = ?", in.ChannelId).Scan(&channelInfo)
	if err != nil {
		return nil, errResponse.DbOperationErrorDefault()
	}
	return &v1.ChannelModifyRes{
		Channel: &channelInfo,
	}, nil
}

func (s *sChannel) ChannelRemove(ctx context.Context, in model.ChannelRemoveInput) (res *v1.ChannelRemoveRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	count, err := dao.Server.Ctx(ctx).Where("server_id = ? AND creator_user_id = ?", in.ServerId, userId).Count()
	if err != nil || count == 0 {
		return nil, errResponse.DbOperationError("权限不足")
	}
	// 删除缓存
	err = cache.DelGroupVoCache(ctx, in.ServerId)
	if err != nil {
		return nil, err
	}
	result, err := dao.Channel.Ctx(ctx).
		Where("channel_id = ?", in.ChannelId).
		Delete()
	if err != nil || result == nil {
		return nil, errResponse.DbOperationError("删除失败")
	}
	return
}
