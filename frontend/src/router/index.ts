import { createRouter, createWebHashHistory } from 'vue-router'
import Home from '@/pages/HomePage.vue'
import Post from '@/pages/PostPage.vue'
import Book from '@/pages/BookPage.vue'
import PostStar from '@/pages/PostStar.vue'
import NotFound from '@/pages/NotFound.vue'

const routes = [
    {
        path: '/',
        name: 'home',
        component: Home
    },

    {
        path: '/book/:category/:book/:locale',
        name: 'book',
        component: Book
    },
    {
        path: '/star',
        name: 'star',
        component: PostStar
    },
    {
        path: '/post/:category/:book/:locale/:chapter?/:section?/:document(.+md)',
        name: 'post',
        component: Post
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'not-found',
        component: NotFound
    }
]

const router = createRouter({
    // history: createWebHistory(),
    history: createWebHashHistory(),
    routes
})

export default router
