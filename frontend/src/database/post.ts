/**
 * 文章页面数据存储
 * @route /post
 * */

import { dbManager } from './manager'
import type { ChapterInfo } from '@/models'
import type { DBData } from './model'
import type { RouteParams } from 'vue-router'
import { StoreConf } from './model'

const storeName = 'post'

const storeConf = new StoreConf(storeName)
dbManager.dbStores.push(storeConf)
dbManager.dbVersion += 1

export async function getPostChapterData(params: RouteParams) {
    const result = (await dbManager.getDataByKey(
        storeName,
        `chapter__${params['category']}_${params['book']}`
    )) as DBData

    if (result != undefined && Date.now() < result.expiration) {
        return result
    } else {
        return undefined
    }
}

export async function addPostChapterData(params: RouteParams, chapters: ChapterInfo[]) {
    await dbManager.addData(
        storeConf,
        `chapter__${params['category']}_${params['book']}`,
        JSON.parse(JSON.stringify(chapters))
    )
}
