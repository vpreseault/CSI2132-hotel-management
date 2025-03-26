<template>
    <div class="w-full m-auto">
      <div class="mt-16 text-center">
        <h1 class="text-2xl mt-4">Employee View</h1>
        <p class="mt-4">You are logged in!</p>
      </div>
  
      <NavBar role="employee" @toggleProfile="toggleProfileModal" @toggleHotel="toggleHotelModal" />
      <Rental :expandedCard="expandedCard" :toggleCard="toggleCard" :isEmployee="true" :employeeRentals="allHotelRentings" />
      <Booking :expandedCard="expandedCard" :toggleCard="toggleCard" :isEmployee="true" @createBooking="handleEmployeeBooking" />
      <Profile v-if="isProfileModalOpen" :toggleProfileModal="toggleProfileModal" />
      <HotelModal v-if="isHotelModalOpen" @close="toggleHotelModal" />
  
      <div class="mt-16 text-center">
        <button 
          class="mt-4 px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition" 
          @click="handleLogout">
          Logout
        </button>
      </div>
    </div>
</template>
  
<script setup lang="ts">
import { ref } from 'vue';
import { removeAuthCookie } from '../utils/auth';
import NavBar from '../components/LandingPage/NavBar.vue';
import Rental from '../components/LandingPage/Rental.vue';
import Booking from '../components/LandingPage/Booking.vue';
import Profile from '../components/LandingPage/Profile.vue';
import HotelModal from '../components/LandingPage/HotelModal.vue';

const expandedCard = ref<{ section: string | null; index: number | null }>({ section: null, index: null });
const isProfileModalOpen = ref(false);
const isHotelModalOpen = ref(false);

function toggleProfileModal() {
  isProfileModalOpen.value = !isProfileModalOpen.value;
}

function toggleHotelModal() {
  isHotelModalOpen.value = !isHotelModalOpen.value;
}

function toggleCard(section: string, index: number) {
expandedCard.value = expandedCard.value.section === section && expandedCard.value.index === index 
    ? { section: null, index: null } 
    : { section, index };
}

function handleLogout() {
  removeAuthCookie();
  window.location.href = '/';
}

function handleEmployeeBooking(booking: {
  customerName: string;
  roomNumber: number;
  checkoutDate: string;
  email : string;
}) { allHotelRentings.value.push(booking); }

const allHotelRentings = ref([
{ customerName: "John Doe", roomNumber: 101, checkoutDate: "2025-03-25", email: "john@example.com" },
{ customerName: "Jane Smith", roomNumber: 202, checkoutDate: "2025-03-30", email: "jane@example.com" },
{ customerName: "Bob Builder", roomNumber: 303, checkoutDate: "2025-04-02", email: "bob@fixit.com" }
]);
</script>
  