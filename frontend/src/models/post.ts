import { MetaData } from './base'

export class Toc {
    id: string = ''
    depth: number = 0
    title: string = ''
}

export class Chapter {
    metadata: MetaData = new MetaData()
    documents: MetaData[] = []
    children: Chapter[] = []
}

export class ChapterData extends Chapter {
    collapsed: boolean = false // 前端独立属性
    ref: HTMLElement = new HTMLElement()
    id: number = 0
}
