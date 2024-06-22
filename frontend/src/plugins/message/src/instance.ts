/**
 * @name: 消息提示插件
 * 样式参考: https://primevue.org/message/ 不过用的是tailwind自带的
 * 本来用的element-ui的 但是好丑
 */
import { createVNode, render } from 'vue'
import VMessage from './VMessage.vue'
import type { MessageProps } from './model'
import { createContainer } from './functions'

type MessageContext = {
    id: string
}

export function Message(props: Partial<MessageProps>) {
    return new Promise((resolve) => {
        // 创建子元素
        const child = document.createElement('div')
        child.setAttribute('style', 'margin-top:1rem')

        const container = createContainer()
    })
}

export const instances: MessageContext[] = shallowReactive([])
