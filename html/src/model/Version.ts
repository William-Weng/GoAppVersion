// Version.ts
interface Version {
    id: string,
    name: string,
    icon: string,
    url: string,
    auto: boolean,
    type: number,
    lang?: string,
    version: string,
    versionProd: string,
    versionDev: string,
    updateVersion: number,
    updateVersionProd: number,
    updateVersionDev: number,
    pushSetting: string,
    pushSettingProduction: string,
    pushSettingDevelop: string,
}

export default Version