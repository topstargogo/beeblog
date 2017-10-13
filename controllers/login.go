package controllers

import(
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"    
    "fmt"
    "time"
	"strconv"
	"star/models"
	"github.com/astaxie/beego/validation"
	

)



func PwGen(pass string) string {
	salt := strconv.FormatInt(time.Now().UnixNano()%9000+1000, 10)
	return Base64Encode(Sha1(Md5(pass)+salt) + salt)
}

func PwCheck(pwd, saved string) bool {
	saved = Base64Decode(saved)
	if len(saved) < 4 {
		return false
	}
	salt := saved[len(saved)-4:]
	return Sha1(Md5(pwd)+salt)+salt == saved
}

func Sha1(s string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(s)))
}

func Md5(s string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(s)))
}

func Base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func Base64Decode(s string) string {
	res, _ := base64.StdEncoding.DecodeString(s)
	return string(res)
}

type LoginController struct{
	BaseController

}

func (c *LoginController) Login(){
    

    loginname := c.GetString("loginname")
	loginpw := c.GetString("loginpw")
	


	valid := validation.Validation{}

	
	valid.Required(loginname, "name")
	valid.Required(loginpw, "pw")
	u := &models.User{
		Username:  loginname,
		
	
	} 
	
	switch {
		
			case valid.HasErrors():
			   c.Ctx.WriteString("用户名或密码不能为空")
			   return
	    	case u.ReadDB() != nil:
			    c.Ctx.WriteString("用户名错误")
	            return
		    case PwCheck(loginpw, u.Password) == false:
				c.Ctx.WriteString("密码错误")
				return
			default: 
				c.DoLogin()
				c.Prepare()
				c.TplName = "home.html"

		
			}


    
}

func (c *LoginController) Logout(){
    
			 c.DoLogout()
			 
  
}