<template>
  <div class="w-full m-auto">
    <div class="mt-16 text-center">
      <h1 class="text-2xl mt-4">{{role}} View</h1>
      <p class="mt-4">You are logged in!</p>
    </div>

    <NavBar :role="role === 'Manager' ? 'manager' : 'employee'" @toggleProfile="toggleProfileModal" @toggleHotel="toggleHotelModal" @toggleCreateEmployee="toggleCreateEmployeeModal" />
    
    <LayoutSection title="Current Rentals">
      <div v-if="hotelRentals.length > 0" class="mt-4 grid grid-cols-1 md:grid-cols-1 lg:grid-cols-1 gap-4 justify-center">
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
      <p v-else class="flex justify-center mt-4">There are no current rentals.</p>
    </LayoutSection>
    <LayoutSection title="Upcoming Bookings">
      <div v-if="hotelBookings.length > 0" class="mt-4 grid grid-cols-1 md:grid-cols-1 lg:grid-cols-1 gap-4 justify-center">
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
      <p v-else class="flex justify-center mt-4">There are no upcoming bookings.</p>
    </LayoutSection>
    <Booking :expandedCard="expandedCard" :toggleCard="toggleCard" :isEmployee="true" @createBooking="handleEmployeeBooking" />
    <EmployeeList v-if="role === 'Manager'" @delete="showEmployeeDeletedToast" />
    <RoomList v-if="role === 'Manager'" @delete="showRoomDeletedToast" />
    <CreateEmployeeModal v-if="isCreateEmployeeModalOpen && role === 'Manager'" @close="toggleCreateEmployeeModal" @created="showEmployeeCreatedToast" />
    <Profile v-if="isProfileModalOpen" role="employee" :toggleProfileModal="toggleProfileModal" />
    <HotelModal v-if="isHotelModalOpen" @close="toggleHotelModal" />

    <div class="mt-16 text-center pb-3">
        <button class="mt-4 px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition" @click="handleLogout">Logout</button>
    </div>
    <Toast position="bottom-center" />
  </div>
</template>
  
<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { getUserID, getUserRole, removeAuthCookie } from '../utils/auth';
import NavBar from '../components/LandingPage/NavBar.vue';
import Profile from '../components/LandingPage/Profile.vue';
import HotelModal from '../components/LandingPage/HotelModal.vue';
import CreateEmployeeModal from '../components/LandingPage/CreateEmployeeModal.vue';
import LayoutSection from '../components/Layout/LayoutSection.vue';
import { useToast } from "primevue/usetoast";
import type { BookingItem, RentalItem } from '../types';
import EmployeeList from '../components/LandingPage/EmployeeList.vue';
import RoomList from '../components/LandingPage/RoomList.vue';
import type { ToastMessageOptions } from 'primevue';

const toast = useToast();

const expandedCard = ref<{ section: string | null; index: number | null }>({ section: null, index: null });
const isProfileModalOpen = ref(false);
const isHotelModalOpen = ref(false);
const isCreateEmployeeModalOpen = ref(false);
const role = ref(getUserRole())

function toggleProfileModal() {
  isProfileModalOpen.value = !isProfileModalOpen.value;
}

function toggleHotelModal() {
  isHotelModalOpen.value = !isHotelModalOpen.value;
}

function toggleCreateEmployeeModal() {
    isCreateEmployeeModalOpen.value = !isCreateEmployeeModalOpen.value;
}

function showEmployeeCreatedToast() {
    toggleCreateEmployeeModal()
    toast.add({ severity: 'success', summary: 'Success', detail: 'Employee account created.', life: 3000});
}

function showEmployeeDeletedToast(severity: ToastMessageOptions["severity"]) {
    toast.add({ 
      severity, 
      summary: severity === 'success' ? 'Success': 'Failed', 
      detail: severity === 'success' ? 'Deleted employee account successfully.': 'Failed to delete employee account.', 
      life: 3000
    });
}

function showRoomDeletedToast(severity: ToastMessageOptions["severity"]) {
  toast.add({ 
    severity, 
    summary: severity === 'success' ? 'Success' : 'Failed', 
    detail: severity === 'success' ? 'Deleted room successfully.' : 'Failed to delete room.', 
    life: 3000 
  })
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

  