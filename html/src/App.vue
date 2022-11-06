<!-- App.vue -->
<template>
  <div class="app" v-loading.fullscreen.lock="isLoading" element-loading-background="rgba(0, 0, 0, 0.7)">

    <h1 align="center">APP管理小工具</h1>

    <!-- 顯示列表 -->
    <el-table :data="versions" :height="tableStyle.height">
      <el-table-column :label="versionLabel.icon">
        <template #default="scope">
          <el-image :src="scope.row.icon" @click="utility.displayDialog(ButtonType.PushSetting, true, scope.row)"/>
        </template>
      </el-table-column>
      <el-table-column :label="versionLabel.type">
        <template #default="scope">
          <a :href="scope.row.url" rel="noreferrer noopenner" target="_blank"><el-image :src='utility.appIcon(scope.row.type)' /></a>
        </template>
      </el-table-column>
      <el-table-column prop="name" :label="versionLabel.name" />
      <el-table-column prop="id" :label="versionLabel.id" />
      <el-table-column :label="versionLabel.version" align="center">
        <template #default="scope">
          <div class="table-action-button"><el-tag :type="utility.updateVersionType(scope.row.updateVersion).type" effect="dark" @click="utility.displayDialog(ButtonType.UpdateVersion, true, scope.row)">{{ utility.updateVersionType(scope.row.updateVersion).label }}</el-tag></div>
          <div class="table-action-button">{{ scope.row.version }}</div>
        </template>
      </el-table-column>
      <el-table-column :label="versionLabel.versionProd" align="center">
        <template #default="scope">
          <div class="table-action-button"><el-tag :type="utility.updateVersionType(scope.row.updateVersionProd).type" effect="dark" @click="utility.displayDialog(ButtonType.UpdateVersionProd, true, scope.row)">{{ utility.updateVersionType(scope.row.updateVersionProd).label }}</el-tag></div>
          <div class="table-action-button">{{ scope.row.versionProd }}</div>
        </template>
      </el-table-column>
      <el-table-column :label="versionLabel.versionDev" align="center">
        <template #default="scope">
          <div class="table-action-button"><el-tag :type="utility.updateVersionType(scope.row.updateVersionDev).type" effect="dark" @click="utility.displayDialog(ButtonType.UpdateVersionDev, true, scope.row)">{{ utility.updateVersionType(scope.row.updateVersionDev).label }}</el-tag></div>
          <div class="table-action-button">{{ scope.row.versionDev }}</div>
        </template>
      </el-table-column>
      <el-table-column prop="supportedDevices" :label="versionLabel.supportedDevices" :formatter="utility.supportedDevicesText"/>
      <el-table-column prop="lang" :label="versionLabel.lang" :formatter="utility.langText"/>
      <el-table-column :label="versionLabel.auto">
        <template #default="scope">
          <el-image :src='utility.autoUpdateIcon(scope.row.auto)' />
        </template>
      </el-table-column>
      <el-table-column fixed="right" width="100">
        <template #default="scope">
          <div class="table-action-button"><el-button :color=versionLabel.fixButton.color round @click="utility.displayDialog(ButtonType.Edit, true, scope.row)">{{ versionLabel.fixButton.text }}</el-button></div>
          <div class="table-action-button"><el-button :color=versionLabel.deleteButton.color round @click="utility.displayDialog(ButtonType.Delete, true, scope.row)">{{ versionLabel.deleteButton.text }}</el-button></div>
        </template>
      </el-table-column>
    </el-table>

    <!-- 使用者名稱 - 等級 -->
    <h1 class="account" v-if="!isForceSinglePage">{{ versionLabel.account }}</h1>

    <!-- 功能按鍵 -->
    <el-button-group v-if="!isForceSinglePage" size="large" class="action-button">
      <el-button type="danger" @click="utility.displayDialog(versionLabel.loginButton.type, true)">{{ versionLabel.loginButton.text }}</el-button>
      <el-button type="primary" :disabled="isDisabled" @click="utility.displayDialog(ButtonType.Append, true)">{{ versionLabel.appendButton.text }}</el-button>
      <el-button type="success" :disabled="isDisabled" @click="utility.refreshData()">{{ versionLabel.refreshButton.text }}</el-button>
      <el-button type="warning" :disabled="isDisabled" @click="utility.downloadExcel()">{{ versionLabel.downloadButton.text }}</el-button>
      <el-button type="primary" :disabled="isDisabled" @click="utility.displayDialog(ButtonType.Registry, true)">{{ versionLabel.registryButton.text }}</el-button>
      <el-upload class="inline-block" :http-request="utility.uploadExcel" :accept="versionLabel.uploadButton.accept"><el-button type="danger" :disabled="isDisabled">{{ versionLabel.uploadButton.text }}</el-button></el-upload>
    </el-button-group>
    
    <!-- 分頁列 -->
    <el-pagination v-if="!isForceSinglePage" class="action-page" :current-page="paginationSetting.currentPage" :hide-on-single-page="paginationSetting.isHideOnSinglePage" background :layout="paginationSetting.layout" :total="paginationSetting.count" :page-size="paginationSetting.pageSize" @current-change="utility.paginationAction"/>
    
    <!-- 強制下載選擇開關 -->
    <span class="action-download"><el-switch v-model="isForceDownload" autocomplete="off" @change="utility.switchAction"/> 強制下載檔案</span>

    <!-- 強制變成單頁顯示 -->
    <span class="action-single-page"><el-switch v-model="isForceSinglePage" autocomplete="off" @change="utility.switchSinglePageAction"/> 強制單頁顯示</span>

    <!-- 刪除確認框 -->
    <el-dialog v-model="confirmDialog.isVisible" :title="confirmDialog.title" :background="confirmDialog.background" @keyup.enter="utility.deleteVersion(utility.tempVersion.value)">
      <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="utility.displayDialog(ButtonType.Cancel, false)">{{ confirmDialog.cancel }}</el-button>
          <el-button type="danger" @click="utility.deleteVersion(utility.tempVersion.value)">{{ confirmDialog.sure }}</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 編譯確認框 -->
    <el-dialog v-model="editorDialog.isVisible" :title="editorDialog.title" :background="editorDialog.background" @keyup.enter="utility.editVersion(utility.tempVersion.value)">
      <el-form-item :label="versionLabel.auto"><el-switch v-model="utility.tempEditorVersion.value.auto" autocomplete="off" /></el-form-item>  
      <el-form-item :label="versionLabel.name"><el-input v-model="utility.tempEditorVersion.value.name" autocomplete="off" /></el-form-item>  
      <el-form-item :label="versionLabel.version"><el-input v-model="utility.tempEditorVersion.value.version.store" autocomplete="off" /></el-form-item>  
      <el-form-item :label="versionLabel.versionProd"><el-input v-model="utility.tempEditorVersion.value.version.prod" autocomplete="off" /></el-form-item>  
      <el-form-item :label="versionLabel.versionDev"><el-input v-model="utility.tempEditorVersion.value.version.dev" autocomplete="off" /></el-form-item>  
      <el-form-item :label="versionLabel.icon"><el-input v-model="utility.tempEditorVersion.value.icon" autocomplete="off" /></el-form-item>  
      <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="utility.displayDialog(ButtonType.Cancel, false)">{{ editorDialog.cancel }}</el-button>
          <el-button type="success" @click="utility.editVersion(utility.tempVersion.value)">{{ editorDialog.sure }}</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 新增確認框 -->
    <el-dialog v-model="appendDialog.isVisible" :title="appendDialog.title" :background="appendDialog.background" @keyup.enter="utility.appendVersion()">
      <el-form-item :label="versionLabel.id"><el-input v-model="utility.tempAppendVersion.value.id" autocomplete="off" /></el-form-item>  
      <el-form-item :label="versionLabel.name"><el-input v-model="utility.tempAppendVersion.value.name" autocomplete="off" /></el-form-item>
      <el-form-item :label="versionLabel.type">
        <el-select v-model="utility.tempAppendVersion.value.type" @change="utility.resetLangType">
          <el-option v-for="item in typeOptions" :key="item.value" :label="item.label" :value="item.value"></el-option>
        </el-select>
      </el-form-item>
      <el-form-item :label="versionLabel.lang">
        <el-select v-model="utility.tempAppendVersion.value._lang">
          <el-option v-for="item in utility.storeTypeTextList(utility.tempAppendVersion.value.type)" :key="item.value" :label="item.label" :value="item.value"></el-option>
        </el-select>
      </el-form-item>
      <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="utility.displayDialog(ButtonType.Cancel, false)">{{ appendDialog.cancel }}</el-button>
          <el-button type="success" @click="utility.appendVersion()">{{ appendDialog.sure }}</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 登入確認框 -->
    <el-dialog v-model="loginDialog.isVisible" :title="loginDialog.title" :background="loginDialog.background" @keyup.enter="utility.userLogin()">
      <el-form-item :label="versionLabel.username"><el-input v-model="utility.tempUser.value.username" autocomplete="off" /></el-form-item>  
      <el-form-item :label="versionLabel.password"><el-input v-model="utility.tempUser.value.password" type="password" autocomplete="off" /></el-form-item>  
      <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="utility.displayDialog(ButtonType.Cancel, false)">{{ loginDialog.cancel }}</el-button>
          <el-button type="success" @click="utility.userLogin()">{{ loginDialog.sure }}</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 使用者註冊確認框 -->
    <el-dialog v-model="registryDialog.isVisible" :title="registryDialog.title" :background="registryDialog.background" @keyup.enter="utility.userRegistry()">
      <el-form-item :label="versionLabel.username"><el-input v-model="utility.tempUser.value.username" autocomplete="off" /></el-form-item>
      <el-form-item :label="versionLabel.password"><el-input v-model="utility.tempUser.value.password" type="password" autocomplete="off" /></el-form-item>
      <el-form-item :label="versionLabel.password2"><el-input v-model="utility.tempUser.value.password2" type="password" autocomplete="off" /></el-form-item>
      <el-form-item :label="versionLabel.level">
        <el-select v-model="utility.tempUser.value.userLevel">
          <el-option v-for="item in userLevelOptions" :key="item.value" :label="item.label" :value="item.value"></el-option>
        </el-select>
      </el-form-item>
      <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="utility.displayDialog(ButtonType.Cancel, false)">{{ registryDialog.cancel }}</el-button>
          <el-button type="success" @click="utility.userRegistry()">{{ registryDialog.sure }}</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 更新樣式Tab -->
    <el-dialog v-model="updateTypeDialog.isVisible" :title="updateTypeDialog.title" :background="updateTypeDialog.background">
      <el-radio-group v-model="utility.tempUpdateType.value.level">
        <el-radio-button v-for="updateType in updateTypes" :key="updateType.label" :label="updateType.value" size="large">{{ updateType.label }}</el-radio-button>
      </el-radio-group>
      <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="utility.displayDialog(ButtonType.Cancel, false)">{{ loginDialog.cancel }}</el-button>
          <el-button type="success" @click="utility.updateType()">{{ loginDialog.sure }}</el-button>
        </span>
      </template>
    </el-dialog>

    <!-- 推播功能單選框 -->
    <el-dialog v-model="pushSettingDialog.isVisible" :title="pushSettingDialog.title" :background="updateTypeDialog.background" @close="utility.displayDialog(ButtonType.Cancel, false)">

      <el-tabs v-model="pushSettingDialog.paneName" type="border-card" stretch @tab-click="utility.pushSettingTabIndex">
        <el-tab-pane v-for="item in pushSettingDialog.types" :label="item.level" :key='item.type' ref="myTab">
          <template v-if="pushSettingDialog.storeType === StoreType.AppStore">
            <el-form-item :label="versionLabel.pushSetting.bundleId.label"><el-input v-model="item.json.bundleId" autocomplete="off" :placeholder="versionLabel.pushSetting.bundleId.placeholder"/></el-form-item>
            <el-form-item :label="versionLabel.pushSetting.keyId.label"><el-input v-model="item.json.keyId" autocomplete="off" :placeholder="versionLabel.pushSetting.keyId.placeholder"/></el-form-item>
            <el-form-item :label="versionLabel.pushSetting.teamId.label"><el-input v-model="item.json.teamId" autocomplete="off" :placeholder="versionLabel.pushSetting.teamId.placeholder"/></el-form-item>
          </template>
          <template v-if="pushSettingDialog.storeType === StoreType.GooglePlay">
            <el-form-item :label="versionLabel.pushSetting.appId.label"><el-input v-model="item.json.appId" autocomplete="off" :placeholder="versionLabel.pushSetting.appId.placeholder"/></el-form-item>
            <el-form-item :label="versionLabel.pushSetting.apiKey.label"><el-input v-model="item.json.apiKey" autocomplete="off" :placeholder="versionLabel.pushSetting.apiKey.placeholder"/></el-form-item>
          </template>
          <el-form-item :label="versionLabel.pushSetting.downloadUrl.label"><el-input v-model="item.json.url" autocomplete="off" :placeholder="versionLabel.pushSetting.image.placeholder"/></el-form-item>
          <el-form-item :label="versionLabel.pushSetting.title.label"><el-input v-model="item.json.title" autocomplete="off" :placeholder="versionLabel.pushSetting.title.placeholder"/></el-form-item>
          <el-form-item :label="versionLabel.pushSetting.message.label"><el-input v-model="item.json.message" autocomplete="off" :placeholder="versionLabel.pushSetting.message.placeholder"/></el-form-item>
          <el-form-item :label="versionLabel.pushSetting.image.label"><el-input v-model="item.json.image" autocomplete="off" :placeholder="versionLabel.pushSetting.image.placeholder"/></el-form-item>
          <el-form-item :label="versionLabel.pushSetting.downloadUrl.QRCode.label"><qrcode-vue :value="item.json.url" :size="versionLabel.pushSetting.downloadUrl.QRCode.size" level="M" /></el-form-item>
        </el-tab-pane>
      </el-tabs>

      <template #footer>
        <span class="dialog-footer">
          <el-button type="primary" @click="utility.displayDialog(ButtonType.Cancel, false)">{{ pushSettingDialog.cancel }}</el-button>
          <el-button type="success" @click="utility.updatePushSetting()">{{ pushSettingDialog.store }}</el-button>
          <el-button type="danger" @click="utility.pushNotification()">{{ pushSettingDialog.sure }}</el-button>
        </span>
      </template>
    </el-dialog>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from 'vue'
import { ElMessage, TabsPaneContext } from 'element-plus'
import axios from 'axios'
import Version from "./model/Version"
import TipsEnum from "./model/TipsEnum"
import UpdateEnum from "./model/UpdateEnum"
import util from "./utility/Utility"
import QrcodeVue from 'qrcode.vue'

declare global {
    interface Window { 
      result: any
      mobileApp: any
      pushTokenRegistry: (json: Map<string, string>) => void
    }
}

export default defineComponent({
  name: 'App',
  components: { QrcodeVue, },
  setup() {

    enum StoreType {
      AppStore, GooglePlay,
    }

    enum ButtonType {
      Edit, Delete, Cancel, Append, Login, Logout, Registry, UpdateVersion, UpdateVersionProd, UpdateVersionDev, PushSetting
    }

    enum LangType {
      NONE, TW, ZH_TW,
    }

    enum UserLevel {
      Account = 0, Admin = 9, Root = 99,
    }

    enum UpdateType {
      General = 0, Always = 1, Force = 2,
    }

    enum PushType {
      Store = 0, Prod = 1, Dev = 2,
    }

    const Constant = {

      baseAPI: "",
      versions: ref<Version[]>([]),

      isLoading: ref<boolean>(false),
      isDisabled: ref<boolean>(true),
      isForceDownload: ref<boolean>(false),
      isForceSinglePage: ref<boolean>(false),

      AppProtocol: {
        home: "app://home",
        download: "app://downloadFile"
      },

      tableStyle: ref<{ height: number | undefined, fixHeight: number, minimumHeight: number }>({
        height: undefined,
        fixHeight: 160,
        minimumHeight: 200,
      }),

      PaginationSetting: ref({
        count: 0,
        currentPage: 1,
        pageSize: 5,
        layout: "total, prev, pager, next",
        isHideOnSinglePage: true,
      }),

      GolangAPI: {
        userLogin: () => { return `${Constant.baseAPI}/login` },
        userRegistry: () => { return `${Constant.baseAPI}/user` },
        appVersionList: (page: number) => { return `${Constant.baseAPI}/versionList/${page}` },
        appendVersion: () => { return `${Constant.baseAPI}/appVersion` },
        deleteVersion: (id: string) => { return `${Constant.baseAPI}/appVersion/${id}`},
        editVersion: (id: string, type: StoreType) => { return `${Constant.baseAPI}/appVersion/${id}/${type}`},
        saveVersion: () => { return `${Constant.baseAPI}/saveAppVersion` },
        uploadVersion: () => { return `${Constant.baseAPI}/uploadExcel` },
        refreshVersion: () => { return `${Constant.baseAPI}/forceUpdate` },
        checkToken: () => { return `${Constant.baseAPI}/checkToken` },
        updateVersionType: (id: string) => { return `${Constant.baseAPI}/updateType/${id}` },
        registerPushToken: (appId: string) => { return `${Constant.baseAPI}/pushToken/${appId}` },
        storePushSetting: (appId: string) => { return `${Constant.baseAPI}/pushSettings/${appId}` },
        pushNotification: (appId: string) => { return `${Constant.baseAPI}/push/${appId}` },
      },
      
      VersionLabel: ref({
        account: "",
        id: "APP ID",
        name: "辨識名稱",
        url: "來源",
        level: "等級",
        icon: "商店圖示",
        type: "商店",
        auto: "自動更新",
        lang: "語言地區",
        supportedDevices: "支援裝置",
        username: "帳號",
        password: "密碼",
        password2: "再次輸入一次",
        version: "商店版本",
        versionProd: "正式版本",
        versionDev: "測試版本",
        pushSetting: {
          appId: { label: "App ID", placeholder: "com.example.appid"},
          bundleId: { label: "Bundle ID", placeholder: "com.example.appid"},
          apiKey: { label: "API Key", placeholder: "AAAAjt-BKyU:APA91bG3cDFUDuMR0nQubZh7cCkiTID7e"},
          keyId: { label: "Key ID", placeholder: "ABCDEFG"},
          teamId: { label: "Team ID", placeholder: "ABCDEFG"},
          title: { label: "標題", placeholder: "iOS好棒棒"},
          message: { label: "內容", placeholder: "不知道要寫什麼？"},
          image: { label: "圖片", placeholder: "https://your.image.png"},
          downloadUrl: { label: "下載點", placeholder: "https://apps.apple.com/app/id3939889", QRCode: { label: "下載Link", size: 128 }},
        },
        loginButton: { text: "登入", color: "#86C8C8", type: ButtonType.Login },
        appendButton: { text: "新增", color: "#FFEB3B", },
        fixButton: { text: "修改", color: "#1565C0", },
        deleteButton: { text: "刪除", color: "#FF3D00", },
        refreshButton: { text: "刷新", color: "#86C8C8", },
        registryButton: { text: "註冊", color: "#86C8C8", },
        downloadButton: { text: "備份", color: "#FF3D00", },
        uploadButton: { text: "匯入", color: "#FF3D00", accept: "application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet" },
      }),

      ConfirmDialog: ref({
        title: "確認框",
        sure: "確認",
        cancel: "取消",
        isVisible: false,
        background: "#558B2F",
      }),

      EditorDialog: ref({
        title: "確認框",
        sure: "確認",
        cancel: "取消",
        isVisible: false,
        background: "#558B2F",
      }),

      AppendDialog: ref({
        title: "確認框",
        sure: "確認",
        cancel: "取消",
        isVisible: false,
        background: "#558B2F",
      }),

      LoginDialog: ref({
        title: "使用者登入",
        sure: "確認",
        cancel: "取消",
        isVisible: false,
        background: "#558B2F",
      }),

      RegistryDialog: ref({
        title: "使用者註冊",
        sure: "確認",
        cancel: "取消",
        isVisible: false,
        background: "#558B2F",
      }),

      UpdateTypeDialog: ref({
        title: "更新等級",
        sure: "確認",
        cancel: "取消",
        isVisible: false,
        background: "#558B2F",
      }),

      PushSettingDialog: ref({
        title: "推播測試",
        sure: "發送",
        cancel: "取消",
        store: "更新",
        isVisible: false,
        paneName: "0",
        storeType: StoreType.AppStore,
        background: "#558B2F",
        index: 1,
        types: [
          { level: "上架版", type: PushType.Store, color: "#FFEBEE", json: { bundleId: "", teamId: "", keyId: "", appId: "", apiKey: "", title: "", message: "", image: "", url: "", }},
          { level: "正式版", type: PushType.Prod, color: "#FFF9C4", json: { bundleId: "", teamId: "", keyId: "", appId: "", apiKey: "", title: "", message: "", image: "", url: "",  }},
          { level: "測試版", type: PushType.Dev, color: "#DCEDC8", json: { bundleId: "", teamId: "", keyId: "", appId: "", apiKey: "", title: "", message: "", image: "", url: "",  }},
        ],
      }),

      TypeOptions: [
        { value: StoreType.AppStore, label: 'App Store' },
        { value: StoreType.GooglePlay, label: 'Google Play' },
      ],

      AreaOptions: [
        { value: LangType.NONE, label: ' ', type: null },
        { value: LangType.TW, label: '台灣', type: StoreType.AppStore },
        { value: LangType.ZH_TW, label: '正體中文', type: StoreType.GooglePlay },
      ],

      UserLevelOptions: [
        { value: UserLevel.Account, label: '使用者' },
        { value: UserLevel.Admin, label: '管理者' },
        { value: UserLevel.Root, label: '老大' },
      ],

      UpdateTypes: {
        0: { value: UpdateType.General, type: "", label: '一般更新' },
        1: { value: UpdateType.Always, type: "success", label: '總是更新' },
        2: { value: UpdateType.Force, type: "danger", label: '強制更新' },
      },
    }

    const API = {

      AuthorizationKey: "Authorization",
      SettingKey: "Setting",

      // 加上認證的headers
      authorizationHeaders: () => {
        const headers = { "Authorization": API.authToken() ?? "" }
        return headers
      },

      // 假的Loading動畫
      fakeLoading: () => {
        Constant.isLoading.value = true
        setTimeout(() => Constant.isLoading.value = false, 500)
      },

      // 取得記在localStorage的token
      authToken: () => {
        return localStorage.getItem(API.AuthorizationKey)
      },

      // 取得App版本列表
      appVersions: async (page: number) => {

        const parameters = {
          limit: Constant.PaginationSetting.value.pageSize,
          offset: (page - 1) * Constant.PaginationSetting.value.pageSize,
        }

        const authAPI = axios.create({ headers: API.authorizationHeaders() })
        const url = Constant.GolangAPI.appVersionList(Constant.PaginationSetting.value.currentPage)          

        try {
          const response = await authAPI.post(url, parameters)
          if (response.status === 401) { Utility.tips("尚未授權", "error"); return }

          const { result } = response.data
          return result

        } catch (error: any) {
          if (error.response.status === 401) { Utility.tips("尚未授權", "error"); return }
        }
      },

      // 刪除Version
      deleteVersion: async (version?: Version) => {

        if (version === undefined) { return }

        const authAPI = axios.create({ headers: API.authorizationHeaders() })
        const url = Constant.GolangAPI.deleteVersion(version.id)

        try {
          const response = await authAPI.delete(url)
          if (response.status === 401) { Utility.tips("尚未授權", "error"); return }

          const { result } = response.data  
          return result as string[]

        } catch(error: any) {
          if (error.response.status === 401) { Utility.tips("尚未授權", "error"); return }
        }
      },

      // 編譯Version
      editVersion: async (version?: Version) => {

        if (version === undefined) { return }

        const authAPI = axios.create({ headers: API.authorizationHeaders() })
        const url = Constant.GolangAPI.editVersion(version.id, version.type)
        const parameters = Utility.tempEditorVersion.value

        try {
          const response = await authAPI.patch(url, parameters)
          if (response.status === 401) { Utility.tips("尚未授權", "error"); return }

          const { result } = response.data
          return result as string[]

        } catch (error: any) {
          if (error.response.status === 401) { Utility.tips("尚未授權", "error"); return }
        }
      },

      // 新增Version
      appendVersion: async () => {

        const authAPI = axios.create({ headers: API.authorizationHeaders() })
        const url = Constant.GolangAPI.appendVersion()
        const parameters = Utility.tempAppendVersion.value

        Utility.tempAppendVersion.value.lang = Utility.langTypeText(Utility.tempAppendVersion.value._lang)

        try {
          const response = await authAPI.post(url, parameters)
          if (response.status === 401) { Utility.tips("尚未授權", "error"); return }

          const { result } = response.data
          return result as string[]

        } catch (error: any) {
          if (error.response.status === 401) { Utility.tips("尚未授權", "error"); return }
        }
      },

      // 儲存Version => Excel
      saveVersion: async () => {

        const authAPI = axios.create({ headers: API.authorizationHeaders() })
        const url = Constant.GolangAPI.saveVersion()

        try {
          const response = await authAPI.post(url)
          if (response.status === 401) { Utility.tips("尚未授權", "error"); return }

          const { result } = response.data
          return result

        } catch (error: any) {
          if (error.response.status === 401) { Utility.tips("尚未授權", "error"); return }
        }
      },

      // 上傳Excel
      uploadExcel: async (info: any) => {        

        const formData = new FormData()
        const authAPI = axios.create({ headers: API.authorizationHeaders() })
        const url = Constant.GolangAPI.uploadVersion()

        formData.append('file', info.file)

        try {
          const response = await authAPI.post(url, formData)
          if (response.status === 401) { Utility.tips("尚未授權", "error"); return }

          const { result } = response.data
          return result

        } catch (error: any) {
          if (error.response.status === 401) { Utility.tips("尚未授權", "error"); return }
        }
      },

      // 使用者登入
      userLogin: async (username: string, password: string) => {
        
        const url = Constant.GolangAPI.userLogin()
        const response = await axios.post(url, {
          "username": username,
          "password": password,
        })

        const { result } = response.data
        return result
      },

      // 使用者註冊
      userRegistry: async (username: string, password: string, userLevel: UserLevel) => {

        const authAPI = axios.create({ headers: API.authorizationHeaders() })
        const url = Constant.GolangAPI.userRegistry()

        try {
          const response = await authAPI.post(url, { "username": username, "password": password, "level": userLevel, })
          if (response.status === 401) { Utility.tips("尚未授權", "error"); return }

          const { result } = response.data
          return result

        } catch (error: any) {
          if (error.response.status === 401) { Utility.tips("尚未授權", "error"); return }
        }
      },

      // 設定更新樣式
      updateVersionType: async (version?: Version) => {

        if (version === undefined) { return }

        const authAPI = axios.create({ headers: API.authorizationHeaders() })
        const url = Constant.GolangAPI.updateVersionType(version.id)

        let type: UpdateEnum = "STORE"

        if (Utility.tempUpdateType.value.type === ButtonType.UpdateVersion) { type = "STORE" }
        if (Utility.tempUpdateType.value.type === ButtonType.UpdateVersionDev) { type = "DEV" }
        if (Utility.tempUpdateType.value.type === ButtonType.UpdateVersionProd) { type = "PROD" }

        try {
          const response = await authAPI.patch(url, { "type": type, "level": Utility.tempUpdateType.value.level, })
          if (response.status === 401) { Utility.tips("尚未授權", "error"); return }

          const { result } = response.data
          return result

        } catch (error: any) {
          if (error.response.status === 401) { Utility.tips("尚未授權", "error"); return }
        }
      },

      // 取得權限Token相關資訊
      checkToken: async () => {

        const url = Constant.GolangAPI.checkToken()
        const parameters = { "token": API.authToken() }
        const response = await axios.post(url, parameters)
        const { result } = response.data

        return result
      },

      // 刷新 => 強制更新
      refreshVersion: async () => {

        const authAPI = axios.create({ headers: API.authorizationHeaders() })
        const url = Constant.GolangAPI.refreshVersion()

        try {
          const response = await authAPI.post(url)
          if (response.status === 401) { Utility.tips("尚未授權", "error"); return }

          const { result } = response.data
          return result

        } catch (error: any) {
          if (error.response.status === 401) { Utility.tips("尚未授權", "error"); return }
        }
      },

      // 註冊推播Token
      registerPushToken: async (appId: string, token: string) => {

        const url = Constant.GolangAPI.registerPushToken(appId)
        const parameters = {
          "token": token,
          "system": window.mobileApp.type,
          "version": window.mobileApp.version,
          "username": Constant.VersionLabel.value.account
        }

        const response = await axios.post(url, parameters)
        const { result } = response.data

        return result
      },

      // 記錄推播相關設定值
      storePushSetting: async (version?: Version) => {

        if (version === undefined) { return }

        const authAPI = axios.create({ headers: API.authorizationHeaders() })
        const url = Constant.GolangAPI.storePushSetting(version.id)
        Utility.clearAllPushMessage()

        try {

          const response = await authAPI.post(url, { "store": Constant.PushSettingDialog.value.types[PushType.Store].json, "prod": Constant.PushSettingDialog.value.types[PushType.Prod].json, "dev": Constant.PushSettingDialog.value.types[PushType.Dev].json, })
          if (response.status === 401) { Utility.tips("尚未授權", "error"); return }

          const { result } = response.data
          return result

        } catch (error: any) {
          if (error.response.status === 401) { Utility.tips("尚未授權", "error"); return }
        }
      },

      // 推播測試
      pushNotification: async (version?: Version) => {

        if (version === undefined) { return }

        const authAPI = axios.create({ headers: API.authorizationHeaders() })
        const url = Constant.GolangAPI.pushNotification(version.id)
        const pushType = parseInt(Constant.PushSettingDialog.value.paneName)

        try {
          const response = await authAPI.post(url, { "pushType": pushType, "store": Constant.PushSettingDialog.value.types[PushType.Store].json, "prod": Constant.PushSettingDialog.value.types[PushType.Prod].json, "dev": Constant.PushSettingDialog.value.types[PushType.Dev].json, })
          if (response.status === 401) { Utility.tips("尚未授權", "error"); return }

          const { result } = response.data
          return result

        } catch (error: any) {
          if (error.response.status === 401) { Utility.tips("尚未授權", "error"); return }
        }
      },
    }

    const Utility = {

      tempVersion: ref<Version>(),
      
      tempUpdateType: ref({
        "level" : UpdateType.General,
        "type": ButtonType.UpdateVersion
      }),

      tempAppendVersion: ref({ 
        "id": "", 
        "name": "", 
        "type": 0, 
        "lang": "",
        "_lang": LangType.NONE, 
      }),

      tempUser: ref({ 
        "username": "", 
        "password": "", 
        "password2": "",
        "userLevel": UserLevel.Account,
      }),

      tempEditorVersion: ref({ 
        "auto": true, 
        "name": "", 
        "icon": "", 
        "version": { store: "1.0", prod: "1.0", dev: "1.0", }, 
      }),

      // 清除一些暫存的東西
      clearTemporary: () => {

        API.fakeLoading()

        Constant.ConfirmDialog.value.isVisible = false
        Constant.EditorDialog.value.isVisible = false
        Constant.AppendDialog.value.isVisible = false
        Constant.LoginDialog.value.isVisible = false
        Constant.RegistryDialog.value.isVisible = false
        Constant.UpdateTypeDialog.value.isVisible = false
        Constant.PushSettingDialog.value.isVisible = false

        Constant.PushSettingDialog.value.storeType = StoreType.AppStore
        Constant.PushSettingDialog.value.paneName = `${PushType.Store}`
        Constant.PushSettingDialog.value.types[PushType.Store].json = Utility.parseTypeJSONString()
        Constant.PushSettingDialog.value.types[PushType.Prod].json = Utility.parseTypeJSONString()
        Constant.PushSettingDialog.value.types[PushType.Dev].json = Utility.parseTypeJSONString()

        Utility.tempUpdateType.value = { level: UpdateType.General, type: ButtonType.UpdateVersion }
        Utility.tempAppendVersion.value = { id: "", name: "", type: 0, lang: "", _lang: LangType.NONE }
        Utility.tempEditorVersion.value = { auto: true, name: "", icon: "", version: { store: "1.0", prod: "1.0", dev: "1.0" }}
        Utility.tempUser.value = { username: "", password: "", password2: "", userLevel: UserLevel.Account }
        Utility.tempVersion.value = undefined
      },

      // 重新讀取資料
      reloadData: () => {

        let versions = ref<Version[]>([])

        Utility.clearTemporary()
        
        API.appVersions(Constant.PaginationSetting.value.currentPage).then((info) => {

          if (info === undefined) { return }

          Constant.PaginationSetting.value.count = info.count

          info.versions.forEach((version: any) => {
            
            const _version = {
              id: version.id, 
              name: version.name,
              icon: version.icon,
              url: version.url,
              lang: version.lang,
              auto: version.autoUpdate,
              type: version.type,
              version: version.version,
              versionProd: version.versionProd,
              versionDev: version.versionDev,
              updateVersion: version.updateVersion,
              updateVersionProd: version.updateVersionProd,
              updateVersionDev: version.updateVersionDev,
              pushSetting: version.pushSetting,
              pushSettingProduction: version.pushSettingProd,
              pushSettingDevelop: version.pushSettingDev,
              supportedDevices: version.supportedDevices
            }
            
            versions.value.push(_version)
          })

          Constant.versions.value = versions.value
        })
      },

      // 更新資料 (強制)
      refreshData: () => {

        Constant.isLoading.value = true

        API.refreshVersion().then((result) => {
          Constant.isLoading.value = false
          if (result.count > 0) { Utility.reloadData() }
        })
      },

      // 下載Excel
      downloadExcel: () => {
        API.saveVersion().then((result) => {
          Utility.downloadFile(result.filePath, Constant.isForceDownload.value)
        })
      },

      // 上傳Excel
      uploadExcel: (info: object) => {        

        API.uploadExcel(info).then((result) => {
          console.log(result)
          Utility.reloadData()
        })
      },

      // 設定更新樣式
      updateType: () => {
        API.updateVersionType(Utility.tempVersion.value).then((result) => {
          console.log(result)
          Utility.reloadData()
        })
      },

      // 更新推播相關設定值
      updatePushSetting: () => {
        API.storePushSetting(Utility.tempVersion.value).then((result) => {
          console.log(result)
          Utility.reloadData()
        })
      },

      // 推播測試
      pushNotification: () => {
        API.pushNotification(Utility.tempVersion.value).then((result) => {
          console.log(result)
          Utility.reloadData()
        })
      },

      // 設定App商店圖示
      appIcon: (type: number) => {        
        if (type == StoreType.AppStore) { return '/assets/AppStore.png' }
        if (type == StoreType.GooglePlay) { return '/assets/GooglePlay.png' }
        return '/assets/Error.png'
      },

      // 設定自動更新圖示
      autoUpdateIcon: (isAuto: boolean) => {
        if (isAuto) { return '/assets/Correct.png' }
        return ''
      },

      // 語言地區文字轉換
      langText: (_: Version, __: Record<string, unknown>, lang: string) => {

        let langType = LangType.NONE

        switch (lang) {
          case "tw": langType = LangType.TW; break
          case "zh_tw": langType = LangType.ZH_TW; break
          default: break 
        }

        return Utility.langTypeLabel(langType)
      },

      // 語系選擇文字
      langTypeLabel: (langType: LangType) => {

        let label = ""

        Constant.AreaOptions.forEach(option => {
          if (option.value === langType) { label = option.label; return }
        });
        return label
      },

      // 語言地區文字轉換
      langTypeText: (langType: LangType) => {
        switch (langType) {
          case LangType.NONE: return ""
          case LangType.TW: return "tw"
          case LangType.ZH_TW: return "zh_tw"
        }
      },

      // 語言地區文字過濾 for StoreType
      storeTypeTextList: (storeType: StoreType) => {

        let _areaOptions: [{ value: LangType, label: string }] = [
          { value: LangType.NONE, label: ' '}
        ]

        Constant.AreaOptions.forEach(option => {
          if (option.type === storeType) { 
            _areaOptions.push({ value: option.value, label: option.label }) 
          }
        });

        return _areaOptions
      },

      // 裝置支援文字轉換 => iPad, iPhone
      supportedDevicesText: (_: Version, __: Record<string, unknown>, supportedDevices: string) => {

        const devices = JSON.parse(supportedDevices)
        let array: string[] = []

        for (const key in devices) {
          const isSupported = Boolean(devices[key])
          if (isSupported) { array.push(key) }
        }

        return array.sort().join(', ')
      },

      // 刪除Version
      deleteVersion: (version?: Version) => {
        API.deleteVersion(version).then(() => {
          Utility.reloadData()
        })
      },

      // 編譯Version
      editVersion: (version?: Version) => {
        API.editVersion(version).then(() => {
          Utility.reloadData()
        })
      },

      // 新增Version
      appendVersion: () => {
        API.appendVersion().then(() => {
          Constant.AppendDialog.value.isVisible = false
          Utility.reloadData()
        })
      },

      // 使用者登入
      userLogin: () => {

        const username = Utility.tempUser.value.username
        const password = Utility.tempUser.value.password

        API.userLogin(username, password).then((result) => {

          if (result.token === undefined) {
            Utility.tips("登入失敗", "error")
            Utility.clearTemporary(); return 
          }

          localStorage.setItem(API.AuthorizationKey, `Bearer ${result.token}`)

          Constant.isDisabled.value = false
          Constant.VersionLabel.value.loginButton.type = ButtonType.Logout
          Constant.VersionLabel.value.loginButton.text = "登出"
          Constant.VersionLabel.value.account = username.toUpperCase()
          Constant.PaginationSetting.value.currentPage = 1

          Utility.clearTemporary()
          Utility.paginationAction(Constant.PaginationSetting.value.currentPage)
        })
      },

      // 註冊推播Token
      pushTokenRegistry: (json: any) => {
        
        const appId = json.appId
        const token = json.token

        if (appId === undefined) { return false }
        if (token === undefined) { return false }

        API.registerPushToken(appId, token).then((result) => {
          console.log(result)
        })
      },

      // 使用者註冊
      userRegistry: () => {

        const username = Utility.tempUser.value.username
        const password = Utility.tempUser.value.password
        const password2 = Utility.tempUser.value.password2
        const userLevel = Utility.tempUser.value.userLevel

        if (password !== password2) { Utility.tips("(╯_╰) 密碼兩次不一樣", "error"); return }

        API.userRegistry(username, password, userLevel).then((result) => {

          if (!result.isSuccess) {Utility.tips("(╥_╥) 註冊失敗", "error"); return }
          
          Utility.tips("≖‿≖ 註冊成功", "success");
          Utility.clearTemporary()
          Utility.reloadData()
        })
      },

      // 語言地區 => 回復成什麼都沒選的狀態
      resetLangType: () => {
        Utility.tempAppendVersion.value._lang = 0
      },

      // 修改Table的高度
      resetTableHeight: (fixHeight: number, miniHeight: number) => {

        if (Constant.isForceSinglePage.value) { return }

        let tableHeight = window.innerHeight - fixHeight
        if (tableHeight < 0) { tableHeight = miniHeight }
        Constant.tableStyle.value.height = tableHeight
      },

      // 根據網址上的頁址來reloadData
      reloadDataFromHref: () => {

        const url = new URL(document.location.href);
        const array = url.pathname.split("/")

        if (array.length === 3) {

          const key = array.at(1) ?? "0"
          const element = array.at(2) ?? "0"
          const page = parseInt(element)

          if (key !== "page") { return }

          Constant.PaginationSetting.value.currentPage = page
        }

        Utility.reloadData()
      },

      // 取得root路徑
      rootPath: () => {
        
        const protocol = document.location.protocol
        const hostname = document.location.hostname
        const port = document.location.port

        let rootPath = `${protocol}//${hostname}`
        if (port.length !== 0) { rootPath = `${rootPath}:${port}` }
        
        return rootPath
      },

      // UpdateType搜尋 (0 => UpdateType)
      updateVersionType: (updateType: UpdateType) => {
        const type = Constant.UpdateTypes[updateType]
        return type
      },

      // 顯示Dialog
      displayDialog: (buttonType: ButtonType, isVisible: boolean, version?: Version) => {

        Utility.tempVersion.value = version

        switch (buttonType) {

          case ButtonType.Cancel:
            Utility.clearTemporary()
            break

          case ButtonType.Delete:

            if (version === undefined) { return }

            Constant.ConfirmDialog.value.isVisible = isVisible
            Constant.ConfirmDialog.value.title = `是否要刪除『${version.name}』`
            break

          case ButtonType.Edit:

            if (version === undefined) { return }

            Constant.EditorDialog.value.isVisible = isVisible
            Constant.EditorDialog.value.title = `修改『${version.name}』相關資訊`

            Utility.tempEditorVersion.value = {
              "auto": version.auto,
              "name": version.name,
              "icon": version.icon,
              "version": { "store": version.version, "prod": version.versionProd, "dev": version.versionDev }
            }
            break

          case ButtonType.Append:
            Constant.AppendDialog.value.isVisible = isVisible
            Constant.AppendDialog.value.title = `新增APP相關訊息`
            break
          
          case ButtonType.Login:
            Constant.LoginDialog.value.isVisible = isVisible
            break

          case ButtonType.Logout:
            
            document.location.href = Utility.rootPath()
            localStorage.removeItem(API.AuthorizationKey)

            Constant.isDisabled.value = true
            Constant.VersionLabel.value.loginButton.type = ButtonType.Login
            Constant.VersionLabel.value.loginButton.text = "登入"
            Constant.VersionLabel.value.account = ""
            Constant.versions.value = []

            Utility.clearTemporary()
            break

          case ButtonType.UpdateVersion:
            Constant.UpdateTypeDialog.value.isVisible = isVisible
            Utility.tempUpdateType.value.type = ButtonType.UpdateVersion
            Utility.tempUpdateType.value.level = version?.updateVersion ?? UpdateType.General
            break

          case ButtonType.UpdateVersionDev:
            Constant.UpdateTypeDialog.value.isVisible = isVisible
            Utility.tempUpdateType.value.type = ButtonType.UpdateVersionDev
            Utility.tempUpdateType.value.level = version?.updateVersionDev ?? UpdateType.General
            break

          case ButtonType.UpdateVersionProd:
            Constant.UpdateTypeDialog.value.isVisible = isVisible
            Utility.tempUpdateType.value.type = ButtonType.UpdateVersionProd
            Utility.tempUpdateType.value.level = version?.updateVersionProd ?? UpdateType.General
            break

          case ButtonType.Registry:
            Constant.RegistryDialog.value.isVisible = isVisible
            break

          case ButtonType.PushSetting:

            Constant.PushSettingDialog.value.isVisible = isVisible
            
            if (version === undefined) { break }

            Constant.PushSettingDialog.value.storeType = version.type
            Constant.PushSettingDialog.value.title = `${version.name} - 推播測試`
            Constant.PushSettingDialog.value.types[PushType.Store].json = Utility.parseTypeJSONString(version.pushSetting)
            Constant.PushSettingDialog.value.types[PushType.Prod].json = Utility.parseTypeJSONString(version.pushSettingProduction)
            Constant.PushSettingDialog.value.types[PushType.Dev].json = Utility.parseTypeJSONString(version.pushSettingDevelop)

            Utility.pushSettingChangeTabColor(Constant.PushSettingDialog.value.paneName)
            break
        }
      },

      // 把資料庫的pushSetting文字 => {"bundleId":"","teamId":"","keyId":"","appId":"","apiKey":"",url:""}
      parseTypeJSONString: (string?: string) => {
        
        let result = JSON.parse('{"bundleId":"","teamId":"","keyId":"","appId":"","apiKey":"","title":"","message":"","image":"","url":""}')

        if (string === undefined) { return result }
        if (string.length === 0) { return result }

        const json = JSON.parse(string)

        result.bundleId = json.bundleId ?? ""
        result.teamId = json.teamId ?? ""
        result.keyId = json.keyId ?? ""
        result.appId = json.appId ?? ""
        result.apiKey = json.apiKey ?? ""
        result.url = json.url ?? ""

        return result
      },

      // 清除所有的推播訊息 (不要存在資料庫裡)
      clearAllPushMessage: () => {
        Constant.PushSettingDialog.value.types.forEach(type => {
          type.json.title = ""
          type.json.message = ""
        })
      },

      // 訊息提示
      tips: (message: string, type: TipsEnum) => {
        ElMessage({
          message: message,
          type: type,
        })
      },

      // 分頁處理
      paginationAction: (page: number) => {
        document.location.href = `/page/${page}`
        Constant.PaginationSetting.value.currentPage = page
      },

      // [下載檔案 / 開新分頁](https://stackoverflow.com/questions/53772331/vue-html-js-how-to-download-a-file-to-browser-using-the-download-tag)
      downloadFile: (urlString: string, isForce: boolean) => {

        const filename = urlString.split('/').pop()
        const url = new URL(urlString, Utility.rootPath())
        const type = "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"

        if (!isForce) { window.open(url); return }
        if (!filename) { window.open(url); return }

        axios.get(url.href, { responseType: 'blob' }).then((response) => {
          const blob = new Blob([response.data], { type: type })
          if (window.mobileApp !== undefined) { Utility.forceDownloadFileWithApp(blob, filename); return }
          util.forceDownloadFile(blob, filename)
        })
      },

      // 下載二進制檔案 for APP
      forceDownloadFileWithApp: (blob: Blob, fullName: string) => {

        util.blobToBase64(blob).then((data) => {

          let _result = {
            data: data,
            name: fullName,
            type: blob.type
          }

          window.result = _result
          util.mobileProtocol(Constant.AppProtocol.download)
        })
      },

      // 記錄個人的switch值
      switchAction: (isOn: any) => {

        const key = "isSwitch"
        const defaultSetting = { [key]: { "": false }}
        const json = JSON.parse(localStorage.getItem(API.SettingKey) ?? JSON.stringify(defaultSetting))
        const isSwitchMap = json[key]

        isSwitchMap[Constant.VersionLabel.value.account] = isOn
        localStorage.setItem(API.SettingKey, JSON.stringify({ [key]: isSwitchMap }))
      },

      // 切換單頁模式
      switchSinglePageAction: (isOn: any) => {

        const elTable = document.querySelectorAll('.el-table.el-table--fit.el-table--layout-fixed')[0] as HTMLElement

        if (isOn) {
          elTable.style.height = ""
          Constant.tableStyle.value.height = undefined
          Constant.PaginationSetting.value.currentPage = 1
          Constant.PaginationSetting.value.pageSize = Constant.PaginationSetting.value.count
          Utility.reloadData()
        } else {
          Constant.PaginationSetting.value.pageSize = 5
          Utility.reloadDataFromHref()
          Utility.resetTableHeight(Constant.tableStyle.value.fixHeight, Constant.tableStyle.value.minimumHeight)
        }
      },

      // 解析記下來的isSwitch的值
      parseSwitchValue: () => {

        const setting = localStorage.getItem(API.SettingKey)
        if (setting === null) { return false }

        const json = JSON.parse(setting)
        const isSwitch = json["isSwitch"]

        return isSwitch[Constant.VersionLabel.value.account] ?? false
      },

      // [記錄選到哪一個Tab](https://element-plus.org/en-US/component/tabs.html)
      pushSettingTabIndex: (tab: TabsPaneContext, event: Event) => {
        Constant.PushSettingDialog.value.paneName = `${tab.paneName}`
        Utility.pushSettingChangeTabColor(Constant.PushSettingDialog.value.paneName )
      },

      // 改變選擇到的Tab的顏色
      pushSettingChangeTabColor: (paneName: string) => {

        const backgroundColor = Constant.PushSettingDialog.value.types[parseInt(paneName)].color
        const tabs = document.querySelectorAll('.el-tabs__item')

        tabs.forEach((_tab) => {
          const __tab = _tab as HTMLElement
          __tab.style.backgroundColor = "white"
        })

        setTimeout(() => {
          const selectedTab = document.querySelectorAll('.el-tabs__item.is-top.is-active')[0] as HTMLElement
          const tabContent = document.getElementsByClassName('el-tabs__content')[0] as HTMLElement

          selectedTab.style.backgroundColor = backgroundColor
          tabContent.style.backgroundColor = backgroundColor
        }, 250);
      },

      // [手機網頁縮放功能開關](https://www.jianshu.com/p/a517989c0f5d)
      mobileResizeSetting: (canResize: boolean) => {
        
        if (canResize) { return }

        let meta = document.createElement('meta')

        meta.setAttribute('name', 'viewport')
        meta.setAttribute('content', 'width=device-width, initial-scale=1.0, maximum-scale=1.0, user-scalable=no')
        document.getElementsByTagName('head')[0].appendChild(meta)
      },

      // [取得環境變數 (.env)](https://blog.driftking.tw/QuickFix/vuejs-dotenv-file-configuration/#:~:text=Vue 使用的是.env,格式儲存環境變數，需要放置在專案根目錄 (非 %2Fsrc)。)
      // [npm install --save @types/node](https://stackoverflow.com/questions/45194598/using-process-env-in-typescript)
      initSetting: () => {
        Constant.baseAPI = `${process.env.VUE_APP_BASE_API}`
        Utility.mobileResizeSetting(false)
      },
    }
    
    // Vue3 Lifecycle
    onMounted(()=> {

      Utility.initSetting()

      API.checkToken().then((result) => {

        if (result === null) { localStorage.removeItem(API.AuthorizationKey); return }

        Constant.isDisabled.value = false
        Constant.VersionLabel.value.account = `${result.username} - LV.${result.level}`
        Constant.VersionLabel.value.loginButton.type = ButtonType.Logout
        Constant.VersionLabel.value.loginButton.text = "登出"
        Constant.isForceDownload.value = Utility.parseSwitchValue()
        
        if (window.mobileApp !== undefined) { util.mobileProtocol(Constant.AppProtocol.home) }
      })

      Utility.resetTableHeight(Constant.tableStyle.value.fixHeight,  Constant.tableStyle.value.minimumHeight)
      Utility.reloadDataFromHref()

      window.onresize = () => { Utility.resetTableHeight(Constant.tableStyle.value.fixHeight, Constant.tableStyle.value.minimumHeight) }
      window.pushTokenRegistry = Utility.pushTokenRegistry
    })

    return {
      utility: Utility,
      versions: Constant.versions,
      versionLabel: Constant.VersionLabel,
      confirmDialog: Constant.ConfirmDialog,
      editorDialog: Constant.EditorDialog,
      appendDialog: Constant.AppendDialog,
      loginDialog: Constant.LoginDialog,
      registryDialog: Constant.RegistryDialog,
      updateTypeDialog: Constant.UpdateTypeDialog,
      pushSettingDialog: Constant.PushSettingDialog,
      typeOptions: Constant.TypeOptions,
      userLevelOptions: Constant.UserLevelOptions,
      updateTypes: Constant.UpdateTypes,
      isLoading: Constant.isLoading,
      isDisabled: Constant.isDisabled,
      isForceDownload: Constant.isForceDownload,
      paginationSetting: Constant.PaginationSetting,
      isForceSinglePage: Constant.isForceSinglePage,
      tableStyle: Constant.tableStyle,
      ButtonType: ButtonType,
      StoreType: StoreType
    }
  }
});
</script>

<style scoped>
#app {
  text-align: center;
  color: #565455;
  margin-top: 40px;
}

.action-button {
  z-index: 87;
  position: fixed;
  bottom: 10px;
  right: 20px;
}

.action-page {
  z-index: 87;
  position: fixed;
  bottom: 20px;
  left: 20px;
}

.account {
  text-align: center;
}

.action-download {
  position: absolute;
  top: 30px;
  left: 40px;
}

.action-single-page {
  position: absolute;
  top: 30px;
  right: 40px;
}

.table-action-button {
  padding: 8px;
}

.inline-block {
  float: left;
}
</style>