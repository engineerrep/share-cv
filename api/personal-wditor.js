import Api from "~/utils/api"

// 新增公司
export const userCompanyAdd = (params) => {
    return Api.post('/api/user/company/add', params)
}

// 自定义属性-新增或修改
export const userCustomAttributeChange = (params) => {
    return Api.post('/api/user/custom/attribute/add', params)
}

// 自定义属性-获取列表
export const apiCustomAttributeList = (params) => {
    return Api.post('/api/user/custom/attribute/list', params)
}

// 自定义属性-删除
export const apiCustomAttributeDelete = (params) => {
    return Api.post('/api/user/custom/attribute/delete', params)
}

// 自定义属性-根据ID获取数据
export const apiCustomAttributeGetId = (params) => {
    return Api.post('/api/user/custom/attribute/get', params)
}

// 邮箱验证码验证
export const apiCheckEmailCode = (params) => {
    return Api.post('/api/account/check/email/code', params)
}

// 更新姓名
export const apiUpdateName = (params) => {
    return Api.post('/api/userbase/update/name', params)
}

// 更新修改技能
export const apiUpdateSkills = (params) => {
    return Api.post('/api/userbase/update/personal/skills', params)
}

// 更新-兴趣
export const apiUpdateInterests = (params) => {
    return Api.post('/api/userbase/update/interests', params)
}

// 更新-交友目的
export const apiUpdateDatingTargets = (params) => {
    return Api.post('/api/userbase/update/dating/targets', params)
}

// 更新目的地城市
export const apiUpdatePurposeCity = (params) => {
    return Api.post('/api/userbase/update/purpose/city', params)
}

// 更新个人实力
export const apiUpdatePersonalStrength = (params) => {
    return Api.post('/api/userbase/update/personal/strength', params)
}

// 教育-新增
export const apiUserEducationAdd = (params) => {
    return Api.post('/api/user/education/add', params)
}

// 教育-删除
export const apiUserEducationDelete = (params) => {
    return Api.post('/api/user/education/delete', params)
}

// 教育-根据ID获取数据
export const apiUserEducationGet = (params) => {
    return Api.post('/api/user/education/get', params)
}

// 教育-获取列表
export const apiUserEducationList = (params) => {
    return Api.post('/api/user/education/list', params)
}

// 是否公开
export const apiUpdateMode = (params) => {
    return Api.post('/api/userbase/update/mode', params)
}

// 公司-新增
export const apiUserCompanyAdd = (params) => {
    return Api.post('/api/user/company/add', params)
}

// 公司-删除
export const apiUserCompanyDelete = (params) => {
    return Api.post('/api/user/company/delete', params)
}

// 公司-根据ID获取数据
export const apiUserCompanyGet = (params) => {
    return Api.post('/api/user/company/get', params)
}

// 公司-获取列表
export const apiUserCompanyList = (params) => {
    return Api.post('/api/user/company/list', params)
}

// 公司技能-删除
export const apiUserCompanyUpdSkills = (params) => {
    return Api.post('/api/user/company/upd/skills', params)
}

// 项目经验-新增
export const apiUserProjectAdd = (params) => {
    return Api.post('/api/user/project/add', params)
}

// 项目经验-删除
export const apiUserProjectDelete = (params) => {
    return Api.post('/api/user/project/delete', params)
}

// 项目经验-根据ID获取数据
export const apiUserProjectGet = (params) => {
    return Api.post('/api/user/project/get', params)
}

// 项目经验-获取列表
export const apiUserProjectList = (params) => {
    return Api.post('/api/user/project/list', params)
}

// 项目经验技能-删除
export const apiUserProjectUpdSkills = (params) => {
    return Api.post('/api/user/project/upd/skills', params)
}

// 更新基本信息
export const apiUpdateBase = (params) => {
    console.log('Api')
    return Api.post('/api/userbase/update/base', params)
}

// 更新-头像
export const apiUpdateAvatar = (params) => {
    return Api.post('/api/userbase/update/avatar', params)
}
// 更新-头像
export const apiUpdateBackground = (params) => {
    return Api.post('/api/userbase/update/background', params)
}

// 上传文件
export const apiOssUpload = (params) => {
    return Api.post('/api/oss/upload', params, 'multipart/form-data')
}

// 获取位置
export const apiConfigInfo = () => {
    return Api.post('/api/config/info')
}

// 退出登录
export const apiLoginLogout = () => {
    return Api.post('/api/login/logout')
}

// 发送手机号验证码
export const apiSendPhoneCode = (params) => {
    return Api.post('/api/account/send/phone/code', params)
}

// 效验手机号验证码
export const apiCheckPhoneCode = (params) => {
    return Api.post('/api/account/check/phone/code', params)
}

// 更新手机信息
export const apiUserbaseUpdatePhone = (params) => {
    return Api.post('/api/userbase/update/phone', params)
}

// 获取语言列表
export const apiCommonLanguage = (params) => {
    return Api.post('/api/common/language', params)
}

// 语言-更新
export const apiUpdateLanguages = (params) => {
    return Api.post('/api/userbase/update/languages', params)
}

// 经验-获取列表
export const apiExperienceList = (params) => {
    return Api.post('/api/user/experience/list', params)
}

// 经验-根据id获取数据
export const apiExperienceGetId = (params) => {
    return Api.post('/api/user/experience/get', params)
}

// 经验-删除
export const apiExperienceDelete = (params) => {
    return Api.post('/api/user/experience/delete', params)
}

// 经验-添加
export const apiExperienceAdd = (params) => {
    return Api.post('/api/user/experience/add', params)
}

// 志愿者经历-获取列表
export const apiVolunteerList = (params) => {
    return Api.post('/api/user/volunteer/list', params)
}

// 志愿者经历-根据id获取数据
export const apiVolunteerGetId = (params) => {
    return Api.post('/api/user/volunteer/get', params)
}

// 志愿者经历-删除
export const apiVolunteerDelete = (params) => {
    return Api.post('/api/user/volunteer/delete', params)
}

// 志愿者经历-添加
export const apiVolunteerAdd = (params) => {
    return Api.post('/api/user/volunteer/add', params)
}

// 相册-添加
export const apiPhotoAdd = (params) => {
    return Api.post('/api/user/photo/add', params)
}

// 相册-删除
export const apiPhotoDelete = (params) => {
    return Api.post('/api/user/photo/delete', params)
}

// 相册-获取列表
export const apiPhotoList = (params) => {
    return Api.post('/api/user/photo/list', params)
}

// 视频-添加
export const apiVideoAdd = (params) => {
    return Api.post('/api/user/video/add', params)
}

// 视频-删除
export const apiVideoDelete = (params) => {
    return Api.post('/api/user/video/delete', params)
}

// 视频-获取列表
export const apiVideoList = (params) => {
    return Api.post('/api/user/video/list', params)
}

// 验证码更改邮箱
export const apiChangeeMailWithCode = (params) => {
    return Api.post('/api/account/change/email/with/code', params)
}
