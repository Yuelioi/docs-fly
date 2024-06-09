/**
 * 数据库数据模型
 */
export class DBData {
    /**
     * 键名, 用于索引
     */
    key: string = ''

    /**
     * 创建时间
     */
    createdTime: number = Date.now()

    /**
     * 过期时间
     */
    expiration: number

    /**
     * 数据, 一般为json对象
     */
    data: any
    constructor(key: string, expiration: number, data: any) {
        this.key = key
        this.expiration = this.createdTime + expiration * (60 * 60 * 1000)
        this.data = data
    }
}

/**
 * 储存库配置
 * @property storeName: 库名
 * @property unique: 键名是否唯一, 不用管, 后续可以使用时间戳作为键
 */
export class StoreConf {
    storeName: string = ''
    unique: boolean = true
    constructor(storeName: string, unique: boolean = true) {
        this.storeName = storeName
        this.unique = unique
    }
}

export class DBConf {
    /**
     * 数据库名
     */
    dbNameConf = 'docs'

    /**
     * 数据库版本号
     */
    dbVersionConf = 1

    /**
     * 数据库表名
     */
    dbStoresConf: StoreConf[] = []

    /**
     * 过期时间 (小时)
     */
    dbDataExpiration = 24
}
