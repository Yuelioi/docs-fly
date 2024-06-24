import { fetchContent } from './utils'

/**
 * Fetches an authentication token by logging in with a username and password.
 *
 * @param {string} username - The username to log in with.
 * @param {string} password - The password to log in with.
 * @returns {Promise<string>} A promise that resolves with the authentication token if the login is successful.
 * @example
 * const token = await fetchAuthLogin('johnDoe', 'mySecretPassword');
 */
export const fetchAuthLogin = async (username: string, password: string) => {
    return await fetchContent(
        '/auth/login',
        {
            username: username,
            password: password
        },
        'post'
    )
}
/**
 * Checks if a given token is valid.
 *
 * @param {string} token - The token to check.
 * @returns {Promise<boolean>} A promise that resolves with `true` if the token is valid, `false` otherwise.
 * @example
 * const isValid = await fetchCheckToken('eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaGFuIjoiMjMwfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c');
 * console.log(isValid); // Output: true
 */
export const fetchCheckToken = async (token: string) => {
    return await fetchContent('/auth/token', {
        token: token
    })
}
