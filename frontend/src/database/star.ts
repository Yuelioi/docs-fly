/**
 * 收藏页面数据存储
 * @route /star
 * */

import type { DBData } from './model'
import type { PostStar } from '@/models/star'
import { StoreConf } from './model'

import { dbManager } from './manager'
const storeName = 'star'

const storeConf = new StoreConf(storeName, false)
dbManager.dbStores.push(storeConf)
dbManager.dbVersion += 1

export async function getPostStarsData() {
    const result = (await dbManager.getAllData(storeName)) as DBData[]
    const stars: PostStar[] = []

    for (const star of result) {
        stars.push(star.data)
    }

    return stars
}

export async function addPostStarData(postStar: PostStar) {
    return await dbManager.addData(storeConf, ``, JSON.parse(JSON.stringify(postStar)))
}

export async function getPostStarData(key: string) {
    return await dbManager.getDataByKey(storeConf.storeName, key)
}

export async function deletePostStarData(key: string) {
    await dbManager.deleteData(storeConf.storeName, key)
}
