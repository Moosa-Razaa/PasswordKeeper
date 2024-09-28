import { createRouter, createWebHistory } from 'vue-router';
import GettingStarted from '../views/GettingStarted.vue';

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'Home',
            component: GettingStarted
        }
    ]
});

export default router;