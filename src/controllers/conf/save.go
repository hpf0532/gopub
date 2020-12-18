package confcontrollers

import (
	"github.com/linclin/gopub/src/controllers"

	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/linclin/gopub/src/models"
)

type SaveController struct {
	controllers.BaseController
}

func (c *SaveController) Post() {
	//projectId,_:=c.GetInt("projectId",0)
	beego.Info(string(c.Ctx.Input.RequestBody))
	var project models.Project
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &project)
	if err != nil {
		c.SetJson(1, nil, "数据格式错误")
		return
	}
	if project.Id != 0 {
		err = models.UpdateProjectById(&project)
	} else {
		// 加入配置创建人
		if c.User != nil || c.User.Id != 0 {
			project.UserId = uint(c.User.Id)
		}
		_, err = models.AddProject(&project)
	}
	if err != nil {
		c.SetJson(1, nil, "数据库更新错误")
		return
	}
	c.SetJson(0, project, "修改成功")
	return
}
