import { fetchContent, fetchContentAdmin } from './utils'

// PostPage

/**
 * @description 获取文章markdown内容/html内容/toc
 * @param slug 分类名
 * @param book 书籍名
 * @param chapter 章节名(可为空)
 * @param document 文档名(.md后缀)
 * @returns {
 *      "content_markdown":"",
 *      "content_html":"",
 *      toc"
 * }
 */
export const getPost = async (slug: string, document: string) => {
    return fetchContent('/post', {
        slug: slug,
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
 * @param slug
 * @param document
 * @param content
 * @returns
 */
export const savePost = async (slug: string, document: string, content: string) => {
    return fetchContentAdmin(
        '/post',
        {
            slug: slug,
            document: document,
            content: content
        },
        'post'
    )
}

export const fetchChapter = async (slug: string, document: string) => {
    return await fetchContent('/post/chapter', {
        slug: slug,
        document: document
    })
}
