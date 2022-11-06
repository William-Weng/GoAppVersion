//
//  HtmlController.swift
//  WebViewApp
//
//  Created by iOS on 2022/10/12.
//

import UIKit
import WebKit
import SafariServices
import WWPrint
import WWNetworking
import StoreKit

final class HtmlViewController: UIViewController {
    
    static let segue: String = "HtmlViewSegue"
    static let appVersionAPI = "http://192.168.1.102:12345/appVersion/mobile"
    
    var urlString: String?
    var positionButton: UIButton!
    var isRegistratedPushToken = false
    
    private var myWebView: WKWebView!
    private var safariViewController: SFSafariViewController?

    override func viewDidLoad() {
        super.viewDidLoad()
        initSetting()
    }
    
    deinit { wwPrint("\(Self.self) deinit.") }
}

// MARK: - WKUIDelegate
extension HtmlViewController: WKUIDelegate {
    
    func webView(_ webView: WKWebView, runJavaScriptAlertPanelWithMessage message: String, initiatedByFrame frame: WKFrameInfo, completionHandler: @escaping () -> Void) {
        webView._showAlertController(target: self, title: nil, message: message, completionHandler: completionHandler)
    }
    
    func webView(_ webView: WKWebView, runJavaScriptConfirmPanelWithMessage message: String, initiatedByFrame frame: WKFrameInfo, completionHandler: @escaping (Bool) -> Void) {
        webView._showComfirmAlertController(target: self, title: nil, message: message, completionHandler: completionHandler)
    }
    
    func webView(_ webView: WKWebView, runJavaScriptTextInputPanelWithPrompt prompt: String, defaultText: String?, initiatedByFrame frame: WKFrameInfo, completionHandler: @escaping (String?) -> Void) {
        webView._showPromptAlertController(target: self, title: nil, prompt: prompt, defaultText: defaultText, completionHandler: completionHandler)
    }
}

// MARK: - WKNavigationDelegate
extension HtmlViewController: WKNavigationDelegate {
    
    func webView(_ webView: WKWebView, didStartProvisionalNavigation navigation: WKNavigation) {}
    
    func webView(_ webView: WKWebView, didCommit navigation: WKNavigation!) {
        Task { let result = await initMobileInformation(webView); wwPrint(result) }
    }
    
    func webView(_ webView: WKWebView, didFinish navigation: WKNavigation) { wwPrint(navigation) }
    func webView(_ webView: WKWebView, didFailProvisionalNavigation navigation: WKNavigation!, withError error: Error) { wwPrint(error) }
    func webView(_ webView: WKWebView, didFail navigation: WKNavigation!, withError error: Error) { wwPrint(navigation) }
    
    func webView(_ webView: WKWebView, decidePolicyFor navigationAction: WKNavigationAction, decisionHandler: @escaping (WKNavigationActionPolicy) -> Void) {
        let actionPolicy = navigationActionPolicyRule(webView, decidePolicyFor: navigationAction)
        decisionHandler(actionPolicy)
    }
    
    func webView(_ webView: WKWebView, decidePolicyFor navigationResponse: WKNavigationResponse, decisionHandler: @escaping ((WKNavigationResponsePolicy) -> Void)) {
        let responsePolicy = WKWebView._navigationResponsePolicy(navigationResponse, mimeTypes: [.xlsx])
        decisionHandler(responsePolicy)
    }
}

// MARK: - WKScriptMessageHandler
extension HtmlViewController: WKScriptMessageHandler {
    
    func userContentController(_ userContentController: WKUserContentController, didReceive message: WKScriptMessage) {
        
        guard let messageKey = Constant.ScriptMessageKey(rawValue: message.name) else { return }
        
        switch messageKey {
        case .alert: myWebView._showAlertController(target: self, title: message.name, message: "\(message.body)") {}
        case .restart: UIApplication.shared._restartApp()
        case .update: updateVersionAction(userContentController, didReceive: message)
        }
    }
}

// MARK: - SFSafariViewControllerDelegate
extension HtmlViewController: SFSafariViewControllerDelegate {
    func safariViewController(_ controller: SFSafariViewController, didCompleteInitialLoad didLoadSuccessfully: Bool) {}
    func safariViewController(_ controller: SFSafariViewController, initialLoadDidRedirectTo URL: URL) {}
}

// MARK: - SKStoreProductViewControllerDelegate
extension HtmlViewController: SKStoreProductViewControllerDelegate {}

// MARK: - 小工具
private extension HtmlViewController {
    
    /// WebView切始化設定
    func initSetting() {
        
        myWebView = WKWebView._build(delegate: self, frame: .zero, canOpenWindows: true, contentInsetAdjustmentBehavior: .always)
        _ = myWebView._load(urlString: urlString)
        _ = myWebView._addScriptMessageKeys(delegate: self, keys: [.alert, .restart, .update])
        
        self.view = myWebView
        
        positionButton = UIButton(frame: .zero)
        positionButton.backgroundColor = .red
        myWebView.addSubview(positionButton)
    }
    
    /// 網頁網址的放行規則
    /// - Parameters:
    ///   - webView: WKWebView
    ///   - navigationAction: WKNavigationAction
    /// - Returns: WKNavigationActionPolicy
    func navigationActionPolicyRule(_ webView: WKWebView, decidePolicyFor navigationAction: WKNavigationAction) -> WKNavigationActionPolicy {
        
        guard let url = navigationAction.request.url,
              let components = URLComponents(string: url.absoluteString),
              let scheme = components.scheme,
              let host = components.host
        else {
            return .cancel
        }
        
        switch (scheme, host) {
        case ("app", "home"): Task { let result = await pushTokenRegistry(with: webView); wwPrint(result) }; return .cancel
        case ("app", "downloadFile"): Task { let result = await storeWebFile(with: webView); wwPrint(result) }; return .cancel
        case ("https", "apps.apple.com"): Task { let result = await url._openAppStoreWithInside(delegate: self); wwPrint(result) }; return .cancel
        case ("https", "play.google.com"): safariViewController = url._openUrlWithInside(delegate: self); return .cancel
        default: break
        }
                
        if (navigationAction._isBlankLink()) { UIApplication.shared.open(url); return .cancel }
        return .allow
    }
    
    /// 測試版本API with WebKit
    /// => window.webkit.messageHandlers.Update.postMessage('{"store":"dev","id":"id443904275"}')
    /// - Parameters:
    ///   - userContentController: WKUserContentController
    ///   - message: WKScriptMessage
    func updateVersionAction(_ userContentController: WKUserContentController, didReceive message: WKScriptMessage) {
        
        guard let body = message.body as? String,
              let dictionary = body._dictionary(encoding: .utf8),
              let store = dictionary["store"] as? String,
              let appId = dictionary["id"] as? String
        else {
            return
        }
        
        let urlString = "\(HtmlViewController.appVersionAPI)/\(appId)"
        
        Task {
            let result = await getRequest(urlString: urlString, paramaters: nil)
            
            switch result {
            case .failure(let error): wwPrint(error)
            case .success(let info): parseResponseInformation(info, store: store)
            }
        }
    }
    
    /// 解析取得的APP版本資訊
    /// - Parameters:
    ///   - info: WWNetworking.ResponseInformation
    ///   - store: String
    func parseResponseInformation(_ info: WWNetworking.ResponseInformation, store: String) {
        
        guard let data = info.data,
              let json = data._jsonObject() as? [String: Any],
              let _result = json["result"] as? [String: Any],
              let _version = _result["version"] as? [String: Any],
              let _storeVersionInfo = _version[store] as? [String: Any]
        else {
            return
        }
        
        var title = ""
        var version = ""
        
        defer {
            self.myWebView._showAlertController(target: self, title: title, message: "現在的測試版本是：\(version)") {}
        }
        
        for _version in _storeVersionInfo.keys {
            
            version = _version
            
            if let updateType = _storeVersionInfo[_version] as? Int {
                switch updateType {
                case 0: title = "一般更新"
                case 1: title = "總是更新"
                case 2: title = "強制更新"
                default: break
                }
            }
        }
    }
}

// MARK: 小工具 (非同步)
private extension HtmlViewController {
    
    /// 送出手機版本資訊 => 網頁
    func initMobileInformation(_ webView: WKWebView) async -> Result<Any?, Error> {
        
        let script = """
        window.mobileApp = {"os":"\(UIDevice.current.systemName)","version":"\(UIDevice.current.systemVersion)","model":"\(UIDevice.current.model)"}
        """
        
        let result = await webView._evaluateJavaScript(script: script)
        return result
    }
    
    /// 向網頁註冊推播Token
    /// - Parameter webView: WKWebView
    func pushTokenRegistry(with webView: WKWebView) async -> Result<Any?, Error> {
        
        if (isRegistratedPushToken) { return .success(nil) }
        
        guard let appId = Bundle.main._appBundleId(),
              let appDelegate = UIApplication.shared.delegate as? AppDelegate,
              let token = appDelegate.pushToken
        else {
            return .success(nil)
        }
        
        let script = """
        window.pushTokenRegistry({"appId":"\(appId)", "token":"\(token)"})
        """
        let result = await webView._evaluateJavaScript(script: script)
        isRegistratedPushToken = true

        return result
    }
    
    /// 儲存網路檔案
    /// - Parameter webView: WKWebView
    func storeWebFile(with webView: WKWebView) async -> Result<Bool, Error> {
        
        let script = "window.result"
        let result = await downloadFile(webView: webView, script: script)
        
        switch result {
        case .failure(let error): return .failure(error)
        case .success(let downloadResult):
            
            guard let downloadResult = downloadResult,
                  let base64Data = downloadResult.base64Data(),
                  let filename = downloadResult.name
            else {
                return .success(false)
            }
            
            let result = await base64DataToFile(with: base64Data, filename: filename)
            return result
        }
    }
    
    /// 下載網路檔案
    /// - Parameters:
    ///   - webView: WKWebView
    ///   - script: 要執行的Key (變數) => downloadFile
    /// - Returns: Result<DownloadResult?, Error>
    func downloadFile(webView: WKWebView, script: String) async -> Result<DownloadResult?, Error> {
        
        let result = await webView._evaluateJavaScript(script: script)
        
        switch result {
        case .failure(let error): return .failure(error)
        case .success(let json):
            
            guard let json = json as? [String: Any],
                  let data = Data._jsonSerialization(with: json),
                  let downloadResult = data._class(type: DownloadResult.self)
            else {
                return .success(nil)
            }
            
            return .success(downloadResult)
        }
    }
    
    /// 下載檔案 (二進制檔) - 使用UIActivityViewController
    /// - Parameters:
    ///   - data: 二進制檔案
    ///   - filename: 檔案名稱
    /// - Returns: Result<Bool, Error>
    func base64DataToFile(with data: Data, filename: String) async -> Result<Bool, Error> {
        
        let _result = FileManager.default._writeTemporaryData(filename: filename, data: data)
        
        switch _result {
        case .failure(let error): return .failure(error)
        case .success(let url):
            
            guard let tempUrl = url else { return .success(false) }
            
            let activityViewController = UIActivityViewController._build(activityItems: [tempUrl], sourceView: positionButton)
            
            self.positionButton.frame = CGRect(origin: CGPoint(x: view.center.x - 388 * 0.5, y: view.frame.midY), size: .zero)
            self.present(activityViewController, animated: true) {}
            
            let result = await activityViewController._completion()
            return result
        }
    }
    
    /// WWNetworking包裝成 => async
    /// - Parameters:
    ///   - urlString: String
    ///   - paramaters: [String : String?]?
    /// - Returns: Result<WWNetworking.ResponseInformation, Error>
    func getRequest(urlString: String, paramaters: [String : String?]? = nil) async -> Result<WWNetworking.ResponseInformation, Error> {
        
        await withCheckedContinuation { continuation in
            WWNetworking.shared.request(with: .GET, urlString: urlString, paramaters: paramaters) { result in
                continuation.resume(returning: result)
            }
        }
    }
}
