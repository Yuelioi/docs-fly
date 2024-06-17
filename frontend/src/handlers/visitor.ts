import { fetchContent } from './utils'

export const AddVisitorLog = async function (url: string) {
    return await fetchContent(
        '/ip',
        {
            url: url
        },
        'post'
    )
}
