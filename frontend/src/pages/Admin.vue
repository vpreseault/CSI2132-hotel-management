<template>
    <div class="w-full m-auto">
        <div class="mt-16 text-center">
            <h1 class="text-2xl mt-4">Admin View</h1>
        </div>

        <NavBar role="admin" @exit-admin="exitAdminView" />
  
        <div class="mt-16 text-center">
            <button class="mt-4 px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition" @click="handleLogout"> Logout </button>
        </div>
        <Toast position="bottom-center" />
    </div>
</template>

<script setup lang="ts">
import { removeAuthCookie } from '../utils/auth';
import NavBar from '../components/LandingPage/NavBar.vue';

function exitAdminView() {
    handleLogout()
}

function handleLogout() {
    removeAuthCookie();
    window.location.href = '/';
}
</script>