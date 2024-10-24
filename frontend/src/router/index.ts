import { createRouter, createWebHashHistory } from 'vue-router'
import Home from '@/pages/HomePage.vue'
import Post from '@/pages/PostPage.vue'
import Book from '@/pages/BookPage.vue'
import PostStar from '@/pages/PostStar.vue'
import Search from '@/pages/SearchPage.vue'
import NotFound from '@/pages/NotFound.vue'
import BooksPage from '@/pages/BooksPage.vue'

const routes = [
    {
        path: '/',
        name: 'home',
        component: Home
    },
    {
        path: '/book/:bookPath+',
        name: 'book',
        component: Book
    },
    {
        path: '/books/',
        name: 'books',
        component: BooksPage
    },
    {
        path: '/star',
        name: 'star',
        component: PostStar
    },
    {
        path: '/search',
        name: 'search',
        component: Search
    },
    {
        path: '/post/:postPath+',
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
