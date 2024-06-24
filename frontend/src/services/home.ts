import { fetchContent } from './utils'

/**
 * Fetches the home page statistics.
 *
 * @returns {Promise<any>} A promise resolving to the home page statistics.
 *
 * @example
 * const stats = await fetchStatisticHome();
 * console.log(stats); // Output: { visits: 100, views: 500,... }
 */
export const fetchStatisticHome = async () => {
    return await fetchContent('/statistic/home')
}

/**
 * Fetches article information based on a keyword.
 *
 * @param {string} bookPath - The path to the book.
 * @param {string} keyword - The keyword to search for.
 * @param {number} page - The page number to fetch.
 * @param {number} pageSize - The number of items per page.
 *
 * @returns {Promise<any>} A promise resolving to the article information.
 *
 * @example
 * const articles = await fetchKeyword('path/to/book', 'javascript', 1, 10);
 * console.log(articles); // Output: [{ title: 'Article 1',... }, { title: 'Article 2',... }]
 */
export const fetchKeyword = async (
    bookPath: string,
    keyword: string,
    page: number,
    pageSize: number
) => {
    return await fetchContent('/query', {
        bookPath: bookPath,
        keyword: keyword,
        page: page,
        pageSize: pageSize
    })
}

/**
 * Fetches the top navigation bar.
 *
 * @returns {Promise<any>} A promise resolving to the navigation bar data.
 *
 * @example
 * const nav = await getNav();
 * console.log(nav); // Output: [{ label: 'Home', url: '/' }, { label: 'About', url: '/about' }]
 */
export const getNav = async () => {
    return await fetchContent('/nav')
}

/**
 * Fetches the search options.
 *
 * @returns {Promise<any>} A promise resolving to the search options data.
 *
 * @example
 * const searchOptions = await fetchSearchOptions();
 * console.log(searchOptions); // Output: [{ label: 'Title', value: 'title' }, { label: 'Author', value: 'author' }]
 */
export const fetchSearchOptions = async () => {
    return await fetchContent('/search_options')
}
