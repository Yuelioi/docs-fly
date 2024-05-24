/**
 * 书籍页面数据存储
 * @route /book
 * */

import type { BookData } from '@/models'
import type { DBData } from './model'
import type { RouteParams } from 'vue-router'

import { dbManager } from './manager'
import { StoreConf } from './model'
const storeName = 'book'

const storeConf = new StoreConf(storeName)
dbManager.dbStores.push(storeConf)
dbManager.dbVersion += 1

export async function getBookData(params: RouteParams) {
    const result = (await dbManager.getDataByKey(
        storeName,
        `book__${params['category']}_${params['book']}`
    )) as DBData

    if (result != undefined && Date.now() < result.expiration) {
        return result
    } else {
        return undefined
    }
}

export async function addBookData(params: RouteParams, chapters: BookData) {
    await dbManager.addData(
        storeConf,
        `book__${params['category']}_${params['book']}`,
        JSON.parse(JSON.stringify(chapters))
    )
}
