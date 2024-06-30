import { createVNode, render } from 'vue'

import VMessage from './VMessage.vue'
import MessageContainer from './MessageContainer.vue'

import { messageContainer } from './model'

export function registerMessageContainer() {
    const container = document.querySelector('#message-container') as HTMLElement
    if (!container) {
        const containerNode = createVNode(MessageContainer)
        const div = document.createElement('div')
        document.body.appendChild(div)
        render(containerNode, div)
        messageContainer.value = div.firstElementChild as HTMLElement
    } else {
        messageContainer.value = container
    }
}
export function registerMessage(props: any) {
    if (messageContainer.value) {
        const child = document.createElement('div')
        const VNode = createVNode(VMessage, props)
        render(VNode, child)
        messageContainer.value.appendChild(child)
    } else {
        console.log('Message container is not registered.')
    }
}
