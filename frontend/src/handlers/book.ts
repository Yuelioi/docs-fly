import type { LocalMetaDatas } from '@/models'
import { fetchContent, fetchContentAdmin } from './utils'

// BookPage

// 获取书籍信息
export const getBookData = async (slug: string, locale: string) => {
    return await fetchContent('/book', {
        slug: slug,
        locale: locale
    })
}

// 获取书籍元信息
export const getBookMeta = async (slug: string, locale: string) => {
    return await fetchContent('/book/meta', {
        slug: slug,
        locale: locale
    })
}

export const saveBookMeta = async (slug: string, locale: string, metas: LocalMetaDatas) => {
    return await fetchContent(
        '/book/meta',
        {
            slug: slug,
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
