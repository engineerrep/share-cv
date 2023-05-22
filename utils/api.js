import {useCookie} from "nuxt/app";
import {createDiscreteApi} from "naive-ui";
import {useRouter} from "vue-router";
import {useToken} from "~/composables/user";

const {message} = createDiscreteApi(["message"]);
const router = useRouter();
const fetch = (path, options, mimeType) => {
    const config = useRuntimeConfig()
    const baseUrl = config.public.VITE_API_URL
    const url = baseUrl + path
    const token = useToken().value || ''
    const lang = useCookie("lang")?.value || 'en'
    let headers = {
        'Content-Type': 'application/json',
        'Content-Language': (lang === 'cn' ? 'zh' : lang),
        'Authorization': token
    }
    if (mimeType) {
        headers = {
            'Content-Type': 'multipart/form-data',
            'Content-Language': (lang === 'cn' ? 'zh' : lang),
            'Authorization': token
        }
    }

    const reqConfig = {
        headers,
        onRequest({request, options}) {

        },
        onResponse({request, options, response}) {
            // console.log(response);
        }
    }

    return new Promise((resolve, reject) => {
        useFetch(url, {...options, ...reqConfig}).then(({data, error}) => {
            if (error.value) {
                reject(error.value)
                return
            }
            const value = data.value

            if (value?.code == 1010) {

                return
                // TODO: - 做注销登录的逻辑处理
                // const token = useCookie("token")
                // token.value = ''
                // document.cookie = 'token=; expires=Thu, 01 Jan 1970 00:00:00 UTC; path=/;'
                useToken().value = ''
                return navigateTo('/')
            } else if (value?.code !== 200) {
                // TODO: - 列表接口会报错记录问题，和找到问题
                console.error(url)
                message.error(value?.msg);
                // if(nc) nc.reset()
                reject(value)
            } else {
                resolve(value)
            }
        }).catch((err) => {
            console.log(err)
        })
    })
}

export default new class Api {

    get(path, params) {
        return fetch(path, {method: 'get', params})
    }

    post(path, body, mimeType) {
        return fetch(path, {method: 'post', body}, mimeType)
    }
}