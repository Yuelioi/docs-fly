/**
 * @name: 消息提示插件
 * 样式参考: https://primevue.org/message/ 不过用的是tailwind自带的
 * 本来用的element-ui的 但是好丑
 */

import type { MessageProps } from './model'

import { messageDefaults } from './model'
import { registerMessageContainer, registerMessage } from './functions'

export function Message(msgProps: Partial<MessageProps>) {
    const props = {
        message: msgProps.message ?? messageDefaults.message,
        type: msgProps.type ?? messageDefaults.type,
        duration: msgProps.duration ?? messageDefaults.duration,
        showClose: msgProps.showClose ?? messageDefaults.showClose
    }

    registerMessageContainer()
    registerMessage(props)
}
