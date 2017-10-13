package controllers

import (
	"star/models"

    
	"github.com/astaxie/beego/validation"
)

func (c *UserController) API_Profile() {
	type user struct {
		Userid string
		Hobby  []string
	}
	u := user{
		"jike",
		[]string{"code", "book"},
	}
	c.Data["json"] = u
	c.ServeJSON()
}


func (c *UserController) Register() {

	uname := c.GetString("username")
	mail := c.GetString("email")
	pwd := c.GetString("password")
	pwd1 := c.GetString("password1")
	mb := c.GetString("mobile")

	

	valid := validation.Validation{}

	valid.Email(mail, "Email")
	valid.Required(mb, "mobile")
	valid.Required(uname, "username")
	valid.Required(pwd, "Password")
	valid.Required(pwd1, "Password1")

    
	switch {

	case valid.HasErrors():
	c.Ctx.WriteString("电话不能为空")
    return
	case pwd != pwd1:
		valid.Error("两次密码不一致")
		c.Ctx.WriteString("两次密码不一致")
        return
	default: 
		u := &models.User{
            Username:  uname,
			Email:     mail,
			Password:  PwGen(pwd),
			Mobile:    mb,
		} 

		switch {
		case u.ExistUsername():
			valid.Error("用户名被占用")
			c.Ctx.WriteString("用户名被占用")
            return
		case u.ExistEmail():
			valid.Error("邮箱被占用")
			c.Ctx.WriteString("邮箱被占用")
            return
		default:
			err := u.Create()
			if err == nil {
				c.Ctx.WriteString("创建成功")
				return
			}
			
		} 

	}
    
	c.TplName="home.html"
	
}
