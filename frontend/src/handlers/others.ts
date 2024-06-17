import { fetchContent } from './utils'

export async function getRandPoem() {
    return await fetchContent('/rand/poem')
}

export async function getRandNickname() {
    return await fetchContent('/rand/nickname')
}

export async function getRandPost() {
    return await fetchContent('/rand/post')
}

export async function getAppVersion() {
    return await fetchContent('/app/version')
}
