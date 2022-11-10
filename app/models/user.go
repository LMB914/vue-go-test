package models

type User struct {
	UserName string `gorm:"type:varchar(255);not null;primary_key" json:"user_name"`
	UserPass string `gorm:"type:varchar(255);not null" json:"user_pass"`
}

func (uc User) CreateTable(){
	var users User
	db.AutoMigrate(&users)
}

func (uc User) InitTable(){
	admin := User{
		UserName: "admin",
		UserPass: "admin12345",
	}
	db.Create(&admin)
}

func (uc User) InsertTable(user *User){
	 db.Create(&user)
}

func (uc User) CheckTable(userName string)  int64{
	var count int64
	var users User
	db.Where("user_name = ?",userName).Find(&users).Count(&count)
	return count
}
