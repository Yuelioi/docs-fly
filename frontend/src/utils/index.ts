/**
 * Adds leading zeros to a number to reach a specified length.
 *
 * @param num The number to add leading zeros to.
 * @param length The desired length of the resulting string.
 * @returns A string representation of the number with leading zeros.
 *
 * Example: `addZero(5, 3)` returns `"005"`
 */
export function addZero(num: number, length: number) {
    let str = num.toString()
    while (str.length < length) {
        str = '0' + str
    }
    return str
}

/**
 * Formats a Date object into a string in the format "YYYY-MM-DD HH:mm:ss".
 *
 * @param date_string The Date object to format.
 * @returns A string representation of the date in the format "YYYY-MM-DD HH:mm:ss".
 *
 * Example: `formatDate(new Date('2022-07-25T14:30:00.000Z'))` returns `"2022-07-25 14:30:00"`
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
