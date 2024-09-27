<script setup lang="ts">

defineProps<{ msg: string }>()

</script>

<template>
  <h1>{{ msg }}</h1>
  <div>
    <h3>Welcome to Password Keeper</h3>
    <p>Enter your password to get started</p>
    <input type="password" placeholder="Enter your password" />
  </div>
</template>