import { fetchContent } from './utils'

export async function getRandPoem() {
    return await fetchContent('/rand/poem')
}

export async function getRandNickname() {
    return await fetchContent('/rand/nickname')
}
