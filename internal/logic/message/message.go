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
	count, err := dao.Member.Ctx(ctx).
		Where("server_id = ? AND user_id = ?", in.ServerId, userId).
		Count()
	if err != nil || count == 0 {
		return nil, errResponse.NotAuthorized("Forbidden")
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

func (s *sMessage) MessagePages(ctx context.Context, in model.MessagePagesRes) (res *v1.MessagePagesRes, err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	count, err := dao.Member.Ctx(ctx).
		Where("server_id = ? AND user_id = ?", in.ServerId, userId).
		Count()
	if err != nil || count == 0 {
		return nil, errResponse.NotAuthorized("Forbidden")
	}
	count, err = dao.Message.Ctx(ctx).
		Where("channel_id = ?", in.ChannelId).
		Count()
	if err != nil {
		return nil, errResponse.DbOperationError("获取消息列表失败")
	}
	var messages []*model.MessageInfo
	if err = dao.Message.Ctx(ctx).
		InnerJoin("user", "message.sender_user_id = user.user_id").
		Fields("message.attachment,message.send_date,message.sender_user_id,message.content,message.server_id,message.message_id,user.avatar_url,user.username").
		Where("channel_id = ?", in.ChannelId).
		Page((count+in.PageSize-1)/in.PageSize-in.Page+1, in.PageSize). // 为了消息的可读性，这里直接从最后一页开始取
		//Page(in.Page, in.PageSize). 										  // 如果按照传统的分页模式应该这样
		Scan(&messages); err != nil {
		return nil, errResponse.DbOperationError("获取消息列表失败")
	}
	return &v1.MessagePagesRes{
		MessagePages: &model.MessagePagesRep{
			MessageInfo: messages,
			PageTotal:   count,
		},
	}, nil
}
