//
//  Extension+.swift
//  WebViewApp
//
//  Created by iOS on 2022/10/12.
//

import WebKit
import MobileCoreServices
import UniformTypeIdentifiers
import StoreKit
import SafariServices
import WWPrint

// MARK: - String (override class function)
extension String {
    
    /// [CountableRange](https://swiftdoc.org/v3.0/type/countablerange/)
    /// "subscript"[2...] => "bscript"
    subscript (bounds: CountablePartialRangeFrom<Int>) -> String {
        let start = index(startIndex, offsetBy: bounds.lowerBound)
        return String(self[start...])
    }
}

extension String {
    
    /// [MIME => UTI](https://gist.github.com/ddeville/1527517)
    /// => "text/html" -> "public.html"
    func _mimeToUTI() -> String? {
        let utiType = UTTypeCreatePreferredIdentifierForTag(kUTTagClassMIMEType, self as CFString, nil)?.takeRetainedValue() as String?
        return utiType
    }
    
    /// [UTI => MIME](https://developer.apple.com/library/archive/documentation/FileManagement/Conceptual/understanding_utis/understand_utis_conc/understand_utis_conc.html)
    /// => "public.html" -> "text/html"
    func _utiToMIME() -> String? {
        let mimeType: String? = UTTypeCopyPreferredTagWithClass(self as CFString, kUTTagClassMIMEType)?.takeRetainedValue() as String?
        return mimeType
    }
    
    /// JSON String => [String: Any]
    /// - Parameters:
    ///   - encoding: 字元編碼
    ///   - options: JSON序列化讀取方式
    /// - Returns: Any?
    func _dictionary(encoding: String.Encoding, options: JSONSerialization.ReadingOptions = .allowFragments) -> [String: Any]? {
        let dictionary = self._jsonObject(encoding: encoding, options: options) as? [String: Any]
        return dictionary
    }
    
    /// JSON String => JSON Object
    /// - Parameters:
    ///   - encoding: 字元編碼
    ///   - options: JSON序列化讀取方式
    /// - Returns: Any?
    func _jsonObject(encoding: String.Encoding = .utf8, options: JSONSerialization.ReadingOptions = .allowFragments) -> Any? {
        
        guard let data = self._data(using: encoding),
              let jsonObject = try? JSONSerialization.jsonObject(with: data, options: options)
        else {
            return nil
        }
        
        return jsonObject
    }

    /// URL編碼 (百分比)
    /// - 是在哈囉 => %E6%98%AF%E5%9C%A8%E5%93%88%E5%9B%89
    /// - Parameter characterSet: 字元的判斷方式
    /// - Returns: String?
    func _encodingURL(characterSet: CharacterSet = .urlQueryAllowed) -> String? { return addingPercentEncoding(withAllowedCharacters: characterSet) }
    
    /// String => Data
    /// - Parameters:
    ///   - encoding: 字元編碼
    ///   - isLossyConversion: 失真轉換
    /// - Returns: Data?
    func _data(using encoding: String.Encoding = .utf8, isLossyConversion: Bool = false) -> Data? {
        let data = self.data(using: encoding, allowLossyConversion: isLossyConversion)
        return data
    }
}

// MARK: - Data (static function)
extension Data {
    
    /// Dictionary => JSON Data
    /// - ["name":"William"] => {"name":"William"} => 7b2268747470223a2022626f6479227d
    /// - Parameter parameters: 參數Dictionary
    /// - Returns: Data?
    static func _jsonSerialization(with parameters: [String: Any]) -> Data? {
        
        guard JSONSerialization.isValidJSONObject(parameters),
              let data = try? JSONSerialization.data(withJSONObject: parameters, options: JSONSerialization.WritingOptions())
        else {
            return nil
        }
        
        return data
    }
}

// MARK: - Data (class function)
extension Data {
    
    /// Data => 16進位文字
    /// - %02x - 推播Token常用
    /// - Returns: String
    func _hexString() -> String {
        let hexString = reduce("") { return $0 + String(format: "%02x", $1) }
        return hexString
    }
    
    /// Data => JSON
    /// - 7b2268747470223a2022626f6479227d => {"http": "body"}
    /// - Returns: Any?
    func _jsonObject(options: JSONSerialization.ReadingOptions = .allowFragments) -> Any? {
        let json = try? JSONSerialization.jsonObject(with: self, options: options)
        return json
    }
    
    /// Data => Class
    /// - Parameter type: 要轉型的Type => 符合Decodable
    /// - Returns: T => 泛型
    func _class<T: Decodable>(type: T.Type) -> T? {
        let modelClass = try? JSONDecoder().decode(type.self, from: self)
        return modelClass
    }
}

// MARK: - UIDevice (static function)
extension UIDevice {
    
    /// [判斷是iPhone / iPad / Apple TV](https://mini.nidbox.com/diary/read/9759417)
    /// - Returns: UIUserInterfaceIdiom
    static func _type() -> UIUserInterfaceIdiom { return UIDevice.current.userInterfaceIdiom }
}

// MARK: - URL (static function)
extension URL {
    
    /// 將URL標準化 (百分比)
    /// - 是在哈囉 => %E6%98%AF%E5%9C%A8%E5%93%88%E5%9B%89
    /// - Parameters:
    ///   - string: url字串
    ///   - characterSet: 字元的判斷方式
    /// - Returns: URL?
    static func _standardization(string: String, characterSet: CharacterSet = .urlQueryAllowed) -> URL? {
        
        var url: URL?
        url = URL(string: string)
        
        guard url == nil,
              let encodeString = string._encodingURL(characterSet: characterSet)
        else {
            return url
        }

        return URL(string: encodeString)
    }
}

// MARK: - URL (class function)
extension URL {
    
    /// 從AppStore網址開啟StoreProductViewController
    /// - Parameters:
    ///   - prefix: String
    ///   - pattern: String
    ///   - delegate: UIViewController & SKStoreProductViewControllerDelegate
    ///   - result: Result<Bool, Error>
    func _openAppStoreWithInside(prefix: String = "id", pattern: String = "id[\\d]{8,12}",  delegate: (UIViewController & SKStoreProductViewControllerDelegate), result: @escaping (Result<Bool, Error>) -> Void) {
        
        let calculateResult = Constant.RegularExpression.extracts(text: self.absoluteString, pattern: pattern).calculate()
        
        switch calculateResult {
        case .failure(let error): result(.failure(error))
        case .success(let list):
            
            guard let list = list,
                  let appKey = list.first
            else {
                return result(.success(false))
            }
            
            UIApplication.shared._presentStoreProductViewController(with: NSDecimalNumber(string: appKey[prefix.count...]), delegate: delegate) { _result in
                switch _result {
                case .failure(let error): result(.failure(error))
                case .success(let isSuccess): result(.success(isSuccess))
                }
            }
        }
    }
    
    /// 從AppStore網址開啟StoreProductViewController
    /// - Parameters:
    ///   - prefix: String
    ///   - pattern: String
    ///   - delegate: UIViewController & SKStoreProductViewControllerDelegate
    /// - Returns: Result<Bool, Error>
    func _openAppStoreWithInside(prefix: String = "id", pattern: String = "id[\\d]{8,12}", delegate: (UIViewController & SKStoreProductViewControllerDelegate)) async -> Result<Bool, Error> {
        
        let result = Constant.RegularExpression.extracts(text: self.absoluteString, pattern: pattern).calculate()
        
        switch result {
        case .failure(let error): return .failure(error)
        case .success(let list):
            
            guard let list = list,
                  let appKey = list.first
            else {
                return .success(false)
            }
            
            do {
                let isSuccess = try await UIApplication.shared._presentStoreProductViewController(with: NSDecimalNumber(string: appKey[prefix.count...]), delegate: delegate)
                return .success(isSuccess)
            } catch {
                return .failure(error)
            }
        }
    }
    
    /// 在APP內部開啟URL (SafariViewController) => window.webkit.messageHandlers.LinkUrl.postMessage("https://www.google.com")
    /// - Parameter urlString: URL網址
    func _openUrlWithInside(delegate: (UIViewController & SFSafariViewControllerDelegate)) -> SFSafariViewController {
        
        let safariViewController = SFSafariViewController(url: self)
        
        safariViewController.delegate = delegate
        safariViewController.modalTransitionStyle = .coverVertical
        safariViewController.modalPresentationStyle = .automatic
        
        delegate.present(safariViewController, animated: true)
        
        return safariViewController
    }
    
    /// 將URL標準化 (百分比)
    /// - 是在哈囉 => %E6%98%AF%E5%9C%A8%E5%93%88%E5%9B%89
    /// - Parameters:
    ///   - characterSet: 字元的判斷方式
    /// - Returns: URL?
    func _standardization(characterSet: CharacterSet = .urlQueryAllowed) -> URL? {
        return Self._standardization(string: self.absoluteString)
    }
}

// MARK: - Bundle (class function)
extension Bundle {
    
    /// 讀取info.plist的欄位資訊
    /// - CFBundleShortVersionString...
    /// - Parameter key: 要取的Key值
    /// - Returns: Any?
    func _infoDictionary(with key: String) -> Any? { return self.infoDictionary?[key] }
    
    /// 取得APP的ID
    /// - Returns: String?
    func _appBundleId() -> String? {
        guard let bundleId = self._infoDictionary(with: "CFBundleIdentifier") as? String else { return nil }
        return bundleId
    }
}

// MARK: - UIApplication (class function)
extension UIApplication {
    
    /// 重新開啟APP
    /// - 重新開啟APP的第一頁
    /// - MainStoryboard -> keyWindow -> rootViewController
    /// - Parameter storyboard: 要重開哪一個Storyboard？
    func _restartApp(from storyboard: UIStoryboard = UIStoryboard(name: "Main", bundle: Bundle.main)) {
        windows.first?.rootViewController = storyboard.instantiateInitialViewController()
    }
    
    /// 彈出AppStroe的產品ViewController
    /// - Parameters:
    ///   - identity: 該APP的identity
    ///   - delegate: SKStoreProductViewControllerDelegate
    ///   - result: Result<Bool, Error>
    func _presentStoreProductViewController(with identity: NSNumber, delegate: (UIViewController & SKStoreProductViewControllerDelegate), result: @escaping (Result<Bool, Error>) -> Void) {
        
        let storeProductViewController = SKStoreProductViewController._build(with: delegate)
        let parameters = SKStoreProductViewController._iTunesItemParametersMaker(with: identity)
        
        storeProductViewController.loadProduct(withParameters: parameters) { (isSuccess, error) in
            if let error = error { result(.failure(error)) }
            result(.success(isSuccess))
        }
        
        delegate.present(storeProductViewController, animated: true, completion: nil)
    }
    
    /// 彈出AppStroe的產品ViewController
    /// - Parameters:
    ///   - identity: 該APP的identity
    ///   - delegate: SKStoreProductViewControllerDelegate
    /// - Returns: Bool
    func _presentStoreProductViewController(with identity: NSNumber, delegate: (UIViewController & SKStoreProductViewControllerDelegate)) async throws -> Bool {
        
        let storeProductViewController = SKStoreProductViewController._build(with: delegate)
        let isSuccess = try await storeProductViewController._loadProduct(with: identity)
        
        delegate.present(storeProductViewController, animated: true, completion: nil)
        return isSuccess
    }
}

// MARK: - UNUserNotificationCenter (static function)
extension UNUserNotificationCenter {
    
    /// 註冊推播 => MainQueue
    /// - application(_:didRegisterForRemoteNotificationsWithDeviceToken:)
    static func _registerForRemoteNotifications() { UIApplication.shared.registerForRemoteNotifications() }
}

// MARK: - UNUserNotificationCenter (class function)
extension UNUserNotificationCenter {
    
    /// [推播權限測試](https://www.jianshu.com/p/1320e74e3a7e)
    /// - 使用在登入後才詢問推播
    /// - Parameters:
    ///   - delegate: UNUserNotificationCenterDelegate
    ///   - authorizationOptions: 允許推播通知的類型
    ///   - grantedHandler: 答應的時候要做什麼？
    ///   - rejectedHandler: 沒答應 / 不答應的時候要做什麼？
    ///   - result: 原來的狀態 => (UNAuthorizationStatus) -> Void
    func _userNotificationSetting(delegate: UNUserNotificationCenterDelegate, authorizationOptions: UNAuthorizationOptions = [.alert, .badge, .sound], grantedHandler: @escaping () -> Void, rejectedHandler: @escaping () -> Void, result: @escaping (UNAuthorizationStatus) -> Void) {

        self._pushNotificationAuthorization { (authorizationStatus) in
            
            switch (authorizationStatus) {
            case .notDetermined:
                self.requestAuthorization(options: authorizationOptions) { (isGranted, error) in
                    guard isGranted else { rejectedHandler(); return }
                    DispatchQueue.main.async { Self._registerForRemoteNotifications() }
                    grantedHandler()
                }
            case .authorized:
                DispatchQueue.main.async { Self._registerForRemoteNotifications() }
                self.delegate = delegate
            case .ephemeral:
                DispatchQueue.main.async { Self._registerForRemoteNotifications() }
                self.delegate = delegate
            case .denied: wwPrint("denied")
            case .provisional: wwPrint("provisional")
            @unknown default: fatalError()
            }
            
            result(authorizationStatus)
        }
    }
    
    /// 推播的授權狀態
    /// - Parameter result: (UNAuthorizationStatus) -> Void
    func _pushNotificationAuthorization(result: @escaping (UNAuthorizationStatus) -> Void) {

        self._pushNotificationSettings { (settings) in
            result(settings.authorizationStatus)
        }
    }
    
    /// 推播的授權狀態
    /// - Returns: UNAuthorizationStatus
    func _pushNotificationAuthorization() async -> UNAuthorizationStatus {
        
        await withCheckedContinuation { continuation in
            self._pushNotificationAuthorization { (settings) in
                continuation.resume(returning: settings)
            }
        }
    }
    
    /// [推播的設定狀態](https://www.jianshu.com/p/61dd9dd431a9)
    /// - Parameter result: (UNNotificationSettings) -> Void
    func _pushNotificationSettings(result: @escaping (UNNotificationSettings) -> Void) {
        self.getNotificationSettings { (settings) in
            result(settings)
        }
    }
    
    /// [推播的設定狀態](https://www.jianshu.com/p/61dd9dd431a9)
    /// - Returns: UNNotificationSettings
    func _pushNotificationSettings() async -> UNNotificationSettings {
        
        await withCheckedContinuation { continuation in
            self.getNotificationSettings { (settings) in
                continuation.resume(returning: settings)
            }
        }
    }
}

// MARK: - FileManager (class function)
extension FileManager {
    
    /// User的「暫存」資料夾
    /// - => ~/tmp
    /// - Returns: URL
    func _temporaryDirectory() -> URL { return temporaryDirectory }
    
    /// 寫入Data - 二進制資料
    /// - Parameters:
    ///   - url: 寫入Data的文件URL
    ///   - data: 要寫入的資料
    /// - Returns: Bool
    func _writeData(to url: URL?, data: Data?) -> Result<Bool, Error> {
        
        guard let url = url,
              let data = data
        else {
            return .success(false)
        }
        
        do {
            try data.write(to: url)
            return .success(true)
        } catch {
            return .failure(error)
        }
    }
    
    /// 寫入Data - 二進制資料
    /// - Parameters:
    ///   - url: 寫入Data的文件檔名
    ///   - data: 要寫入的資料
    /// - Returns: Result<URL?, Error>
    func _writeTemporaryData(filename: String?, data: Data?) -> Result<URL?, Error> {
        
        guard let filename = filename,
              let data = data
        else {
            return .success(nil)
        }
        
        let tempUrl = temporaryDirectory.appendingPathComponent(filename)
        
        do {
            try data.write(to: tempUrl)
            return .success(tempUrl)
        } catch {
            return .failure(error)
        }
    }
}

// MARK: - WKWebView (static function)
extension WKWebView {
    
    /// 產生WKWebView (WKNavigationDelegate & WKUIDelegate)
    /// - Parameters:
    ///   - delegate: WKNavigationDelegate & WKUIDelegate
    ///   - frame: WKWebView的大小
    ///   - canOpenWindows: [window.open(url)](https://www.jianshu.com/p/561307f8aa9e) for  [webView(_:createWebViewWith:for:windowFeatures:)](https://developer.apple.com/documentation/webkit/wkuidelegate/1536907-webview)
    ///   - configuration: WKWebViewConfiguration
    ///   - contentInsetAdjustmentBehavior: scrollView是否為全畫面
    /// - Returns: WKWebView
    static func _build(delegate: (WKNavigationDelegate & WKUIDelegate)?, frame: CGRect, canOpenWindows: Bool = false, configuration: WKWebViewConfiguration = WKWebViewConfiguration(), contentInsetAdjustmentBehavior: UIScrollView.ContentInsetAdjustmentBehavior = .never) -> WKWebView {
        
        let webView = WKWebView(frame: frame, configuration: configuration)
        configuration.preferences.javaScriptCanOpenWindowsAutomatically = canOpenWindows
        
        webView.backgroundColor = .white
        webView.navigationDelegate = delegate
        webView.uiDelegate = delegate
        webView.scrollView.contentInsetAdjustmentBehavior = contentInsetAdjustmentBehavior
        
        return webView
    }
    
    /// [在內部開啟網址與否的設定](https://qiita.com/vc7/items/90bd7da89aee54936bfd)
    /// - _blank，在外部開網頁，但內部的不要有反應 => .cancel
    /// - 比如說網址的副檔名為「.PDF」，開啟外部的Browser，其它就開在WKWebView上
    /// - [webView(_:decidePolicyFor:decisionHandler:)](https://developer.apple.com/documentation/webkit/wknavigationdelegate/1455641-webview)
    /// - Parameters:
    ///   - action: WKNavigationAction
    ///   - extensions: 要過濾的「副檔名」
    /// - Returns: WKNavigationActionPolicy
    static func _navigationActionPolicy(_ action: WKNavigationAction, extensions strings: [String]) -> WKNavigationActionPolicy {
        
        guard let url = action.request.url,
              let fileExtension = Optional.some(url.lastPathComponent),
              UIApplication.shared.canOpenURL(url)
        else {
            return .cancel
        }
                
        for string in strings {
            if (fileExtension.uppercased() != string.uppercased()) { continue }
            return .cancel
        }
        
        return .allow
    }
    
    /// [在外面收到的網址開啟與否的設定](https://www.itread01.com/content/1545674242.html)
    /// - [比如說MIME為「application/pdf」不允許在內部開啟](https://www.runoob.com/http/http-content-type.html)
    /// - [webView(_:decidePolicyFor:decisionHandler:)](https://developer.apple.com/documentation/webkit/wknavigationdelegate/1455643-webview)
    /// - Parameters:
    ///   - response: WKNavigationResponse
    ///   - mimeTypes: 要過濾的「MIME Type」
    /// - Returns: WKNavigationActionPolicy
    static func _navigationResponsePolicy(_ response: WKNavigationResponse, mimeTypes: [Constant.MimeType]) -> WKNavigationResponsePolicy {
        
        guard let mimeType = response.response.mimeType else { return .cancel }
        
        for type in mimeTypes {
            if (type.mimeType != mimeType) { continue }
            return .cancel
        }
        
        return .allow
    }
}

// MARK: - WKWebView (class function)
extension WKWebView {
    
    /// 載入WebView網址
    /// - Parameters:
    ///   - urlString: URL網址
    ///   - cachePolicy: 快取規則
    ///   - timeoutInterval: 過時時效
    /// - Returns: WKNavigation?
    func _load(urlString: String?, cachePolicy: URLRequest.CachePolicy = .reloadIgnoringCacheData, timeoutInterval: TimeInterval = 60.0) -> WKNavigation? {
        
        guard let urlString = urlString,
              let url = URL(string: urlString),
              let urlRequest = Optional.some(URLRequest(url: url, cachePolicy: cachePolicy, timeoutInterval: timeoutInterval))
        else {
            return nil
        }
        
        return self.load(urlRequest)
    }
    
    /// [加上網頁端要傳訊息給APP的名稱](https://ithelp.ithome.com.tw/articles/10207572)
    /// - userContentController(_:didReceive:)
    /// - window.webkit.messageHandlers.<key>.postMessage("Send message to Swift")
    /// - Parameters:
    ///   - delegate: WKScriptMessageHandler
    ///   - keys: 傳訊息給APP的Key => "ToApp"
    func _addScriptMessageKeys(delegate: WKScriptMessageHandler, keys: [Constant.ScriptMessageKey]?) -> Bool {
        
        guard let keys = keys, !keys.isEmpty, let keySet = Optional.some(Set(keys)) else { return false }
        
        keySet.forEach { (key) in self.configuration.userContentController.add(delegate, name: key.rawValue) }
        return true
    }
    
    /// 顯示提示視窗 => window.webkit.messageHandlers.Alert.postMessage("iOS好棒棒")
    /// => [webView(_:runJavaScriptAlertPanelWithMessage:initiatedByFrame:completionHandler:)](https://developer.apple.com/documentation/webkit/wkuidelegate/1537406-webview)
    /// - Parameters:
    ///   - title: [標題](https://www.jianshu.com/p/e4c274248a78)
    ///   - message: [內容](https://www.w3schools.com/jsref/tryit.asp?filename=tryjsref_alert)
    ///   - completionHandler: 彈出完成後所要做的事
    func _showAlertController(target: UIViewController, title: String?, message: String?, completionHandler: @escaping () -> Void) {
        
        let alertController = UIAlertController(title: title, message: message, preferredStyle: .alert)
        let action = UIAlertAction(title: "確認", style: .default) { (_) in completionHandler() }
        
        alertController.addAction(action)
        
        target.present(alertController, animated: true, completion: nil)
    }
    
    /// [顯示確認視窗](https://www.w3schools.com/jsref/tryit.asp?filename=tryjsref_confirm)
    /// => [webView(_:runJavaScriptConfirmPanelWithMessage:initiatedByFrame:completionHandler:)https://developer.apple.com/documentation/webkit/wkuidelegate/1536489-webview]()
    /// - Parameters:
    ///   - title: 標題
    ///   - message: [內容](https://www.w3schools.com/jsref/tryit.asp?filename=tryjsref_confirm)
    ///   - completionHandler: 彈出完成後所要做的事 (true / false)
    func _showComfirmAlertController(target: UIViewController, title: String?, message: String?, completionHandler: @escaping (Bool) -> Void) {
        
        let alertController = UIAlertController(title: title, message: message, preferredStyle: .alert)
        
        let comfirmAction = UIAlertAction(title: "確認", style: .default) { (_) in completionHandler(true) }
        let cancelAction = UIAlertAction(title: "取消", style: .cancel) { (_) in completionHandler(false) }
        
        alertController.addAction(comfirmAction)
        alertController.addAction(cancelAction)
        
        target.present(alertController, animated: true, completion: nil)
    }
    
    /// [顯示文字框](https://www.w3schools.com/jsref/tryit.asp?filename=tryjsref_confirm)
    /// => [webView(_:runJavaScriptTextInputPanelWithPrompt:defaultText:initiatedByFrame:completionHandler:)](https://developer.apple.com/documentation/webkit/wkuidelegate/1538086-webview)
    /// - Parameters:
    ///   - prompt: [提示](https://www.w3schools.com/jsref/tryit.asp?filename=tryjsref_prompt)
    ///   - defaultText: 預設文字
    ///   - completionHandler: 彈出完成後所要做的事 (String)
    func _showPromptAlertController(target: UIViewController, title: String?, prompt: String?, defaultText: String?, completionHandler: @escaping (String?) -> Void) {
        
        let alertController = UIAlertController(title: title, message: prompt, preferredStyle: .alert)
        alertController.addTextField { (textField) in textField.text = defaultText }
        
        let action = UIAlertAction(title: "完成", style: .default) { (_) in completionHandler(alertController.textFields?.first?.text) }
        alertController.addAction(action)
        
        target.present(alertController, animated: true, completion: nil)
    }
    
    /// 執行JavaScript
    /// - Parameters:
    ///   - script: JavaScript程式碼
    ///   - result: Result<Any?, Error>
    func _evaluateJavaScript(script: String?, result: @escaping (Result<Any?, Error>) -> Void) {
        
        guard let script = script else { result(.success(nil)); return }
        
        DispatchQueue.main.async {
            self.evaluateJavaScript(script) { (object, error) in
                if let error = error { result(.failure(error)); return }
                result(.success(object))
            }
        }
    }
    
    /// [執行JavaScript - async](https://www.hackingwithswift.com/swift/5.5/continuations)
    /// - Parameter script: JavaScript程式碼
    /// - Returns: Result<Any?, Error>
    func _evaluateJavaScript(script: String?) async -> Result<Any?, Error> {
        
        await withCheckedContinuation { continuation in
            _evaluateJavaScript(script: script) { result in
                continuation.resume(returning: result)
            }
        }
    }
}

// MARK: - WKNavigationAction (class function)
extension WKNavigationAction {
    
    /// [測試link是不是target="_blank" => 開新網頁](https://www.w3school.com.cn/html5/att_a_target.asp)
    func _isBlankLink() -> Bool { return self.targetFrame == nil }
}

// MARK: - UTTypeReference (static function)
extension UTTypeReference {
    
    /// [副檔名 (pdf) => 統一類型標識符 (com.adobe.pdf)](https://www.jianshu.com/p/d6fe1e7af9b6)
    /// - Parameter filenameExtension: 副檔名
    /// - Returns: UTTypeReference?
    static func _find(with filenameExtension: String) -> UTTypeReference? { return UTTypeReference(filenameExtension: filenameExtension) }
}

// MARK: - UIActivityViewController (static function)
extension UIActivityViewController {
    
    /// [產生UIActivityViewController分享功能](https://jjeremy-xue.medium.com/swift-玩玩-uiactivityviewcontroller-5995bb80ff68 )
    /// - Parameters:
    ///   - activityItems: [Any]
    ///   - applicationActivities: [UIActivity]?
    ///   - tintColor: tintColor
    ///   - sourceView: 要貼在哪個View上 (for iPad)
    /// - Returns: UIActivityViewController
    static func _build(activityItems: [Any], applicationActivities: [UIActivity]? = nil, tintColor: UIColor = .white, sourceView: UIView? = nil) -> UIActivityViewController {
        
        let activityViewController = UIActivityViewController(activityItems: activityItems, applicationActivities: applicationActivities)
        
        activityViewController.view.tintColor = tintColor
        activityViewController.popoverPresentationController?.sourceView = sourceView
        
        return activityViewController
    }
}

// MARK: - UIActivityViewController (class function)
extension UIActivityViewController {
    
    /// 完成後的結果
    /// - Returns: Result<Bool, Error>
    func _completion() async -> Result<Bool, Error> {
        
        await withCheckedContinuation { continuation in
            self.completionWithItemsHandler = { (activityType: UIActivity.ActivityType?, completed: Bool, returnedItems: [Any]?, error: Error?) in
                if let error = error { continuation.resume(returning: .failure(error)) }
                return continuation.resume(returning: .success(completed))
            }
        }
    }
}

// MARK: - SKStoreProductViewController (static function)
extension SKStoreProductViewController {
    
    /// [建立內彈的AppStore](https://juejin.cn/post/6844903912248623111)
    /// - Parameter delegate: UIViewController & SKStoreProductViewControllerDelegate
    /// - Returns: SKStoreProductViewController
    static func _build(with delegate: (UIViewController & SKStoreProductViewControllerDelegate)) -> SKStoreProductViewController {
        
        let storeProductViewController = SKStoreProductViewController()
        storeProductViewController.delegate = delegate
        
        return storeProductViewController
    }
    
    /// iTunes的identifier形式
    /// - Parameter identifier: NSNumber
    /// - Returns: [String: NSNumber]
    static func _iTunesItemParametersMaker(with identifier: NSNumber) -> [String: NSNumber] {
        return [SKStoreProductParameterITunesItemIdentifier: identifier]
    }
}

// MARK: - SKStoreProductViewController (class function)
extension SKStoreProductViewController {
    
    /// 載入AppStore
    /// - Parameter identity: NSNumber
    /// - Returns: Bool
    func _loadProduct(with identity: NSNumber) async throws -> Bool {
        let parameters = Self._iTunesItemParametersMaker(with: identity)
        let isSuccess = try await self.loadProduct(withParameters: parameters)
        return isSuccess
    }
}
