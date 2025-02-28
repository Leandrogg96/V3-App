import { createRouter, createWebHistory } from "vue-router";
import AppBody from "@/components/App-Body.vue";
import AppLogin from "@/components/App-Login.vue";
import Books from "@/components/BooksList.vue";
import Book from "@/components/BookItem.vue";
import BooksAdmin from "@/components/BooksAdmin.vue";
import BookEdit from "@/components/BookEdit.vue";
import Users from "@/components/UsersList.vue";
import User from "@/components/UserEdit.vue";
import Security from "@/components/security"; 

//import { compile } from "vue";

const routes = [
    {
        path: '/',
        name: 'Home',
        component: AppBody, 
    },
    {
        path: '/login',
        name: 'Login',
        component: AppLogin, 
    },
    {
        path: '/books',
        name: 'BooksList',
        component: Books,
    },
    {
        path: '/books/:bookName',
        name: 'Book',
        component: Book,
    },
    {
        path: '/admin/books',
        name: 'BooksAdmin',
        component: BooksAdmin,
    },
    {
        path: '/admin/books/:bookId',
        name: 'BookEdit',
        component: BookEdit,
    },
    {
        path: '/admin/users',
        name: 'Users',
        component: Users,
    },
    {
        path: '/admin/users/:userId',
        name: 'User',
        component: User,
    },
]

const router = createRouter({history: createWebHistory(), routes})
router.beforeEach(() => {
    Security.checkToken();
})
export default router