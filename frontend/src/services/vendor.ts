import { fetchContent } from './utils'

export async function fetchYiYan() {
    return await fetchContent(
        '/',
        {
            c: 'b'
        },
        'get',
        'https://v1.hitokoto.cn'
    )
}
