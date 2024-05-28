import { fetchContent } from './utils'

export async function getBookComment(category: string, book: string) {
    return await fetchContent('/comment', {
        category: category,
        book: book
    })
}

export async function postBookComment(
    category: string,
    book: string,
    nickname: string,
    parent: number = 0
) {
    return await fetchContent(
        '/comment',
        {
            category: category,
            book: book,
            nickname: nickname,
            parent: parent
        },
        'post'
    )
}

export const getDocumentComment = async (
    category: string,
    book: string,
    locale: string,
    chapter: string,
    document: string
) => {
    return fetchContent('/comment', {
        category: category,
        book: book,
        locale: locale,
        chapter: chapter,
        document: document
    })
}

export const postDocumentComment = async (
    category: string,
    book: string,
    locale: string,
    chapter: string,
    document: string,
    parent: number = 0
) => {
    return fetchContent('/comment', {
        category: category,
        book: book,
        locale: locale,
        chapter: chapter,
        document: document,
        parent: parent
    })
}
