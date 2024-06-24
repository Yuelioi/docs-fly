import { fetchContent, fetchContentAdmin } from './utils'

// PostPage

/**
 * Fetches the post content, HTML, and TOC.
 *
 * @param {object} options
 * @param {string} options.postPath - The path to the post.
 * @param {string} options.document - The document type (e.g. markdown).
 * @returns {Promise<{ content_markdown: string, content_html: string, toc: any }>}
 * @example
 * const result = await getPost({ postPath: 'path/to/post', document: 'markdown' });
 * console.log(result); // { content_markdown: '', content_html: '', toc: [] }
 */
export const getPost = async ({ postPath, document }: { postPath: string; document: string }) => {
    return fetchContent('/post', {
        postPath: postPath,
        document: document
    })
}

/**
 * Converts markdown content to HTML.
 *
 * @param {string} content - The markdown content.
 * @returns {Promise<string>} The HTML content.
 * @example
 * const html = await fetchPostHtml('# Hello World!');
 * console.log(html); // <h1>Hello World!</h1>
 */
export const fetchPostHtml = async (content: string) => {
    return fetchContent('/post', {
        content: content
    })
}

/**
 * Saves a post with the given content.
 *
 * @param {string} postPath - The path to the post.
 * @param {string} content - The post content.
 * @returns {Promise<any>} The result of the save operation.
 * @example
 * await savePost('path/to/post', '# Hello World!');
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

/**
 * Fetches the chapter content for a given post path.
 *
 * @param {string} postPath - The path to the post.
 * @returns {Promise<any>} The chapter content.
 * @example
 * const chapter = await getChapter('path/to/post');
 * console.log(chapter); // { ... }
 */
export const getChapter = async (postPath: string) => {
    return await fetchContent('/post/chapter', {
        postPath: postPath
    })
}
