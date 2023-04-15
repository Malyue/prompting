package model

import (
	"gorm.io/gorm"
	"prompting/pkg/auth"
	"time"
)

type UserM struct {
	ID        int64     `gorm:"column:id;primary_key"`
	Username  string    `gorm:"column:username;not null"`
	Password  string    `gorm:"column:password;not null"`
	Nickname  string    `gorm:"column:nickname"`
	Email     string    `gorm:"column:email"`
	Phone     string    `gorm:"column:phone"`
	Introduce string    `gorm:"column:introduce"`
	Avatar    string    `gorm:"column:avatar"`
	Role      string    `gorm:"column:role"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdateAt  time.Time `gorm:"column:updatedAt"`
}

// TableName 用来指定映射的 MySQL 表名.
func (u *UserM) TableName() string {
	return "t_user"
}

// BeforeCreate 在创建数据库记录之前加密明文密码.
func (u *UserM) BeforeCreate(tx *gorm.DB) (err error) {
	// Encrypt the user password
	u.Password, err = auth.Encrypt(u.Password)
	return err
}
