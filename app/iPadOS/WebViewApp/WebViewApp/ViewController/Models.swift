//
//  Models.swift
//  WebViewApp
//
//  Created by iOS on 2022/10/13.
//

import Foundation

class DownloadResult: Codable {
    
    let data: String?
    let name: String?
    let type: String?
    
    func base64Data() -> Data? {
        
        guard let data = data,
              let type = type,
              let prefixString = Optional.some("data:\(type);base64,")
        else {
            return nil
        }
        
        return Data(base64Encoded: data[prefixString.count...])
    }
}
