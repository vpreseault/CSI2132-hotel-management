<script setup lang="ts">
defineProps<{
    title?: string
}>()
</script>

<template>
    <div class="bg-green-100 p-6 rounded-lg shadow-md border border-gray-300 mb-10 max-w-5xl mx-auto">
        <h2 v-if="title" class="text-center text-xl text-gray-600 font-semibold mb-5"> 
            {{ title }} 
        </h2>
        <slot />
    </div>
</template>