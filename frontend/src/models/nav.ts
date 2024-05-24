import { MetaData } from './base'

/**
 * 顶部导航信息
 * category_identity: string 分类标识符
 * category_display_name: string

 * book_identity: string
 * book_display_name: string
 **/

export class NavData extends MetaData {
    children: MetaData[] = []
}
