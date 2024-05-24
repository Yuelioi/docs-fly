import { fetchContent } from './utils'

// HomePage

// 根据关键词获取文章信息
export const fetchKeyword = async (category: string, book: string, keyword: string) => {
    return await fetchContent('/query', { category: category, book: book, keyword: keyword })
}

// 获取顶部导航栏
export const fetchNav = async () => {
    return await fetchContent('/nav', {})
}
// 获取顶部导航栏
export const fetchSearchOptions = async () => {
    return await fetchContent('/search_options', {})
}
