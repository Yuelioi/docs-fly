import type { App, Plugin } from 'vue'
import VDialog from './src/VDialog.vue'

export { VDialog }

export const VDialogPlugin: Plugin = {
    install(app: App) {
        app.component('VDialog', VDialog)
    }
}
