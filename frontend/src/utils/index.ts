// 补零
export function addZero(num: number, length: number) {
    let str = num.toString()
    while (str.length < length) {
        str = '0' + str
    }
    return str
}

export function getCat(filepath: string): string[] {
    const pathSegments = filepath.split('/') // 按斜杠分割路径
    return pathSegments.filter((segment) => !segment.endsWith('.md'))
}

export function getDocument(filepath: string): string {
    const pathSegments = filepath.split('/') // 按斜杠分割路径

    const filteredDocument = pathSegments.filter((segment) => segment.endsWith('.md')).join('')
    return filteredDocument ? filteredDocument : 'xxxxx'
}

/**
 *  日期转字符串
 * @param date_string
 * @returns
 */
export function formatDate(date_string: Date): string {
    const date = new Date(date_string)

    const year = date.getFullYear()
    const month = String(date.getMonth() + 1).padStart(2, '0') // 月份从0开始，所以需要+1
    const day = String(date.getDate()).padStart(2, '0')
    const hours = String(date.getHours()).padStart(2, '0')
    const minutes = String(date.getMinutes()).padStart(2, '0')
    const seconds = String(date.getSeconds()).padStart(2, '0')

    const formattedDate = `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`

    return formattedDate
}
