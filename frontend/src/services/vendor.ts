import { fetchContent } from './utils'

export async function fetchYiYan() {
    return await fetchContent('/vendor/yiyan')
}
