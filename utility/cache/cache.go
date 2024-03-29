package cache

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"voichatter/internal/consts"
	"voichatter/internal/dao"
	"voichatter/utility/errResponse"
)

// DelGroupVoCache 删除缓存
func DelGroupVoCache(ctx context.Context, serverId uint64) error {
	ids, err := dao.Member.Ctx(ctx).
		Fields("DISTINCT(user_id)").
		Where("server_id = ?", serverId).
		Array()
	if err != nil {
		return errResponse.DbOperationErrorDefault()
	}
	for _, id := range ids {
		_, err = g.Redis().Del(ctx,
			fmt.Sprintf("%s-%d-%s-%d", consts.ServerId, serverId, consts.GroupList, gconv.Uint64(id)))
		if err != nil {
			return errResponse.DbOperationError("删除GroupList缓存失败")
		}
		_, err = g.Redis().Del(ctx,
			fmt.Sprintf("%s-%d-%s-%d", consts.ServerId, serverId, consts.ChannelList, gconv.Uint64(id)))
		if err != nil {
			return errResponse.DbOperationError("删除ChannelList缓存失败")
		}
	}
	return err
}

// DelServerListsCache 删除缓存
func DelServerListsCache(ctx context.Context, serverId uint64) error {
	ids, err := dao.Member.Ctx(ctx).
		Fields("DISTINCT(user_id)").
		Where("server_id = ?", serverId).
		Array()
	if err != nil {
		return errResponse.DbOperationErrorDefault()
	}
	for _, id := range ids {
		_, err = g.Redis().Del(ctx,
			fmt.Sprintf("%s-%d", consts.ServerList, gconv.Uint64(id)))
		if err != nil {
			return errResponse.DbOperationError("删除ServerLists缓存失败")
		}
	}
	return err
}

func DelServerUsersCache(ctx context.Context, serverId uint64) error {
	_, err := g.Redis().Del(ctx,
		fmt.Sprintf("%s-%d", consts.ServerUsers, serverId))
	if err != nil {
		return errResponse.DbOperationError("删除ServerList缓存失败")
	}
	return err
}

func DelServerListCache(ctx context.Context, userId uint64) error {
	_, err := g.Redis().Del(ctx,
		fmt.Sprintf("%s-%d", consts.ServerList, userId))
	if err != nil {
		return errResponse.DbOperationError("删除ServerList缓存失败")
	}
	return err
}
