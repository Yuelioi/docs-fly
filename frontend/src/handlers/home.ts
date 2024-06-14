import { fetchContent } from './utils'

// HomePage

// 获取主页统计信息
export const fetchStatisticHome = async () => {
    return await fetchContent('/statistic/home')
}

// 根据关键词获取文章信息
export const fetchKeyword = async (category: string, book: string, keyword: string) => {
    return await fetchContent('/query', { category: category, book: book, keyword: keyword })
}

// 获取顶部导航栏
export const getNav = async () => {
    return await fetchContent('/nav', {})
}
// 获取顶部导航栏
export const fetchSearchOptions = async () => {
    return await fetchContent('/search_options', {})
}
