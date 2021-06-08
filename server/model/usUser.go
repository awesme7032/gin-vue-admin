// 自动生成模板UsUser
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type UsUser struct {
      global.GVA_MODEL
      Email  string `json:"email" form:"email" gorm:"column:email;comment:;type:varchar(255);size:255;"`
      Gender  *bool `json:"gender" form:"gender" gorm:"column:gender;comment:;type:tinyint;"`
      Icon  string `json:"icon" form:"icon" gorm:"column:icon;comment:;type:varchar(255);size:255;"`
      IdCard  string `json:"idCard" form:"idCard" gorm:"column:id_card;comment:;type:varchar(255);size:255;"`
      InvitationId  string `json:"invitationId" form:"invitationId" gorm:"column:invitation_id;comment:;type:varchar(255);size:255;"`
      JoinIp  string `json:"joinIp" form:"joinIp" gorm:"column:join_ip;comment:;type:varchar(255);size:255;"`
      LastLoginIp  string `json:"lastLoginIp" form:"lastLoginIp" gorm:"column:last_login_ip;comment:;type:varchar(255);size:255;"`
      LastLoginTime  time.Time `json:"lastLoginTime" form:"lastLoginTime" gorm:"column:last_login_time;comment:;type:timestamp;"`
      LevelId  int `json:"levelId" form:"levelId" gorm:"column:level_id;comment:;type:tinyint;"`
      Nickname  string `json:"nickname" form:"nickname" gorm:"column:nickname;comment:;type:varchar(255);size:255;"`
      Password  string `json:"password" form:"password" gorm:"column:password;comment:;type:varchar(255);size:255;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:;type:varchar(255);size:255;"`
      Pid  int `json:"pid" form:"pid" gorm:"column:pid;comment:;type:bigint;size:19;"`
      Realname  string `json:"realname" form:"realname" gorm:"column:realname;comment:;type:varchar(255);size:255;"`
      State  int `json:"state" form:"state" gorm:"column:state;comment:;type:tinyint;"`
      Type  int `json:"type" form:"type" gorm:"column:type;comment:;type:tinyint;"`
      Username  string `json:"username" form:"username" gorm:"column:username;comment:;type:varchar(255);size:255;"`
}


func (UsUser) TableName() string {
  return "us_user"
}

