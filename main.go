package main

import (
	_ "demo/routers"
	"fmt"
	"github.com/astaxie/beego"
)
func main(){
    Congig :=	beego.AppConfig
    beego.SetStaticPath("../static/js","../static/js")
	beego.SetStaticPath("../static/img","../static/css")
	beego.SetStaticPath("../static/img","../static/img")
	port,err:= Congig.Int("httpport")
	if err!=nil{
		panic("项目配置文件解析失败,请检查配置文件")
	}
	fmt.Println(port)
	beego.Run()
}


