// 基础信息

export class MetaData {
    name = ''
    title = ''
    depth = 0
    order = 0
    icon = ''
    status = 1
    filepath = ''
}

export class LocalMetaDatas {
    documents: MetaData[] = []
    categorys: MetaData[] = []
}
