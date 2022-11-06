package model

import (
	"errors"
	"fmt"
	"strings"
	"time"
	"william/utility"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type UserLevel uint

const (
	Account UserLevel = 0
	Admin   UserLevel = 9
	Root    UserLevel = 99
)

const (
	UsernameMinimumLength int = 3
	PasswordMinimumLength int = 6
)

type User struct {
	gorm.Model
	Username string    `json:"username" gorm:"index:idx_username,unique"`
	Password string    `json:"password"`
	Level    UserLevel `json:"level"`
}

type authClaims struct {
	jwt.RegisteredClaims
	UserID uint      `json:"userId"`
	Level  UserLevel `json:"level"`
}

// 私鑰
var jwtKey = []byte("FDr1VjVQiSiybYJrQZNt8Vfd7bFEsKP6vNX1brOSiWl0mAIVCxJiR4/T3zpAlBKc2/9Lw2ac4IwMElGZkssfj3dqwa7CQC7IIB+nVxiM1c9yfowAZw4WQJ86RCUTXaXvRX8JoNYlgXcRrK3BK0E/fKCOY1+izInW3abf0jEeN40HJLkXG6MZnYdhzLnPgLL/TnIFTTAbbItxqWBtkz6FkZTG+dkDSXN7xNUxlg==")

// 新增資料
func (_user *User) Insert(database *gorm.DB, dictionary map[string]interface{}) (map[string]interface{}, error) {

	var error error
	isSuccess := false
	result := map[string]interface{}{"isSuccess": isSuccess}

	_username := dictionary["username"]
	_password := dictionary["password"]
	_level := numberToLevel(dictionary["level"].(float64))

	if _username == nil || _password == nil {
		return result, errors.New("username or password is nil")
	}

	if len(_username.(string)) < UsernameMinimumLength || len(_password.(string)) < PasswordMinimumLength {
		errorMessage := fmt.Sprintf("username is lower %d, or password is lower %d", UsernameMinimumLength, PasswordMinimumLength)
		return result, errors.New(errorMessage)
	}

	encryptPassword := PasswordEncryptRule(_password.(string))
	error = database.Create(&User{Username: strings.ToUpper(_username.(string)), Password: encryptPassword, Level: _level}).Error

	if error == nil {
		isSuccess = true
	}

	result["isSuccess"] = isSuccess
	return result, error
}

// 搜尋單一User
func (_user *User) Select(database *gorm.DB, username string) *User {

	var user User
	database.Take(&user, "username=?", strings.ToUpper(username))

	if user.ID == 0 {
		return nil
	}

	return &user
}

// 登入 (成功後會回傳Token)
func (_user *User) Login(database *gorm.DB, username string, password string, expiresTime time.Time) (map[string]interface{}, error) {

	var user User
	var authToken AuthToken
	result := map[string]interface{}{}

	encryptPassword := PasswordEncryptRule(password)
	database.Take(&user, "username=? AND password=?", strings.ToUpper(username), encryptPassword)

	if user.ID == 0 {
		return result, errors.New("record not found")
	}

	token, error := authToken.JwtToken(user, expiresTime)

	if error != nil {
		return result, error
	}

	result["token"] = token
	result["expiresTime"] = expiresTime

	return result, error
}

// 密碼加密的演算法 => openssl rand -base64 32
func PasswordEncryptRule(word string) string {

	_randKey := "Es1CLMWoDepb4Uy9w3vyDiXhwrmgDYncaNvNVJdxnKw="
	_word := utility.Md5(word)
	_key := utility.Md5(_randKey)
	_password := fmt.Sprintf("%v%v", _word, _key)

	return _password
}

// float64 => UserLevel
func numberToLevel(number float64) UserLevel {

	level := Account

	if number == 9 {
		level = Admin
	}

	if number == 99 {
		level = Root
	}

	return level
}
