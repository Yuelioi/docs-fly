import { fetchContent, fetchContentAdmin } from './utils'

// PostPage

/**
 * @description 获取文章markdown内容/html内容/toc
 * @param postPath 书籍路径
 * @returns {
 *      "content_markdown":"",
 *      "content_html":"",
 *      toc"
 * }
 */
export const getPost = async (postPath: string, document: string) => {
    return fetchContent('/post', {
        postPath: postPath,
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

 * @returns
 */
export const savePost = async (postPath: string, content: string) => {
    return fetchContentAdmin(
        '/post',
        {
            postPath: postPath,
            content: content
        },
        'post'
    )
}

export const getChapter = async (postPath: string) => {
    return await fetchContent('/post/chapter', {
        postPath: postPath
    })
}
