package controllers

import (
	"github.com/astaxie/beego"
)

type ComeinController struct {
	beego.Controller
}

func (c *ComeinController) Get(){
	c.TplName = "login.html"
}