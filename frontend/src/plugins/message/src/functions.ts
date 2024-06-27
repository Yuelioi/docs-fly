import { createVNode, render, type VNodeTypes } from 'vue'
import { instances } from './instance'

export function register(
    container: HTMLElement,
    component: VNodeTypes | ClassComponent | unique symbol
) {
    render(null, container)

    // 挂载到父元素
    const VNode = createVNode(component, {
        type: props.type,
        message: props.message
    })

    render(VNode, container.appendChild(child))
}

export function getInstance() {}

export function removeInstance(id: string) {}
