package utility

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/sideshow/apns2"
	"github.com/sideshow/apns2/token"
)

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

// iOS推播功能
func PushNotification(info APNSInformation, apnsMessage APNSMessage) (*apns2.Response, error) {

	var authKey *ecdsa.PrivateKey = nil
	var response *apns2.Response = nil
	var error error = nil
	var client *apns2.Client = nil

	context := fmt.Sprintf(`{"aps":{"alert":{"title":"%s","body":"%s","badge":87},"category":"%s"},"payload":{"image":"%s"}}`, apnsMessage.Title, apnsMessage.Message, apnsMessage.Category, apnsMessage.Image)
	payload := []byte(context)

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
