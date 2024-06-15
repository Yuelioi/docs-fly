import { MetaData } from './base'

export class BookData {
    url = ''
    is_dir = false
    metadata: MetaData = new MetaData()
}

export class BookStatistic {
    bookTitle = ''
    readCount = 0
    chapterCount = 0
    documentCount = 0
}
