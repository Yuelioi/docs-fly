export class Comment {
    id: number = 0
    createdAt: Date = new Date()
    nickname: string = ''
    content: string = ''
    parent: number = 0
    replies: Comment[] = []
    url: string = ''
}
