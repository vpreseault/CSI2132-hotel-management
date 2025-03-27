<script setup lang="ts">
import AuthPage from './pages/Auth.vue';
import LandingPage from './pages/LandingPage.vue';
import EmployeeDashboard from './pages/EmployeeDashboard.vue';
import ManagerDasboard from './pages/ManagerDashboard.vue';
import { computed } from 'vue'
import { getAuthCookie, getUserRole } from './utils/auth';


const authenticated = computed(() => getAuthCookie())
const userRole = computed(() => getUserRole())
</script>

<template>
    <div>
        <AuthPage v-if="!authenticated" />
        <EmployeeDashboard v-else-if="userRole === 'Employee'" />
        <ManagerDasboard v-else-if="userRole === 'Manager'" />
        <LandingPage v-else />
    </div>
</template>