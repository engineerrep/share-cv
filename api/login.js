import Api from "~/utils/api"

// 谷歌授权登录
export const googleAuthLogin = (params) => {
    return Api.post('/api/web/google/login', params)
}

// 获取邮箱验证码
export const sendCode = (params) => {
    return Api.post('/api/account/send/email/code', params)
}

/*
* 邮箱注册
 */
export const registerEmail = (params) => {
    return Api.post('/api/register/email', params)
}

/*
* 邮箱注册
 */
export const apiEmailRegister = (params) => {
    return Api.post('/api/web/email/register', params)
}

// 注销
export const logout = (params) => {
    return Api.post('/api/login/logout', params)
}

// 邮箱密码登录
export const emailLogin = (params) => {
    return Api.post('/api/web/email/password', params)
}

// 获取国家列表
export const apiCommonCountry = (params) => {
    return Api.post('/api/common/country', params)
}
// 更新国家
export const apiUpdateCountry = (params) => {
    return Api.post('/api/userbase/update/country', params)
}

// 更新详情地址
export const apiUpdateAddress = (params) => {
    return Api.post('/api/userbase/update/address', params)
}
// 更新-城市
export const apiUpdateCity = (params) => {
    return Api.post('/api/userbase/update/city', params)
}

// 验证邮箱是否注册
export const apiRegisterCheckEmail = (params) => {
    return Api.post('/api/register/check/email', params)
}
