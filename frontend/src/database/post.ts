/**
 * 文章页面数据存储
 * @route /post
 * */

import { dbManager } from './manager'
import type { Chapter } from '@/models'
import type { DBData } from './model'
import type { RouteParams } from 'vue-router'
import { StoreConf } from './model'

const storeName = 'post'

const storeConf = new StoreConf(storeName)
dbManager.dbStores.push(storeConf)
dbManager.dbVersion += 1

export async function getPostChapterData(slug: string) {
    const result = (await dbManager.getDataByKey(storeName, `chapter__${slug}`)) as DBData

    if (result != undefined && Date.now() < result.expiration) {
        return result
    } else {
        return undefined
    }
}

export async function addPostChapterData(slug: string, chapters: Chapter) {
    await dbManager.addData(storeConf, `chapter__${slug}`, JSON.parse(JSON.stringify(chapters)))
}
