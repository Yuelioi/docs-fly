// package handlers: 用于管理api接口

import { getBookData, getBookMeta, saveBookMeta, updateBookMeta } from './book'
import { getNav, fetchKeyword, fetchSearchOptions } from './home'
import { fetchChapter, getPost, fetchPostHtml, savePost } from './post'
import { AddVisitorLog } from './visitor'
import { fetchStatisticHome, fetchStatisticBook } from './statistic'
import { fetchAuthLogin, fetchCheckToken } from './auth'

import { fetchYiYan } from './vendor'
import { postBookComment } from './comment'

import { getRandNickname, getRandPoem } from './others'

export {
    getBookData,
    getBookMeta,
    saveBookMeta,
    updateBookMeta,
    fetchStatisticHome,
    fetchStatisticBook,
    getNav,
    fetchKeyword,
    fetchChapter,
    postBookComment,
    fetchPostHtml,
    fetchCheckToken,
    savePost,
    getPost,
    AddVisitorLog,
    fetchYiYan,
    fetchSearchOptions,
    fetchAuthLogin,
    getRandNickname,
    getRandPoem
}
