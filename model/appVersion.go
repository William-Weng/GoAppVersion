package model

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"william/utility"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
	"gorm.io/gorm"
)

type StoreType int8
type UpdateType int8
type PushType int8

// [APNS的.p8檔相關設定資訊](https://medium.com/彼得潘的-swift-ios-app-開發教室/推播-p8憑證匯出-d195dd6b801b)
type APNSInformation struct {
	AuthKeyP8 string
	KeyID     string
	TeamID    string
	Topic     string
}

// [APNS的傳送內容](https://blog.toright.com/posts/2806/ios-apns-訊息推撥-apple-push-notification-service-介紹.html)
type APNSMessage struct {
	IsProduction bool
	DeviceToken  string
	Title        string
	Message      string
	Image        string
	Category     string
}

// [FCM推播的相關設定資訊](https://firebase.google.com/)
type FCMInformation struct {
	ApiKey string
}

const (
	AppStore   StoreType = 0
	GooglePlay StoreType = 1
)

const (
	General UpdateType = 0
	Always  UpdateType = 1
	Force   UpdateType = 2
)

const (
	Store      PushType = 0
	Production PushType = 1
	Develop    PushType = 2
)

// 資料表的長相
type AppVersion struct {
	gorm.Model
	AppStoreType            StoreType  `json:"type"`
	AutoUpdate              bool       `json:"autoUpdate"`
	AppId                   string     `json:"id" gorm:"index:idx_name,unique"`
	AppName                 string     `json:"name"`
	AppIcon                 *string    `json:"icon"`
	AppUrl                  string     `json:"url" gorm:"-:all"`
	Lang                    string     `json:"lang"`
	Version                 string     `json:"version"`
	VersionProduction       string     `json:"versionProd"`
	VersionDevelop          string     `json:"versionDev"`
	UpdateVersion           UpdateType `json:"updateVersion"`
	UpdateVersionProduction UpdateType `json:"updateVersionProd"`
	UpdateVersionDevelop    UpdateType `json:"updateVersionDev"`
	PushSetting             string     `json:"pushSetting"`
	PushSettingProduction   string     `json:"pushSettingProd"`
	PushSettingDevelop      string     `json:"pushSettingDev"`
	SupportedDevices        *string    `json:"supportedDevices"`
}

// 新增資料
// => https://apps.apple.com/tw/app/line/id443904275
// => https://play.google.com/store/apps/details?id=com.ubercab.driver&hl=zh_tw
func (_appVersion *AppVersion) Insert(database *gorm.DB, dictionary map[string]interface{}) (map[string]interface{}, error) {

	isSuccess := false

	var storeType = AppStore

	_storeType, _ := dictionary["type"].(float64)
	id := dictionary["id"].(string)
	name := dictionary["name"].(string)
	lang := ""
	version := "1.0"
	versionProd := "1.0"
	versionDev := "1.0"
	autoUpdate := true

	if _storeType == 0 {
		storeType = AppStore
	}
	if _storeType == 1 {
		storeType = GooglePlay
	}
	if dictionary["lang"] != nil {
		lang = dictionary["lang"].(string)
	}

	error := database.Create(&AppVersion{AppStoreType: storeType, AutoUpdate: autoUpdate, Lang: lang, AppId: id, AppName: name, Version: version, VersionProduction: versionProd, VersionDevelop: versionDev}).Error
	if error == nil {
		isSuccess = true
	}

	result := map[string]interface{}{"isSuccess": isSuccess}
	return result, error
}

// 搜尋
func (_appVersion *AppVersion) Select(database *gorm.DB, id string) AppVersion {

	var version AppVersion

	database.Take(&version, "app_id=?", id)
	version.AppUrl = _appStoreUrlMaker(&version)

	return version
}

// 搜尋 (迷你版)
func (_appVersion *AppVersion) SelectMini(database *gorm.DB, id string) map[string]interface{} {

	version := _appVersion.Select(database, id)

	if version.ID == 0 {
		return nil
	}

	jsonTemplate := `{"url":"%v","version":{"store":{"%v":%v},"prod":{"%v":%v},"dev":{"%v":%v}}}`
	json := fmt.Sprintf(jsonTemplate, version.AppUrl, version.Version, version.UpdateVersion, version.VersionProduction, version.UpdateVersionProduction, version.VersionDevelop, version.UpdateVersionDevelop)

	result := utility.RawStringToMap(json)
	return result
}

// 搜尋
func (_appVersion *AppVersion) SelectWhere(database *gorm.DB, dictionary map[string]interface{}) ([]AppVersion, int64) {

	var appVersionList []AppVersion
	var _database = database
	var count int64 = 0

	limit := dictionary["limit"]
	offset := dictionary["offset"]
	database.Model(&AppVersion{}).Count(&count)

	if limit != nil {
		_database = _database.Limit(int(limit.(float64)))
	}
	if offset != nil {
		_database = _database.Offset(int(offset.(float64)))
	}

	_database = _database.Where("id > 0")
	_database.Order("app_name asc, app_store_type asc").Find(&appVersionList)

	for index, version := range appVersionList {
		appVersionList[index].AppUrl = _appStoreUrlMaker(&version)
	}

	return appVersionList, count
}

// 更新
func (_appVersion *AppVersion) Update(database *gorm.DB, dictionary map[string]interface{}) (map[string]interface{}, error) {

	info := map[string]interface{}{}

	isSuccess := false
	id := dictionary["id"].(string)
	version := dictionary["version"].(map[string]interface{})
	appIcon := dictionary["icon"]
	appName := dictionary["name"]
	autoUpdate := dictionary["auto"]
	supportedDevices := dictionary["supportedDevices"]

	if version["store"] != nil {
		info["version"] = version["store"]
	}
	if version["prod"] != nil {
		info["version_production"] = version["prod"]
	}
	if version["dev"] != nil {
		info["version_develop"] = version["dev"]
	}
	if appIcon != nil {
		info["app_icon"] = appIcon
	}
	if appName != nil {
		info["app_name"] = appName
	}
	if autoUpdate != nil {
		info["auto_update"] = autoUpdate.(bool)
	}
	if supportedDevices != nil {

		devices, error := utility.MapToJSON(supportedDevices.(map[string]bool))

		if error == nil {
			info["supported_devices"] = devices
			utility.Println(devices)
		}
	}

	appVersion := _appVersion.Select(database, id)
	error := database.Model(&appVersion).Updates(info).Error

	if error == nil {
		isSuccess = true
	}

	result := map[string]interface{}{}
	result["isSuccess"] = isSuccess

	return result, error
}

// 設定更新樣式
func (_appVersion *AppVersion) UpdateType(database *gorm.DB, dictionary map[string]interface{}) (map[string]interface{}, error) {

	info := map[string]interface{}{}

	isSuccess := false

	_id := dictionary["id"].(string)
	_type := strings.ToUpper(dictionary["type"].(string))
	_level := dictionary["level"].(float64)

	if _type == "STORE" {
		info["update_version"] = _level
	}

	if _type == "PROD" {
		info["update_version_production"] = _level
	}

	if _type == "DEV" {
		info["update_version_develop"] = _level
	}

	appVersion := _appVersion.Select(database, _id)
	error := database.Model(&appVersion).Updates(info).Error

	if error == nil {
		isSuccess = true
	}

	result := map[string]interface{}{}
	result["isSuccess"] = isSuccess

	return result, error
}

// 刪除
func (_appVersion *AppVersion) Delete(database *gorm.DB, id string) map[string]interface{} {

	appVersion := _appVersion.Select(database, id)
	database.Unscoped().Delete(&appVersion, appVersion.ID)

	result := map[string]interface{}{}
	result["isSuccess"] = (appVersion.ID != 0)

	return result
}

// 取得APP在架上的版本號等相關資訊
func (_appVersion *AppVersion) CheckAppVersion(dictionary map[string]interface{}) (map[string]interface{}, error) {

	var storeType = AppStore
	var err error = nil
	var result map[string]interface{}

	id := dictionary["id"].(string)
	lang := ""
	_storeType, error := strconv.Atoi(dictionary["type"].(string))

	if error != nil {
		return result, error
	}
	if _storeType == 0 {
		storeType = AppStore
	}
	if _storeType == 1 {
		storeType = GooglePlay
	}
	if dictionary["lang"] != nil {
		lang = dictionary["lang"].(string)
	}

	switch storeType {
	case AppStore:
		result, err = _appVersion.AppStoreAppVersion(id, lang)
	case GooglePlay:
		result, err = _appVersion.GooglePlayAppVersion(id, lang)
	}

	return result, err
}

// 取得APP在架上的版本號 (AppStore)
// => https://itunes.apple.com/lookup?id=443904275
func (_appVersion *AppVersion) AppStoreAppVersion(id string, country string) (map[string]interface{}, error) {

	var list = map[string]interface{}{}

	_id := strings.Trim(id, "id")
	number, error := utility.StringToInt(_id)

	if error != nil {
		return list, error
	}

	url := fmt.Sprintf("https://itunes.apple.com/lookup?id=%d", number)
	if len(country) > 0 {
		url = fmt.Sprintf("https://itunes.apple.com/lookup?id=%d&country=%v", number, country)
	}

	response, error := http.Get(url)

	if error != nil {
		return list, error
	}

	defer response.Body.Close()

	bytes, error := io.ReadAll(response.Body)
	if error != nil {
		return list, error
	}

	dictionary := utility.RawBodyToMap(bytes)
	_results := dictionary["results"]

	if _results == nil {
		return list, error
	}

	results := _results.([]interface{})

	if len(results) == 0 {
		return list, error
	}

	result := results[0].(map[string]interface{})
	supportedDevices := map[string]bool{
		"iPhone": false,
		"iPad":   false,
	}

	_supportedDevices := fmt.Sprintf("%v", result["supportedDevices"])

	rule := `iPhone`
	regular, _ := regexp.Compile(rule)
	resultString := regular.FindString(_supportedDevices)

	if len(resultString) != 0 {
		supportedDevices["iPhone"] = true
	}

	rule = `iPad`
	regular, _ = regexp.Compile(rule)
	resultString = regular.FindString(_supportedDevices)

	if len(resultString) != 0 {
		supportedDevices["iPad"] = true
	}

	list["version"] = result["version"].(string)
	list["icon"] = result["artworkUrl512"].(string)
	list["supportedDevices"] = supportedDevices

	return list, error
}

// 取得APP在架上的版本號 (GooglePlay)
// => https://play.google.com/store/apps/details?id=com.ubercab.driver&hl=zh_tw
func (_appVersion *AppVersion) GooglePlayAppVersion(id string, hl string) (map[string]interface{}, error) {

	var list = map[string]interface{}{}

	url := fmt.Sprintf("https://play.google.com/store/apps/details?id=%s", id)
	if len(hl) > 0 {
		url = fmt.Sprintf("https://play.google.com/store/apps/details?id=%s&hl=%v", id, hl)
	}

	response, error := http.Get(url)
	if error != nil {
		return list, error
	}

	defer response.Body.Close()

	bytes, error := io.ReadAll(response.Body)
	if error != nil {
		return list, error
	}

	body := string(bytes)

	rule := `\[\["([\d.]{1,}\d)"]]`
	regular, _ := regexp.Compile(rule)
	version := regular.FindString(body)

	rule = `([\d.]{1,}\d)`
	regular, _ = regexp.Compile(rule)
	version = regular.FindString(version)

	rule = `(srcset="https://play-lh.googleusercontent.com/)(.+)(=w480-h960 2x" class=")`
	regular, _ = regexp.Compile(rule)
	icon := regular.FindString(body)

	rule = `(srcset="https://play-lh.googleusercontent.com/)(.+)(=w480-h960)(.+)(alt=)`
	regular, _ = regexp.Compile(rule)
	icon = regular.FindString(icon)

	rule = `(https://play-lh.googleusercontent.com/)(.+)(=w480-h960)`
	regular, _ = regexp.Compile(rule)
	icon = regular.FindString(icon)

	list["version"] = version
	list["icon"] = icon
	list["supportedDevices"] = nil

	return list, error
}

// 產生App在Store上的網址
func _appStoreUrlMaker(version *AppVersion) string {

	var appUrl = ""

	switch version.AppStoreType {
	case AppStore:
		appUrl = fmt.Sprintf("https://apps.apple.com/app/%s", version.AppId)
		if len(version.Lang) > 0 {
			appUrl = fmt.Sprintf("https://apps.apple.com/%s/app/%s", version.Lang, version.AppId)
		}
	case GooglePlay:
		appUrl = fmt.Sprintf("https://play.google.com/store/apps/details?id=%s", version.AppId)
		if len(version.Lang) > 0 {
			appUrl = fmt.Sprintf("https://play.google.com/store/apps/details?id=%s&hl=%s", version.AppId, version.Lang)
		}
	}

	return appUrl
}

// 儲存推播所需要的資訊
func (_appVersion *AppVersion) PushSettings(database *gorm.DB, dictionary map[string]interface{}) (map[string]interface{}, error) {

	result := map[string]interface{}{}
	info := map[string]interface{}{}
	isSuccess := false

	_id := dictionary["id"]
	_storeSetting := dictionary["store"]
	_prodSetting := dictionary["prod"]
	_devSetting := dictionary["dev"]

	result["isSuccess"] = isSuccess

	if _id == nil {
		return result, errors.New("id is null")
	}

	appVersion := _appVersion.Select(database, _id.(string))

	if appVersion.ID == 0 {
		return result, errors.New("version is null")
	}

	storeJSON, error := utility.MapToJSON(_storeSetting.(map[string]interface{}))
	if error != nil {
		return result, error
	}

	prodJSON, error := utility.MapToJSON(_prodSetting.(map[string]interface{}))
	if error != nil {
		return result, error
	}

	devJSON, error := utility.MapToJSON(_devSetting.(map[string]interface{}))
	if error != nil {
		return result, error
	}

	info["push_setting"] = storeJSON
	info["push_setting_production"] = prodJSON
	info["push_setting_develop"] = devJSON

	error = database.Model(&appVersion).Updates(info).Error

	if error == nil {
		isSuccess = true
	}

	result["isSuccess"] = isSuccess
	return result, error
}

// 發送推播
func (_appVersion *AppVersion) PushNotification(database *gorm.DB, dictionary map[string]interface{}) (map[string]interface{}, error) {

	var _pushToken PushToken

	result, error := _appVersion.PushSettings(database, dictionary)
	isSuccess := result["isSuccess"].(bool)
	category := "AppNotificationCategory"
	isProduction := false

	if !isSuccess {
		return result, error
	}

	_id := dictionary["id"]
	_pushType := dictionary["pushType"]
	pushType := Store

	if _id == nil {
		return result, errors.New("id is null")
	}

	if _pushType == nil {
		result["isSuccess"] = false
		return result, errors.New("pushType is null")
	}

	if _pushType.(float64) == 0 {
		pushType = Store
	}
	if _pushType.(float64) == 1 {
		pushType = Production
	}
	if _pushType.(float64) == 2 {
		pushType = Develop
	}

	switch pushType {
	case Store:
		category = "AppNotificationCategory"
		isProduction = true
	case Develop:
		category = "DemoNotificationCategory"
		isProduction = false
	case Production:
		category = "DemoNotificationCategory"
		isProduction = false
	}

	appVersion := _appVersion.Select(database, _id.(string))

	if appVersion.ID == 0 {
		return result, errors.New("version is null")
	}

	info := parseAPNSInformation(appVersion, pushType)
	payload := parsePayload(dictionary, pushType)

	if info == nil {
		result["isSuccess"] = false
		return result, errors.New("payload is null")
	}

	tokens := _pushToken.TokenList(database, info.Topic)

	apnsMessage := APNSMessage{
		IsProduction: isProduction,
		DeviceToken:  "",
		Title:        payload["title"].(string),
		Message:      payload["message"].(string),
		Image:        payload["image"].(string),
		Category:     category,
	}

	for index := 0; index < len(tokens); index++ {

		apnsMessage.DeviceToken = tokens[index]

		response, _ := _appVersion.APNS_PushNotification(*info, apnsMessage)
		result["isSuccess"] = isSuccess
		result["response"] = response
	}

	return result, error
}

// iOS推播功能
func (_appVersion *AppVersion) APNS_PushNotification(info APNSInformation, apnsMessage APNSMessage) (*apns2.Response, error) {

	var authKey *ecdsa.PrivateKey = nil
	var response *apns2.Response = nil
	var error error = nil
	var client *apns2.Client = nil

	context := fmt.Sprintf(`{"aps":{"alert":{"title":"%s","body":"%s","badge":87},"category":"%s"},"payload":{"image":"%s"}}`, apnsMessage.Title, apnsMessage.Message, apnsMessage.Category, apnsMessage.Image)
	payload := []byte(context)

	utility.Println(context)

	authKey, error = token.AuthKeyFromFile(info.AuthKeyP8)
	if error != nil {
		return response, error
	}

	notification := &apns2.Notification{}
	notification.Topic = info.Topic
	notification.DeviceToken = apnsMessage.DeviceToken
	notification.Payload = payload

	token := &token.Token{
		AuthKey: authKey,
		KeyID:   info.KeyID,
		TeamID:  info.TeamID,
	}

	if apnsMessage.IsProduction {
		client = apns2.NewTokenClient(token).Production()
	} else {
		client = apns2.NewTokenClient(token).Development()
	}

	response, error = client.Push(notification)
	return response, error
}

// 將iOS的JSON => APNSInformation
func parseAPNSInformation(version AppVersion, pushType PushType) *APNSInformation {

	rawString := version.PushSetting

	if pushType == Store {
		rawString = version.PushSetting
	}
	if pushType == Production {
		rawString = version.PushSettingProduction
	}
	if pushType == Develop {
		rawString = version.PushSettingDevelop
	}

	json := utility.RawStringToMap(rawString)
	_bundleId := json["bundleId"]
	_teamId := json["teamId"]
	_keyId := json["keyId"]

	if _bundleId == nil {
		return nil
	}
	if _teamId == nil {
		return nil
	}
	if _keyId == nil {
		return nil
	}

	authKeyP8 := fmt.Sprintf(`./material/AuthKey_%v.p8`, _keyId.(string))

	info := APNSInformation{
		Topic:     _bundleId.(string),
		KeyID:     _keyId.(string),
		TeamID:    _teamId.(string),
		AuthKeyP8: authKeyP8,
	}

	return &info
}

// 將推播訊息的String解開
func parsePayload(info map[string]interface{}, pushType PushType) map[string]interface{} {

	dictionary := info["store"].(map[string]interface{})

	if pushType == Store {
		dictionary = info["store"].(map[string]interface{})
	}
	if pushType == Production {
		dictionary = info["prod"].(map[string]interface{})
	}
	if pushType == Develop {
		dictionary = info["dev"].(map[string]interface{})
	}

	result := map[string]interface{}{
		"title":   dictionary["title"],
		"message": dictionary["message"],
		"image":   dictionary["image"],
	}

	return result
}
