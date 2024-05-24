import { DocumentLinkMeta } from '@/models'
import type { RouteParams } from 'vue-router'

// 补零
export function addZero(num: number, length: number) {
    let str = num.toString()
    while (str.length < length) {
        str = '0' + str
    }
    return str
}

// 创建链接数据 用于router link to
export function createLinkMeta(params: RouteParams) {
    return new DocumentLinkMeta(
        params['category'] as string,
        params['book'] as string,
        params['locale'] as string,
        params['chapter'] as string,
        params['section'] as string,
        params['document'] as string
    )
}

// 生成key 用于router link to 以及 star删除
export function generateKey(params: RouteParams) {
    const linkParams = createLinkMeta(params)
    const keyList = []
    for (const values of Object.entries(linkParams)) {
        keyList.push(values[1])
    }
    return keyList.join('')
}

export function formatDate(date_string: Date) {
    const date = new Date(date_string)

    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, '0') // 月份从0开始，所以需要+1
    const day = String(date.getDate()).padStart(2, '0')
    const hours = String(date.getHours()).padStart(2, '0')
    const minutes = String(date.getMinutes()).padStart(2, '0')
    const seconds = String(date.getSeconds()).padStart(2, '0')

    const formattedDate = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`

    return formattedDate
}
