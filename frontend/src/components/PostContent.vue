<template>
    <div
        id="content"
        ref="fullElement"
        class="max-w-6xl m-auto relative mb-12 bg-white-base dark:bg-dark-light p-16 rounded-2xl">
        <div
            class="toolbar py-4 fixed flex-col top-16 right-[2rem] lg:right-[17rem] xl:right-[21rem]">
            <div class="p-2">
                <BIconArrowsFullscreen v-if="!isFullscreen" @click="enter"></BIconArrowsFullscreen>
                <BIconFullscreenExit v-else @click="exit"></BIconFullscreenExit>
            </div>
            <div class="p-2">
                <i class="pi pi-star" v-if="!isStared" @click="starPost"></i>
                <i v-else class="pi pi-star-fill" @click="unStarPost"></i>
            </div>
            <div v-if="isAdmin" class="flex flex-col p-2">
                <i class="pi pi-file-edit" v-if="!isEditing" @click="isEditing = !isEditing"></i>
                <i class="pi pi-file" v-else @click="isEditing = !isEditing"></i>
            </div>
            <div v-if="isAdmin && isEditing" class="flex flex-col p-2">
                <i class="pi pi-save" @click="save"></i>
            </div>
            <div class="p-2"><BIconQuestionCircle></BIconQuestionCircle></div>
        </div>
        <!-- 显示模式/编辑模式 -->
        <div class="">
            <div
                v-show="!isEditing"
                ref="postContainer"
                class="main sm:mx-[1rem] md:mx-[2rem] lg:mx-[2.5rem] xl:mx-[3.5rem] 2xl:mx-[5rem]"
                v-html="postHtml"></div>
            <div v-show="isEditing" id="vditor" class="w-6xl"></div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { savePost } from '@/handlers/index'
import { basicStore, keyStore } from '@/stores/index'
import { watch, ref, computed, onMounted, nextTick } from 'vue'
import { useRoute } from 'vue-router'

import { createLinkMeta, generateKey } from '@/utils'
import { useFullscreen } from '@vueuse/core'
const { isFullscreen, enter, exit } = useFullscreen()

import Vditor from 'vditor'

import { Message } from '@/plugins/message'

import { addPostStarData, deletePostStarData } from '@/database/star'
import { PostStar } from '@/models'

const route = useRoute()

const basic = basicStore()
const keys = keyStore()
const isAdmin = computed(() => basic.isAdmin)
const isEditing = ref(false)
const isStared = ref(false)

const fullElement = ref<HTMLElement | null>(null)

const postContainer = ref()
const postContent = defineModel('postContent', { type: String, required: true })
const postHtml = defineModel('postHtml', { type: String, required: true })

const contentEditor = ref<Vditor>()

function switchFullscreen() {
    const content = document.querySelector('#content')
    if (content) {
        if (!isFullscreen.value) {
            content.requestFullscreen()
        } else {
            document.exitFullscreen()
        }
    }
}

// 保存文章
async function save() {
    const params = route.params
    const mdContent = contentEditor.value?.getValue()

    if (mdContent) {
        const [ok, data] = await savePost(createLinkMeta(params), mdContent)

        if (ok) {
            Message('已保存', 'success')
            postContent.value = mdContent
            postHtml.value = data['content_html']
        } else {
            Message('保存失败', 'error')
        }
    }
}

async function starPost() {
    isStared.value = !isStared.value
    const params = route.params
    const postStar = new PostStar()

    postStar.key = generateKey(params)
    postStar.params = createLinkMeta(params)

    await addPostStarData(postStar)
    Message('收藏成功')
}
async function unStarPost() {
    isStared.value = !isStared.value
    const params = route.params
    const key = generateKey(params)

    await deletePostStarData(key)
    Message('已取消收藏', 'contrast')
}

watch(postContent, async () => {
    contentEditor.value?.setValue(postContent.value)
})
watch(postHtml, async () => {
    // 需要使用nextTick等待页面渲染
    await nextTick()
    clipboardListener()
})

watch(isEditing, async () => {
    if (contentEditor.value) {
        return
    }
    // https://www.cnblogs.com/fzu221801127/p/14939429.html#2%E5%88%9B%E5%BB%BA%E4%B8%80%E4%B8%AAvditor%E5%AE%9E%E4%BE%8B%E5%B9%B6%E8%B5%8B%E5%80%BC%E7%BB%99contenteditor
    if (isEditing.value) {
        contentEditor.value = new Vditor('vditor', {
            height: 'auto',
            width: '100%',
            mode: 'ir',
            toolbarConfig: {
                pin: true
            },
            preview: {
                hljs: {
                    style: 'base16-snazzy',
                    lineNumber: true
                }
            },

            cache: {
                enable: false
            },
            toolbar: [
                'emoji',
                'headings',
                'bold',
                'italic',
                'strike',
                '|',
                'line',
                'quote',
                'list',
                'ordered-list',
                'check',
                'outdent',
                'indent',
                '|',
                'code',
                'inline-code',
                'insert-after',
                'insert-before',
                'undo',
                'redo',
                '|',
                'upload',
                'link',
                'table',

                'edit-mode',
                'both',
                'fullscreen',
                'code-theme',
                'content-theme',
                'export',
                'help',
                'br'
            ],
            after: () => {
                contentEditor.value?.setValue(postContent.value)

                // 保存后 写回数据库
            }
        })
    }
})
function handleFullscreenChange() {
    if (document.fullscreenElement) {
        document.body.classList.add('fullscreen')
        isFullscreen.value = true
    } else {
        document.body.classList.remove('fullscreen')
        isFullscreen.value = false
    }
}

function copyCodeBlock(codeBlock: HTMLElement) {
    navigator.clipboard.writeText(codeBlock.innerText).then(
        () => {
            Message('已复制到剪切板', 'success')
        },
        () => {
            console.log('拒绝复制', 'error')
            Message('复制失败,请手动复制', 'error')
        }
    )
}

function clipboardListener() {
    const container = postContainer.value
    if (container) {
        var copyButtons = container.querySelectorAll('.copy-button')
        copyButtons.forEach(function (button: any) {
            button.addEventListener('click', function () {
                var codeBlock = (button.parentElement as HTMLElement).querySelector('code')
                if (codeBlock) {
                    copyCodeBlock(codeBlock)
                }
            })
        })
    }
}

onMounted(() => {
    document.addEventListener('fullscreenchange', handleFullscreenChange)

    document.addEventListener('keydown', function (event: KeyboardEvent) {
        if (keys.isFullScreenKey(event)) {
            event.preventDefault()
            switchFullscreen()
        }

        if (keys.isEditKey(event)) {
            event.preventDefault()
            isEditing.value = !isEditing.value
        }

        if (keys.isSaveKey(event)) {
            event.preventDefault()
            if (isEditing.value == true) {
                save()
            }
        }
    })
})
</script>
<style scoped>
code:hover::-webkit-scrollbar {
    background-color: #5858584d;
}
#content code:hover::-webkit-scrollbar-thumb {
    background-color: #27272791 !important;
}

#content code::-webkit-scrollbar-thumb {
    background: transparent;
}
</style>
