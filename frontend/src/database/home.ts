/**
 * 主页数据存储
 * @route /
 * */

import type { SearchOption, LocalMetaDatas } from '@/models'
import type { DBData } from './model'
import { StoreConf } from './model'

import { dbManager } from './manager'

const storeName = 'home'

const storeConf = new StoreConf(storeName)
dbManager.dbStores.push(storeConf)
dbManager.dbVersion += 1

export async function getSearchOption() {
    const result = (await dbManager.getDataByKey(storeName, `search_option`)) as DBData

    if (result != undefined && Date.now() < result.expiration) {
        return result
    } else {
        return undefined
    }
}
export async function addSearchOption(data: SearchOption[]) {
    await dbManager.addData(storeConf, `search_option`, JSON.parse(JSON.stringify(data)))
}

export async function getNav() {
    const result = (await dbManager.getDataByKey(storeName, `nav`)) as DBData
    if (result != undefined && Date.now() < result.expiration) {
        return result
    } else {
        return undefined
    }
}

export async function addNav(datas: LocalMetaDatas[]) {
    await dbManager.addData(storeConf, `nav`, JSON.parse(JSON.stringify(datas)))
}
