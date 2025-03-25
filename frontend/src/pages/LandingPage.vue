<template>
    <div class="w-full m-auto">
      <NavBar role="customer" @toggleProfile="toggleProfileModal" />
  
      <div class="mt-16 text-center">
        <h1 class="text-2xl mt-4">Welcome to your one-stop shop for all things Hotel Room Rentings!</h1>
        <p class="mt-4">You are logged in!</p>
      </div>
  
      <Booking :expandedCard="expandedCard" :toggleCard="toggleCard" />
      <Rental :expandedCard="expandedCard" :toggleCard="toggleCard" />
      <Profile v-if="isProfileModalOpen" :toggleProfileModal="toggleProfileModal" />
  
      <div class="mt-16 text-center">
        <button class="mt-4 px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition" @click="handleLogout"> Logout </button>
      </div>
    </div>
</template>
  
<script setup lang="ts">
import { ref } from 'vue';
import { removeAuthCookie } from '../utils/auth';
import NavBar from '../components/LandingPage/NavBar.vue';
import Booking from '../components/LandingPage/Booking.vue';
import Rental from '../components/LandingPage/Rental.vue';
import Profile from '../components/LandingPage/Profile.vue';

const expandedCard = ref<{ section: string | null; index: number | null }>({ section: null, index: null });

const isProfileModalOpen = ref(false);

function toggleCard(section: string, index: number) {
expandedCard.value = expandedCard.value.section === section && expandedCard.value.index === index 
  ? { section: null, index: null } 
  : { section, index };
}

function toggleProfileModal() {
  isProfileModalOpen.value = !isProfileModalOpen.value;
}

function handleLogout() {
  removeAuthCookie();
  window.location.href = '/';
}
</script>
  