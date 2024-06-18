/**
 * 文章页面数据存储
 * @route /post
 * */

import { dbManager } from './manager'
import type { Chapter } from '@/models/post'
import type { DBData } from './model'

import { StoreConf } from './model'

const storeName = 'post'

const storeConf = new StoreConf(storeName)
dbManager.dbStores.push(storeConf)
dbManager.dbVersion += 1

export async function getPostChapterData(postPath: string) {
    const result = (await dbManager.getDataByKey(storeName, `chapter__${postPath}`)) as DBData

    if (result != undefined && Date.now() < result.expiration) {
        return result
    } else {
        return undefined
    }
}

export async function addPostChapterData(postPath: string, chapters: Chapter) {
    await dbManager.addData(storeConf, `chapter__${postPath}`, JSON.parse(JSON.stringify(chapters)))
}
