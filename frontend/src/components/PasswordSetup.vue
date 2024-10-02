<script lang="ts">
import { defineComponent, ref } from 'vue';

export default defineComponent({
    name: 'PasswordSetup',
    props: {
        loading: {
            type: Boolean,
            required: true
        }
    }
});

</script>

<template>
    <p v-if="!loading" class="text-white mt-5">It seems you are using our product for the first time. You need to
            set a default entrance key password.</p>
        <div v-if="!loading" class="flex justify-center items-center flex-col m-10 pt-5 border-t-2 border-t-lime-600">
            <p class="text-white mb-5">-- please enter the key entrance password --</p>
            <input type="password" class="p-1.5 rounded-lg w-80 text-center"
                placeholder="enter the key entrance password" />
        </div>
</template>