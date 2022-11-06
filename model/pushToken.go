package model

import (
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type SystemType int8

const (
	WINDOWS SystemType = 0
	LINUX   SystemType = 1
	MACOS   SystemType = 2
	IOS     SystemType = 3
	ANDROID SystemType = 4
	UNKNOWN SystemType = -128
)

// 紀錄推播Token
type PushToken struct {
	gorm.Model
	AppId    string     `json:"id"`
	Token    string     `json:"token" gorm:"index:idx_push,unique"`
	Version  string     `json:"version"`
	Username string     `json:"username"`
	IP       string     `json:"ip"`
	System   SystemType `json:"system"`
}

// 新增 / 更新資料
func (_pushToken *PushToken) InsertOrUpdate(database *gorm.DB, dictionary map[string]interface{}) (map[string]interface{}, error) {

	var error error
	isSuccess := false
	result := map[string]interface{}{"isSuccess": isSuccess}
	system := UNKNOWN

	_IP := dictionary["IP"]
	_appId := dictionary["id"]
	_token := dictionary["token"]
	_system := dictionary["system"]
	_username := dictionary["username"]
	_version := dictionary["version"]

	if _appId == nil {
		return result, errors.New("appId is nil")
	}
	if _token == nil {
		return result, errors.New("token is nil")
	}
	if _system == nil {
		return result, errors.New("system is nil")
	}
	if _username == nil {
		return result, errors.New("username is nil")
	}
	if _version == nil {
		return result, errors.New("version is nil")
	}

	if strings.ToUpper(_system.(string)) == "WINDOWS" {
		system = WINDOWS
	}
	if strings.ToUpper(_system.(string)) == "LINUX" {
		system = LINUX
	}
	if strings.ToUpper(_system.(string)) == "MACOS" {
		system = MACOS
	}
	if strings.ToUpper(_system.(string)) == "IOS" {
		system = IOS
	}
	if strings.ToUpper(_system.(string)) == "ANDROID" {
		system = ANDROID
	}

	error = database.Create(&PushToken{AppId: _appId.(string), Token: _token.(string), System: system, Username: _username.(string), Version: _version.(string), IP: _IP.(string)}).Error

	if error == nil {
		isSuccess = true
		result["isSuccess"] = isSuccess
		return result, error
	}

	info := map[string]interface{}{
		"app_id":   _appId,
		"token":    _token,
		"version":  _version,
		"username": _username,
		"ip":       _IP,
		"system":   system,
	}

	result, error = _pushToken.Update(database, info)
	return result, error
}

// 搜尋
func (_pushToken *PushToken) Select(database *gorm.DB, token string) PushToken {

	var pushToken PushToken
	database.Take(&pushToken, "token=?", token)

	return pushToken
}

// [搜尋](https://5xruby.tw/posts/what-is-jwt)
func (_pushToken *PushToken) SelectWhere(database *gorm.DB, where string) []PushToken {

	var pushTokenList []PushToken
	database.Where(where).Find(&pushTokenList)

	return pushTokenList
}

// 更新
func (_pushToken *PushToken) Update(database *gorm.DB, dictionary map[string]interface{}) (map[string]interface{}, error) {

	isSuccess := false
	token := dictionary["token"].(string)

	pushToken := _pushToken.Select(database, token)
	error := database.Model(&pushToken).Updates(dictionary).Error

	if error == nil {
		isSuccess = true
	}

	result := map[string]interface{}{}
	result["isSuccess"] = isSuccess

	return result, error
}

// 尋找該AppId的PushToken們
func (_pushToken *PushToken) DeviceTokenList(database *gorm.DB, bundleId string) []PushToken {

	var list []PushToken

	where := fmt.Sprintf(`app_id = "%v"`, bundleId)
	database.Where(where).Find(&list)

	return list
}

// 尋找該AppId的Token們
func (_pushToken *PushToken) TokenList(database *gorm.DB, bundleId string) []string {

	var tokenList []string
	pushTokenList := _pushToken.DeviceTokenList(database, bundleId)

	for index := 0; index < len(pushTokenList); index++ {
		token := pushTokenList[index]
		tokenList = append(tokenList, token.Token)
	}

	return tokenList
}
