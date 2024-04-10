package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"voichatter/internal/model"
)

type ActivityCreateReq struct {
	g.Meta        `path:"/activities" method:"post" tags:"ActivityService" summary:""`
	ActivityTitle string      `p:"activityTitle" v:"required|length:1,10#请输入活动标题|活动标题长度为:{min}到:{max}位"`
	ActivityDesc  string      `p:"activityDesc"  v:"required|length:1,99#请输入活动描述|活动描述长度为:{min}到:{max}位" `
	ServerId      uint64      `p:"serverId"      v:"required" `
	StartDate     *gtime.Time `p:"startDate"     v:"required" `
	EndDate       *gtime.Time `p:"endDate"       v:"required" `
}

type ActivityCreateRes struct {
	Activities *model.Activity `json:"activities"`
}
