package MySqldriver

import (
	"database/sql"
	_ "demo/mysql"
	. "demo/userdata"
	"fmt"
	"github.com/astaxie/beego"
	"log"
)

/*
2020年9月25日16:47:02
用于打开数据库
*/

func OpenSql()(*sql.DB,error){
	Congig :=	beego.AppConfig
	DriverName := Congig.String("DriverName")
	SqlUser := Congig.String("SqlUser")
	Sqlpassword := Congig.String("Sqlpassword")
	Sqldatabase:= Congig.String("Sqldatabase")
	Sqltcp:= Congig.String("Sqltcp")
	database,err :=sql.Open(DriverName,SqlUser+":"+Sqlpassword+"@tcp("+Sqltcp+")/"+Sqldatabase+"?charset=utf8")
	if err != nil{
		fmt.Println("数据库打开失败")
		log.Fatal(err.Error())
		return database,err
	}
	return database,nil
}

/*
2020年9月25日16:47:48
查询数据库是否有此表格
*/

func User_DataTableInit()error{
	database,_ := OpenSql()
	defer database.Close()
	row:=database.QueryRow("select COUNT(Userdemo.index_id) count from Userdemo")
	var count int
	err2 := row.Scan(&count)
	fmt.Println("用户数量:",count)
	if err2 != nil {
		_, err :=database.Exec("create table "+"Userdemo"+"( index_id int not null primary key auto_increment,UserName varchar(20),User_Phonenumber varchar(20),User_Password blob)"+"charset=utf8;")
		if err != nil{
			fmt.Println("数据库操作执行失败")
			log.Fatal(err.Error())
			return err
		}
		return nil
	}else{
		fmt.Println("表格已存在")
		return nil
	}
}

/*
读取数据库所以用户信息
*/

func Get_AllUser()([]User,error){
	var users []User
	var user User
	var index int
	var pas []byte
	database,_:=OpenSql()
	defer database.Close()
	qur,err2:= database.Query("select *from Userdemo")
	if err2 != nil {
		log.Fatal(err2.Error())
		return users,err2
	}
	for {
		if qur.Next(){
			qur.Scan(&index,&user.UserName,&user.Phonenumber,&pas)
			user.Password = string(pas)
			users = append(users,user)
		}else{
			break
		}
	}
	return users,nil
}


/*
通过id查找数据库存储的某一用户
*/
func Switch_TheUser(id string)(User,bool){
         users,err :=  Get_AllUser()
         if err!=nil{
         	return User{},false
		 }
		for i:=0;i<len(users);i++{
			if users[i].Phonenumber==id{
				return users[i],true
			}
		}
		return User{},false
}

/*
保存新的用户信息
*/

func Save_Data(user User)string{
	database,_:=OpenSql()
	defer database.Close()
    _,bo := Switch_TheUser(user.Phonenumber)
    if bo{
    	fmt.Println("该用户已存在!")
    	return ""
	}
	fmt.Println([]byte(user.Password))
	_,err1 := database.Exec("insert into " +
		"Userdemo(UserName,User_Phonenumber,User_Password)"+
		"values(?,?,?)",user.UserName,user.Phonenumber,[]byte(user.Password))
	if err1 != nil {
		log.Fatal(err1.Error())
		return "插入保存失败"
	}
	return ""
}