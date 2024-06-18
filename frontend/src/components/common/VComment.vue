<template>
    <div class="comment-top">
        <div class="w-full">
            <textarea
                name=""
                id=""
                cols="30"
                rows="3"
                v-model="commentContent"
                class="w-full py-3 px-4 rounded-br-md min-h-12"
                :placeholder="poem"></textarea>
        </div>
        <div class="mt-2 flex">
            <div type="text" class="items-center ml-auto flex gap-2 py-2 text-right">
                <span class="select-none text-sm">昵称:</span>
                <span class="select-none text-sm">{{ nickname }}</span>
                <BIconArrowClockwise @click="refreshNickname"> </BIconArrowClockwise>
            </div>
            <button
                class="btn bg-theme-primary-base hover:bg-theme-primary-hover ml-4 px-2 py-0"
                @click="postNewComment">
                发布
            </button>
        </div>
    </div>

    <div class="comment-body">
        <div class="border-b border-theme-text-muted" v-for="comment in comments" :key="comment.id">
            <div class="my-4">
                <div class="flex">
                    <div class="font-bold">{{ comment.nickname }}</div>
                    <div class="ml-4">{{ comment.content }}</div>
                    <div class="ml-auto text-theme-text-muted">
                        {{ formatDate(comment.createdAt) }}
                    </div>
                </div>
            </div>

            <div class="" v-for="reply in comment.replies" :key="reply.id">
                <div class="my-4 ml-8">
                    <div class="flex">
                        <div class="font-bold">{{ reply.nickname }}</div>
                        <div class="ml-4">{{ reply.content }}</div>
                        <div class="ml-auto text-theme-text-muted">
                            {{ formatDate(reply.createdAt) }}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted, watch } from 'vue'
import { storeToRefs } from 'pinia'
import { Comment } from '@/models/comment'
import { getRandNickname, getRandPoem, getComments, postComment } from '@/services/index'

import { BIconArrowClockwise } from 'bootstrap-icons-vue'
const commentContent = ref('')
const poem = ref('')
import { fetchBasic, formatDate } from '@/utils'

const route = useRoute()

import { basicStore } from '@/stores'
import { useRoute } from 'vue-router'
import { Message } from '@/plugins/message'

const basic = basicStore()
const locale = computed(() => basic.locale)

let { nickname } = storeToRefs(basic)
const comments = ref<Comment[]>([])
async function postNewComment() {
    const comment = new Comment()
    comment.nickname = nickname.value
    comment.parent = 0
    comment.url = (route.params['bookPath'] as string[]).join('/') + '/' + locale.value
    comment.content = commentContent.value

    // fetchHandler(comments,[],getComments,"data",await Message('发布成功'),await Message('发布失败', 'warn')

    const [ok] = await postComment(comment)

    if (ok) {
        await fetchBasic(
            comments,
            [],
            getComments,
            (route.params['bookPath'] as string[]).join('/') + '/' + locale.value
        )
        await Message({ message: '发布成功' })
    } else {
        await Message({ message: '发布失败', type: 'warn' })
    }
}
watch(route, () => {
    refresh()
})

async function refreshNickname() {
    await fetchBasic(nickname, nickname.value, getRandNickname)
    localStorage.setItem('nickname', nickname.value)
}

async function refresh() {
    await fetchBasic(
        comments,
        [],
        getComments,
        (route.params['bookPath'] as string[]).join('/') + '/' + locale.value
    )
    await fetchBasic(poem, '山重水复疑无路，柳暗花明又一村。', getRandPoem)

    if (nickname.value == '') {
        await fetchBasic(nickname, '匿名用户', getRandNickname)
        localStorage.setItem('nickname', nickname.value)
    }
}

onMounted(async () => {
    await refresh()
})
</script>