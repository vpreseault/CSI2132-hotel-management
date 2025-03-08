<script setup lang="ts">
import AuthPage from './pages/Auth.vue';
import { computed } from 'vue'
import { getAuthCookie, removeAuthCookie } from './utils/auth';

const authenticated = computed(() => getAuthCookie())
</script>

<template>
    <div>
        <AuthPage v-if="!authenticated" />
        <Button v-else @click="removeAuthCookie">Logout</Button>
    </div>
</template>