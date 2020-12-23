package confcontrollers

import (
	"github.com/linclin/gopub/src/controllers"
	"github.com/linclin/gopub/src/models"
	"time"
)

type CopyController struct {
	controllers.BaseController
}

func (c *CopyController) Get() {
	if c.User == nil || c.User.Id == 0 {
		c.SetJson(2, nil, "not login")
		return
	}
	if c.Project == nil || c.Project.Id == 0 {
		c.SetJson(1, nil, "Parameter error")
		return
	}
	c.Project.Name = c.Project.Name + " - copy"
	// 复制等价于新增，修改ID和时间为默认值
	c.Project.Id = 0
	c.Project.CreatedAt = time.Time{}
	c.Project.UpdatedAt = time.Time{}
	c.Project.UserId = uint(c.User.Id)
	_, err := models.AddProject(c.Project)
	if err != nil {
		c.SetJson(1, nil, "复制失败")
	}
	c.SetJson(0, c.Project, "")
	return

}
