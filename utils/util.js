// uuid


import {useState} from "#app";
import {userStore} from "@/composables/store/user.js"

let userInfoStore = userStore()

export function uuid() {
    var s = [];
    var hexDigits = "0123456789abcdef";
    for (var i = 0; i < 36; i++) {
        s[i] = hexDigits.substr(Math.floor(Math.random() * 0x10), 1);
    }
    s[14] = "4"; // bits 12-15 of the time_hi_and_version field to 0010
    s[19] = hexDigits.substr((s[19] & 0x3) | 0x8, 1); // bits 6-7 of the clock_seq_hi_and_reserved to 01
    s[8] = s[13] = s[18] = s[23] = "-";

    var uuid = s.join("");
    return uuid;
}

export function isMobile() {
    if (process.client) {
        if (/Android|webOS|iPhone|iPad|iPod|navigator|IEMobile|Opera Mini/i.test(window.navigator.userAgent) || window.innerWidth <= 770) {
            return true
        }
    }
    return false
}

// 浏览器信息
export function getBrowserInfo() {
    let agent = navigator.userAgent.toLowerCase();

    let regStr_ie = /msie [\d.]+;/gi;
    let regStr_ff = /firefox\/[\d.]+/gi
    let regStr_chrome = /chrome\/[\d.]+/gi;
    let regStr_saf = /safari\/[\d.]+/gi;
    //IE
    if (agent.indexOf("msie") > 0) {
        return agent.match(regStr_ie);
    }

    //firefox
    if (agent.indexOf("firefox") > 0) {
        return agent.match(regStr_ff);
    }

    //Safari
    if (agent.indexOf("safari") > 0 && agent.indexOf("chrome") < 0) {
        return agent.match(regStr_saf);
    }

    //Chrome
    if (agent.indexOf("chrome") > 0) {
        return agent.match(regStr_chrome);
    }

}

export const getUrlParam = (search, key) => {
    if (!search) {
        return ''
    }
    const _search = search.split('?').length > 1 ? search.split('?')[1] : ''
    if (_search === '') {
        return ''
    }
    const params = decodeURIComponent(_search).split('&')
    let value = ''
    for (let i = 0; i < params.length; i++) {
        const keyValueArray = params[i].split('=')
        if (keyValueArray[0] === key) {
            value = keyValueArray[1]
        }
    }
    return value
}
export const googleData = async (state, code) => {
    if (code) {
        console.log('code', code)
        return code
        // let url = window.location.host;
        //
        // // let url = window.location.href.replace(/\/$/, '');
        // await googleAuthLogin({
        //     code,
        //     state,
        //     callbackURL: url,
        //     deviceId: uuid(),
        //     browserName: getBrowserInfo()[0].replace('/', 'V'),
        // }).then((res) => {
        //     console.log('sss', res);
        // });
    }
};

export const googleDataError = (error) => {
    if (error) {

    }
};

// 获取指定cookie
export const getCookie = (cname) => {
    let name = cname + "=";
    let ca = document.cookie.split(";");
    for (let i = 0; i < ca.length; i++) {
        let c = ca[i].trim();
        if (c.indexOf(name) == 0) return c.substring(name.length, c.length);
    }
    return "";
}

//  获取浏览器名称和版本
export const getExplorerInfo = () => {
    let explorer = window.navigator.userAgent;
    explorer = explorer.toLowerCase();
    //ie
    if (explorer.indexOf("msie") >= 0) {
        let ver = explorer.match(/msie ([\d.]+)/)[1] || "";
        return "IE" + ver;
    }
    //firefox
    else if (explorer.indexOf("firefox") >= 0) {
        let ver = explorer.match(/firefox\/([\d.]+)/)[1] || "";
        return "Firefox" + ver;
    }
    //Chrome
    else if (explorer.indexOf("chrome") >= 0) {
        let ver = explorer.match(/chrome\/([\d.]+)/)[1] || "";
        return "Chrome" + ver;
    }
    //Opera
    else if (explorer.indexOf("opera") >= 0) {
        let ver = explorer.match(/opera.([\d.]+)/)[1] || "";
        return "Opera" + ver;
    }
    //Safari
    else if (explorer.indexOf("safari") >= 0) {
        let ver = explorer.match(/version\/([\d.]+)/)[1] || "";
        return "Safari" + ver;
    }
    if (explorer.indexOf("edge") >= 0) {
        let ver = explorer.match(/edge\/([\d.]+)/)[1] || "";
        return "edge" + ver;
    }
    //遨游浏览器
    if (explorer.indexOf("maxthon") >= 0) {
        let ver = explorer.match(/maxthon\/([\d.]+)/)[1] || "";
        return "傲游浏览器" + ver;
    }
    //QQ浏览器
    if (explorer.indexOf("qqbrowser") >= 0) {
        let ver = explorer.match(/qqbrowser\/([\d.]+)/)[1] || "";
        return "QQ浏览器" + ver;
    }
    //搜狗浏览器
    if (explorer.indexOf("se 2.x") >= 0) {
        return "搜狗浏览器";
    }
    return "";
}

/* 修改标签 */
export const reviseDom = (dom, txt) => {
    if (document.querySelector(dom)) document.querySelector(dom).innerHTML = txt
}

// 回到顶部
export const backToTheTop = () => {
    let currentScrollPosition = document.documentElement.scrollTop || document.body.scrollTop;
    if (currentScrollPosition > 0) {
        window.requestAnimationFrame(backToTheTop);
        window.scrollTo(0, 0);
    }
}

// 计算多少岁
export const calculateHowOld = (time = new Date()) => {
    const birthday = new Date(time);
    const today = new Date();
    const diffTime = Math.abs(today - birthday);
    const diffDays = Math.ceil(diffTime / (1000 * 60 * 60 * 24));
    let oldNum = Math.floor(diffDays / 365)
    return oldNum
}
/**
 * 防抖
 * @param fn
 * @param delay
 * @returns {(function(...[*]): void)|*}
 */
export const debounce = (fn, delay = 300) => {
    let timeout

    return (...args) => {
        if (timeout) {
            clearTimeout(timeout)
        }

        timeout = setTimeout(() => {
            fn(...args)
        }, delay)
    }
}

export const compress = async (file, quality) => {
    if (file[0]) {
        return Promise.all(Array.from(file).map(e => compress(e, quality))) // 如果是 file 数组返回 Promise 数组
    } else {
        return new Promise((resolve) => {
            const reader = new FileReader() // 创建 FileReader
            reader.onload = ({target: {result: src}}) => {
                const image = new Image() // 创建 img 元素
                image.onload = async () => {
                    const canvas = document.createElement('canvas') // 创建 canvas 元素
                    canvas.width = image.width
                    canvas.height = image.height
                    canvas.getContext('2d').drawImage(image, 0, 0, image.width, image.height) // 绘制 canvas
                    const canvasURL = canvas.toDataURL('image/jpeg', quality)
                    const buffer = atob(canvasURL.split(',')[1])
                    let length = buffer.length
                    const bufferArray = new Uint8Array(new ArrayBuffer(length))
                    while (length--) {
                        bufferArray[length] = buffer.charCodeAt(length)
                    }
                    const miniFile = new File([bufferArray], file.name, {type: 'image/jpeg'})
                    miniFile.uid = file?.uid || ''
                    resolve({
                        file: miniFile,
                        origin: file,
                        beforeSrc: src,
                        afterSrc: canvasURL,
                        beforeKB: Number((file.size / 1024).toFixed(2)),
                        afterKB: Number((miniFile.size / 1024).toFixed(2))
                    })
                }
                image.src = src
            }
            reader.readAsDataURL(file)
        })
    }
}


// 显示弹框
export const showAddToBulletBox = (type, operateType = null) => {
    if (operateType === 'edit') {
        useState("_dialog").value = true
        useState("_formActionType").value = operateType
    } else if (operateType === 'add') {
        useState("_formActionType").value = operateType
        let data = {orderBy: 0}
        if (type === 'education') userInfoStore.addToEducationForm = data
        if (type === 'experience') userInfoStore.experienceFormData = data
        else if (type === 'jobExperience') userInfoStore.addToJobExperienceForm = data
        else if (type === 'projectExperience') userInfoStore.addToprojectExperience = data
        else if (type === 'volunteerExperience') userInfoStore.volunteerExperienceForm = data
        else if (type === 'customItem') userInfoStore.addToCustomAttributes = data
    }
    useState('_dialogBoxType').value = type
    useState("_dialogBox").value = true
}