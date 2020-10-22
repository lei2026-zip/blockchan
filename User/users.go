package User

import (
	. "demo/MySqldriver"
	. "demo/userdata"
)

type Consumer struct {
	User
}


func (user Consumer)Login()string{
     User_DataTableInit()
     u,bo:= Switch_TheUser(user.Phonenumber)
     if bo{
         if u.Password==user.Password{
         	return ""
		 }else {
		 	return "输入密码错误"
		 }
	 }else{
	 	//Save_Data(user)
	 	return "用户不存在!"
	 }
}

func (user Consumer)Register()string{
	User_DataTableInit()
	return Save_Data(user.User)
}

//func (user Consumer)Repassword()string{
//
//}

