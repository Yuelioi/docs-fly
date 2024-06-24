import axios from 'axios'

import type { InternalAxiosRequestConfig, AxiosResponse, AxiosRequestConfig } from 'axios'

/**
 * Base URL for API requests
 */
const baseurl = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8088/v1'

/**
 * Axios client instance with default config
 */
const apiClient = axios.create({
    baseURL: baseurl, // API基础URL
    timeout: 1000 // 请求超时时间
})

/**
 * Request interceptor to add token to headers
 */
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

/**
 * Make a request to the API
 *
 * @param query - API endpoint URL
 * @param params - Request parameters
 * @param method - Request method (get, post, put, delete)
 * @param withCookie - Whether to send cookies with the request
 * @param data - Request body data
 * @returns A promise resolving to an array with two elements: [success, data]
 */
const makeRequest = async (
    query: string,
    params: any = {},
    method: 'get' | 'post' | 'put' | 'delete' = 'get',
    withCookie: boolean = false,
    data: any = ''
): Promise<[boolean, any]> => {
    try {
        if (!['get', 'post', 'put', 'delete'].includes(method)) {
            return [false, `无效的请求方法: ${method}`]
        }

        const config: AxiosRequestConfig = {
            method: method,
            url: `${baseurl}${query}`,
            params: params,
            withCredentials: withCookie,
            data: data
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

/**
 * Make a request to the API without sending cookies
 *
 * @param query - API endpoint URL
 * @param params - Request parameters
 * @param method - Request method (get, post, put, delete)
 * @param data - Request body data
 * @returns A promise resolving to an array with two elements: [success, data]
 */
export const fetchContent = (
    query: string,
    params: any = {},
    method: 'get' | 'post' | 'put' | 'delete' = 'get',
    data: any = ''
): Promise<[boolean, any]> => {
    return makeRequest(query, params, method, false, data)
}

/**
 * Make a request to the API with sending cookies
 *
 * @param query - API endpoint URL
 * @param params - Request parameters
 * @param method - Request method (get, post, put, delete)
 * @returns A promise resolving to an array with two elements: [success, data]
 */
export const fetchContentAdmin = (
    query: string,
    params: any = '',
    method: 'get' | 'post' | 'put' | 'delete' = 'get'
): Promise<[boolean, any]> => {
    return makeRequest(query, params, method, true)
}

import type { Ref } from 'vue'

/**
 * Fetch data from API and update a ref value
 *
 * @param refValue - Ref value to update
 * @param defaultValue - Default value to set if request fails
 * @param fetchFunction - Fetch function to use (e.g. fetchContent or fetchContentAdmin)
 * @param params - Request parameters
 * @param prop - Property to extract from response data
 */
export async function fetchHandleBasic(
    refValue: Ref<any>,
    defaultValue: any,
    fetchFunction: any,
    params: any = {},
    prop: string = 'data'
) {
    const [ok, data] = await fetchFunction(params)
    if (ok) {
        refValue.value = data[prop]
    } else {
        refValue.value = defaultValue
    }
}

/**
 * Fetch data from API, update a ref value, and call success/error callbacks
 *
 * @param refValue - Ref value to update
 * @param defaultValue - Default value to set if request fails
 * @param fetchFunction - Fetch function to use
 * @param params - Request parameters
 * @param prop - Property to extract from response data
 * @param success_callback - Callback to call on success
 * @param error_callback - Callback to call on error
 */
export async function fetchHandleBasicCallback(
    refValue: Ref<any>,
    defaultValue: any,
    fetchFunction: any,
    params: any = {},
    prop: string = 'data',
    success_callback: () => Promise<void>,
    error_callback: () => Promise<void> = async () => {}
) {
    const [ok, data] = await fetchFunction(params)
    if (ok) {
        ;(refValue.value = data[prop]), await success_callback
    } else {
        ;(refValue.value = defaultValue), await error_callback
    }
}
