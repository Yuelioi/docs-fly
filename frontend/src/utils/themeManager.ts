import { reactive } from 'vue'

/**
 * Formats a Date object into a string in the format "YYYY-MM-DD HH:mm:ss".
 *
 * @param date_string The Date object to format.
 * @returns A string representation of the date in the format "YYYY-MM-DD HH:mm:ss".
 *
 * Example: `formatDate(new Date('2022-07-25T14:30:00.000Z'))` returns `"2022-07-25 14:30:00"`
 */
export const themeState = reactive({
    currentTheme: 'default', // 初始主题
    availableThemes: [] as string[] // 可用主题列表
})

/**
 * Sets the current theme.
 *
 * @param {string} theme - The theme to set as the current theme.
 * @example
 * setTheme('dark'); // Sets the current theme to 'dark'
 */
export function setTheme(theme: string) {
    themeState.currentTheme = theme
}
