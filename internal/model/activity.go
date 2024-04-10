package model

import "github.com/gogf/gf/v2/os/gtime"

type Activity struct {
	ActivityId    uint64      `json:"activityId"     ` // 活动id
	ServerId      uint64      `json:"serverId"       ` // 服务器id
	ActivityTitle string      `json:"activityTitle"  ` // 活动主题/内容
	ActivityDesc  string      `json:"activityDesc"   ` // 活动描述
	CreatorUserId uint64      `json:"creatorUserId"  ` // 活动创建者id
	StartDate     *gtime.Time `json:"startDate"      ` // 开始日期
	EndDate       *gtime.Time `json:"endDate"        ` // 结束日期
}

type ActivityPages struct {
	ActivityId    uint64      `json:"activityId"     ` // 活动id
	ServerId      uint64      `json:"serverId"       ` // 服务器id
	ActivityTitle string      `json:"activityTitle"  ` // 活动主题/内容
	ActivityDesc  string      `json:"activityDesc"   ` // 活动描述
	CreatorUserId uint64      `json:"creatorUserId"  ` // 活动创建者id
	StartDate     *gtime.Time `json:"startDate"      ` // 开始日期
	EndDate       *gtime.Time `json:"endDate"        ` // 结束日期
}

type ActivityPagesInput struct {
	ActivityTitle string `json:"activityTitle"  `
	Page          int    `json:"page"           `
	PageSize      int    `json:"pageSize"       `
}

type ActivityCreateInput struct {
	ServerId      uint64      `json:"serverId"       `
	ActivityTitle string      `json:"activityTitle"  `
	ActivityDesc  string      `json:"activityDesc"   `
	CreatorUserId uint64      `json:"creatorUserId"  `
	StartDate     *gtime.Time `json:"startDate"      `
	EndDate       *gtime.Time `json:"endDate"        `
}

type ActivitySearchInput struct {
	ActivityTitle string `json:"activityTitle"  `
}

type ActivityUpdateInput struct {
	ActivityId    uint64      `json:"activityId"     `
	ServerId      uint64      `json:"serverId"       `
	ActivityTitle string      `json:"activityTitle"  `
	ActivityDesc  string      `json:"activityDesc"   `
	CreatorUserId uint64      `json:"creatorUserId"  `
	StartDate     *gtime.Time `json:"startDate"      `
	EndDate       *gtime.Time `json:"endDate"        `
}

type ActivityDeleteInput struct {
	ActivityId uint64 `json:"activityId" `
}
