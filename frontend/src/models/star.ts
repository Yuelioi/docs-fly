/**
 * 收藏文章
 */
export class PostStar {
    key: string = Date.now().toString()
    slug: string[] = []
    document: string = ''
    link: string = ''
    createdTime: Date = new Date() // 收藏时间
    mark: string = '' // 批注
    params = ''
}
