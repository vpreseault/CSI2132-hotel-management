<template>
  <div class="w-full m-auto">
    <div class="mt-16 text-center">
      <h1 class="text-2xl mt-4">{{role}} View</h1>
      <p class="mt-4">You are logged in!</p>
    </div>

    <NavBar :role="role === 'Manager' ? 'manager' : 'employee'" @toggleProfile="toggleProfileModal" @toggleHotel="toggleHotelModal" @toggleCreateEmployee="toggleCreateEmployeeModal" />
    
    <LayoutSection title="Current Rentals">
      <div v-if="hotelRentals?.length" class="mt-4 grid grid-cols-1 md:grid-cols-1 lg:grid-cols-1 gap-4 justify-center">
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
      <div v-if="hotelBookings?.length" class="mt-4 grid grid-cols-1 md:grid-cols-1 lg:grid-cols-1 gap-4 justify-center">
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
        >
        <Button class="mt-3" label="Activate Booking" size="small" @click="openPaymentModal(booking)"/>
        </ActivityCard>
      </div>
      <p v-else class="flex justify-center mt-4">There are no upcoming bookings.</p>
    </LayoutSection>
    <Booking :expandedCard="expandedCard" :toggleCard="toggleCard" :isEmployee="true" @createBooking="handleEmployeeBooking" />
    <EmployeeList v-if="role === 'Manager'" @delete="showEmployeeDeletedToast" />
    <RoomList v-if="role === 'Manager'" @delete="showRoomDeletedToast" @update="showRoomUpdatedToast" @create="showRoomCreatedToast" />
    <CreateEmployeeModal v-if="isCreateEmployeeModalOpen && role === 'Manager'" @close="toggleCreateEmployeeModal" @created="showEmployeeCreatedToast" />
    <Profile v-if="isProfileModalOpen" role="employee" :toggleProfileModal="toggleProfileModal" />
    <HotelModal v-if="isHotelModalOpen" @close="toggleHotelModal" />
    <PaymentModal v-if="isPaymentModalOpen" :booking="selectedBooking" @close="closePaymentModal" />
    <!--@confirm="confirmPayment"-->

    <Toast position="bottom-center" />
  </div>
</template>
  
<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { getUserID, getUserRole } from '../utils/auth';
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
import PaymentModal from '../components/LandingPage/PaymentModal.vue';


const toast = useToast();

const expandedCard = ref<{ section: string | null; index: number | null }>({ section: null, index: null });
const isProfileModalOpen = ref(false);
const isHotelModalOpen = ref(false);
const isCreateEmployeeModalOpen = ref(false);
const role = ref(getUserRole())

const isPaymentModalOpen = ref(false);
const selectedBooking = ref<BookingItem | null>(null);

function openPaymentModal(booking: BookingItem) {
  selectedBooking.value = booking;
  isPaymentModalOpen.value = true;
}

function closePaymentModal() {
  isPaymentModalOpen.value = false;
  selectedBooking.value = null;
}

// function confirmPayment(updatedBooking: BookingItem & {
//   room_number: number;
//   start_date: string;
//   end_date: string;
//   total_price: number;
//   card_type: string;
//   card_number: string;
// }) {
//   hotelBookings.value = hotelBookings.value.filter(b => b.booking_id !== updatedBooking.booking_id);

//   hotelRentals.value.push({
//     customer_name: updatedBooking.customer_name,
//     hotel_name: updatedBooking.hotel_name,
//     room_number: updatedBooking.room_number,
//     start_date: updatedBooking.start_date,
//     end_date: updatedBooking.end_date,
//     employee_name: "Current Employee",
//     total_price: updatedBooking.total_price,
//     payment: true,  
//   });

//   closePaymentModal();

//   toast.add({
//     severity: 'success',
//     summary: 'Booking Activated',
//     detail: 'Moved to current rentals.',
//     life: 3000
//   });
// }

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

function showRoomUpdatedToast(severity: ToastMessageOptions["severity"]) {
  toast.add({ 
    severity, 
    summary: severity === 'success' ? 'Success' : 'Failed', 
    detail: severity === 'success' ? 'Updated room successfully.' : 'Failed to update room.', 
    life: 3000 
  })
}

function showRoomCreatedToast(severity: ToastMessageOptions["severity"]) {
  toast.add({ 
    severity, 
    summary: severity === 'success' ? 'Success' : 'Failed', 
    detail: severity === 'success' ? 'Created room successfully.' : 'Failed to create room.', 
    life: 3000 
  })
}

function toggleCard(section: string, index: number) {
  expandedCard.value = expandedCard.value.section === section && expandedCard.value.index === index 
    ? { section: null, index: null } 
    : { section, index };
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

  