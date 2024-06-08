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

    category_name: string = ''
    category_title: string = ''

    book_name: string = ''
    book_title: string = ''

    chapter_name: string = ''
    chapter_title: string = ''

    section_name = ''
    section_title = ''

    document_name: string = ''
    document_title: string = ''

    content: string = ''
}
