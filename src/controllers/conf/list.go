package confcontrollers

import (
	"github.com/astaxie/beego/orm"
	taskcontrollers "github.com/hpf0532/gopub/src/controllers/task"
	"github.com/linclin/gopub/src/controllers"
	"github.com/linclin/gopub/src/library/common"
)

type ListController struct {
	controllers.BaseController
}

func (c *ListController) Get() {
	page, _ := c.GetInt("page", 0)
	start := 0
	length, _ := c.GetInt("length", 200000)
	if page > 0 {
		start = (page - 1) * length
	}
	selectInfo := c.GetString("select_info")
	where := ""
	if selectInfo != "" {
		where = "  and(`name` LIKE '%" + selectInfo + "%' )"
	}
	var projects []orm.Params
	o := orm.NewOrm()

	o.Raw("SELECT *, (SELECT realname FROM `user` WHERE `user`.id=project.user_id LIMIT 1) as realname FROM `project`  WHERE 1=1 "+where+" ORDER BY id LIMIT ?,?", start, length).Values(&projects)
	for _, item := range projects {
		item["level_name"] = taskcontrollers.GetProjectLevel(common.GetInt(item["level"]))
	}
	var count []orm.Params
	total := 0
	o.Raw("SELECT count(id) FROM `project` WHERE 1=1 " + where).Values(&count)
	if len(count) > 0 {
		total = common.GetInt(count[0]["count(id)"])
	}
	c.SetJson(0, map[string]interface{}{"total": total, "currentPage": page, "table_data": projects}, "")

	return

}
