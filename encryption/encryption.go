package encryption

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/ioutil"
	"fmt"
)

func EncryptByMd5(str string)string{
  wir := md5.New()
  wir.Write([]byte(str))
  cipherStr := wir.Sum(nil)
  return hex.EncodeToString(cipherStr)
}

func EncrytFileByMd5( reader io.Reader)(string,error){
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}
	md5Hash := md5.New()
	md5Hash.Write(bytes)
	hashBytes := md5Hash.Sum(nil)
	return hex.EncodeToString(hashBytes), nil
}

func ConfirmPassword(password string,cipherStr string)bool{
	if cipherStr == EncryptByMd5(password){
		return true
	}
	return false
}


