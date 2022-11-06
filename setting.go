package main

import "william/model"

const (
	AutoUpdate              bool   = true
	ValidityDays            int    = 14
	SqliteDatabasePath      string = "./material/database.sqlite3" // 把資料庫存放在material資料夾下
	AutoUpdateCronRule      string = "*/60 * * * *"
	AutoRemoveTokenCronRule string = "*/60 * * * *"
	Port                    string = ":12345"
	Sheet                   string = "App版本相關訊息"
	AuthorizationKey        string = "Authorization"
)

const (
	UserMinimumLength     int = 8
	PasswordMinimumLength int = 8
)

var (
	// 預先設定的帳號
	SuperUsers = map[string]map[string]model.UserLevel{
		"root":    {"3939889": model.Root},
		"admin":   {"28825252": model.Admin},
		"william": {"987987": model.Account},
	}

	// API的通行規則
	ApiRules = map[model.UserLevel]map[string][]string{
		model.Account: {
			"GET":  {"/appVersion/:id"},
			"POST": {"/checkToken", "/versionList/:page", "/saveAppVersion", "/forceUpdate", "/pushToken/:appId"},
		},
		model.Admin: {"GET": {"*"}, "POST": {"*"}, "PUT": {"*"}, "PATCH": {"*"}},
		model.Root:  {"GET": {"*"}, "POST": {"*"}, "PUT": {"*"}, "PATCH": {"*"}, "DELETE": {"*"}},
	}
)
