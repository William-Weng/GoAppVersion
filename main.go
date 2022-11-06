package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"william/model"
	"william/utility"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseType int8

const (
	Sqlite3 DatabaseType = 0 // 支援Sqlite3
	MySql   DatabaseType = 1 // 支援MySql
)

var gormDatabase *gorm.DB

// 建立資料庫
func (dbType DatabaseType) createDatabase(path string) (*gorm.DB, error) {

	var database *gorm.DB
	var error error

	switch dbType {
	case Sqlite3:
		database, error = gorm.Open(sqlite.Open(path), &gorm.Config{})
	case MySql:
		break
	}

	return database, error
}

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func main() {
	error := initSetting()

	if error != nil {
		return
	}
}

// 初始設定
func initSetting() error {

	database, error := Sqlite3.createDatabase(SqliteDatabasePath)
	if error != nil {
		return error
	}

	gormDatabase = database

	error = createTables(database)
	if error != nil {
		return error
	}

	initRouter(database)

	return nil
}

// 建立資料表
func createTables(database *gorm.DB) error {

	error := database.AutoMigrate(&model.SystemLog{})
	if error != nil {
		return error
	}

	error = database.AutoMigrate(&model.User{})
	if error != nil {
		return error
	}

	error = database.AutoMigrate(&model.AuthToken{})
	if error != nil {
		return error
	}

	error = database.AutoMigrate(&model.AppVersion{})
	if error != nil {
		return error
	}

	error = database.AutoMigrate(&model.PushToken{})
	if error != nil {
		return error
	}

	return error
}

// 初始化Router相關設定
func initRouter(database *gorm.DB) {

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20
	router.Use(cors.New(corsConfig()))

	registerWebAPI(router, database)
	router.Run(Port)
}

// 註冊API
func registerWebAPI(router *gin.Engine, database *gorm.DB) {

	initAdminUser(database)
	checkToken(router, database)
	selectAppVersion(router, database)
	selectAppVersionMini(router, database)
	userLogin(router, database)
	pushToken(router, database)
	cronSchedule(database)

	router.Use(checkTokenMiddleware)

	checkAppVersion(router)
	insertAppVersion(router, database)
	selectAppVersionList(router, database)
	updateAppVersion(router, database)
	deleteAppVersion(router, database)
	saveAppVersions(router, database)
	uploadAppVersions(router, database)
	forceUpdateVersions(router, database)
	updateVersionType(router, database)
	pushTokenList(router, database)
	pushSetting(router, database)
	pushNotification(router, database)

	insertUser(router, database)
	selectUser(router, database)

	authToken(router, database)
}

// [加入管理者帳號](https://www.ez2o.com/App/Coder/RandomPassword)
func initAdminUser(database *gorm.DB) {

	var systemLog model.SystemLog
	var _user model.User

	for username, info := range SuperUsers {

		dictionary := map[string]interface{}{
			"username": username,
		}

		for _password, _level := range info {

			level := 0.0

			if _level == model.Account {
				level = 0
			}
			if _level == model.Admin {
				level = 9
			}
			if _level == model.Root {
				level = 99
			}

			dictionary["password"] = _password
			dictionary["level"] = level

			result, error := _user.Insert(database, dictionary)

			systemLog.Insert(database, nil, result, error)
		}
	}
}

// [CORS相關設定](https://pjchender.dev/golang/gin-cors/)
func corsConfig() cors.Config {

	corsConfig := cors.DefaultConfig() // 防止CORS的設定
	corsConfig.AllowAllOrigins = true  // 防止CORS的設定
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}

	return corsConfig
}

// MARK: - WebAPI (AppVersion)
// <POST> 新增 (localhost:12345/appVersion)
// => {"type":0,"id":"id443904275","name":"LINE"}
// => {"type":1,"id":"com.ubercab.driver","name":"Uber Driver"}
func insertAppVersion(router *gin.Engine, database *gorm.DB) {

	var appVersion model.AppVersion

	router.POST("/appVersion", func(context *gin.Context) {

		dictionary := utility.RequestBodyToMap(context)
		result, error := appVersion.Insert(database, dictionary)

		utility.ContextJSON(context, http.StatusOK, result, error)
	})
}

// <GET> 搜尋
// => localhost:12345/appVersion/com.ubercab.driver
// => localhost:12345/appVersion/id443904275
func selectAppVersion(router *gin.Engine, database *gorm.DB) {

	var appVersion model.AppVersion

	router.GET("/appVersion/:id", func(context *gin.Context) {

		id := context.Param("id")
		result := appVersion.Select(database, id)

		if result.ID == 0 {
			utility.ContextJSON(context, http.StatusOK, nil, nil)
			return
		}

		utility.ContextJSON(context, http.StatusOK, result, nil)
	})
}

// <GET> 搜尋 (簡單版)
// => localhost:12345/appVersion/mobile/com.ubercab.driver
// => localhost:12345/appVersion/mobile/id443904275
func selectAppVersionMini(router *gin.Engine, database *gorm.DB) {

	var appVersion model.AppVersion

	router.GET("/appVersion/mobile/:id", func(context *gin.Context) {

		id := context.Param("id")
		result := appVersion.SelectMini(database, id)

		utility.ContextJSON(context, http.StatusOK, result, nil)
	})
}

// <POST> 搜尋
// => localhost:12345/appVersion
func selectAppVersionList(router *gin.Engine, database *gorm.DB) {

	var appVersion model.AppVersion

	router.POST("/versionList/:page", func(context *gin.Context) {

		dictionary := utility.RequestBodyToMap(context)
		versions, count := appVersion.SelectWhere(database, dictionary)

		result := map[string]interface{}{
			"count":    count,
			"versions": versions,
		}

		utility.ContextJSON(context, http.StatusOK, result, nil)
	})
}

// <PATCH> 更新APP版本號
// => localhost:12345/appVersion/id443904275/0 + {version":{"store":"1.0.0","prod":"1.0.0","dev":"1.0.0"}}
// => localhost:12345/com.ubercab.driver/1 + {"auto":false,"name":"優食 - Android","version":{"store":"1.2.3"}}
func updateAppVersion(router *gin.Engine, database *gorm.DB) {

	var appVersion model.AppVersion

	router.PATCH("/appVersion/:id/:type", func(context *gin.Context) {

		dictionary := utility.RequestBodyToMap(context)

		_id := context.Param("id")
		_type := context.Param("type")
		dictionary["id"] = _id
		dictionary["type"] = _type

		result, error := appVersion.Update(database, dictionary)

		utility.ContextJSON(context, http.StatusOK, result, error)
	})
}

// <PATCH> 設定更新樣式
// => localhost:12345/updateType/id304878510 + {"type":"Store","value":0}
func updateVersionType(router *gin.Engine, database *gorm.DB) {

	var appVersion model.AppVersion

	router.PATCH("/updateType/:id", func(context *gin.Context) {

		dictionary := utility.RequestBodyToMap(context)

		_id := context.Param("id")
		dictionary["id"] = _id

		result, error := appVersion.UpdateType(database, dictionary)
		utility.ContextJSON(context, http.StatusOK, result, error)
	})
}

// <PUT> 取得上架版的APP版本號
// => localhost:12345/appVersion/id443904275/0
// => localhost:12345/com.ubercab.driver/1
func checkAppVersion(router *gin.Engine) {

	var appVersion model.AppVersion

	router.PUT("/appVersion/:id/:type", func(context *gin.Context) {

		dictionary := utility.RequestBodyToMap(context)

		_id := context.Param("id")
		_type := context.Param("type")
		dictionary["id"] = _id
		dictionary["type"] = _type

		list, error := appVersion.CheckAppVersion(dictionary)
		utility.ContextJSON(context, http.StatusOK, list, error)
	})
}

// <DELETE> 刪除該APP資訊
// => localhost:12345/appVersion/com.ubercab.driver
// => localhost:12345/appVersion/id443904275
func deleteAppVersion(router *gin.Engine, database *gorm.DB) {

	var appVersion model.AppVersion

	router.DELETE("/appVersion/:id", func(context *gin.Context) {

		_id := context.Param("id")
		result := appVersion.Delete(database, _id)

		utility.ContextJSON(context, http.StatusOK, result, nil)
	})
}

// <POST> 上傳檔案文件
func uploadAppVersions(router *gin.Engine, database *gorm.DB) {

	var excelTools model.ExcelTools

	router.POST("/uploadExcel", func(context *gin.Context) {

		_, header, error := context.Request.FormFile("file")
		result := map[string]interface{}{"filename": header.Filename, "size": header.Size, "filePath": nil}

		defer func() {
			utility.ContextJSON(context, http.StatusOK, result, error)
		}()

		if error != nil {
			return
		}

		filePath := fmt.Sprintf(`./html/public/file/%v_%v`, time.Now().Unix(), header.Filename)
		result["filePath"] = filePath

		error = context.SaveUploadedFile(header, filePath)

		if error != nil {
			return
		}

		result, error = excelTools.InsertExcel(database, Sheet, filePath)
	})
}

// <POST> 強制更新Versions
func forceUpdateVersions(router *gin.Engine, database *gorm.DB) {

	router.POST("/forceUpdate", func(context *gin.Context) {
		result := updateAllAppVersion(database)
		utility.ContextJSON(context, http.StatusOK, result, nil)
	})
}

// MARK: - WebAPI (User)
// <POST> 使用者登入 (localhost:12345/login)
// => {"username":"william","password":"123456"}
func userLogin(router *gin.Engine, database *gorm.DB) {

	var user model.User

	router.POST("/login", func(context *gin.Context) {

		var result map[string]interface{}
		var error error

		defer func() {
			utility.ContextJSON(context, http.StatusOK, result, nil)
		}()

		dictionary := utility.RequestBodyToMap(context)
		expiresTime := time.Now().AddDate(0, 0, ValidityDays)
		username := dictionary["username"].(string)
		password := dictionary["password"].(string)

		result, error = user.Login(database, username, password, expiresTime)

		if error != nil {
			return
		}

		token := result["token"].(*string)

		if token != nil {
			storeAuthToken(database, username, *token, expiresTime)
		}
	})
}

// <POST> 新增使用者 (localhost:12345/user)
// => {"username":"william","password":"123456"}
func insertUser(router *gin.Engine, database *gorm.DB) {

	var user model.User

	router.POST("/user", func(context *gin.Context) {

		dictionary := utility.RequestBodyToMap(context)
		result, error := user.Insert(database, dictionary)

		utility.ContextJSON(context, http.StatusOK, result, error)
	})
}

// <GET> 搜尋單一使用者 (localhost:12345/user/william)
func selectUser(router *gin.Engine, database *gorm.DB) {

	var user model.User

	router.GET("/user/:username", func(context *gin.Context) {

		username := context.Param("username")
		result := user.Select(database, username)

		utility.ContextJSON(context, http.StatusOK, result, nil)
	})
}

// <POST> 測試Token合法性 (http://localhost:12345/checkToken)
// => Header: {"Authorization": "Bearer <Token>"}
func checkToken(router *gin.Engine, database *gorm.DB) {

	var authToken model.AuthToken

	router.POST("/checkToken", func(context *gin.Context) {

		dictionary := utility.RequestBodyToMap(context)

		_token := dictionary["token"]

		if _token == nil {
			utility.ContextJSON(context, http.StatusOK, nil, nil)
			return
		}

		token := phaseBearerToken(_token.(string))
		result, error := authToken.CheckToken(database, token)

		utility.ContextJSON(context, http.StatusOK, result, error)
	})
}

// MARK: - WebAPI (AuthToken)
// <POST> Token列表
// => {"where":"id > 0"}
func authToken(router *gin.Engine, database *gorm.DB) {

	var authToken model.AuthToken

	router.POST("/authToken", func(context *gin.Context) {

		dictionary := utility.RequestBodyToMap(context)
		where := dictionary["where"]

		if where == nil {
			where = "id > 0"
		}

		result := authToken.SelectWhere(database, where.(string))
		utility.ContextJSON(context, http.StatusOK, result, nil)
	})
}

// MARK: - Tools
// [定時排程任務](https://www.readfog.com/a/1637371620314157056)
func cronSchedule(database *gorm.DB) {

	schedule := cron.New()

	schedule.AddFunc(AutoUpdateCronRule, func() {
		updateAllAppVersion(database)
	})

	schedule.AddFunc(AutoRemoveTokenCronRule, func() {
		removeExpiresToken(database)
	})

	schedule.Start()
}

// <POST> 儲存成Excel
func saveAppVersions(router *gin.Engine, database *gorm.DB) {

	var appVersion model.AppVersion

	router.POST("/saveAppVersion", func(context *gin.Context) {

		var excelTools model.ExcelTools

		versions, _ := appVersion.SelectWhere(database, nil)
		filename := fmt.Sprintf("AppId_%v.xlsx", time.Now().Unix())

		excelTools.Sheet = Sheet
		excelTools.CreateExcel()

		result, error := excelTools.SaveAppVersions(versions, filename)

		utility.ContextJSON(context, http.StatusOK, result, error)
	})
}

// MARK: - WebAPI (AuthToken)
// <POST> 註冊 / 推播PushToken
func pushToken(router *gin.Engine, database *gorm.DB) {

	var pushToken model.PushToken

	router.POST("/pushToken/:appId", func(context *gin.Context) {

		_appId := context.Param("appId")

		dictionary := utility.RequestBodyToMap(context)
		dictionary["id"] = _appId
		dictionary["IP"] = context.ClientIP()

		result, error := pushToken.InsertOrUpdate(database, dictionary)
		utility.ContextJSON(context, http.StatusOK, result, error)
	})
}

// MARK: - WebAPI (AuthToken)
// <POST> Token列表
// => {"where":"id > 0"}
func pushTokenList(router *gin.Engine, database *gorm.DB) {

	var pushToken model.PushToken

	router.POST("/pushToken", func(context *gin.Context) {

		dictionary := utility.RequestBodyToMap(context)
		where := dictionary["where"]

		if where == nil {
			where = "id > 0"
		}

		result := pushToken.SelectWhere(database, where.(string))
		utility.ContextJSON(context, http.StatusOK, result, nil)
	})
}

// <POST> 更新推播相關設定值
func pushSetting(router *gin.Engine, database *gorm.DB) {

	var version model.AppVersion

	router.POST("/pushSettings/:appId", func(context *gin.Context) {

		_appId := context.Param("appId")

		dictionary := utility.RequestBodyToMap(context)
		dictionary["id"] = _appId

		result, error := version.PushSettings(database, dictionary)
		utility.ContextJSON(context, http.StatusOK, result, error)
	})
}

// <POST> 發送推播
func pushNotification(router *gin.Engine, database *gorm.DB) {

	var version model.AppVersion

	router.POST("/push/:appId", func(context *gin.Context) {

		_appId := context.Param("appId")

		dictionary := utility.RequestBodyToMap(context)
		dictionary["id"] = _appId

		result, error := version.PushNotification(database, dictionary)
		utility.ContextJSON(context, http.StatusOK, result, error)
	})
}

// MARK: - 小工具
// 更新資料庫內(auto_update = true)的版本號
func updateAllAppVersion(database *gorm.DB) map[string]int {

	var versionList []model.AppVersion
	var _database = database

	_database = _database.Where("auto_update == ?", true)
	_database.Find(&versionList)

	for index := 0; index < len(versionList); index++ {
		versionInfo := versionList[index]
		autoUpdateAppVersion(database, versionInfo)
	}

	return map[string]int{"count": len(versionList)}
}

// 刪除過期的Token
func removeExpiresToken(database *gorm.DB) {

	var authToken model.AuthToken

	where := fmt.Sprintf(`expires_time < "%v"`, time.Now().AddDate(0, 0, 0))
	authToken.DeleteWhere(database, where)
}

// 更新資料庫內的版本號
func autoUpdateAppVersion(database *gorm.DB, info model.AppVersion) {

	var appVersion model.AppVersion

	var id = info.AppId
	var lang = info.Lang
	var allError error = nil
	var result map[string]interface{}

	dictionary := map[string]interface{}{}
	dictionary["type"] = fmt.Sprintf(`%v`, info.AppStoreType)
	dictionary["id"] = id
	dictionary["lang"] = lang

	defer func() {
		var systemLog model.SystemLog
		systemLog.Insert(database, dictionary, result, allError)
	}()

	list, error := appVersion.CheckAppVersion(dictionary)

	if error != nil {
		allError = error
		return
	}

	dictionary["version"] = map[string]interface{}{"store": list["version"]}
	dictionary["icon"] = list["icon"]
	dictionary["supportedDevices"] = list["supportedDevices"]

	_result, error := appVersion.Update(database, dictionary)
	if error != nil {
		allError = error
		return
	}

	result = _result
}

// 記錄Username / AuthToken
func storeAuthToken(database *gorm.DB, username string, token string, expiresTime time.Time) error {
	error := database.Create(&model.AuthToken{Username: strings.ToUpper(username), Token: token, ExpiresTime: expiresTime}).Error
	return error
}

// {"Authorization": "Bearer <Token>"} => <Token>
func phaseAuthorizationToken(context *gin.Context) string {
	token := context.GetHeader(AuthorizationKey)
	return phaseBearerToken(token)
}

// "Bearer <Token>" => "<Token>""
func phaseBearerToken(token string) string {

	const tokenPrefix = "Bearer "

	if len(token) > len(tokenPrefix) {
		return token[len(tokenPrefix):]
	}

	return ""
}

// 中間件 (卡在中間擋路的) => 測試是否登入？沒有沒Token
func checkTokenMiddleware(context *gin.Context) {

	var authToken model.AuthToken

	token := phaseAuthorizationToken(context)
	result, error := authToken.CheckToken(gormDatabase, token)
	level := result["level"]

	if error != nil || level == nil {
		utility.AbortWithStatusJSON(context, http.StatusUnauthorized, result, error)
		return
	}

	isMatch := authToken.CheckLevel(level.(model.UserLevel), context.Request.URL, context.Request.Method, ApiRules)

	if !isMatch {
		utility.AbortWithStatusJSON(context, http.StatusUnauthorized, result, error)
		return
	}

	context.Next()
}
