import { defineStore } from 'pinia'

export const keyStore = defineStore('key', () => {
    function isSearchKey(event: KeyboardEvent) {
        return event.ctrlKey && event.key == '1'
    }
    function isEditKey(event: KeyboardEvent) {
        return event.ctrlKey && event.key == '1'
    }
    function isFullScreenKey(event: KeyboardEvent) {
        return event.ctrlKey && event.key == '2'
    }

    function isSaveKey(event: KeyboardEvent) {
        return event.ctrlKey && event.key == 's'
    }

    return { isSearchKey, isEditKey, isFullScreenKey, isSaveKey }
})
