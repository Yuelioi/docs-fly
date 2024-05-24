import { MetaData } from './base'

// 主页统计信息
export class HomeStatistic {
    book_count: number = 0
    document_count: number = 0
    historical_visitor_count: number = 0
    today_visitor_count: number = 0
}

// 搜索选项 分类+书籍[]
export class SearchOption extends MetaData {
    children: MetaData[] = []
}

export class SearchResult {
    search_time: number = 0
    result: SearchData[] = []
}

// 主页搜索显示数据
export class SearchData {
    locale: string = ''

    category_identity: string = ''
    category_display_name: string = ''

    book_identity: string = ''
    book_display_name: string = ''

    chapter_identity: string = ''
    chapter_display_name: string = ''

    section_identity = ''
    section_display_name = ''

    document_identity: string = ''
    document_display_name: string = ''

    content: string = ''
}
