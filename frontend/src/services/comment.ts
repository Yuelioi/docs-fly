import { fetchContent } from './utils'

export async function getComments(url: string) {
    return await fetchContent('/comment', {
        url: url
    })
}

export async function postComment(data: any) {
    return await fetchContent('/comment', {}, 'post', '', data)
}
