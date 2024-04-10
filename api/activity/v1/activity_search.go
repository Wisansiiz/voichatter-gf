package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type ActivitySearchReq struct {
	g.Meta        `path:"/activities/search" method:"get" tags:"ActivityService" summary:""`
	ActivityTitle string `p:"activityTitle" v:""`
}

type ActivitySearchRes struct {
	Activities []*model.Activity `json:"activities"`
}
