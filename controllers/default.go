package controllers



type UserController struct {
	BaseController
}


func (c *UserController) Get(){
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.Data["IsHome"] = true
	c.Prepare()
	c.TplName = "home.html"

    // type u struct{ 
	// 	Name string
	// 	Age  int
	// 	Sex  string   
    // }
    // user := &u{
    // 	Name:"joe",
    // 	Age: 20,
    // 	Sex: "man",
    	
    // } 
	// c.Data["User"] = user
	
	// nums := []int{1,2,3,4,5,6,7,8}
	// c.Data["Nums"]=nums

	// c.Data["Html"]="<p>hello world<p>"
 }


func (c *UserController) Category(){
	c.Data["userid"] = "beego.me"
	c.Data["IsCategory"] = true
	c.Data["hobby"] = []string{"code","book"}
	c.Prepare()
	c.TplName = "Category.html"

}

func (c *UserController) Pageinfo(){
	c.Data["userid"] = "beego.me"
	c.Data["IsCategory"] = true
	c.Data["hobby"] = []string{"code","book"}
	c.Prepare()
	c.TplName = "Pageinfo.html"

}

func (c *UserController) Setting(){
	c.Data["userid"] = "beego.me"
	c.Data["IsCategory"] = true
	c.Data["hobby"] = []string{"code","book"}
	c.Prepare()
	c.TplName = "Category.html"

}

