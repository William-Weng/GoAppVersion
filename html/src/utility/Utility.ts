// [blobToBase64](https://stackoverflow.com/questions/18650168/convert-blob-to-base64)
function blobToBase64(blob: Blob): Promise<string | ArrayBuffer | null> {

    return new Promise((resolve, _) => {
        const reader = new FileReader();
        reader.onloadend = () => resolve(reader.result);
        reader.readAsDataURL(blob);
    });
}

// 傳送給APP的自定義協定 => 讓APP有反應 (.src = app://downloadFile)
function mobileProtocol(src: string) {
    const iframe = document.createElement('iframe')
    iframe.src = src
    iframe.style.display = "none"
    document.body.appendChild(iframe)
    iframe.remove()
}

// [強制下載檔案 => Convert the Byte Data to BLOB object.](https://www.aspsnippets.com/Articles/Automatically-download-PDF-File-using-JavaScript.aspx)
function forceDownloadFile(blob: Blob, fullName: string) {
        
    const url = window.URL || window.webkitURL
    const link = url.createObjectURL(blob)
    const a = document.createElement("a")
        
    a.setAttribute("download", fullName)
    a.setAttribute("href", link)
        
    document.body.appendChild(a)
    a.click()
    document.body.removeChild(a)
}

class Utility {
    
    readonly blobToBase64: (blob: Blob) => Promise<string | ArrayBuffer | null>
    readonly mobileProtocol: (src: string) => void
    readonly forceDownloadFile: (blob: Blob, fullName: string) => void

    constructor() {
        this.blobToBase64 = blobToBase64
        this.mobileProtocol = mobileProtocol
        this.forceDownloadFile = forceDownloadFile
    }
}

export default new Utility()