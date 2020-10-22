package MySqldriver

import (
	"crypto/md5"
	"demo/userdata"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"
)



//区块链概念 -  -  - 数据存储链
func Input_DataTableInit()string{
	database,_ := OpenSql()
	defer database.Close()
	row:=database.QueryRow("select COUNT(DataChan.index_id) count from DataChan")
	var count int
	err2 := row.Scan(&count)
	fmt.Println("数据链存储数量:",count)
	if err2 != nil {//  id,Phonenumber,tag,file_name,file_url,encryption,time
		_, err :=database.Exec("create table "+"DataChan"+"( index_id int not null primary key auto_increment,Phonenumber varchar(20),tag varchar(60),file_name varchar(50),file_url varchar(60),encryption blob,time varchar(60))"+"charset=utf8;")
		if err != nil{
			fmt.Println("数据库操作执行失败")
			log.Fatal(err.Error())
			return ""
		}
		_,err1 := database.Exec("insert into " +
			"DataChan(Phonenumber,tag,file_name,file_url,encryption,time)"+
			"values(?,?,?,?,?,?)","88888888","初始区块","晖","暂无",[]byte("SSR"),time.Now().String())
		if err1 != nil {
			log.Fatal(err1.Error())
			return "初始链创建失败!"
		}
		return ""
	}else{
		fmt.Println("表格已存在")
		return ""
	}
}

func Input_Data(user userdata.InputData)string{
	database,_:=OpenSql()
	defer database.Close()
	//user.Time = time.Now().UnixNano()
	_,err1 := database.Exec("insert into " +
		"DataChan(Phonenumber,tag,file_name,file_url,encryption,time)"+
		"values(?,?,?,?,?,?)",user.Phonenumber,user.Tag,user.File_name,user.File_url,[]byte(user.Encryption),user.Time)
	if err1 != nil {
		log.Fatal(err1.Error())
		return "插入保存失败"
	}
	return ""
}

func QuerylastData()string{
	database,_ := OpenSql()
	defer database.Close()
	row:=database.QueryRow("select COUNT(DataChan.index_id) count from DataChan")
	var count int
	var data userdata.InputData
	row.Scan(&count)
    cout := strconv.Itoa(count)
	Row := database.QueryRow("select * from DataChan where index_id="+cout+";")
	err:= Row.Scan(&data.Id,&data.Phonenumber,&data.Tag,&data.File_name,&data.File_url,&data.Encryption,&data.Time)
	if err!=nil{
		fmt.Println("Scan failed Err:",err)
		return ""
	}
	//对表数据进行MD5计算获取对应的hax值应用于下一节链的打包
	wir := md5.New()
	str,err :=  json.Marshal(data)
	wir.Write(str)
	cipherStr := wir.Sum(nil)
	return hex.EncodeToString(cipherStr)
}