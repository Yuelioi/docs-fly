import { MetaData } from './base'

// 后端拿到的初始数据
export class ChapterInfo {
    category: MetaData = new MetaData()
    chapter: MetaData = new MetaData()
    sections: MetaData[] = []
    document: MetaData = new MetaData()
    documents: MetaData[] = []
}

// 前端用于展示的数据
export class ChapterData extends ChapterInfo {
    collapsed: boolean = false // 前端独立属性
    ref: HTMLElement = new HTMLElement()
    id: number = 0
}

export class Toc {
    id: string = ''
    depth: number = 0
    title: string = ''
}
