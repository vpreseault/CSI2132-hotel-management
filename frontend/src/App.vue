<script setup lang="ts">
import AuthPage from './pages/Auth.vue';
import LandingPage from './pages/LandingPage.vue';
import EmployeeDashboard from './pages/EmployeeDashboard.vue';
import AdminPage from './pages/Admin.vue'
import { computed } from 'vue'
import { getAuthCookie, getUserRole } from './utils/auth';


const authenticated = computed(() => getAuthCookie())
const userRole = computed(() => getUserRole())
</script>

<template>
    <div>
        <AuthPage v-if="!authenticated" />
        <LandingPage v-else-if="userRole === 'Customer'" />
        <EmployeeDashboard v-else-if="userRole === 'Employee'" />
        <AdminPage v-else />
    </div>
</template>