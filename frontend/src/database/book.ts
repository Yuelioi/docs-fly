/**
 * 书籍页面数据存储
 * @route /book
 * */

import type { LocalMetaDatas } from '@/models'
import type { DBData } from './model'
import type { RouteParams } from 'vue-router'

import { dbManager } from './manager'
import { StoreConf } from './model'
const storeName = 'book'

const storeConf = new StoreConf(storeName)
dbManager.dbStores.push(storeConf)
dbManager.dbVersion += 1

export async function getDBBookData(bookName: string[], locale: string) {
    const result = (await dbManager.getDataByKey(
        storeName,
        `book__${bookName.join('_')}_${locale}`
    )) as DBData

    if (result != undefined && Date.now() < result.expiration) {
        return result
    } else {
        return undefined
    }
}

export async function addDBBookData(bookName: string[], locale: string, chapters: LocalMetaDatas) {
    await dbManager.addData(
        storeConf,
        `book__${bookName.join('_')}_${locale}`,
        JSON.parse(JSON.stringify(chapters))
    )
}
