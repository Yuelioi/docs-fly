import type { LocalMetaDatas } from '@/models/base'
import { fetchContent, fetchContentAdmin } from './utils'

/**
 * Fetches book page statistic information.
 *
 * @param {string} bookPath - The path of the book.
 * @param {string} locale - The locale of the book.
 * @returns {Promise<[boolean, any]>} - A promise resolving to the statistic information.
 *
 * @example
 * const bookPath = 'path/to/book';
 * const locale = 'en';
 * fetchStatisticBook(bookPath, locale).then((statistic) => console.log(statistic));
 */
export const fetchStatisticBook = async (bookPath: string, locale: string) => {
    return await fetchContent('/statistic/book', {
        bookPath: bookPath,
        locale: locale
    })
}

/**
 * Fetches book information.
 *
 * @param {string} path - The path of the book.
 * @param {string} locale - The locale of the book.
 * @returns {Promise<[boolean, any]>} - A promise resolving to the book information.
 *
 * @example
 * const path = 'path/to/book';
 * const locale = 'en';
 * getBookData(path, locale).then((bookData) => console.log(bookData));
 */
export const getBookData = async (path: string, locale: string) => {
    return await fetchContent('/book', {
        bookPath: path,
        locale: locale
    })
}

/**
 * Fetches book metadata.
 *
 * @param {string} bookPath - The path of the book.
 * @param {string} locale - The locale of the book.
 * @returns {Promise<[boolean, any]>} - A promise resolving to the book metadata.
 *
 * @example
 * const bookPath = 'path/to/book';
 * const locale = 'en';
 * getBookMeta(bookPath, locale).then((meta) => console.log(meta));
 */
export const getBookMeta = async (bookPath: string, locale: string) => {
    return await fetchContent('/book/meta', {
        bookPath: bookPath,
        locale: locale
    })
}

/**
 * Saves book metadata to the database.
 *
 * @param {string} bookPath - The path of the book.
 * @param {string} locale - The locale of the book.
 * @param {LocalMetaDatas} metas - The metadata to save.
 * @returns {Promise<[boolean, any]>} - A promise resolving to the saved metadata.
 *
 * @example
 * const bookPath = 'path/to/book';
 * const locale = 'en';
 * const metas: LocalMetaDatas = { title: 'Book Title', author: 'Author Name' };
 * saveBookMeta(bookPath, locale, metas).then((savedMeta) => console.log(savedMeta));
 */
export const saveBookMeta = async (bookPath: string, locale: string, metas: LocalMetaDatas) => {
    return await fetchContentAdmin(
        '/book/meta',
        {
            bookPath: bookPath,
            locale: locale,
            metas: JSON.stringify(metas)
        },
        'put'
    )
}

/**
 * Updates book metadata based on local changes.
 *
 * @returns {Promise<[boolean, any]>} - A promise resolving to the updated metadata.
 *
 * @example
 * updateBookMeta().then((updatedMeta) => console.log(updatedMeta));
 */
export const updateBookMeta = async () => {
    return await fetchContentAdmin('/book/meta', {}, 'post')
}
