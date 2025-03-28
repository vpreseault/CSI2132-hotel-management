<template>
    <div class="w-full m-auto">
      <div class="mt-16 text-center">
        <h1 class="text-2xl mt-4">Employee View</h1>
        <p class="mt-4">You are logged in!</p>
      </div>
  
      <NavBar role="employee" @toggleProfile="toggleProfileModal" @toggleHotel="toggleHotelModal" />
      
      <LayoutSection title="Current Rentals">
        <div class="mt-4 grid grid-cols-1 md:grid-cols-1 lg:grid-cols-1 gap-4 justify-center">
          <ActivityCard
            v-for="(rental, index) in hotelRentals"
            :key="`rental-${index}`"
            :customerName="rental.customer_name"
            :hotelName="rental.hotel_name"
            :startDate="new Date(rental.start_date)"
            :endDate="new Date(rental.end_date)"
            :employeeName="rental.employee_name"
            :roomNumber="rental.room_number"
            :price="rental.total_price"
            :payment="rental.payment"
            :section="'rental'"
            :index="index"
            :expandedCard="expandedCard"
            @toggle="toggleCard"
          />
        </div>
      </LayoutSection>
      <LayoutSection title="Upcomming Bookings">
        <div class="mt-4 grid grid-cols-1 md:grid-cols-1 lg:grid-cols-1 gap-4 justify-center">
          <ActivityCard
            v-for="(booking, index) in hotelBookings"
            :key="`booking-${index}`"
            :customerName="booking.customer_name"
            :hotelName="booking.hotel_name"
            :startDate="new Date(booking.start_date)"
            :endDate="new Date(booking.end_date)"
            :roomNumber="booking.room_number"
            :price="booking.total_price"
            :section="'booking'"
            :index="index"
            :expandedCard="expandedCard"
            @toggle="toggleCard"
          />
        </div>
      </LayoutSection>
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
import { onMounted, ref } from 'vue';
import { getUserID, removeAuthCookie } from '../utils/auth';
import NavBar from '../components/LandingPage/NavBar.vue';
import Profile from '../components/LandingPage/Profile.vue';
import HotelModal from '../components/LandingPage/HotelModal.vue';
import LayoutSection from '../components/Layout/LayoutSection.vue';

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
  email: string;
}) {
  // allHotelRentings.value.push({
  //   customer_name: booking.customerName,
  //   room_number: booking.roomNumber,
  //   start_date: new Date().toISOString().split('T')[0],
  //   end_date: booking.checkoutDate,
  //   employee_name: "Current Employee",
  //   cardType: "renting",
  //   total_price: 0,
  //   payment: false
  // });
}

type RentalItem = {
  customer_name: string;
  hotel_name: string;
  start_date: string;
  end_date: string;
  employee_name: string;
  room_number: number;
  total_price: number;
  payment: boolean;
};
type BookingItem = {
  customer_name: string;
  hotel_name: string;
  room_number: number;
  start_date: string;
  end_date: string;
  total_price: number;
};
const hotelRentals = ref<Array<RentalItem>>([])
  const hotelBookings = ref<Array<BookingItem>>([])
onMounted(async () => {
  const employeeID = getUserID()
  
  try {
      const rentalResponse = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/rentings?employee_ID=${employeeID}`)
      if (rentalResponse.ok) {
        hotelRentals.value = await rentalResponse.json()
      }

      const bookingResponse = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/bookings?employee_ID=${employeeID}`)
      if (bookingResponse.ok) {
        hotelBookings.value = await bookingResponse.json()
      }

  } catch (error) {
      console.error('Error calling API:', error);
  }
})
</script>

  