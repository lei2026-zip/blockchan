package controllers

import (
	"demo/User"
	"demo/encryption"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (c *LoginController) Post() {
	var pers User.Consumer
	fmt.Println("Post请求:")
	c.Ctx.Request.ParseForm()
	body :=   c.Ctx.Request.Form
	fmt.Println(body)
	//	fmt.Println("body",body)
	pers.Phonenumber = body["Phonenumber"][0]
	pers.Password = encryption.EncryptByMd5(body["Password"][0])
	fmt.Println(pers.Password)
	str := pers.Login()
	if str!=""{
		fmt.Println(str)
		c.Ctx.WriteString("登录失败:"+str)
		c.Data["Status"]="密码错误"
		c.TplName = "login.html"
	}else{
		c.Data["Status"]="登入成功!"
		c.Data["Phonenumber"] = pers.Phonenumber
		c.Data["UserPhonenumber"] = pers.Phonenumber
		c.TplName = "input.html"
	}
	return
}