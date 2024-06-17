<template>
    <div
        id="content"
        ref="fullElement"
        class="max-w-6xl m-auto relative mb-12 bg-theme-card p-16 rounded-2xl">
        <div
            class="toolbar py-4 fixed flex-col top-16 text-[1.25rem] right-[2rem] lg:right-[17rem] xl:right-[21rem]">
            <div class="p-2">
                <BIconArrowsFullscreen
                    v-if="!isFullscreen"
                    @click="switchFullscreen"></BIconArrowsFullscreen>
                <BIconFullscreenExit v-else @click="switchFullscreen"></BIconFullscreenExit>
            </div>
            <div class="p-2">
                <BIconBookmarkPlus v-if="!isStared" @click="starPost"></BIconBookmarkPlus>
                <BIconBookmarkDash v-else @click="unStarPost"></BIconBookmarkDash>
            </div>
            <div v-if="isAdmin" class="flex flex-col p-2">
                <BIconPencilSquare
                    v-if="!isEditing"
                    @click="isEditing = !isEditing"></BIconPencilSquare>
                <BIconEye v-else @click="isEditing = !isEditing"></BIconEye>
            </div>
            <div v-if="isAdmin && isEditing" class="flex flex-col p-2">
                <BIconSave2 @click="save"></BIconSave2>
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

import Vditor from 'vditor'

import { Message } from '@/plugins/message'

import { addPostStarData, deletePostStarData, getPostStarData } from '@/database/star'
import { PostStar } from '@/models'
import {
    BIconBookmarkPlus,
    BIconBookmarkDash,
    BIconArrowsFullscreen,
    BIconFullscreenExit,
    BIconPencilSquare,
    BIconQuestionCircle,
    BIconSave2,
    BIconEye
} from 'bootstrap-icons-vue'

const route = useRoute()

const basic = basicStore()
const shortcuts = keyStore()
const isAdmin = computed(() => basic.isAdmin)
const isEditing = ref(false)
const isStared = ref(false)
const isFullscreen = ref(false)

const fullElement = ref<HTMLElement | null>(null)
const postContainer = ref()
const postContent = defineModel('postContent', { type: String, required: true })
const postHtml = defineModel('postHtml', { type: String, required: true })

const contentEditor = ref<Vditor>()

function switchFullscreen() {
    const content = document.querySelector('article')
    if (content) {
        if (!document.fullscreenElement) {
            content
                .requestFullscreen()
                .then(() => {
                    isFullscreen.value = true
                    content.style.overflow = 'auto' // 允许滚动
                    content.style.height = '100vh' // 确保高度覆盖视窗高度
                })
                .catch((err) => {
                    console.error(
                        `Error attempting to enable full-screen mode: ${err.message} (${err.name})`
                    )
                })
        } else {
            document
                .exitFullscreen()
                .then(() => {
                    isFullscreen.value = false
                    content.style.overflow = '' // 还原滚动设置
                    content.style.height = '' // 还原高度设置
                })
                .catch((err) => {
                    console.error(
                        `Error attempting to exit full-screen mode: ${err.message} (${err.name})`
                    )
                })
        }
    }
}

// 保存文章
async function save() {
    const params = route.params
    const mdContent = contentEditor.value?.getValue()

    if (mdContent) {
        const [ok, data] = await savePost((params['postPath'] as string[]).join('/'), mdContent)

        if (ok) {
            await Message('已保存', 'success')
            postContent.value = mdContent
            postHtml.value = data['content_html']
        } else {
            await Message('保存失败', 'error')
        }
    }
}

async function starPost() {
    isStared.value = !isStared.value

    const postStar = new PostStar()

    postStar.key = route.fullPath
    postStar.document = route.params['document'] as string
    postStar.postPath = route.params['postPath'] as string[]
    postStar.params = (route.params['postPath'] as string[]).slice(0, 3).join('/')

    await addPostStarData(postStar)
    await Message('收藏成功')
}
async function unStarPost() {
    isStared.value = !isStared.value
    const key = route.fullPath

    await deletePostStarData(key)
    await Message('已取消收藏')
}

async function refreshStarStatus() {
    const data = await getPostStarData(route.fullPath)

    if (data) {
        isStared.value = true
    } else {
        isStared.value = false
    }
}

watch(route, async () => {
    await refreshStarStatus()
})

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

async function copyCodeBlock(codeBlock: HTMLElement) {
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

onMounted(async () => {
    document.addEventListener('keydown', function (event: KeyboardEvent) {
        if (shortcuts.isFullScreenKey(event)) {
            event.preventDefault()
            switchFullscreen()
        }

        if (shortcuts.isEditKey(event)) {
            event.preventDefault()
            isEditing.value = !isEditing.value
        }

        if (shortcuts.isSaveKey(event)) {
            event.preventDefault()
            if (isEditing.value == true) {
                save()
            }
        }
    })

    await refreshStarStatus()
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
