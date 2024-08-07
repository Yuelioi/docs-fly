// 基础信息

export class MetaData {
    name = ''
    title = ''
    depth = 0
    order = 0
    icon = ''
    is_dir = true
    status = 1
    url = ''
}

export class LocalMetaDatas {
    documents: MetaData[] = []
    categories: MetaData[] = []
}

export class ResponseData {
    clientTime: Date = new Date()
    ip: string = ''
    serverTime: Date = new Date()
    statusCode: number = 0
    data: any = null
}
