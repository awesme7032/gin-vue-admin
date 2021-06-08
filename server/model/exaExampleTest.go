// 自动生成模板ExaExampleTest
package model

import (
	"gin-vue-admin/global"
      "time"
)

// 如果含有time.Time 请自行import time包
type ExaExampleTest struct {
      global.GVA_MODEL
      AdminId  int `json:"adminId" form:"adminId" gorm:"column:admin_id;comment:管理员ID;type:int;size:10;"`
      Username string `json:"username" from:"username" gorm:"column:username;comment:用户名称;type:varchar(255);""`
      CategoryId  int `json:"categoryId" form:"categoryId" gorm:"column:category_id;comment:分类ID(单选);type:int;size:10;"`
      CategoryIds  string `json:"categoryIds" form:"categoryIds" gorm:"column:category_ids;comment:分类ID(多选);type:varchar(100);size:100;"`
      Week  string `json:"week" form:"week" gorm:"column:week;comment:星期(单选):monday=星期一,tuesday=星期二,wednesday=星期三;type:varchar(100);size:100;"`
      Flag  string `json:"flag" form:"flag" gorm:"column:flag;comment:标志(多选):hot=热门,index=首页,recommend=推荐;type:varchar(100);size:100;"`
      Genderdata  string `json:"genderdata" form:"genderdata" gorm:"column:genderdata;comment:性别(单选):male=男,female=女;type:varchar(100);size:100;"`
      Hobbydata  string `json:"hobbydata" form:"hobbydata" gorm:"column:hobbydata;comment:爱好(多选):music=音乐,reading=读书,swimming=游泳;type:varchar(100);size:100;"`
      Title  string `json:"title" form:"title" gorm:"column:title;comment:标题;type:varchar(50);size:50;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:内容;type:text;"`
      Image  string `json:"image" form:"image" gorm:"column:image;comment:图片;type:varchar(100);size:100;"`
      Images  string `json:"images" form:"images" gorm:"column:images;comment:图片组;type:varchar(1500);size:1500;"`
      Attachfile  string `json:"attachfile" form:"attachfile" gorm:"column:attachfile;comment:附件;type:varchar(100);size:100;"`
      Keywords  string `json:"keywords" form:"keywords" gorm:"column:keywords;comment:关键字;type:varchar(100);size:100;"`
      Description  string `json:"description" form:"description" gorm:"column:description;comment:描述;type:varchar(255);size:255;"`
      City  string `json:"city" form:"city" gorm:"column:city;comment:省市;type:varchar(100);size:100;"`
      Json  string `json:"json" form:"json" gorm:"column:json;comment:配置:key=名称,value=值;type:varchar(255);size:255;"`
      Price  float64 `json:"price" form:"price" gorm:"column:price;comment:价格;type:float;"`
      Views  int `json:"views" form:"views" gorm:"column:views;comment:点击;type:int;size:10;"`
      Startdate  time.Time `json:"startdate" form:"startdate" gorm:"column:startdate;comment:开始日期;type:date;"`
      Activitytime  time.Time `json:"activitytime" form:"activitytime" gorm:"column:activitytime;comment:活动时间(datetime);type:datetime;"`
      Year  int64 `json:"year" form:"year" gorm:"column:year;comment:年;type:int;"`
      Times  time.Time `json:"times" form:"times" gorm:"column:times;comment:时间;type:datetime;"`
      Refreshtime  time.Time `json:"refreshtime" form:"refreshtime" gorm:"column:refreshtime;comment:刷新时间(int);type:int;size:10;"`
      Weigh  int `json:"weigh" form:"weigh" gorm:"column:weigh;comment:权重;type:int;size:10;"`
      Switch  int64 `json:"switch" form:"switch" gorm:"column:switch;comment:开关;type:tinyint;"`
      Status  int64 `json:"status" form:"status" gorm:"column:status;comment:状态;type:varchar(100);size:100;"`
      State  int64 `json:"state" form:"state" gorm:"column:state;comment:状态值:0=禁用,1=正常,2=推荐;type:tinyint;"`
}


func (ExaExampleTest) TableName() string {
  return "exa_example_test"
}

