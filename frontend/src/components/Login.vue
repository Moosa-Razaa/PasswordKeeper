<script lang="ts">
import { defineComponent } from 'vue';
import Loader from '@/components/Loader.vue';

export default defineComponent({
    name: 'Login',
    components: {
        Loader
    },
    setup() {
        const loading = false;

        return {
            loading
        };
    }
});

</script>

<template>
    <div class="mainDiv h-screen w-screen bg-gray-900 flex justify-center items-center flex-col">
        <h1 class="text-4xl text-lime-600 mb-5">Login</h1>
        <Loader :loading="loading" />
        <p v-if="!loading" class="text-white mt-5">Please enter your credentials to login.</p>
        <div v-if="!loading" class="flex justify-center items-center flex-col m-10 pt-5 border-t-2 border-t-lime-600">
            <p class="text-white mb-5">-- please enter your credentials --</p>
            <input type="text" class="p-1.5 rounded-lg w-80 text-center"
                placeholder="enter your username" />
            <input type="password" class="p-1.5 rounded-lg w-80 text-center"
                placeholder="enter your password" />
            <button class="bg-lime-600 text-white p-2 rounded-lg w-80 mt-5">Login</button>
        </div>
    </div>
</template>