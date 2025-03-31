<script setup lang="ts">
defineProps<{
    title?: string
}>()
</script>

<template>
    <div class="my-16 p-8 bg-green-100 rounded-lg shadow-md mx-25 xl:mx-50">
        <h2 v-if="title" class="text-center text-xl text-gray-600 font-semibold mb-5"> 
            {{ title }} 
        </h2>
        <slot />
    </div>
</template>