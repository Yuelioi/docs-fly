import axios from 'axios'

export const baseurl = 'http://localhost:8088'

const apiClient = axios.create({
    baseURL: baseurl, // 你的API基础URL
    timeout: 1000 // 请求超时时间
})

// 请求拦截器
apiClient.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token')

        // 如果token存在，则添加到请求头
        if (token) {
            config.headers.Authorization = `Bearer ${token}`
        }

        return config
    },
    (error) => {
        // 错误处理
        return Promise.reject(error)
    }
)

export const fetchContent = async (
    query: string,
    params: any,
    method: string = 'get',
    base: string = ''
) => {
    try {
        if (['get', 'post', 'put', 'delete'].includes(method)) {
            const response = await axios({
                method: method,
                url: base ? `${base}${query}` : `${baseurl}${query}`,
                params: params
            })

            return [true, response.data]
        } else {
            return [false, `无效的请求方法: ${method}`]
        }
    } catch (error) {
        return [false, error]
    }
}

export const fetchContentAdmin = async (
    query: string,
    params: any,
    method: string = 'get',
    base: string = ''
) => {
    try {
        if (['get', 'post', 'put', 'delete'].includes(method)) {
            const response = await apiClient({
                method: method,
                url: base ? `${base}${query}` : `${baseurl}${query}`,
                params: params
            })

            return [true, response.data]
        } else {
            return [false, `无效的请求方法: ${method}`]
        }
    } catch (error) {
        return [false, error]
    }
}
