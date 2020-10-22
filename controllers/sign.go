package controllers

import (
	"demo/User"
	"demo/encryption"
	"fmt"
	"github.com/astaxie/beego"
)

type SignController struct {
	beego.Controller
}

func (c *SignController) Post(){
	var pers User.Consumer
	fmt.Println("Post请求:")
	c.Ctx.Request.ParseForm()
	body :=   c.Ctx.Request.Form
	fmt.Println(body)
	//	fmt.Println("body",body)
	pers.Phonenumber = body["Phonenumber"][0]
	pers.Password =encryption.EncryptByMd5(body["Password"][0])
	fmt.Println(pers.Password)
	str := pers.Register()
	if str!=""{
		fmt.Println(str)
		c.Ctx.WriteString("注册失败:"+str)
	}else{
		c.Ctx.WriteString("注册成功!")
	}
	return
}