/**
 * 主页数据存储
 * @route /
 * */

import type { SearchOption, Nav } from '@/models'
import type { DBData } from './model'
import { StoreConf } from './model'

import { dbManager } from './manager'

const storeName = 'home'

const storeConf = new StoreConf(storeName)
dbManager.dbStores.push(storeConf)
dbManager.dbVersion += 1

export async function getDBSearchOption() {
    const result = (await dbManager.getDataByKey(storeName, `search_option`)) as DBData

    if (result != undefined && Date.now() < result.expiration) {
        return result
    } else {
        return undefined
    }
}
export async function addDBSearchOption(data: SearchOption[]) {
    await dbManager.addData(storeConf, `search_option`, JSON.parse(JSON.stringify(data)))
}

export async function getDBNav() {
    const result = (await dbManager.getDataByKey(storeName, `nav`)) as DBData
    if (result != undefined && Date.now() < result.expiration) {
        return result
    } else {
        return undefined
    }
}

export async function addDBNav(datas: Nav[]) {
    await dbManager.addData(storeConf, `nav`, JSON.parse(JSON.stringify(datas)))
}
