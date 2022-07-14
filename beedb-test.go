package main

import (
	"fmt"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

type Userinfo struct {
	Uid        int `orm:"column:(uid);pk"` //如果表的主鍵不是 id，那麼需要加上 pk 註釋，明確的說這個欄位是主鍵
	Username   string
	Departname string
	Created    time.Time
}

type User struct {
	Uid        int `orm:"column:(uid);pk"` //如果表的主鍵不是 id，那麼需要加上 pk 註釋，明確的說這個欄位是主鍵
	Name       string
	Profile    *Profile `orm:"rel(one)"`      // OneToOne relation
	Post       []*Post  `orm:"reverse(many)"` // 設定一對多的反向關係
	Departname string
}

type Profile struct {
	Id   int
	Age  int16
	User *User `orm:"reverse(one)"` // 設定一對一反向關係(可選)
}

type Post struct {
	Id    int
	Title string
	User  *User  `orm:"rel(fk)"`
	Tags  []*Tag `orm:"rel(m2m)"` //設定一對多關係
}

type Tag struct {
	Id    int
	Name  string
	Posts []*Post `orm:"reverse(many)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:root@/test?charset=utf8", 30)
	orm.RegisterModel(new(Userinfo), new(User), new(Profile), new(Post), new(Tag)) // will create table in db ; if not existed skip
	orm.RunSyncdb("default", false, true)
	orm.Debug = false //
}

func main() {
	o := orm.NewOrm()

	var user User
	user.Name = "zxxx"
	user.Departname = "zxxx"

	id, err := o.Insert(&user)
	if err == nil {
		fmt.Println(id)
	} else {
		fmt.Println(err)
	}

	// user := User{Name: "slene"}

	// // insert record
	// id, err := o.Insert(&user)
	// fmt.Printf("ID: %d, ERR: %v\n", id, err)

	// // update user record
	// user.Name = "astaxie"
	// num, err := o.Update(&user)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)

	// // read user record
	// u := User{Id: user.Id}
	// err = o.Read(&u)
	// fmt.Println(u)
	// fmt.Printf("ERR: %v\n", err)

	// // delete record
	// num, err = o.Delete(&u)
	// fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
