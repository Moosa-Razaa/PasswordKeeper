<script lang="ts">
import CheckFileExist from '@/api/CheckFileExist';
import { defineComponent, onMounted, ref } from 'vue';
import Loader from '@/components/Loader.vue';
import PasswordSetup from '@/components/PasswordSetup.vue';
import Login from '@/components/Login.vue';

export default defineComponent({
    name: 'GettingStarted',
    components: {
        Loader
    },
    setup() {
        const error = ref<string | null>();
        const loading = ref<boolean>(true);
        const state = ref<"default" | "new" | "login">("default");
        const resource = "repository";

        function ReturnMainHeading(): string {
            switch (state.value) {
                case "new":
                    return "Getting Started";
                case "login":
                    return "Login";
                default:
                    return "Password Keeper";
            }
        }

        onMounted(async () => {
            await CheckFileExist(resource)
                .then((res) => {
                    loading.value = false;
                    if (res) {
                        if(res === "") {
                            state.value = "login";
                        } else {
                            state.value = "new";
                        }
                        error.value = res;
                    }
                })
                .catch((err) => {
                    loading.value = false;
                    error.value = err;
                    state.value = "new";
                })
                .finally(() => {
                    loading.value = false;
                });
        });

        return {
            error,
            loading,
            state,
            ReturnMainHeading
        };
    }
});

</script>

<template>
    <div class="mainDiv h-screen w-screen bg-gray-900 flex justify-center items-center flex-col">
        <h1 class="text-4xl text-lime-600 mb-5">{{ ReturnMainHeading() }}</h1>
        <Loader :loading="loading" />
        <p v-if="state === 'default'" class="text-white">Welcome! There seems to be problem in starting the application.</p>
        <PasswordSetup v-else-if="state === 'new'" :loading="loading" />
        <Login v-else-if="state === 'login'" :loading="loading" />
    </div>
</template>