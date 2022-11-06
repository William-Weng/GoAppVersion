package model

import (
	"fmt"
	"william/utility"

	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

// [Excelize 繁體字文檔](https://xuri.me/excelize/zh-tw/)
type ExcelTools struct {
	Sheet string
	excel *excelize.File
}

type excelField struct {
	Index int
	Key   string
	Value interface{}
}

// 建立工作表
func (_excelTools *ExcelTools) CreateExcel() int {

	_excelTools.excel = excelize.NewFile()

	index := _excelTools.excel.NewSheet(_excelTools.Sheet)
	_excelTools.excel.DeleteSheet("Sheet1")
	_excelTools.excel.SetActiveSheet(index)

	return index
}

// 儲存工作表
func (_excelTools *ExcelTools) SaveAppVersions(versions []AppVersion, filename string) (map[string]interface{}, error) {

	_excel := _excelTools.excel
	_sheet := _excelTools.Sheet

	fields := []excelField{
		{Index: 0, Key: "A", Value: "App Name"},
		{Index: 0, Key: "B", Value: "App Id"},
		{Index: 0, Key: "C", Value: "App Store"},
		{Index: 0, Key: "D", Value: "App Lang"},
	}

	for _index := 0; _index < len(fields); _index++ {

		fields[_index].Index += 1

		key := fmt.Sprintf("%v%v", fields[_index].Key, fields[_index].Index)
		value := fmt.Sprintf("%v", fields[_index].Value)

		_excel.SetCellValue(_sheet, key, value)
	}

	for index := 0; index < len(versions); index++ {

		_version := versions[index]
		_values := []interface{}{_version.AppName, _version.AppId, _version.AppStoreType, _version.Lang}

		for _index := 0; _index < len(fields); _index++ {

			fields[_index].Index += 1

			key := fmt.Sprintf("%v%v", fields[_index].Key, fields[_index].Index)
			value := fmt.Sprintf("%v", _values[_index])

			_excel.SetCellValue(_sheet, key, value)
		}
	}

	filePath := fmt.Sprintf("/excel/%s", filename)
	savePath := fmt.Sprintf("./html/public%s", filePath)
	error := _excel.SaveAs(savePath)

	result := map[string]interface{}{}
	result["filePath"] = filePath

	return result, error
}

// 將Excel匯入DB
func (_excelTools *ExcelTools) InsertExcel(database *gorm.DB, sheet string, filePath string) (map[string]interface{}, error) {

	var appVersion AppVersion
	result := map[string]interface{}{"count": 0}
	count := 0

	excel, error := excelize.OpenFile(filePath)

	if error != nil {
		return result, error
	}

	rows, _error := excel.GetRows(sheet)
	if _error != nil {
		return result, _error
	}

	for index, row := range rows {

		if index == 0 {
			continue
		}

		info := map[string]interface{}{}
		info["name"] = row[0]
		info["id"] = row[1]
		info["type"], _ = utility.StringToFloat64(row[2])

		if len(row) > 3 {
			info["lang"] = row[3]
		}

		_result, _ := appVersion.Insert(database, info)

		if _result["isSuccess"] == true {
			count += 1
		}
	}

	result["count"] = count
	error = excel.Close()

	return result, error
}
