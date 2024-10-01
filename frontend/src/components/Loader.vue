<script lang="ts">
import { defineComponent } from 'vue';

export default defineComponent({
    name: 'Loader',
    props: {
        loading: {
            type: Boolean,
            required: true
        }
    }
});

</script>

<template>
    <div v-if="loading" class="loader"></div>
</template>

<style lang="css">
.loader
{
    width: 5em;
    height: 1em;

    border: 0.3em solid #283739;
    border-right: 0.3em solid #2C5D63;    
    background-color: #A2C11C;

    animation: 3s infinite example-three;
}

@keyframes example-three{
    from { transform: rotateY(0deg);}
    50% { transform: rotateY(180deg);}
    to { transform: rotateY(360deg);}
}
</style>