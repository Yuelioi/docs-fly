import { createVNode, render } from 'vue'
import VMessage from './VMessage.vue'

/**
 * @name: 消息提示插件
 * 样式参考: https://primevue.org/message/ 不过用的是tailwind自带的
 * 本来用的element-ui的 但是好丑
 */

export type MessageType = 'success' | 'info' | 'warn' | 'error' | 'secondary' | 'contrast'

/**
 *  @property:message: string 消息内容
 *  @property:type: MessageType 5种显示类型
 *  @property:duration: number 消息持续时间(ms)
 */
export function Message(message: string, severity: MessageType = 'info', duration: number = 2000) {
    let mask
    const current = document.getElementById('message-dialog')

    // 创建/获取消息元素
    if (current) {
        mask = current
    } else {
        mask = document.createElement('div')
        mask.setAttribute(
            'style',
            'position: fixed; top:4rem;  width: 100%; display: flex; flex-direction: column; center; align-items: center'
        )
        mask.setAttribute('key', Date.now().toString())
        mask.setAttribute('id', 'message-dialog')
        document.body.appendChild(mask)
    }

    // 创建子元素
    const child = document.createElement('div')
    child.setAttribute('style', 'margin-top:1rem')

    // 定时删除子元素
    setTimeout(() => {
        render(null, mask)
        child.remove()
    }, duration)

    // 挂载到父元素
    const VNode = createVNode(VMessage, {
        severity: severity,
        message: message
    })

    render(VNode, mask.appendChild(child))
}
