import type { LocalMetaDatas } from '@/models'
import { fetchContent, fetchContentAdmin } from './utils'

// BookPage

// 获取书籍页统计信息
export const fetchStatisticBook = async (bookPath: string, locale: string) => {
    return await fetchContent('/statistic/book', {
        bookPath: bookPath,
        locale: locale
    })
}

// 获取书籍信息
export const getBookData = async (path: string, locale: string) => {
    return await fetchContent('/book', {
        bookPath: path,
        locale: locale
    })
}

// 获取书籍元信息
export const getBookMeta = async (bookPath: string, locale: string) => {
    return await fetchContent('/book/meta', {
        bookPath: bookPath,
        locale: locale
    })
}

// 更新数据库数据
export const saveBookMeta = async (bookPath: string, locale: string, metas: LocalMetaDatas) => {
    return await fetchContentAdmin(
        '/book/meta',
        {
            bookPath: bookPath,
            locale: locale,
            metas: JSON.stringify(metas)
        },
        'put'
    )
}

/**
 * 本地修改后, 基于本地更新meta, 会修改数据库
 */
export const updateBookMeta = async () => {
    return await fetchContentAdmin('/book/meta', {}, 'post')
}
