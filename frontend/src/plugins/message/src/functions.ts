import { createVNode, render, type VNodeTypes } from 'vue'
import { instances } from './instance'

export function createContainer() {
    let container
    const current = document.getElementById('message-dialogs')

    // 创建/获取消息元素
    if (current) {
        container = current
    } else {
        container = document.createElement('div')
        container.setAttribute(
            'style',
            'position: fixed; top:4rem;  width: 100%; display: flex; flex-direction: column; center; align-items: center;z-index:51;'
        )
        container.setAttribute('key', Date.now().toString())
        container.setAttribute('id', 'message-dialogs')
        document.body.appendChild(container)
    }
    return container
}

export function register(
    container: HTMLElement,
    component: VNodeTypes | ClassComponent | unique symbol
) {
    render(null, container)

    // // 挂载到父元素
    // const VNode = createVNode(component, {
    //     type: props.type,
    //     message: props.message
    // })

    // render(VNode, container.appendChild(child))
}

export function getInstance() {}

export function removeInstance(id: string) {}
