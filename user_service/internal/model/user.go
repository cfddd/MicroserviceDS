package model

import (
	"gorm.io/gorm"
	"sync"
	"time"
	"user_service/pkg/encryption"
	"utils/snowFlake"
)

type COMMON struct {
	ID        int64          `gorm:"primarykey" json:"id"`    // 主键ID
	CreatedAt time.Time      `json:"created_at"`              // 创建时间
	UpdatedAt time.Time      `json:"updated_at"`              // 更新时间
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"` // 删除时间
}

type User struct {
	COMMON
	Name       string `gorm:"unique" json:"name"`      //名字
	Password   string `gorm:"notnull" json:"password"` //密码
	Douyin_num string `json:"douyin_num"`              //抖音号
}

type UserModel struct {
}

var userModel *UserModel
var userOnce sync.Once //单例模式

// GetInstance 获取单例实例
func GetInstance() *UserModel {
	userOnce.Do(
		func() {
			userModel = &UserModel{}
		},
	)
	return userModel
}

// Create 创建用户
func (*UserModel) Create(user *User) error {
	flake, _ := snowFlake.NewSnowFlake(7, 1)
	user.ID = flake.NextId()
	user.Password = encryption.HashPassword(user.Password)
	res := DB.Create(&user)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

// FindUserByName 根据用户名称查找用户,并返回对象
func (*UserModel) FindUserByName(username string) (*User, error) {
	user := User{}
	res := DB.Where("name=?", username).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

// FindUserById 根据id查找用户,并返回对象
func (*UserModel) FindUserById(id int64) (*User, error) {
	user := User{}
	res := DB.Where("id=?", id).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

// FindUserByNum 根据抖音号查找用户,并返回对象
func (*UserModel) FindUserByNum(num string) (*User, error) {
	user := User{}
	res := DB.Where("douyin_num=?", num).First(&user)
	if res.Error != nil {
		return nil, res.Error
	}
	return &user, nil
}

// CheckUserExist 检查User是否存在（已经被注册过了）
func (*UserModel) CheckUserExist(username string) bool {
	user := User{}
	err := DB.Where("name=?", username).First(&user).Error
	if err == gorm.ErrRecordNotFound {
		return true //用户不存在
	}
	return false //用户存在
}

// CheckPassWord 检查密码是否正确
func (*UserModel) CheckPassWord(password string, storePassword string) bool {
	return encryption.VerifyPasswordWithHash(password, storePassword)
}
