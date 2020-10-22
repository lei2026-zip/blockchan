package input

import (
	"demo/MySqldriver"
	. "demo/userdata"
	"path"
)

func InputToChan(user InputData)string{
	 MySqldriver.Input_DataTableInit()
	 encoder :=  MySqldriver.QuerylastData()
	if encoder!="" {
	   user.Encryption = encoder
	}else {
		return "上一链hax值获取失败"
	}
	err:= MySqldriver.Input_Data(user)
	if err!="" {
		return err
	}
    return ""
}



//给上传的文件大致分类
func Queryfileclass(str string)string{
	//index := strings.Index(str,".")
	//return str[index+1:]
	return path.Ext(str)  //检索后缀名
}

func FileClassify(str string)string{
	switch Queryfileclass(str) {
	case ".jpg": return "img" ;
	case ".jpeg": return "img" ;
	case ".png": return "img" ;
	case ".gif": return "img" ;
	case ".mp4": return "vido" ;
	case ".doc": return  "doc" ;
	case ".docx": return  "docx";
	default:
		return "anther"; break
	}
	return ""
}