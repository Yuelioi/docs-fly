// 基础信息
export class MetaData {
    identity = ''
    display_name = ''
    order = 0
    icon = ''
    hidden = false
}

/**
 * 文章链接参数, 用于直接route link
 */
export class DocumentLinkMeta {
    category = ''
    book = ''
    locale = ''
    chapter = ''
    section = ''
    document = ''
    constructor(category = '', book = '', locale = '', chapter = '', section = '', document = '') {
        this.category = category
        this.book = book
        this.locale = locale
        this.chapter = chapter
        this.section = section
        this.document = document
    }
}
