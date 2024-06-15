import axios from 'axios'

import type { InternalAxiosRequestConfig, AxiosResponse, AxiosRequestConfig } from 'axios'

export const baseurl = 'http://localhost:8088/api/v1'

const apiClient = axios.create({
    baseURL: baseurl, // API基础URL
    timeout: 1000 // 请求超时时间
})

// 请求拦截器
apiClient.interceptors.request.use(
    (config: InternalAxiosRequestConfig) => {
        const token = localStorage.getItem('token')
        // 如果token存在，则添加到请求头
        if (token) {
            config.headers.Authorization = `Bearer ${token}`
        }
        return config
    },
    (error: any) => {
        // 错误处理
        return Promise.reject(error)
    }
)

// 通用请求函数
const makeRequest = async (
    query: string,
    params: any,
    method: 'get' | 'post' | 'put' | 'delete' = 'get',
    base: string = '',
    withCookie: boolean = false
): Promise<[boolean, any]> => {
    try {
        if (!['get', 'post', 'put', 'delete'].includes(method)) {
            return [false, `无效的请求方法: ${method}`]
        }

        const config: AxiosRequestConfig = {
            method: method,
            url: base ? `${base}${query}` : `${baseurl}${query}`,
            params: params,
            withCredentials: withCookie
        }
        let response: AxiosResponse
        if (withCookie) {
            response = await apiClient(config)
        } else {
            response = await axios(config)
        }

        if (response.status >= 200 && response.status < 300) {
            return [true, response.data]
        } else {
            return [false, response.data]
        }
    } catch (error: any) {
        let errorMessage: string

        if (error.response) {
            errorMessage = `请求失败，状态码: ${error.response.status}, 响应: ${error.response.data}`
        } else if (error.request) {
            errorMessage = '请求已发出，但没有收到服务器响应'
        } else {
            errorMessage = `请求出错: ${error.message}`
        }

        return [false, errorMessage]
    }
}

// 不需要cookie的请求
export const fetchContent = (
    query: string,
    params: any = '',
    method: 'get' | 'post' | 'put' | 'delete' = 'get',
    base: string = ''
): Promise<[boolean, any]> => {
    return makeRequest(query, params, method, base, false)
}

// 需要cookie的请求
export const fetchContentAdmin = (
    query: string,
    params: any = '',
    method: 'get' | 'post' | 'put' | 'delete' = 'get',
    base: string = ''
): Promise<[boolean, any]> => {
    return makeRequest(query, params, method, base, true)
}
