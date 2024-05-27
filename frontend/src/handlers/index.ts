// package handlers: 用于管理api接口

import { fetchBook, fetchBookMeta, saveBookMeta, updateBookMeta } from './book'
import { fetchNav, fetchKeyword, fetchSearchOptions } from './home'
import { fetchChapter, fetchPost, fetchPostHtml, savePost } from './post'
import { AddVisitorLog } from './visitor'
import { fetchStatisticHome, fetchStatisticBook } from './statistic'
import { fetchAuthLogin, fetchCheckToken } from './auth'

import { fetchYiYan } from './vendor'

export {
    fetchBook,
    fetchBookMeta,
    saveBookMeta,
    updateBookMeta,
    fetchStatisticHome,
    fetchStatisticBook,
    fetchNav,
    fetchKeyword,
    fetchChapter,
    fetchPostHtml,
    fetchCheckToken,
    savePost,
    fetchPost,
    AddVisitorLog,
    fetchYiYan,
    fetchSearchOptions,
    fetchAuthLogin
}
