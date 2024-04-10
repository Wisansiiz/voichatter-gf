package auth

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	"voichatter/internal/dao"
	"voichatter/utility/errResponse"
)

func IsServerCreator(ctx context.Context, serverId uint64) (err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	count, err := dao.Server.Ctx(ctx).Where("server_id = ? AND creator_user_id = ?", serverId, userId).Count()
	if err != nil {
		return errResponse.DbOperationError("操作失败")
	}
	if count == 0 {
		return errResponse.OperationFailed("权限不足")
	}
	return nil
}

// IsActivityCreator 活动创建者
func IsActivityCreator(ctx context.Context, activityId uint64) (err error) {
	userId := gconv.Uint64(ctx.Value("userId"))
	count, err := dao.Activity.Ctx(ctx).Where("activity_id = ? AND creator_user_id = ?", activityId, userId).Count()
	if err != nil {
		return errResponse.DbOperationError("操作失败")
	}
	if count == 0 {
		return errResponse.OperationFailed("权限不足")
	}
	return nil
}
