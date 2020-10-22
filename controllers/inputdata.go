package controllers

import (
	"demo/input"
	"demo/userdata"
	"fmt"
	"github.com/astaxie/beego"
	"log"
	"time"
)

type InputController struct {
	beego.Controller
}

func (c *InputController) Post() {
	var data userdata.InputData
	file,head,err:=c.GetFile("file")
	if err!= nil{
		log.Fatal("Get file error:",err)
	}
	defer file.Close()  //及时关闭上传的文件,不然会出现临时文件不能清除的请况
	url := "Localtion/"+input.FileClassify(head.Filename)+"/"
	c.SaveToFile("file",url+head.Filename)
	c.Ctx.Request.ParseForm()
	body := c.Ctx.Request.Form
	url += head.Filename
	data.File_name = head.Filename
	data.File_url = url
	data.Phonenumber = body["Phonenumber"][0]
	data.Tag = body["tag"][0]
	t := time.Now()
	data.Time = t.Format("2006年01月02日 15:04:05")
	err1:= input.InputToChan(data)
	if err1!="" {
		fmt.Println(err1)
		return
	}
	fmt.Println(head.Filename)
	fmt.Println(head.Size)
	fmt.Println(head.Header)
	c.Data["dara"] = data
	c.Data["UserPhonenumber"]  = data.Phonenumber
	c.Data["Status"] = "成功上链!"
	c.TplName = "input.html"

}
/*file, header, err := u.GetFile("upload_file")
if err != nil {
	u.Ctx.WriteString("抱歉，用户文件解析失败，请重试")
	return
}
//3、关闭文件
defer file.Close()

fmt.Println("自定义的文件标题:", fileTitle)
fmt.Println("文件名称:", header.Filename)
fmt.Println("文件的大小:", header.Size) //字节大小

//2、将文件保存在本地的一个目录中
//文件全路径： 路径 + 文件名 + "." + 扩展名
//要的文件的路径
uploadDir := "./static/img/" + header.Filename
//文件权限：a+b+c
//a:文件所有者拥有的权限，读4、写2、执行1
//b:文件所有者所在的组的用户对文件拥有的权限，读4、写2、执行1
//c:其他用户对文件拥有的权限，读4、写2、执行1
//eg:某个文件m，其权限是985(错误)
saveFile, err := os.OpenFile(uploadDir, os.O_RDWR|os.O_CREATE, 777)

//创建一个writer: 用于向硬盘上写一个文件
writer := bufio.NewWriter(saveFile)
file_size, err := io.Copy(writer, file)
if err != nil { //invalid argument
	fmt.Println(err.Error())
	u.Ctx.WriteString("抱歉，保存电子数据失败，请重试")
	return
}
defer saveFile.Close()

//2、计算文件的hash
hashFile, err := os.Open(uploadDir)
defer hashFile.Close()*/