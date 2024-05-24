import { fetchContent } from './utils'

// 获取主页统计信息
export const fetchStatisticHome = async () => {
    return await fetchContent('/statistic/home', {})
}
// 获取书籍页统计信息
export const fetchStatisticBook = async (category: string, book: string) => {
    return await fetchContent('/statistic/book', {
        category: category,
        book: book
    })
}
