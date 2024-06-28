import { createVNode, render } from 'vue'

import VMessage from './VMessage.vue'
import MessageContainer from './MessageContainer.vue'

export function registerMessageContainer() {
    const container = document.querySelector('#message-container')
    if (!container) {
        const containerNode = createVNode(MessageContainer)
        containerNode.ref
        containerNode.target
        render(containerNode, document.body)
        return containerNode.target as Element
    }
    return container as Element
}
export function registerMessage(container: Element, props: any) {
    const child = document.createElement('div')
    const VNode = createVNode(VMessage, props)
    render(VNode, container.appendChild(child))
}
