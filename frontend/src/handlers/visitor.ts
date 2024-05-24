import { fetchContent } from './utils'

import type { RouteParams } from 'vue-router'

export const AddVisitorLog = async function (params: RouteParams, url: string) {
    return await fetchContent(
        '/ip',
        {
            category: params['category'],
            book: params['book'],
            url: url
        },
        'post'
    )
}
