<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
    name: 'GettingStarted',
});

</script>

<template>
    <div class="mainDiv h-screen w-screen bg-gray-900 flex justify-center items-center flex-col">
        <h1 class="text-4xl text-lime-600">Getting Started</h1>
        <p class="text-white mt-5">It seems you are using our product for the first time. You need to set a default entrance key password.</p>
        <div class="flex justify-center items-center flex-col m-10 pt-5 border-t-2 border-t-lime-600">
            <p class="text-white mb-5">-- please enter the key entrance password --</p>
            <input type="password" class="p-1.5 rounded-lg w-80 text-center" placeholder="enter the key entrance password" />
        </div>
    </div>
</template>