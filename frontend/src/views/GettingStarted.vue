<script lang="ts">
import CheckFileExist from '@/api/CheckFileExist';
import { defineComponent, onMounted, ref } from 'vue';
import Loader from '@/components/Loader.vue';

export default defineComponent({
    name: 'GettingStarted',
    setup() {
        const error = ref<string | null>();
        const loading = ref<boolean>(true);

        onMounted(async () => {
            await CheckFileExist("repository")
            .then((res) => 
            {
                loading.value = false;
                if (res) {
                    error.value = res;
                }
            })
            .catch((err) => 
            {
                loading.value = false;
                error.value = err;
            })
            .finally(() => 
            {
                loading.value = false;
            });
        });

        return {
            error,
            loading
        };
    }
});

</script>

<template>
    <div class="mainDiv h-screen w-screen bg-gray-900 flex justify-center items-center flex-col">
        <h1 class="text-4xl text-lime-600">Getting Started</h1>
        <Loader :isVisible="loading" />
        <p v-if="!loading" class="text-white mt-5">It seems you are using our product for the first time. You need to set a default entrance key password.</p>
        <div v-if="!loading" class="flex justify-center items-center flex-col m-10 pt-5 border-t-2 border-t-lime-600">
            <p class="text-white mb-5">-- please enter the key entrance password --</p>
            <input type="password" class="p-1.5 rounded-lg w-80 text-center" placeholder="enter the key entrance password" />
        </div>
    </div>
</template>