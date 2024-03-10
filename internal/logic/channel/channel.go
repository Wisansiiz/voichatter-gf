package channel

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "voichatter/api/channel/v1"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/model/entity"
	"voichatter/internal/service"
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
		return nil, err
	}
	if count == 0 {
		return nil, gerror.New("权限不足")
	}
	_, err = dao.Channel.Ctx(ctx).
		Insert(&entity.Channel{
			ServerId:     in.ServerId,
			ChannelName:  in.ChannelName,
			Type:         in.Type,
			CreateUserId: userId,
			CreationDate: gtime.Now(),
		})
	return
}
