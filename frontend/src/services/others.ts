/**
 * Fetches a random poem from the server.
 *
 * @returns {Promise<string>} A promise that resolves with a random poem.
 *
 * @example
 * const poem = await getRandPoem();
 * console.log(poem); // Output: A random poem
 */
export async function getRandPoem() {
    return await fetchContent('/rand/poem')
}

/**
 * Fetches a random nickname from the server.
 *
 * @returns {Promise<string>} A promise that resolves with a random nickname.
 *
 * @example
 * const nickname = await getRandNickname();
 * console.log(nickname); // Output: A random nickname
 */
export async function getRandNickname() {
    return await fetchContent('/rand/nickname')
}

/**
 * Fetches a random post from the server.
 *
 * @returns {Promise<string>} A promise that resolves with a random post.
 *
 * @example
 * const post = await getRandPost();
 * console.log(post); // Output: A random post
 */
export async function getRandPost() {
    return await fetchContent('/rand/post')
}

/**
 * Fetches the current application version from the server.
 *
 * @returns {Promise<string>} A promise that resolves with the current application version.
 *
 * @example
 * const version = await getAppVersion();
 * console.log(version); // Output: The current application version
 */
export async function getAppVersion() {
    return await fetchContent('/app/version')
}
