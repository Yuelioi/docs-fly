<template>
    <div class="comment-top">
        <div class="w-full">
            <textarea
                name=""
                id=""
                cols="30"
                rows="3"
                v-model="commentContent"
                class="w-full px-4 py-3 rounded-br-md min-h-12"
                :placeholder="poem"></textarea>
        </div>
        <div class="flex mt-2">
            <div type="text" class="flex items-center gap-2 py-2 ml-auto text-right">
                <span class="text-sm select-none">昵称:</span>
                <span class="text-sm select-none">{{ nickname }}</span>
                <BIconArrowClockwise @click="refreshNickname"> </BIconArrowClockwise>
            </div>
            <button class="px-2 py-0 ml-4 btn hover:" @click="postNewComment">发布</button>
        </div>
    </div>

    <div class="comment-body">
        <div class="border-b" v-for="comment in comments" :key="comment.id">
            <div class="my-4">
                <div class="flex">
                    <div class="font-bold">{{ comment.nickname }}</div>
                    <div class="ml-4">{{ comment.content }}</div>
                    <div class="ml-auto">
                        {{ formatDate(comment.createdAt) }}
                    </div>
                </div>
            </div>

            <div class="" v-for="reply in comment.replies" :key="reply.id">
                <div class="my-4 ml-8">
                    <div class="flex">
                        <div class="font-bold">{{ reply.nickname }}</div>
                        <div class="ml-4">{{ reply.content }}</div>
                        <div class="ml-auto">
                            {{ formatDate(reply.createdAt) }}
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { Comment } from '@/models/comment'

import { BIconArrowClockwise } from 'bootstrap-icons-vue'

import { formatDate } from '@/utils'

const commentContent = ref('')
const poem = ref('')

const route = useRoute()
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

    const [ok] = await postComment(comment)

    if (ok) {
        await fetchHandleBasic(
            comments,
            [],
            getComments,
            (route.params['bookPath'] as string[]).join('/') + '/' + locale.value
        )
        Message({ message: '发布成功' })
    } else {
        Message({ message: '发布失败', type: 'warn' })
    }
}
watch(route, () => {
    refresh()
})

async function refreshNickname() {
    await fetchHandleBasic(nickname, nickname.value, getRandNickname)
    localStorage.setItem('nickname', nickname.value)
}

async function refresh() {
    await fetchHandleBasic(
        comments,
        [],
        getComments,
        (route.params['bookPath'] as string[]).join('/') + '/' + locale.value
    )

    await fetchHandleBasic(poem, '山重水复疑无路，柳暗花明又一村。', getRandPoem)

    if (nickname.value == '') {
        await fetchHandleBasic(nickname, '匿名用户', getRandNickname)
        localStorage.setItem('nickname', nickname.value)
    }
}

onMounted(async () => {
    await refresh()
})
</script>
