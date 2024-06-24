/**
 * Fetches comments for a given URL.
 *
 * @param {string} url - The URL for which to fetch comments.
 * @returns {Promise<any>} A promise resolving to the fetched comments.
 * @example
 * const comments = await getComments('https://example.com/article');
 */
export async function getComments(url: string) {
    return await fetchContent('/comment', {
        url: url
    })
}

/**
 * Posts a new comment.
 *
 * @param {any} data - The comment data to be posted.
 * @returns {Promise<any>} A promise resolving to the result of the post operation.
 * @example
 * const result = await postComment({
 *   author: 'John Doe',
 *   text: 'This is a sample comment',
 *   url: 'https://example.com/article'
 * });
 */
export async function postComment(data: any) {
    return await fetchContent('/comment', {}, 'post', data)
}
