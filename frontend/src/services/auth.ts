import { fetchContent } from './utils'

// 登录验证,如果成功则返回token
export const fetchAuthLogin = async (username: string, password: string) => {
    return await fetchContent(
        '/auth/login',
        {
            username: username,
            password: password
        },
        'post'
    )
}

export const fetchCheckToken = async (token: string) => {
    return await fetchContent('/auth/token', {
        token: token
    })
}
