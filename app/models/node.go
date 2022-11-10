package models

type NodeInfo struct {
	NodeIp string	`gorm:"type:varchar(255);not null" json:"node_ip"`
	NodePort int	`gorm:"type:int;" json:"node_port"`
	NodeUser string	`gorm:"type:varchar(255)" json:"node_user"`
	NodePassword string	`gorm:"type:varchar(255)" json:"node_password"`
	NodeStatus int	`gorm:"type:int" json:"node_status"`
}

func (nc NodeInfo) CreateTable(){
	var nodes NodeInfo
	db.AutoMigrate(&nodes)
}
