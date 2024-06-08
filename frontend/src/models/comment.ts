export class Comment {
    ip: string = ''
    created_at: Date = new Date()
    nickname: string = ''
    content: string = ''
    parent: number = 0
    replies: Comment[] = []
    comment_type: string = ''
}
