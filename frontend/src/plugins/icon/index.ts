import {
    BIconArrowReturnRight,
    BIconBook,
    BIconJournal,
    BIconGraphUpArrow,
    BIconFiletypeDoc,
    BIconArrowsFullscreen,
    BIconFullscreenExit,
    BIconQuestionCircle,
    BIconCaretRight,
    BIconCaretLeft
} from 'bootstrap-icons-vue'

import type { Plugin, App } from 'vue'

import 'primeicons/primeicons.css'

/**
 * 图标管理, 需要手动导入!
 */

// https://icons.getbootstrap.com/#usage
// https://primevue.org/icons

export const IconPlugin: Plugin = {
    install(app: App) {
        app.component('BIconArrowReturnRight', BIconArrowReturnRight)
        app.component('BIconBook', BIconBook)
        app.component('BIconJournal', BIconJournal)
        app.component('BIconGraphUpArrow', BIconGraphUpArrow)
        app.component('BIconFiletypeDoc', BIconFiletypeDoc)
        app.component('BIconArrowsFullscreen', BIconArrowsFullscreen)
        app.component('BIconFullscreenExit', BIconFullscreenExit)
        app.component('BIconQuestionCircle', BIconQuestionCircle)
        app.component('BIconCaretLeft', BIconCaretLeft)
        app.component('BIconCaretRight', BIconCaretRight)
    }
}
