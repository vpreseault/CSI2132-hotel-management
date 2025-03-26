<script setup lang="ts">
import AuthPage from './pages/Auth.vue';
import LandingPage from './pages/LandingPage.vue';
import EmployeeDashboard from './pages/EmployeeDashboard.vue';
import ManagerDasboard from './pages/ManagerDashboard.vue';
import { computed } from 'vue'
import { getAuthCookie } from './utils/auth';


const authenticated = computed(() => getAuthCookie())
</script>

<template>
    <div>
        <AuthPage v-if="!authenticated" />
        <EmployeeDashboard v-else-if="authenticated === 'John Doe'" />
        <ManagerDasboard v-else-if="authenticated === 'Jane Smith'" />
        <LandingPage v-else />
    </div>
</template>