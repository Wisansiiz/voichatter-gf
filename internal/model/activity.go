package model

import "github.com/gogf/gf/v2/os/gtime"

type Activity struct {
	ActivityId    int64       `json:"activityId"     ` // 活动id
	ServerId      int64       `json:"serverId"       ` // 服务器id
	ActivityTitle string      `json:"activityTitle"  ` // 活动主题/内容
	ActivityDesc  string      `json:"activityDesc"   ` // 活动描述
	CreatorUserId int64       `json:"creatorUserId" `  // 活动创建者id
	StartDate     *gtime.Time `json:"startDate"      ` // 开始日期
	EndDate       *gtime.Time `json:"endDate"        ` // 结束日期
}

type ActivityPages struct {
	ActivityId    int64       `json:"activityId"     ` // 活动id
	ServerId      int64       `json:"serverId"       ` // 服务器id
	ActivityTitle string      `json:"activityTitle"  ` // 活动主题/内容
	ActivityDesc  string      `json:"activityDesc"   ` // 活动描述
	CreatorUserId int64       `json:"creatorUserId" `  // 活动创建者id
	StartDate     *gtime.Time `json:"startDate"      ` // 开始日期
	EndDate       *gtime.Time `json:"endDate"        ` // 结束日期
}

type ActivityPagesInput struct {
	Page     int `json:"page" `
	PageSize int `json:"pageSize"`
}
