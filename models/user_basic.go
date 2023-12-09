package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Name          string
	Password      string
	Phone         string
	Email         string
	Identity      string
	ClentIp       string
	ClientPort    string
	LoginTime     uint64
	HeartbeatTime uint64
	LogOutTime    uint64
	IsLogout      bool
	DeviceInfo    string
}

func (table *UserBasic) TableName() string{
	return "user_basic"
}