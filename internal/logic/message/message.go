package message

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "voichatter/api/message/v1"
	"voichatter/internal/dao"
	"voichatter/internal/model"
	"voichatter/internal/service"
	"voichatter/utility/errResponse"
)

type (
	sMessage struct{}
)

func init() {
	service.RegisterMessage(New())
}

func New() service.IMessage {
	return &sMessage{}
}

func (s *sMessage) MessageList(ctx context.Context, in model.Message) (res *v1.MessageListRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	ids, err := dao.Member.Ctx(ctx).
		Fields("server_id").
		Where("user_id = ?", userId).
		Array()
	if err != nil {
		return nil, errResponse.DbOperationErrorDefault()
	}
	var m = make(map[uint64]bool)
	for _, id := range ids {
		m[id.Uint64()] = true
	}
	if !m[in.ServerId] {
		return nil, errResponse.NotAuthorized("Forbidden")
	}
	if err != nil {
		return nil, err
	}
	var messages []model.MessageInfo
	err = dao.Message.Ctx(ctx).
		InnerJoin("user", "message.sender_user_id = user.user_id").
		Fields("message.attachment,message.send_date,message.sender_user_id,message.content,message.server_id,message.message_id,user.avatar_url,user.username").
		Where("channel_id = ?", in.ChannelId).
		Scan(&messages)
	if err != nil {
		return nil, errResponse.DbOperationError("获取消息列表失败")
	}
	return &v1.MessageListRes{
		MessageList: &messages,
	}, nil
}
