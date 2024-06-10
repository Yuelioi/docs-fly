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

export async function getDBBookData(params: RouteParams) {
    const result = (await dbManager.getDataByKey(
        storeName,
        `book__${(params['slug'] as string[]).join('_')}`
    )) as DBData

    if (result != undefined && Date.now() < result.expiration) {
        return result
    } else {
        return undefined
    }
}

export async function addDBBookData(params: RouteParams, chapters: LocalMetaDatas) {
    await dbManager.addData(
        storeConf,
        `book__${(params['slug'] as string[]).join('_')}`,
        JSON.parse(JSON.stringify(chapters))
    )
}
