import { fetchContent } from './utils'

// BookPage

// 获取书籍信息
export const fetchBook = async (category: string, book: string, locale: string) => {
    return await fetchContent('/book', {
        category: category,
        book: book,
        locale: locale
    })
}
