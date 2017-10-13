package controllers

import (
	"github.com/astaxie/beego"
)


type BaseController struct {
	beego.Controller
}



// func (c *BaseController) Prepare() {
// 	loginname := c.GetString("loginname")


// 	c.SetSession("username", loginname)
// 	v := c.GetSession("username")
// 	fmt.Println(v)
// 	c.TplName = "home.html" 

// }




func (c *BaseController) Prepare() {

	if c.IsLogin() {
		
		c.Data["username"] = c.GetSession("username")
	}
}
func (c *BaseController) DoLogin() {
	loginname := c.GetString("loginname")
	c.SetSession("username", loginname)
}

func (c *BaseController) DoLogout() {
	c.DestroySession()
	c.Redirect("/", 302)
}

func (c *BaseController) IsLogin() bool {
	return c.GetSession("username") != nil
}

func (c *BaseController) CheckLogin() {
	if !c.IsLogin() {
		c.Redirect("/home.html", 302)
		c.Abort("302")
	}
}




