package pwdutil

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
)

func GenPwd(pwd string) string {
	password, _ := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	return string(password)
}

// ComparePwd pwd 数据库里面的密码 rePwd 用户输入的密码
func ComparePwd(pwd, reqPwd string) error {
	if e := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(reqPwd)); e != nil {
		return errors.New("用户名或密码不正确")
	}
	return nil
}
