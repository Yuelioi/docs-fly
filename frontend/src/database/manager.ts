import { StoreConf, DBData } from './model'
import { DBConf } from './model'

/**
 * 创建一个新的数据库配置
 */
const dbConf = new DBConf()

/**
 * 管理浏览器数据库
 */
class IndexDBManager {
    dbName: string = dbConf.dbNameConf
    dbVersion: number = dbConf.dbVersionConf
    dbStores: StoreConf[] = dbConf.dbStoresConf
    db: IDBDatabase | null = null
    request: IDBOpenDBRequest | null = null

    /**
     * 初始化数据库
     * key: 索引值, 请用 用途__具体索引指定, 比如 chapter__AE_EFFECTS 意思是AE_EFFECTS的章节
     */
    initDB() {
        return new Promise((resolve, reject) => {
            const indexDB = window.indexedDB

            this.request = indexDB.open(this.dbName, this.dbVersion)

            this.request.onupgradeneeded = (event: IDBVersionChangeEvent) => {
                this.db = (event.target as IDBOpenDBRequest).result

                // 在初始化时 "注册"所有store
                for (const storeConf of dbConf.dbStoresConf) {
                    if (!this.db.objectStoreNames.contains(storeConf.storeName)) {
                        const objectStore = this.db.createObjectStore(storeConf.storeName, {
                            keyPath: 'key',
                            autoIncrement: true
                        })
                        objectStore.createIndex('key', 'key', { unique: storeConf.unique })
                    }
                }
            }

            this.request.onsuccess = async (event: Event) => {
                this.db = (event.target as IDBOpenDBRequest).result
                resolve(this.db)
            }

            this.request.onerror = () => {
                reject(`数据库打开错误`)
            }
        })
    }
    /**
     * key: 索引值, 请用 用途__具体索引指定, 比如 chapter__AE_EFFECTS 意思是AE_EFFECTS的章节
     * data: 具体写入的内容,可以是对象
     *
     * @description 如果需要创建非unique 传入的data需要有 key 属性
     */
    async addData(storeConf: StoreConf, key: string, data: any) {
        await this.initDB()

        if (!this.db) {
            return Promise.reject('Database is not initialized')
        }

        try {
            const transaction = this.db.transaction([storeConf.storeName], 'readwrite')
            const objectStore = transaction.objectStore(storeConf.storeName)

            let dataObj
            if (storeConf.unique) {
                dataObj = new DBData(key, 24 * 30, data)
            } else {
                dataObj = new DBData(data['key'], 24 * 30, data)
            }

            const request = objectStore.put(dataObj)

            request.onsuccess = () => {
                transaction.commit() // 提交事务
            }

            request.onerror = () => {
                transaction.abort() // 终止事务
                console.error('数据添加或更新失败')
            }
        } catch (error) {
            console.error('发生错误:', error)
        }
    }

    /**
     * 通过主键读取数据,数据存储在 result['data]
     * key: 索引值, 请用 用途__具体索引指定, 比如 chapter__AE_EFFECTS 意思是AE_EFFECTS的章节
     */
    async getDataByKey(storeName: string, key: string) {
        await this.initDB()

        return new Promise((resolve, reject) => {
            if (!this.db) {
                reject('Database is not initialized')
                return
            }

            const transaction = this.db.transaction([storeName], 'readonly')
            const objectStore = transaction.objectStore(storeName)

            const request = objectStore.get(key)

            request.onsuccess = () => {
                resolve(request.result)
            }

            request.onerror = (event) => {
                reject(event)
            }
        })
    }

    async deleteData(storeName: string, key: string) {
        await this.initDB()

        return new Promise((resolve, reject) => {
            if (!this.db) {
                reject('Database is not initialized')
                return
            }

            const transaction = this.db.transaction([storeName], 'readwrite')
            const objectStore = transaction.objectStore(storeName)
            console.log(key)

            const request = objectStore.delete(key)

            request.onsuccess = () => {
                resolve('Data deleted successfully')
            }

            request.onerror = (event) => {
                reject(event)
            }
        })
    }

    async getAllData(storeName: string) {
        await this.initDB()

        return new Promise((resolve, reject) => {
            if (!this.db) {
                reject('Database is not initialized')
                return
            }

            const transaction = this.db.transaction([storeName], 'readonly')
            const objectStore = transaction.objectStore(storeName)
            const request = objectStore.getAll()

            request.onsuccess = () => {
                resolve(request.result)
            }

            request.onerror = (event) => {
                reject(event)
            }
        })
    }
}

const dbManager = new IndexDBManager()

export { dbManager }
