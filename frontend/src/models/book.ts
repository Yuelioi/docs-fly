import { MetaData } from './base'

export class BookChapter extends MetaData {
    locale: string = ''
    chapter: string = ''
    section: string = ''
    document: string = ''
}

// book 页面数据 书籍+章节信息
export class BookData {
    category: MetaData = new MetaData()
    book: MetaData = new MetaData()
    children: BookChapter[] = []
}
