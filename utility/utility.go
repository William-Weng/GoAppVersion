package utility

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// [仿forEach()](https://openhome.cc/Gossip/Go/Closure.html) => forEach(numbers, func(element int, index int) { sum += element })
func ForEach[T any](elements []T, closure func(T, int)) {
	for index, element := range elements {
		closure(element, index)
	}
}

// [仿map()](https://medium.com/starbugs/why-go-need-generics-c8f1495ef00a) => title := Map(numbers, func(number int, index int) string { return strconv.Itoa(number * 2) })
func Map[T, U any](elements []T, closure func(T, int) U) []U {

	newElements := make([]U, len(elements))

	ForEach(elements, func(element T, index int) {
		newElements[index] = closure(element, index)
	})

	return newElements
}

// [印出結果](https://dotblogs.com.tw/DizzyDizzy/2019/11/29/ColorfulCLI) => log.SetFlags(log.LstdFlags | log.Lshortfile)
func Println(value ...any) {

	now := time.Now()
	pc, filename, line, _ := runtime.Caller(1)
	functionName := runtime.FuncForPC(pc).Name()

	content := fmt.Sprintf("\033[37;103m[LOG] %v\033[0m", now)
	fmt.Println(content)

	content = fmt.Sprintf("\033[37;103m%s: %v => %v\033[0m", filename, line, functionName)
	fmt.Println(content)
	fmt.Println(value...)
}

// 文字 => 數字
func StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// 文字 => float64
func StringToFloat64(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

// 文字是否是空字串
func StringIsEmpty(str string) bool {
	_string := strings.TrimSpace(str)
	return len(_string) == 0
}

// 讀取出類別
func TypeOf(object any) reflect.Type {
	return reflect.TypeOf(object)
}

// 將網路傳來的json文字 => map
func RawStringToMap(json string) map[string]interface{} {
	dictionary := RawBodyToMap([]byte(json))
	return dictionary
}

// 將網路傳來的[]byte => map
func RawBodyToMap(rawJSON []byte) map[string]interface{} {

	dictionary := map[string]interface{}{}
	json.Unmarshal(rawJSON, &dictionary)

	return dictionary
}

// 將Map => JSON String
func MapToJSON[dict map[string]interface{} | map[string]bool](dictionary dict) (string, error) {
	bytes, error := json.Marshal(dictionary)
	return string(bytes), error
}

// [將字串 => MD5(字串)](https://www.perfcode.com/p/golang-md5.html)
func Md5(word string) string {

	_md5 := md5.New()
	_bytes := []byte(word)
	_md5.Write(_bytes)

	return hex.EncodeToString(_md5.Sum(nil))
}
