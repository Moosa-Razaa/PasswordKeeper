<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
    name: 'GettingStarted',
});

</script>

<template>
    <div class="mainDiv">
        <h2 class="">Password Keeper</h2>
        <div>
            <h3>Welcome to Password Keeper</h3>
            <p>Enter your password to get started</p>
            <input type="password" placeholder="Enter your password" />
        </div>
    </div>
</template>