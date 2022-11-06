(function(){"use strict";var e={1455:function(e,t,a){var l=a(9242),o=a(3396),i=a(7139);const n={class:"app","element-loading-background":"rgba(0, 0, 0, 0.7)"},r=(0,o._)("h1",null,"手機App版本設定介面",-1),s=["href"],u={class:"table-action-button"},p={class:"table-action-button"},d={class:"table-action-button"},c={class:"table-action-button"},g={class:"table-action-button"},m={class:"table-action-button"},v={class:"table-action-button"},y={class:"table-action-button"},b={class:"action-download"},V=(0,o.Uk)(" 強制下載檔案"),f={class:"dialog-footer"},w={class:"dialog-footer"},D={class:"dialog-footer"},h={class:"dialog-footer"},k={class:"dialog-footer"},U={class:"dialog-footer"};function T(e,t,a,T,L,_){const W=(0,o.up)("el-image"),A=(0,o.up)("el-table-column"),P=(0,o.up)("el-tag"),C=(0,o.up)("el-button"),S=(0,o.up)("el-table"),z=(0,o.up)("el-upload"),B=(0,o.up)("el-button-group"),H=(0,o.up)("el-pagination"),E=(0,o.up)("el-switch"),x=(0,o.up)("el-dialog"),I=(0,o.up)("el-form-item"),O=(0,o.up)("el-input"),F=(0,o.up)("el-option"),G=(0,o.up)("el-select"),R=(0,o.up)("el-radio-button"),$=(0,o.up)("el-radio-group"),N=(0,o.Q2)("loading");return(0,o.wy)(((0,o.wg)(),(0,o.iD)("div",n,[r,(0,o.Wm)(S,{data:e.versions,height:e.tableStyle.height},{default:(0,o.w5)((()=>[(0,o.Wm)(A,{label:e.versionLabel.icon},{default:(0,o.w5)((e=>[(0,o.Wm)(W,{src:e.row.icon},null,8,["src"])])),_:1},8,["label"]),(0,o.Wm)(A,{label:e.versionLabel.type},{default:(0,o.w5)((t=>[(0,o._)("a",{href:t.row.url,target:"_blank"},[(0,o.Wm)(W,{src:e.utility.appIcon(t.row.type)},null,8,["src"])],8,s)])),_:1},8,["label"]),(0,o.Wm)(A,{prop:"name",label:e.versionLabel.name},null,8,["label"]),(0,o.Wm)(A,{prop:"id",label:e.versionLabel.id},null,8,["label"]),(0,o.Wm)(A,{label:e.versionLabel.version,align:"center"},{default:(0,o.w5)((t=>[(0,o._)("div",u,[(0,o.Wm)(P,{type:e.utility.updateVersionType(t.row.updateVersion).type,effect:"dark",onClick:a=>e.utility.displayDialog(e.ButtonType.UpdateVersion,!0,t.row)},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.utility.updateVersionType(t.row.updateVersion).label),1)])),_:2},1032,["type","onClick"])]),(0,o._)("div",p,(0,i.zw)(t.row.version),1)])),_:1},8,["label"]),(0,o.Wm)(A,{label:e.versionLabel.versionProd,align:"center"},{default:(0,o.w5)((t=>[(0,o._)("div",d,[(0,o.Wm)(P,{type:e.utility.updateVersionType(t.row.updateVersionProd).type,effect:"dark",onClick:a=>e.utility.displayDialog(e.ButtonType.UpdateVersionProd,!0,t.row)},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.utility.updateVersionType(t.row.updateVersionProd).label),1)])),_:2},1032,["type","onClick"])]),(0,o._)("div",c,(0,i.zw)(t.row.versionProd),1)])),_:1},8,["label"]),(0,o.Wm)(A,{label:e.versionLabel.versionDev,align:"center"},{default:(0,o.w5)((t=>[(0,o._)("div",g,[(0,o.Wm)(P,{type:e.utility.updateVersionType(t.row.updateVersionDev).type,effect:"dark",onClick:a=>e.utility.displayDialog(e.ButtonType.UpdateVersionDev,!0,t.row)},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.utility.updateVersionType(t.row.updateVersionDev).label),1)])),_:2},1032,["type","onClick"])]),(0,o._)("div",m,(0,i.zw)(t.row.versionDev),1)])),_:1},8,["label"]),(0,o.Wm)(A,{prop:"lang",label:e.versionLabel.lang,formatter:e.utility.langText},null,8,["label","formatter"]),(0,o.Wm)(A,{label:e.versionLabel.auto},{default:(0,o.w5)((t=>[(0,o.Wm)(W,{src:e.utility.autoUpdateIcon(t.row.auto)},null,8,["src"])])),_:1},8,["label"]),(0,o.Wm)(A,{fixed:"right",width:"100"},{default:(0,o.w5)((t=>[(0,o._)("div",v,[(0,o.Wm)(C,{color:e.versionLabel.fixButton.color,round:"",onClick:a=>e.utility.displayDialog(e.ButtonType.Edit,!0,t.row)},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.versionLabel.fixButton.text),1)])),_:2},1032,["color","onClick"])]),(0,o._)("div",y,[(0,o.Wm)(C,{color:e.versionLabel.deleteButton.color,round:"",onClick:a=>e.utility.displayDialog(e.ButtonType.Delete,!0,t.row)},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.versionLabel.deleteButton.text),1)])),_:2},1032,["color","onClick"])])])),_:1})])),_:1},8,["data","height"]),(0,o._)("h1",null,(0,i.zw)(e.versionLabel.account),1),(0,o.Wm)(B,{size:"large",class:"action-button"},{default:(0,o.w5)((()=>[(0,o.Wm)(C,{type:"danger",onClick:t[0]||(t[0]=t=>e.utility.displayDialog(e.versionLabel.loginButton.type,!0))},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.versionLabel.loginButton.text),1)])),_:1}),(0,o.Wm)(C,{type:"primary",disabled:e.isDisabled,onClick:t[1]||(t[1]=t=>e.utility.displayDialog(e.ButtonType.Append,!0))},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.versionLabel.appendButton.text),1)])),_:1},8,["disabled"]),(0,o.Wm)(C,{type:"success",disabled:e.isDisabled,onClick:t[2]||(t[2]=t=>e.utility.refreshData())},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.versionLabel.refreshButton.text),1)])),_:1},8,["disabled"]),(0,o.Wm)(C,{type:"warning",disabled:e.isDisabled,onClick:t[3]||(t[3]=t=>e.utility.downloadExcel())},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.versionLabel.downloadButton.text),1)])),_:1},8,["disabled"]),(0,o.Wm)(C,{type:"primary",disabled:e.isDisabled,onClick:t[4]||(t[4]=t=>e.utility.displayDialog(e.ButtonType.Registry,!0))},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.versionLabel.registryButton.text),1)])),_:1},8,["disabled"]),(0,o.Wm)(z,{class:"inline-block","http-request":e.utility.uploadExcel,accept:e.versionLabel.uploadButton.accept},{default:(0,o.w5)((()=>[(0,o.Wm)(C,{type:"danger",disabled:e.isDisabled},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.versionLabel.uploadButton.text),1)])),_:1},8,["disabled"])])),_:1},8,["http-request","accept"])])),_:1}),(0,o.Wm)(H,{class:"action-page","current-page":e.paginationSetting.currentPage,"hide-on-single-page":e.paginationSetting.isHideOnSinglePage,background:"",layout:"total, prev, pager, next",total:e.paginationSetting.count,"page-size":e.paginationSetting.pageSize,onCurrentChange:e.utility.paginationAction},null,8,["current-page","hide-on-single-page","total","page-size","onCurrentChange"]),(0,o._)("span",b,[(0,o.Wm)(E,{modelValue:e.isForceDownload,"onUpdate:modelValue":t[5]||(t[5]=t=>e.isForceDownload=t),autocomplete:"off",onChange:e.utility.switchAction},null,8,["modelValue","onChange"]),V]),(0,o.Wm)(x,{modelValue:e.confirmDialog.isVisible,"onUpdate:modelValue":t[8]||(t[8]=t=>e.confirmDialog.isVisible=t),title:e.confirmDialog.title,background:e.confirmDialog.background,onKeyup:t[9]||(t[9]=(0,l.D2)((t=>e.utility.deleteVersion(e.utility.tempVersion.value)),["enter"]))},{footer:(0,o.w5)((()=>[(0,o._)("span",f,[(0,o.Wm)(C,{type:"primary",onClick:t[6]||(t[6]=t=>e.utility.displayDialog(e.ButtonType.Cancel,!1))},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.confirmDialog.cancel),1)])),_:1}),(0,o.Wm)(C,{type:"danger",onClick:t[7]||(t[7]=t=>e.utility.deleteVersion(e.utility.tempVersion.value))},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.confirmDialog.sure),1)])),_:1})])])),_:1},8,["modelValue","title","background"]),(0,o.Wm)(x,{modelValue:e.editorDialog.isVisible,"onUpdate:modelValue":t[18]||(t[18]=t=>e.editorDialog.isVisible=t),title:e.editorDialog.title,background:e.editorDialog.background,onKeyup:t[19]||(t[19]=(0,l.D2)((t=>e.utility.editVersion(e.utility.tempVersion.value)),["enter"]))},{footer:(0,o.w5)((()=>[(0,o._)("span",w,[(0,o.Wm)(C,{type:"primary",onClick:t[16]||(t[16]=t=>e.utility.displayDialog(e.ButtonType.Cancel,!1))},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.editorDialog.cancel),1)])),_:1}),(0,o.Wm)(C,{type:"success",onClick:t[17]||(t[17]=t=>e.utility.editVersion(e.utility.tempVersion.value))},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.editorDialog.sure),1)])),_:1})])])),default:(0,o.w5)((()=>[(0,o.Wm)(I,{label:e.versionLabel.auto},{default:(0,o.w5)((()=>[(0,o.Wm)(E,{modelValue:e.utility.tempEditorVersion.value.auto,"onUpdate:modelValue":t[10]||(t[10]=t=>e.utility.tempEditorVersion.value.auto=t),autocomplete:"off"},null,8,["modelValue"])])),_:1},8,["label"]),(0,o.Wm)(I,{label:e.versionLabel.name},{default:(0,o.w5)((()=>[(0,o.Wm)(O,{modelValue:e.utility.tempEditorVersion.value.name,"onUpdate:modelValue":t[11]||(t[11]=t=>e.utility.tempEditorVersion.value.name=t),autocomplete:"off"},null,8,["modelValue"])])),_:1},8,["label"]),(0,o.Wm)(I,{label:e.versionLabel.version},{default:(0,o.w5)((()=>[(0,o.Wm)(O,{modelValue:e.utility.tempEditorVersion.value.version.store,"onUpdate:modelValue":t[12]||(t[12]=t=>e.utility.tempEditorVersion.value.version.store=t),autocomplete:"off"},null,8,["modelValue"])])),_:1},8,["label"]),(0,o.Wm)(I,{label:e.versionLabel.versionProd},{default:(0,o.w5)((()=>[(0,o.Wm)(O,{modelValue:e.utility.tempEditorVersion.value.version.prod,"onUpdate:modelValue":t[13]||(t[13]=t=>e.utility.tempEditorVersion.value.version.prod=t),autocomplete:"off"},null,8,["modelValue"])])),_:1},8,["label"]),(0,o.Wm)(I,{label:e.versionLabel.versionDev},{default:(0,o.w5)((()=>[(0,o.Wm)(O,{modelValue:e.utility.tempEditorVersion.value.version.dev,"onUpdate:modelValue":t[14]||(t[14]=t=>e.utility.tempEditorVersion.value.version.dev=t),autocomplete:"off"},null,8,["modelValue"])])),_:1},8,["label"]),(0,o.Wm)(I,{label:e.versionLabel.icon},{default:(0,o.w5)((()=>[(0,o.Wm)(O,{modelValue:e.utility.tempEditorVersion.value.icon,"onUpdate:modelValue":t[15]||(t[15]=t=>e.utility.tempEditorVersion.value.icon=t),autocomplete:"off"},null,8,["modelValue"])])),_:1},8,["label"])])),_:1},8,["modelValue","title","background"]),(0,o.Wm)(x,{modelValue:e.appendDialog.isVisible,"onUpdate:modelValue":t[26]||(t[26]=t=>e.appendDialog.isVisible=t),title:e.appendDialog.title,background:e.appendDialog.background,onKeyup:t[27]||(t[27]=(0,l.D2)((t=>e.utility.appendVersion()),["enter"]))},{footer:(0,o.w5)((()=>[(0,o._)("span",D,[(0,o.Wm)(C,{type:"primary",onClick:t[24]||(t[24]=t=>e.utility.displayDialog(e.ButtonType.Cancel,!1))},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.appendDialog.cancel),1)])),_:1}),(0,o.Wm)(C,{type:"success",onClick:t[25]||(t[25]=t=>e.utility.appendVersion())},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.appendDialog.sure),1)])),_:1})])])),default:(0,o.w5)((()=>[(0,o.Wm)(I,{label:e.versionLabel.id},{default:(0,o.w5)((()=>[(0,o.Wm)(O,{modelValue:e.utility.tempAppendVersion.value.id,"onUpdate:modelValue":t[20]||(t[20]=t=>e.utility.tempAppendVersion.value.id=t),autocomplete:"off"},null,8,["modelValue"])])),_:1},8,["label"]),(0,o.Wm)(I,{label:e.versionLabel.name},{default:(0,o.w5)((()=>[(0,o.Wm)(O,{modelValue:e.utility.tempAppendVersion.value.name,"onUpdate:modelValue":t[21]||(t[21]=t=>e.utility.tempAppendVersion.value.name=t),autocomplete:"off"},null,8,["modelValue"])])),_:1},8,["label"]),(0,o.Wm)(I,{label:e.versionLabel.type},{default:(0,o.w5)((()=>[(0,o.Wm)(G,{modelValue:e.utility.tempAppendVersion.value.type,"onUpdate:modelValue":t[22]||(t[22]=t=>e.utility.tempAppendVersion.value.type=t),onChange:e.utility.resetLangType},{default:(0,o.w5)((()=>[((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(e.typeOptions,(e=>((0,o.wg)(),(0,o.j4)(F,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue","onChange"])])),_:1},8,["label"]),(0,o.Wm)(I,{label:e.versionLabel.lang},{default:(0,o.w5)((()=>[(0,o.Wm)(G,{modelValue:e.utility.tempAppendVersion.value._lang,"onUpdate:modelValue":t[23]||(t[23]=t=>e.utility.tempAppendVersion.value._lang=t)},{default:(0,o.w5)((()=>[((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(e.utility.storeTypeTextList(e.utility.tempAppendVersion.value.type),(e=>((0,o.wg)(),(0,o.j4)(F,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1},8,["label"])])),_:1},8,["modelValue","title","background"]),(0,o.Wm)(x,{modelValue:e.loginDialog.isVisible,"onUpdate:modelValue":t[32]||(t[32]=t=>e.loginDialog.isVisible=t),title:e.loginDialog.title,background:e.loginDialog.background,onKeyup:t[33]||(t[33]=(0,l.D2)((t=>e.utility.userLogin()),["enter"]))},{footer:(0,o.w5)((()=>[(0,o._)("span",h,[(0,o.Wm)(C,{type:"primary",onClick:t[30]||(t[30]=t=>e.utility.displayDialog(e.ButtonType.Cancel,!1))},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.loginDialog.cancel),1)])),_:1}),(0,o.Wm)(C,{type:"success",onClick:t[31]||(t[31]=t=>e.utility.userLogin())},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.loginDialog.sure),1)])),_:1})])])),default:(0,o.w5)((()=>[(0,o.Wm)(I,{label:e.versionLabel.username},{default:(0,o.w5)((()=>[(0,o.Wm)(O,{modelValue:e.utility.tempUser.value.username,"onUpdate:modelValue":t[28]||(t[28]=t=>e.utility.tempUser.value.username=t),autocomplete:"off"},null,8,["modelValue"])])),_:1},8,["label"]),(0,o.Wm)(I,{label:e.versionLabel.password},{default:(0,o.w5)((()=>[(0,o.Wm)(O,{modelValue:e.utility.tempUser.value.password,"onUpdate:modelValue":t[29]||(t[29]=t=>e.utility.tempUser.value.password=t),type:"password",autocomplete:"off"},null,8,["modelValue"])])),_:1},8,["label"])])),_:1},8,["modelValue","title","background"]),(0,o.Wm)(x,{modelValue:e.registryDialog.isVisible,"onUpdate:modelValue":t[40]||(t[40]=t=>e.registryDialog.isVisible=t),title:e.registryDialog.title,background:e.registryDialog.background,onKeyup:t[41]||(t[41]=(0,l.D2)((t=>e.utility.userRegistry()),["enter"]))},{footer:(0,o.w5)((()=>[(0,o._)("span",k,[(0,o.Wm)(C,{type:"primary",onClick:t[38]||(t[38]=t=>e.utility.displayDialog(e.ButtonType.Cancel,!1))},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.registryDialog.cancel),1)])),_:1}),(0,o.Wm)(C,{type:"success",onClick:t[39]||(t[39]=t=>e.utility.userRegistry())},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.registryDialog.sure),1)])),_:1})])])),default:(0,o.w5)((()=>[(0,o.Wm)(I,{label:e.versionLabel.username},{default:(0,o.w5)((()=>[(0,o.Wm)(O,{modelValue:e.utility.tempUser.value.username,"onUpdate:modelValue":t[34]||(t[34]=t=>e.utility.tempUser.value.username=t),autocomplete:"off"},null,8,["modelValue"])])),_:1},8,["label"]),(0,o.Wm)(I,{label:e.versionLabel.password},{default:(0,o.w5)((()=>[(0,o.Wm)(O,{modelValue:e.utility.tempUser.value.password,"onUpdate:modelValue":t[35]||(t[35]=t=>e.utility.tempUser.value.password=t),type:"password",autocomplete:"off"},null,8,["modelValue"])])),_:1},8,["label"]),(0,o.Wm)(I,{label:e.versionLabel.password2},{default:(0,o.w5)((()=>[(0,o.Wm)(O,{modelValue:e.utility.tempUser.value.password2,"onUpdate:modelValue":t[36]||(t[36]=t=>e.utility.tempUser.value.password2=t),type:"password",autocomplete:"off"},null,8,["modelValue"])])),_:1},8,["label"]),(0,o.Wm)(I,{label:e.versionLabel.level},{default:(0,o.w5)((()=>[(0,o.Wm)(G,{modelValue:e.utility.tempUser.value.userLevel,"onUpdate:modelValue":t[37]||(t[37]=t=>e.utility.tempUser.value.userLevel=t)},{default:(0,o.w5)((()=>[((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(e.userLevelOptions,(e=>((0,o.wg)(),(0,o.j4)(F,{key:e.value,label:e.label,value:e.value},null,8,["label","value"])))),128))])),_:1},8,["modelValue"])])),_:1},8,["label"])])),_:1},8,["modelValue","title","background"]),(0,o.Wm)(x,{modelValue:e.updateTypeDialog.isVisible,"onUpdate:modelValue":t[45]||(t[45]=t=>e.updateTypeDialog.isVisible=t),title:e.updateTypeDialog.title,background:e.updateTypeDialog.background},{footer:(0,o.w5)((()=>[(0,o._)("span",U,[(0,o.Wm)(C,{type:"primary",onClick:t[43]||(t[43]=t=>e.utility.displayDialog(e.ButtonType.Cancel,!1))},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.loginDialog.cancel),1)])),_:1}),(0,o.Wm)(C,{type:"success",onClick:t[44]||(t[44]=t=>e.utility.updateType())},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.loginDialog.sure),1)])),_:1})])])),default:(0,o.w5)((()=>[(0,o.Wm)($,{modelValue:e.utility.tempUpdateType.value.level,"onUpdate:modelValue":t[42]||(t[42]=t=>e.utility.tempUpdateType.value.level=t)},{default:(0,o.w5)((()=>[((0,o.wg)(!0),(0,o.iD)(o.HY,null,(0,o.Ko)(e.updateTypes,(e=>((0,o.wg)(),(0,o.iD)("div",{key:e.label},[(0,o.Wm)(R,{label:e.value,size:"large"},{default:(0,o.w5)((()=>[(0,o.Uk)((0,i.zw)(e.label),1)])),_:2},1032,["label"])])))),128))])),_:1},8,["modelValue"])])),_:1},8,["modelValue","title","background"])])),[[N,e.isLoading,void 0,{fullscreen:!0,lock:!0}]])}a(7658),a(2262),a(4506);var L=a(4870),_=a(7178),W=a(6265),A=a.n(W),P=(0,o.aZ)({name:"App",components:{},setup(){let e,t,a,l,i;(function(e){e[e["AppStore"]=0]="AppStore",e[e["GooglePlay"]=1]="GooglePlay"})(e||(e={})),function(e){e[e["Edit"]=0]="Edit",e[e["Delete"]=1]="Delete",e[e["Cancel"]=2]="Cancel",e[e["Append"]=3]="Append",e[e["Login"]=4]="Login",e[e["Logout"]=5]="Logout",e[e["Registry"]=6]="Registry",e[e["UpdateVersion"]=7]="UpdateVersion",e[e["UpdateVersionProd"]=8]="UpdateVersionProd",e[e["UpdateVersionDev"]=9]="UpdateVersionDev"}(t||(t={})),function(e){e[e["NONE"]=0]="NONE",e[e["TW"]=1]="TW",e[e["ZH_TW"]=2]="ZH_TW"}(a||(a={})),function(e){e[e["Account"]=0]="Account",e[e["Admin"]=9]="Admin",e[e["Root"]=99]="Root"}(l||(l={})),function(e){e[e["General"]=0]="General",e[e["Always"]=1]="Always",e[e["Force"]=2]="Force"}(i||(i={}));const n={baseAPI:"",versions:(0,L.iH)([]),isLoading:(0,L.iH)(!1),isDisabled:(0,L.iH)(!0),isForceDownload:(0,L.iH)(!1),tableStyle:(0,L.iH)({height:0,fixHeight:200,minimumHeight:200}),PaginationSetting:(0,L.iH)({count:0,currentPage:1,pageSize:3,layout:"total, prev, pager, next",isHideOnSinglePage:!0}),GolangAPI:{userLogin:()=>`${n.baseAPI}/login`,userRegistry:()=>`${n.baseAPI}/user`,appVersionList:e=>`${n.baseAPI}/versionList/${e}`,appendVersion:()=>`${n.baseAPI}/appVersion`,deleteVersion:e=>`${n.baseAPI}/appVersion/${e}`,editVersion:(e,t)=>`${n.baseAPI}/appVersion/${e}/${t}`,saveVersion:()=>`${n.baseAPI}/saveAppVersion`,uploadVersion:()=>`${n.baseAPI}/uploadExcel`,refreshVersion:()=>`${n.baseAPI}/forceUpdate`,checkToken:()=>`${n.baseAPI}/checkToken`,updateVersionType:e=>`${n.baseAPI}/updateType/${e}`},VersionLabel:(0,L.iH)({account:"",id:"APP ID",name:"辨識名稱",url:"來源",level:"等級",icon:"商店圖示",type:"商店",auto:"自動更新",lang:"語言地區",username:"帳號",password:"密碼",password2:"再次輸入一次",version:"商店版本",versionProd:"正式版本",versionDev:"測試版本",loginButton:{text:"登入",color:"#86C8C8",type:t.Login},appendButton:{text:"新增",color:"#FFEB3B"},fixButton:{text:"修改",color:"#1565C0"},deleteButton:{text:"刪除",color:"#FF3D00"},refreshButton:{text:"刷新",color:"#86C8C8"},registryButton:{text:"註冊",color:"#86C8C8"},downloadButton:{text:"備份",color:"#FF3D00"},uploadButton:{text:"匯入",color:"#FF3D00",accept:"application/vnd.ms-excel, application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"}}),ConfirmDialog:(0,L.iH)({title:"確認框",sure:"確認",cancel:"取消",isVisible:!1,background:"#558B2F"}),EditorDialog:(0,L.iH)({title:"確認框",sure:"確認",cancel:"取消",isVisible:!1,background:"#558B2F"}),AppendDialog:(0,L.iH)({title:"確認框",sure:"確認",cancel:"取消",isVisible:!1,background:"#558B2F"}),LoginDialog:(0,L.iH)({title:"使用者登入",sure:"確認",cancel:"取消",isVisible:!1,background:"#558B2F"}),RegistryDialog:(0,L.iH)({title:"使用者註冊",sure:"確認",cancel:"取消",isVisible:!1,background:"#558B2F"}),UpdateTypeDialog:(0,L.iH)({title:"更新等級",sure:"確認",cancel:"取消",isVisible:!1,background:"#558B2F"}),TypeOptions:[{value:e.AppStore,label:"App Store"},{value:e.GooglePlay,label:"Google Play"}],AreaOptions:[{value:a.NONE,label:" ",type:null},{value:a.TW,label:"台灣",type:e.AppStore},{value:a.ZH_TW,label:"正體中文",type:e.GooglePlay}],UserLevelOptions:[{value:l.Account,label:"使用者"},{value:l.Admin,label:"管理者"},{value:l.Root,label:"老大"}],UpdateTypes:{0:{value:i.General,type:"",label:"一般更新"},1:{value:i.Always,type:"success",label:"總是更新"},2:{value:i.Force,type:"danger",label:"強制更新"}}},r={AuthorizationKey:"Authorization",SettingKey:"Setting",authorizationHeaders:()=>{const e={Authorization:r.authToken()??""};return e},fakeLoading:()=>{n.isLoading.value=!0,setTimeout((()=>n.isLoading.value=!1),500)},authToken:()=>localStorage.getItem(r.AuthorizationKey),appVersions:async e=>{const t={limit:n.PaginationSetting.value.pageSize,offset:(e-1)*n.PaginationSetting.value.pageSize},a=A().create({headers:r.authorizationHeaders()}),l=n.GolangAPI.appVersionList(n.PaginationSetting.value.currentPage);try{const e=await a.post(l,t);if(401===e.status)return void s.tips("尚未授權","error");const{result:o}=e.data;return o}catch(o){if(401===o.response.status)return void s.tips("尚未授權","error")}},deleteVersion:async e=>{if(void 0===e)return;const t=A().create({headers:r.authorizationHeaders()}),a=n.GolangAPI.deleteVersion(e.id);try{const e=await t.delete(a);if(401===e.status)return void s.tips("尚未授權","error");console.log(e.status);const{result:l}=e.data;return l}catch(l){if(401===l.response.status)return void s.tips("尚未授權","error")}},editVersion:async e=>{if(void 0===e)return;const t=A().create({headers:r.authorizationHeaders()}),a=n.GolangAPI.editVersion(e.id,e.type),l=s.tempEditorVersion.value;try{const e=await t.patch(a,l);if(401===e.status)return void s.tips("尚未授權","error");const{result:o}=e.data;return o}catch(o){if(401===o.response.status)return void s.tips("尚未授權","error")}},appendVersion:async()=>{const e=A().create({headers:r.authorizationHeaders()}),t=n.GolangAPI.appendVersion(),a=s.tempAppendVersion.value;s.tempAppendVersion.value.lang=s.langTypeText(s.tempAppendVersion.value._lang);try{const l=await e.post(t,a);if(401===l.status)return void s.tips("尚未授權","error");const{result:o}=l.data;return o}catch(l){if(401===l.response.status)return void s.tips("尚未授權","error")}},saveVersion:async()=>{const e=A().create({headers:r.authorizationHeaders()}),t=n.GolangAPI.saveVersion();try{const a=await e.post(t);if(401===a.status)return void s.tips("尚未授權","error");const{result:l}=a.data;return l}catch(a){if(401===a.response.status)return void s.tips("尚未授權","error")}},uploadExcel:async e=>{const t=new FormData,a=A().create({headers:r.authorizationHeaders()}),l=n.GolangAPI.uploadVersion();t.append("file",e.file);try{const e=await a.post(l,t);if(401===e.status)return void s.tips("尚未授權","error");const{result:o}=e.data;return o}catch(o){if(401===o.response.status)return void s.tips("尚未授權","error")}},userLogin:async(e,t)=>{const a=n.GolangAPI.userLogin(),l=await A().post(a,{username:e,password:t}),{result:o}=l.data;return o},userRegistry:async(e,t,a)=>{const l=A().create({headers:r.authorizationHeaders()}),o=n.GolangAPI.userRegistry();try{const i=await l.post(o,{username:e,password:t,level:a});if(401===i.status)return void s.tips("尚未授權","error");const{result:n}=i.data;return n}catch(i){if(401===i.response.status)return void s.tips("尚未授權","error")}},updateVersionType:async e=>{if(void 0===e)return;const a=A().create({headers:r.authorizationHeaders()}),l=n.GolangAPI.updateVersionType(e.id);let o="STORE";s.tempUpdateType.value.type===t.UpdateVersion&&(o="STORE"),s.tempUpdateType.value.type===t.UpdateVersionDev&&(o="DEV"),s.tempUpdateType.value.type===t.UpdateVersionProd&&(o="PROD");try{const e=await a.patch(l,{type:o,level:s.tempUpdateType.value.level});if(401===e.status)return void s.tips("尚未授權","error");const{result:t}=e.data;return t}catch(i){if(401===i.response.status)return void s.tips("尚未授權","error")}},checkToken:async()=>{const e=n.GolangAPI.checkToken(),t={token:r.authToken()},a=await A().post(e,t),{result:l}=a.data;return l},refreshVersion:async()=>{const e=A().create({headers:r.authorizationHeaders()}),t=n.GolangAPI.refreshVersion();try{const a=await e.post(t);if(401===a.status)return void s.tips("尚未授權","error");const{result:l}=a.data;return l}catch(a){if(401===a.response.status)return void s.tips("尚未授權","error")}}},s={tempVersion:(0,L.iH)(),tempUpdateType:(0,L.iH)({level:i.General,type:t.UpdateVersion}),tempAppendVersion:(0,L.iH)({id:"",name:"",type:0,lang:"",_lang:a.NONE}),tempUser:(0,L.iH)({username:"",password:"",password2:"",userLevel:l.Account}),tempEditorVersion:(0,L.iH)({auto:!0,name:"",icon:"",version:{store:"1.0",prod:"1.0",dev:"1.0"}}),reloadData:()=>{let e=(0,L.iH)([]);s.clearTemporary(),r.appVersions(n.PaginationSetting.value.currentPage).then((t=>{n.PaginationSetting.value.count=t.count,t.versions.forEach((t=>{const a={id:t.id,name:t.name,icon:t.icon,url:t.url,lang:t.lang,auto:t.autoUpdate,type:t.type,version:t.version,versionProd:t.versionProd,versionDev:t.versionDev,updateVersion:t.updateVersion,updateVersionProd:t.updateVersionProd,updateVersionDev:t.updateVersionDev};e.value.push(a)})),n.versions.value=e.value}))},refreshData:()=>{r.refreshVersion().then((e=>{e.count>0&&s.reloadData()}))},downloadExcel:()=>{r.saveVersion().then((e=>{s.downloadFile(e.filePath,n.isForceDownload.value)}))},uploadExcel:e=>{r.uploadExcel(e).then((e=>{s.reloadData()}))},updateType:()=>{r.updateVersionType(s.tempVersion.value).then((e=>{s.reloadData()}))},clearTemporary:()=>{r.fakeLoading(),n.ConfirmDialog.value.isVisible=!1,n.EditorDialog.value.isVisible=!1,n.AppendDialog.value.isVisible=!1,n.LoginDialog.value.isVisible=!1,n.RegistryDialog.value.isVisible=!1,n.UpdateTypeDialog.value.isVisible=!1,s.tempUpdateType.value={level:i.General,type:t.UpdateVersion},s.tempAppendVersion.value={id:"",name:"",type:0,lang:"",_lang:a.NONE},s.tempEditorVersion.value={auto:!0,name:"",icon:"",version:{store:"1.0",prod:"1.0",dev:"1.0"}},s.tempUser.value={username:"",password:"",password2:"",userLevel:l.Account},s.tempVersion.value=void 0},appIcon:t=>t==e.AppStore?"/assets/AppStore.png":t==e.GooglePlay?"/assets/GooglePlay.png":"/assets/Error.png",autoUpdateIcon:e=>e?"/assets/Correct.png":"",langText:(e,t,l)=>{let o=a.NONE;switch(l){case"tw":o=a.TW;break;case"zh_tw":o=a.ZH_TW;break;default:break}return s.langTypeLabel(o)},langTypeLabel:e=>{let t="";return n.AreaOptions.forEach((a=>{a.value!==e||(t=a.label)})),t},langTypeText:e=>{switch(e){case a.NONE:return"";case a.TW:return"tw";case a.ZH_TW:return"zh_tw"}},storeTypeTextList:e=>{let t=[{value:a.NONE,label:" "}];return n.AreaOptions.forEach((a=>{a.type===e&&t.push({value:a.value,label:a.label})})),t},deleteVersion:e=>{r.deleteVersion(e).then((()=>{s.reloadData()}))},editVersion:e=>{r.editVersion(e).then((()=>{s.reloadData()}))},appendVersion:()=>{r.appendVersion().then((()=>{n.AppendDialog.value.isVisible=!1,s.reloadData()}))},userLogin:()=>{const e=s.tempUser.value.username,a=s.tempUser.value.password;r.userLogin(e,a).then((a=>{if(void 0===a.token)return s.tips("登入失敗","error"),void s.clearTemporary();localStorage.setItem(r.AuthorizationKey,`Bearer ${a.token}`),n.isDisabled.value=!1,n.VersionLabel.value.loginButton.type=t.Logout,n.VersionLabel.value.loginButton.text="登出",n.VersionLabel.value.account=e.toUpperCase(),n.PaginationSetting.value.currentPage=1,s.clearTemporary(),s.paginationAction(n.PaginationSetting.value.currentPage)}))},userRegistry:()=>{const e=s.tempUser.value.username,t=s.tempUser.value.password,a=s.tempUser.value.password2,l=s.tempUser.value.userLevel;t===a?r.userRegistry(e,t,l).then((e=>{e.isSuccess?(s.tips("≖‿≖ 註冊成功","success"),s.clearTemporary(),s.reloadData()):s.tips("(╥_╥) 註冊失敗","error")})):s.tips("(╯_╰) 密碼兩次不一樣","error")},resetLangType:()=>{s.tempAppendVersion.value._lang=0},resetTableHeight:(e,t)=>{let a=window.innerHeight-e;a<0&&(a=t),n.tableStyle.value.height=a},reloadDataFromHref:()=>{const e=new URL(document.location.href),t=e.pathname.split("/");if(3===t.length){const e=t.at(1)??"0",a=t.at(2)??"0",l=parseInt(a);if("page"!==e)return;n.PaginationSetting.value.currentPage=l}s.reloadData()},rootPath:()=>{const e=document.location.protocol,t=document.location.hostname,a=document.location.port;let l=`${e}//${t}`;return 0!==a.length&&(l=`${l}:${a}`),l},updateVersionType:e=>{const t=n.UpdateTypes[e];return t},displayDialog:(e,a,l)=>{switch(e){case t.Cancel:s.clearTemporary();break;case t.Delete:if(void 0===l)return;n.ConfirmDialog.value.isVisible=a,n.ConfirmDialog.value.title=`是否要刪除『${l.name}』`,s.tempVersion.value=l;break;case t.Edit:if(void 0===l)return;n.EditorDialog.value.isVisible=a,n.EditorDialog.value.title=`修改『${l.name}』相關資訊`,s.tempVersion.value=l,s.tempEditorVersion.value={auto:l.auto,name:l.name,icon:l.icon,version:{store:l.version,prod:l.versionProd,dev:l.versionDev}};break;case t.Append:n.AppendDialog.value.isVisible=a,n.AppendDialog.value.title="新增APP相關訊息";break;case t.Login:n.LoginDialog.value.isVisible=a;break;case t.Logout:document.location.href=s.rootPath(),localStorage.removeItem(r.AuthorizationKey),n.isDisabled.value=!0,n.VersionLabel.value.loginButton.type=t.Login,n.VersionLabel.value.loginButton.text="登入",n.VersionLabel.value.account="",n.versions.value=[],s.clearTemporary();break;case t.UpdateVersion:n.UpdateTypeDialog.value.isVisible=a,s.tempVersion.value=l,s.tempUpdateType.value.type=t.UpdateVersion,s.tempUpdateType.value.level=l?.updateVersion??i.General;break;case t.UpdateVersionDev:n.UpdateTypeDialog.value.isVisible=a,s.tempVersion.value=l,s.tempUpdateType.value.type=t.UpdateVersionDev,s.tempUpdateType.value.level=l?.updateVersionDev??i.General;break;case t.UpdateVersionProd:n.UpdateTypeDialog.value.isVisible=a,s.tempVersion.value=l,s.tempUpdateType.value.type=t.UpdateVersionProd,s.tempUpdateType.value.level=l?.updateVersionProd??i.General;break;case t.Registry:n.RegistryDialog.value.isVisible=a;break}},tips:(e,t)=>{(0,_.z8)({message:e,type:t})},paginationAction:e=>{document.location.href=`/page/${e}`,n.PaginationSetting.value.currentPage=e},downloadFile:(e,t)=>{const a=e.split("/").pop(),l=new URL(e,s.rootPath()),o="application/vnd.openxmlformats-officedocument.wordprocessingml.document";t&&a?A().get(l.href,{responseType:"blob"}).then((e=>{const t=new Blob([e.data],{type:o});s.forceDownloadFileWithBlob(t,a)})):window.open(l)},forceDownloadFileWithBlob:(e,t)=>{const a=window.URL||window.webkitURL,l=a.createObjectURL(e),o=document.createElement("a");o.setAttribute("download",t),o.setAttribute("href",l),document.body.appendChild(o),o.click(),document.body.removeChild(o)},switchAction:e=>{const t={isSwitch:e};localStorage.setItem(r.SettingKey,JSON.stringify(t))},parseSwitchValue:()=>{const e=localStorage.getItem(r.SettingKey);if(null===e)return!1;const t=JSON.parse(e),a=t["isSwitch"]??!1;return a},initSetting:()=>{n.baseAPI="http://192.168.1.102:12345",alert(n.baseAPI)}};return(0,o.bv)((()=>{s.initSetting(),r.checkToken().then((e=>{null!==e?(n.isDisabled.value=!1,n.VersionLabel.value.account=`${e.username} - LV.${e.level}`,n.VersionLabel.value.loginButton.type=t.Logout,n.VersionLabel.value.loginButton.text="登出",n.isForceDownload.value=s.parseSwitchValue()):localStorage.removeItem(r.AuthorizationKey)})),s.resetTableHeight(n.tableStyle.value.fixHeight,n.tableStyle.value.minimumHeight),s.reloadDataFromHref(),window.onresize=()=>{s.resetTableHeight(n.tableStyle.value.fixHeight,n.tableStyle.value.minimumHeight)}})),{utility:s,versions:n.versions,versionLabel:n.VersionLabel,confirmDialog:n.ConfirmDialog,editorDialog:n.EditorDialog,appendDialog:n.AppendDialog,loginDialog:n.LoginDialog,registryDialog:n.RegistryDialog,updateTypeDialog:n.UpdateTypeDialog,typeOptions:n.TypeOptions,userLevelOptions:n.UserLevelOptions,updateTypes:n.UpdateTypes,isLoading:n.isLoading,isDisabled:n.isDisabled,isForceDownload:n.isForceDownload,paginationSetting:n.PaginationSetting,tableStyle:n.tableStyle,ButtonType:t}}}),C=a(89);const S=(0,C.Z)(P,[["render",T]]);var z=S,B=a(7086),H=a(6229);a(2834);(0,l.ri)(z).use(B.Z,{locale:H.Z}).mount("#app")}},t={};function a(l){var o=t[l];if(void 0!==o)return o.exports;var i=t[l]={exports:{}};return e[l].call(i.exports,i,i.exports,a),i.exports}a.m=e,function(){var e=[];a.O=function(t,l,o,i){if(!l){var n=1/0;for(p=0;p<e.length;p++){l=e[p][0],o=e[p][1],i=e[p][2];for(var r=!0,s=0;s<l.length;s++)(!1&i||n>=i)&&Object.keys(a.O).every((function(e){return a.O[e](l[s])}))?l.splice(s--,1):(r=!1,i<n&&(n=i));if(r){e.splice(p--,1);var u=o();void 0!==u&&(t=u)}}return t}i=i||0;for(var p=e.length;p>0&&e[p-1][2]>i;p--)e[p]=e[p-1];e[p]=[l,o,i]}}(),function(){a.n=function(e){var t=e&&e.__esModule?function(){return e["default"]}:function(){return e};return a.d(t,{a:t}),t}}(),function(){a.d=function(e,t){for(var l in t)a.o(t,l)&&!a.o(e,l)&&Object.defineProperty(e,l,{enumerable:!0,get:t[l]})}}(),function(){a.g=function(){if("object"===typeof globalThis)return globalThis;try{return this||new Function("return this")()}catch(e){if("object"===typeof window)return window}}()}(),function(){a.o=function(e,t){return Object.prototype.hasOwnProperty.call(e,t)}}(),function(){var e={826:0};a.O.j=function(t){return 0===e[t]};var t=function(t,l){var o,i,n=l[0],r=l[1],s=l[2],u=0;if(n.some((function(t){return 0!==e[t]}))){for(o in r)a.o(r,o)&&(a.m[o]=r[o]);if(s)var p=s(a)}for(t&&t(l);u<n.length;u++)i=n[u],a.o(e,i)&&e[i]&&e[i][0](),e[i]=0;return a.O(p)},l=self["webpackChunkhtml"]=self["webpackChunkhtml"]||[];l.forEach(t.bind(null,0)),l.push=t.bind(null,l.push.bind(l))}();var l=a.O(void 0,[998],(function(){return a(1455)}));l=a.O(l)})();
//# sourceMappingURL=index.bb9940c9.js.map