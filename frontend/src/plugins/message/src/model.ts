// import { createVNode, render } from 'vue'
// import VNotification from './VMessage.vue'

import type { ExtractPropTypes } from 'vue'

export const messageTypes = ['success', 'info', 'warn', 'error', 'secondary', 'contrast'] as const
export type messageType = (typeof messageTypes)[number]

export const messageDefaults = {
    message: '',
    type: 'info',
    duration: 3000,
    showClose: true
} as const

export const messageProps = {
    message: {
        type: [String, Object, Function] as PropType<string | VNode | (() => VNode)>,
        default: messageDefaults.message
    },
    type: {
        type: String,
        values: messageTypes,
        default: messageDefaults.type
    },
    duration: {
        type: Number,
        default: messageDefaults.duration
    },
    showClose: {
        type: Boolean,
        default: messageDefaults.showClose
    }
}

export type MessageProps = ExtractPropTypes<typeof messageProps>
