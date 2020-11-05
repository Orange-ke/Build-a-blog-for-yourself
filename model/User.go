package model

import (
	"encoding/base64"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/scrypt"
	"log"
	"myblog/utils/errmsg"
)

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(20);not null" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type: varchar(50);not null" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type: int;DEFAULT:2" json:"role" validate:"required,gte=2" label:"角色"` // gte 表示 >=2
}

// 数据库操作
// 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED // 1001
	}
	return errmsg.SUCCESS
}

// 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id, username").Where("username = ?", name).First(&user)
	if user.ID == uint(id) { // 说明未修改用户名
		return errmsg.SUCCESS
	}
	if user.ID > 0 { // 此时说明修改的用户名重复了
		return errmsg.ERROR_USERNAME_USED // 1001
	}
	return errmsg.SUCCESS
}

// 新增用户
func CreateUser(data *User) (code int) {
	//data.Password = ScryptPw(data.Password)
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR // 500
	}
	return errmsg.SUCCESS
}

// 查询单个用户
func GetUser(id int) (User, int) {
	var user User
	err= db.Where("ID = ?", id).First(&user).Error
	if err != nil {
		return user, errmsg.ERROR
	}
	return user, errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(userName string, pageSize int, pageNum int) ([]User, int) {
	var users []User
	var total int

	if userName == "" {
		err = db.Find(&users).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	} else {
		err = db.Where("username LIKE ?", userName + "%").Find(&users).Count(&total).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	}

	if err != nil || err == gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// 编辑用户 (修改密码单独设定，需要验证)
func EditUser(id int, data *User) int {
	// gorm使用struct更新时，零值默认不会更新，因此使用map来更新
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role
	err = db.Model(&user).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 重置密码
func ResetPassword(id int, data *User) (code int) {
	// gorm使用struct更新时，零值默认不会更新，因此使用map来更新
	var user User
	var maps = make(map[string]interface{})
	maps["password"] = ScryptPw(data.Password)
	err = db.Model(&user).Where("id = ?", id).Update(maps).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 密码加密
func (u *User) BeforeSave() {
	u.Password = ScryptPw(u.Password)
}

func ScryptPw(password string) string {
	const KeyLen = 32
	salt := make([]byte, 8)
	salt = []byte{12, 32, 4, 8, 10, 222, 11, 122}
	hashPw, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, KeyLen)
	if err != nil {
		log.Fatal(err)
	}
	saltedPw := base64.StdEncoding.EncodeToString(hashPw)
	return saltedPw
}

// 登陆验证
func CheckLogin(username string, password string) int {
	var user User
	var code int
	db.Where("username = ?", username).First(&user)
	if user.ID == 0 {
		code = errmsg.ERROR_USER_NOT_EXIST
		return code
	}
	if ScryptPw(password) != user.Password {
		code = errmsg.ERROR_PASSWORD_WRONG
		return code
	}
	if user.Role != 1 {
		code = errmsg.ERROR_USER_NO_RIGHT
		return code
	}
	return errmsg.SUCCESS
}
