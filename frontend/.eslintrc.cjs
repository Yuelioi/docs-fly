/* eslint-env node */
require('@rushstack/eslint-patch/modern-module-resolution')

module.exports = {
    root: true,
    extends: [
        'plugin:vue/vue3-essential',
        'eslint:recommended',
        '@vue/eslint-config-typescript',
        '@vue/eslint-config-prettier/skip-formatting'
    ],
    overrides: [
        {
            files: ['e2e/**/*.{test,spec}.{js,ts,jsx,tsx}'],
            extends: ['plugin:playwright/recommended']
        }
    ],
    parserOptions: {
        ecmaVersion: 'latest'
    },
    rules: {
        'no-unused-vars': 'off'
    },
    env: {
        browser: true,
        es6: true
    }
}
