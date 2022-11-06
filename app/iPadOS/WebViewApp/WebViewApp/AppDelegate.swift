//
//  AppDelegate.swift
//  WebViewApp
//
//  Created by iOS on 2022/10/12.
//

import UIKit
import WWPrint

@main
final class AppDelegate: UIResponder, UIApplicationDelegate {

    var window: UIWindow?
    var pushToken: String?
    
    private lazy var guardImageView: UIImageView = guardImageViewMaker(image: #imageLiteral(resourceName: "Logo"))
    
    func application(_ application: UIApplication, didFinishLaunchingWithOptions launchOptions: [UIApplication.LaunchOptionsKey: Any]?) -> Bool {
        
        UNUserNotificationCenter.current()._userNotificationSetting(delegate: self) {
            UIApplication.shared._restartApp()
        } rejectedHandler: {
            UIApplication.shared._restartApp()
        } result: { status in
            wwPrint(status)
        }
        
        return true
    }
    
    func application(_ application: UIApplication, didRegisterForRemoteNotificationsWithDeviceToken deviceToken: Data) {
        pushToken = deviceToken._hexString()
        pushTokenTest(pushToken ?? "")
    }
    
    func applicationWillResignActive(_ application: UIApplication) {
        addGuardImageView()
    }
    
    func applicationDidBecomeActive(_ application: UIApplication) {
        removeGuardImageView()
    }
}

// MARK: - UNUserNotificationCenterDelegate
extension AppDelegate: UNUserNotificationCenterDelegate {

    func userNotificationCenter(_ center: UNUserNotificationCenter, willPresent notification: UNNotification, withCompletionHandler completionHandler: @escaping (UNNotificationPresentationOptions) -> Void) {
        
        let options : UNNotificationPresentationOptions = [.badge, .sound, .list, .banner]
        completionHandler([options])
    }
    
    func userNotificationCenter(_ center: UNUserNotificationCenter, didReceive response: UNNotificationResponse, withCompletionHandler completionHandler: @escaping () -> Void) {
        completionHandler()
    }
}

// MARK: - 小工具
private extension AppDelegate {
        
    /// 加上防偷窺的ImageView
    func addGuardImageView() { window?.addSubview(guardImageView) }
    
    /// 移除防偷窺的ImageView
    func removeGuardImageView() { guardImageView.removeFromSuperview() }
    
    /// 保護APP的ImageView
    /// - Parameter image: Logo圖示
    /// - Returns: UIImageView
    func guardImageViewMaker(image: UIImage) -> UIImageView {
        
        let imageView = UIImageView(image: image)
        imageView.contentMode = .scaleAspectFill
        imageView.frame = window?.frame ?? .zero
        
        return imageView
    }
    
    /// 推播測試 => 會震動 + 複製Token
    /// - Parameter pushToken: String
    func pushTokenTest(_ pushToken: String) {
        
        #if DEBUG
        
        let feedBackGenertor = UIImpactFeedbackGenerator(style: .heavy)
        let pasteboard = UIPasteboard.general
          
        pasteboard.string = pushToken
        feedBackGenertor.impactOccurred()
        
        #endif
    }
}
