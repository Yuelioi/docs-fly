// handlers: 管理api接口
export { fetchStatisticHome, getNav, fetchKeyword, fetchSearchOptions } from './home'
export { fetchYiYan } from './vendor'
export { fetchStatisticBook, getBookData, getBookMeta, saveBookMeta, updateBookMeta } from './book'
export { getChapter, getPost, fetchPostHtml, savePost } from './post'

export { AddVisitorLog } from './visitor'
export { fetchAuthLogin, fetchCheckToken } from './auth'

export { postBookComment } from './comment'
export { getRandNickname, getRandPoem } from './others'
