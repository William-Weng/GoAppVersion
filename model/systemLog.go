package model

import (
	"fmt"
	"william/utility"

	"gorm.io/gorm"
)

// 紀錄Log
type SystemLog struct {
	gorm.Model
	Log    string `json:"log"`
	Result string `json:"result"`
	Error  string `json:"error"`
}

// 紀錄Log
func (_systemLog *SystemLog) Insert(database *gorm.DB, log interface{}, result interface{}, err error) error {
	_log := fmt.Sprintf(`%v`, log)
	_error := fmt.Sprintf(`%v`, err)
	_result := fmt.Sprintf(`%v`, result)

	dbError := database.Create(&SystemLog{Log: _log, Result: _result, Error: _error}).Error
	resultInfo := fmt.Sprintf("[Log] => %v, [Result] => %v, [Error] => %v", _log, _result, _error)

	utility.Println(resultInfo)
	return dbError
}
