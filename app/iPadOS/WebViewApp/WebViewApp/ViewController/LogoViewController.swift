//
//  LogoViewController.swift
//  WebViewApp
//
//  Created by iOS on 2022/10/12.
//

import UIKit
import WWPrint
import WWNetworking

final class LogoViewController: UIViewController {
    
    static let urlString = "http://192.168.1.102:8080/"
    
    override func viewDidLoad() {
        super.viewDidLoad()
    }
    
    override func viewDidAppear(_ animated: Bool) {
        super.viewDidAppear(animated)
        checkNextPage()
    }
    
    override func prepare(for segue: UIStoryboardSegue, sender: Any?) {
        
        guard let viewController = segue.destination as? HtmlViewController else { return }
        viewController.urlString = Self.urlString
    }
    
    deinit { wwPrint("\(Self.self) deinit.") }
}

// MARK: - 小工具
extension LogoViewController {
    
    /// 利用Segue切換UIViewController
    func gotoNextPageWithSegue(_ segue: String) {
        self.performSegue(withIdentifier: segue, sender: nil)
    }
    
    /// 測試網頁還活著才進入WebView
    func checkNextPage() {
        
        WWNetworking.shared.header(urlString: Self.urlString) { result in

            switch result {
            case .failure(let error): wwPrint(error)
            case .success(let info):

                guard let statusCode = info.response?.statusCode,
                      statusCode == 200
                else {
                    return
                }

                DispatchQueue.main.asyncAfter(deadline: .now() + 2) { self.gotoNextPageWithSegue(HtmlViewController.segue) }
            }
        }
    }
}
