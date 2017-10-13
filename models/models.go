package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// 分类
type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

// 文章
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}

type User struct {
	Username string `orm:"pk"`
	Email    string
	Password string
	Mobile   string
}



func init() {
    orm.Debug = true
	orm.RegisterModel(new(Category), new(Topic), new(User))	
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:123456@/beeblog?charset=utf8")
	orm.RunSyncdb("default",false,true)
}


func (u *User) ReadDB() (err error) {
	o := orm.NewOrm()
	err = o.Read(u)
	return
}

func (u User) Create() (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(&u)
	return
}

func (u User) Update() (err error) {
	o := orm.NewOrm()
	_, err = o.Update(&u)
	return
}


func (u User) ExistUsername() bool {
	o := orm.NewOrm()
	if err := o.Read(&u); err == orm.ErrNoRows {
		return false
	}
	return true
}

func (u User) ExistEmail() bool {
	o := orm.NewOrm()
	return o.QueryTable("user").Filter("Email", u.Email).Exist()
}
