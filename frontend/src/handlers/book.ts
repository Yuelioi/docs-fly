import type { MetaData } from '@/models'
import { fetchContent, fetchContentAdmin } from './utils'

// BookPage

// 获取书籍信息
export const fetchBook = async (category: string, book: string, locale: string) => {
    return await fetchContent('/book', {
        category: category,
        book: book,
        locale: locale
    })
}

// 获取书籍元信息
export const fetchBookMeta = async (category: string, book: string, locale: string) => {
    return await fetchContent('/book/meta', {
        category: category,
        book: book,
        locale: locale
    })
}

export const saveBookMeta = async (
    category: string,
    book: string,
    locale: string,
    metas: MetaData[]
) => {
    return await fetchContent(
        '/book/meta',
        {
            category: category,
            book: book,
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
    return await fetchContent('/book/meta', {}, 'post')
}
