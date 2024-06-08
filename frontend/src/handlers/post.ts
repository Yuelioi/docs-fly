import { fetchContent, fetchContentAdmin } from './utils'

// PostPage

/**
 * @description 获取文章markdown内容/html内容/toc
 * @param category 分类名
 * @param book 书籍名
 * @param chapter 章节名(可为空)
 * @param document 文档名(.md后缀)
 * @returns {
 *      "content_markdown":"",
 *      "content_html":"",
 *      toc"
 * }
 */
export const fetchPost = async (
    category: string,
    book: string,
    locale: string,
    chapter: string,
    document: string
) => {
    return fetchContent('/post', {
        category: category,
        book: book,
        locale: locale,
        chapter: chapter,
        document: document
    })
}

/**
 * @param content markdown内容
 * @returns
 */
export const fetchPostHtml = async (content: string) => {
    return fetchContent('/post', {
        content: content
    })
}

/**
 * @description 保存文章
 * @param category
 * @param book
 * @param chapter
 * @param document
 * @param content
 * @returns
 */
export const savePost = async (category: string, content: string) => {
    return fetchContentAdmin(
        '/post',
        {
            category: category,

            content: content
        },
        'post'
    )
}

export const fetchChapter = async (category: string, book: string, locale: string) => {
    return await fetchContent('/chapter', {
        category: category,
        book: book,
        locale: locale
    })
}
