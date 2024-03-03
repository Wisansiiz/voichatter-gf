package message

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "voichatter/api/message/v1"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/service"
)

type (
	sMessage struct{}
)

func (s *sMessage) MessageList(ctx context.Context, in model.Message) (res *v1.MessageListRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	ids, err := dao.Member.Ctx(ctx).
		Fields("server_id").
		Where("user_id = ?", userId).
		Array()
	var m = make(map[uint64]bool)
	for _, id := range ids {
		m[id.Uint64()] = true
	}
	if !m[in.ServerId] {
		return nil, gerror.New("Forbidden")
	}
	if err != nil {
		return nil, err
	}
	var messages []model.MessageInfo
	err = dao.Message.Ctx(ctx).
		Where("channel_id = ?", in.ChannelId).
		Scan(&messages)
	if err != nil {
		return nil, err
	}
	return &v1.MessageListRes{
		MessageList: &messages,
	}, nil
}

func init() {
	service.RegisterMessage(New())
}

func New() service.IMessage {
	return &sMessage{}
}
