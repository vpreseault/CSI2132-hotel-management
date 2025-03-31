<template>
    <div class="w-full m-auto">
      <div class="mt-16 text-center">
        <h1 class="text-2xl mt-4">Admin View</h1>
      </div>
  
      <NavBar role="admin" @exit-admin="exitAdminView" />
  
      <div class="mt-12 px-6">
        <HotelList />
      </div>
  
      <Toast position="bottom-center" />
    </div>
</template>

<script setup lang="ts">
import { removeAuthCookie } from '../utils/auth'
import NavBar from '../components/LandingPage/NavBar.vue'
import HotelList from '../components/Admin/HotelList.vue'  

function exitAdminView() {
handleLogout()
}

function handleLogout() {
removeAuthCookie()
window.location.href = '/'
}
</script>
