package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"voichatter/internal/model"
)

type ActivityPagesReq struct {
	g.Meta        `path:"/activities" method:"get" tags:"ActivityService" summary:"ActivityPages"`
	ActivityTitle string `p:"activityTitle" v:""`
	Page          int    `p:"page"     example:"1" d:"1" v:"min:1#页码最小值不能低于1"  dc:"当前页码"`
	PageSize      int    `p:"pageSize" example:"5" d:"5" v:"min:1|max:200#每页数量最小值不能低于1|最大值不能大于200" dc:"每页数量"`
}

type ActivityPagesRes struct {
	ActivityPages []*model.ActivityPages `json:"activityPages"`
	PageTotal     int                    `json:"pageTotal"`
}
