package model

import (
	"encoding/base64"
	"fmt"
	"go-blog/utils/result"
	"golang.org/x/crypto/scrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	gorm.Model
	Username string `gorm:"type: varchar(20); notnull" json:"username" validate:"required,min=4,max=12" label:"用户名"`
	Password string `gorm:"type: varchar(20); notnull" json:"password" validate:"required,min=6,max=20" label:"密码"`
	Role     int    `gorm:"type: int; DEFAULT: 2" json:"role" validate:"required,gte=2"`
}

func CheckUser(name string) result.Code {
	var user User
	db.Select("id").Where("username = ?", name).First(&user)
	if user.ID > 0 {
		return result.ERROR_USERNAME_USED //1001
	}
	return result.SUCCESS
}

func AddUser(user *User) result.Code {
	user.Password = ScryptPwd(user.Password)
	if err := db.Create(user).Error; err != nil {
		return result.ERROR
	}
	fmt.Println(*user)
	return result.SUCCESS
}

func DeleteUser(id int) result.Code {
	err := db.Where("id = ?", id).Delete(&User{}).Error
	if err != nil {
		return result.ERROR
	}
	return result.SUCCESS
}

func EditUser(id int, user *User) result.Code {
	var maps = make(map[string]interface{})
	maps["username"] = user.Username
	maps["role"] = user.Role
	err := db.Model(&User{}).Where("id = ?", id).Updates(maps).Error
	if err != nil {
		return result.ERROR
	}
	return result.SUCCESS
}

func GetUsers(pageSize, pageNum int) ([]User, int64) {
	var users []User
	var total int64
	if err := db.Limit(pageSize).
		Offset((pageNum - 1) * pageSize).
		Find(&users).Count(&total).Error; err != nil && err != gorm.ErrRecordNotFound {
		return nil, 0
	}
	return users, total
}

// ScryptPwd 加密
func ScryptPwd(password string) string {
	const keyLen = 10
	salt := make([]byte, 8)
	salt = []byte{6, 12, 22, 32, 43, 58, 72, 33}

	HashPwd, err := scrypt.Key([]byte(password), salt, 1<<15, 8, 1, keyLen)
	if err != nil {
		log.Fatal(err)
	}

	return base64.StdEncoding.EncodeToString(HashPwd)
}

func Login(username, password string) result.Code {
	var user User
	db.Where("username = ?", username).First(&user)

	if user.ID == 0 {
		return result.ERROR_USER_NOT_EXIST
	}

	if ScryptPwd(password) != user.Password {
		return result.ERROR_PASSWORD_WRONG
	}

	if user.Role != 0 {
		return result.ERROR_USER_NO_RIGHT
	}

	return result.SUCCESS
}
