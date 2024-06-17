import type { Ref } from 'vue'

// 补零
export function addZero(num: number, length: number) {
    let str = num.toString()
    while (str.length < length) {
        str = '0' + str
    }
    return str
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

// 基础异步请求api 并根据状态赋值/初始化
export async function fetchBasic(
    refValue: Ref<any>,
    defaultValue: any,
    fetchFunction: any,
    params: any = {},
    prop: string = 'data'
) {
    const [ok, data] = await fetchFunction(params)
    if (ok) {
        refValue.value = data[prop]
    } else {
        refValue.value = defaultValue
    }
}

// export async function fetchHandler(
//     refValue: Ref<any>,
//     defaultValue: any,
//     fetchFunction: (params:any) => Promise<[boolean, any]>,

//     prop: string = 'data',
//     successMessageHandler: () => Promise<void>,
//     failureMessageHandler: () => Promise<void>
// ) {
//     const [ok, data] = await fetchFunction()
//     if (ok) {
//         refValue.value = data[prop]
//         await successMessageHandler()
//     } else {
//         refValue.value = defaultValue
//         await failureMessageHandler()
//     }
// }
