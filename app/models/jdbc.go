package models

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"io/ioutil"
	"time"
)

type JdbcConf struct {
	App App `yaml:"app" json:"app"`
}


type App struct {
	JdbcIp	string	`yaml:"jdbc_ip" json:"jdbc_ip"`
	JdbcPort	int `yaml:"jdbc_port" json:"jdbc_port"`
	JdbcUsername string `yaml:"jdbc_username" json:"jdbc_username"`
	JdbcPassword string  `yaml:"jdbc_password" json:"jdbc_password"`
	JdbcDatabase string  `yaml:"jdbc_database" json:"jdbc_database"`
}

var db *gorm.DB
func  init() {
	var jdbcInfo JdbcConf
	var err error
	yamlFile,err := ioutil.ReadFile("conf/jdbc.yaml")
	if err != nil{
		fmt.Println(err.Error())
	}
	err = yaml.Unmarshal(yamlFile,&jdbcInfo)
	if err != nil{
		fmt.Println(err.Error())
	}
	jdbcUrl:= fmt.Sprintf("%s:%s@(%s:%d)/%s?&parseTime=True&loc=Local",
		jdbcInfo.App.JdbcUsername,
		jdbcInfo.App.JdbcPassword,
		jdbcInfo.App.JdbcIp,
		jdbcInfo.App.JdbcPort,
		jdbcInfo.App.JdbcDatabase)
	db,err = gorm.Open(mysql.Open(jdbcUrl),&gorm.Config{})
	if err != nil{
		panic("数据库连接失败：" + err.Error())
	}
	sqlDB,_ := db.DB()
	sqlDB.SetMaxOpenConns(1000)
	sqlDB.SetConnMaxLifetime(0)
	sqlDB.SetConnMaxLifetime(time.Hour*4)

	//预制建表
	db.Debug().AutoMigrate(&User{})
	admin := User{
		UserName: "admin",
		UserPass: "admin12345",
	}
	db.Create(&admin)
	db.Debug().AutoMigrate(&NodeInfo{})
}



