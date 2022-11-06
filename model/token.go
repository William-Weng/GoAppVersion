package model

import (
	"errors"
	"fmt"
	"net/url"
	"regexp"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

// 紀錄Token
type AuthToken struct {
	gorm.Model
	Username    string    `json:"username"`
	Token       string    `json:"token" gorm:"index:idx_token,unique"`
	ExpiresTime time.Time `json:"expiresTime"`
}

// 搜尋
func (_authToken *AuthToken) Select(database *gorm.DB, token string) AuthToken {

	var authToken AuthToken
	database.Take(&authToken, "token=?", token)

	return authToken
}

// [搜尋](https://5xruby.tw/posts/what-is-jwt)
func (_authToken *AuthToken) SelectWhere(database *gorm.DB, where string) []AuthToken {

	var authTokenList []AuthToken
	database.Where(where).Order("expires_time asc, username asc").Find(&authTokenList)

	return authTokenList
}

// 刪除
func (_authToken *AuthToken) DeleteWhere(database *gorm.DB, where string) map[string]interface{} {

	var authToken AuthToken
	result := map[string]interface{}{}

	database.Where(where).Unscoped().Delete(&authToken)
	result["isSuccess"] = true

	return result
}

// [測試Token (資料庫要有存 => 合法的)](https://www.jianshu.com/p/50ade6f2e4fd)
func (_authToken *AuthToken) CheckToken(database *gorm.DB, token string) (map[string]interface{}, error) {

	authToken := _authToken.Select(database, token)

	if authToken.ID == 0 {
		return nil, errors.New("token is null")
	}

	result, error := _authToken.validateToken(token)
	return result, error
}

// [用 Go 實現 JWT Based 驗證](https://waynestalk.com/go-jwt-tutorial/)
func (_authToken *AuthToken) JwtToken(user User, expiresTime time.Time) (*string, error) {

	claims := jwt.RegisteredClaims{
		Subject:   user.Username,
		ExpiresAt: jwt.NewNumericDate(expiresTime),
	}

	authClaims := authClaims{
		RegisteredClaims: claims,
		UserID:           user.ID,
		Level:            user.Level,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims)
	tokenString, error := token.SignedString(jwtKey)

	if error != nil {
		return nil, error
	}

	return &tokenString, nil
}

// [驗證Token](https://blog.csdn.net/default7/article/details/123530116)
func (_authToken *AuthToken) validateToken(jwtToken string) (map[string]interface{}, error) {

	var claims authClaims

	result := map[string]interface{}{
		"id":       0,
		"username": nil,
		"level":    Account,
	}

	token, error := jwt.ParseWithClaims(jwtToken, &claims, func(token *jwt.Token) (interface{}, error) {

		_, isSuccess := token.Method.(*jwt.SigningMethodHMAC)

		if !isSuccess {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})

	if error != nil {
		return result, error
	}

	if !token.Valid {
		return result, errors.New("invalid token")
	}

	result = map[string]interface{}{
		"id":          claims.UserID,
		"username":    claims.Subject,
		"level":       claims.Level,
		"expiresTime": claims.ExpiresAt.Time,
	}

	return result, nil
}

// [設定API權限](https://www.jianshu.com/p/fee89fbb5f27)
// => [將uri轉成正規式過濾](https://www.wjsljc.com/document/ts-axios/axios/chapter1/require.html)
func (_authToken *AuthToken) CheckLevel(level UserLevel, url *url.URL, method string, rules map[UserLevel]map[string][]string) bool {

	isMatch := false
	urlString := url.String()

	urls := rules[level][method]

	for _, _url := range urls {

		if _url == "*" {
			isMatch = true
			return isMatch
		}

		rule := `:[\w]+`
		regular := regexp.MustCompile(rule)
		newRule := regular.ReplaceAllString(_url, `[\w]+`)

		regular = regexp.MustCompile(newRule)
		isMatch = regular.MatchString(urlString)

		if isMatch {
			return isMatch
		}
	}

	return isMatch
}
