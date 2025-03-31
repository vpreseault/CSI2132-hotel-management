<template>
  <div class="w-full m-auto">
    <div class="mt-16 text-center">
      <h1 class="text-2xl mt-4">{{ role }} View</h1>
      <p class="mt-4">You are logged in!</p>
    </div>

    <NavBar :role="role === 'Manager' ? 'manager' : 'employee'" @toggleProfile="toggleProfileModal"
      @toggleHotel="toggleHotelModal" @toggleCreateEmployee="toggleCreateEmployeeModal" />

    <LayoutSection title="Current Rentals">
      <div v-if="hotelRentals.length > 0"
        class="mt-4 grid grid-cols-1 md:grid-cols-1 lg:grid-cols-1 gap-4 justify-center">
        <ActivityCard v-for="(rental, index) in hotelRentals" :key="`rental-${index}`"
          :customerName="rental.customer_name" :hotelName="rental.hotel_name" :startDate="new Date(rental.check_in_date)"
          :endDate="new Date(rental.check_out_date)" :employeeName="rental.employee_name" :roomNumber="rental.room_number"
          :price="rental.total_price" :payment="rental.payment" :section="'rental'" :index="index"
          :expandedCard="expandedCard" @toggle="toggleCard" />
      </div>
      <p v-else class="flex justify-center mt-4">There are no current rentals.</p>
    </LayoutSection>
    <LayoutSection title="Upcoming Bookings">
      <div v-if="hotelBookings.length > 0"
        class="mt-4 grid grid-cols-1 md:grid-cols-1 lg:grid-cols-1 gap-4 justify-center">
        <ActivityCard v-for="(booking, index) in hotelBookings" :key="`booking-${index}`"
          :customerName="booking.customer_name" :hotelName="booking.hotel_name"
          :startDate="new Date(booking.start_date)" :endDate="new Date(booking.end_date)"
          :roomNumber="booking.room_number" :price="booking.total_price" :section="'booking'" :index="index"
          :expandedCard="expandedCard" @toggle="toggleCard">
          <Button class="mt-3" label="Activate Booking" size="small" @click="openPaymentModal(booking)" />
        </ActivityCard>
      </div>
      <p v-else class="flex justify-center mt-4">There are no upcoming bookings.</p>
    </LayoutSection>

    <LayoutSection title="Search For Room">
      <div class="mt-4">
        <EmployeeSearchSection @bookingSubmitted="handleBookingSubmitted" />
      </div>
    </LayoutSection>

    <EmployeeList v-if="role === 'Manager'" @delete="showEmployeeDeletedToast" />
    <RoomList v-if="role === 'Manager'" @delete="showRoomDeletedToast" @update="showRoomUpdatedToast" @create="showRoomCreatedToast" />
    <CreateEmployeeModal v-if="isCreateEmployeeModalOpen && role === 'Manager'" @close="toggleCreateEmployeeModal"
      @created="showEmployeeCreatedToast" />
    <Profile v-if="isProfileModalOpen" role="employee" :toggleProfileModal="toggleProfileModal" />
    <HotelModal v-if="isHotelModalOpen" @close="toggleHotelModal" />
    <PaymentModal v-if="isPaymentModalOpen" :booking="selectedBooking" @confirm="confirmPayment" @close="closePaymentModal" />

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
import type { BookingItem, RentalItem, RentalWithBookingPayload } from '../types';
import EmployeeList from '../components/LandingPage/EmployeeList.vue';
import RoomList from '../components/LandingPage/RoomList.vue';
import type { ToastMessageOptions } from 'primevue';
import PaymentModal from '../components/LandingPage/PaymentModal.vue';
import EmployeeSearchSection from '../components/Search/EmployeeSearchSection.vue';


const toast = useToast();

const expandedCard = ref<{ section: string | null; index: number | null }>({ section: null, index: null });
const isProfileModalOpen = ref(false);
const isHotelModalOpen = ref(false);
const isCreateEmployeeModalOpen = ref(false);
const role = ref(getUserRole())

const isPaymentModalOpen = ref(false);
const selectedBooking = ref<BookingItem | null>(null);

const hotelRentals = ref<Array<RentalItem>>([])
const hotelBookings = ref<Array<BookingItem>>([])
const employeeID = getUserID()

function openPaymentModal(booking: BookingItem) {
  selectedBooking.value = booking;
  isPaymentModalOpen.value = true;
}

function closePaymentModal() {
  isPaymentModalOpen.value = false;
  selectedBooking.value = null;
}

async function confirmPayment(booking: BookingItem) {
  hotelBookings.value = hotelBookings.value.filter(b => b.booking_ID !== booking.booking_ID);

  const payload: RentalWithBookingPayload = {
    booking_ID: booking.booking_ID,
    employee_ID: employeeID,
  }

  closePaymentModal();

  try {
    const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/renting-from-booking`,
      {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(payload)
      }
    )

    if (res.ok) {
      toast.add({
        severity: 'success',
        summary: 'Booking Activated',
        detail: 'Moved to current rentals.',
        life: 3000
      });
      await fetchRentals()
      return
    }
    toast.add({
      severity: 'error',
      summary: 'Booking Activation Failed',
      detail: 'Booking could not be activated to a rental.',
      life: 3000
    });
  } catch (error) {
    console.error('Error calling API:', error);
    toast.add({
      severity: 'error',
      summary: 'Booking Activation Failed',
      detail: 'Booking could not be activated to a rental.',
      life: 3000
    });
  }
}

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
  toast.add({ severity: 'success', summary: 'Success', detail: 'Employee account created.', life: 3000 });
}

function showEmployeeDeletedToast(severity: ToastMessageOptions["severity"]) {
  toast.add({
    severity,
    summary: severity === 'success' ? 'Success' : 'Failed',
    detail: severity === 'success' ? 'Deleted employee account successfully.' : 'Failed to delete employee account.',
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
  roomNumber: string;
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

async function fetchRentals() {
  try {
    const rentalResponse = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/rentings?employee_ID=${employeeID}`)
    if (rentalResponse.ok) {
      hotelRentals.value = await rentalResponse.json()
    }
  } catch (error) {
    console.error('Error calling API:', error);
  }
}

async function fetchBookings() {
  try {
    const bookingResponse = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/bookings?employee_ID=${employeeID}`)
    if (bookingResponse.ok) {
      hotelBookings.value = await bookingResponse.json()
    }

  } catch (error) {
    console.error('Error calling API:', error);
  }
}

async function handleBookingSubmitted(severity: ToastMessageOptions["severity"], isRental: boolean) {
  if (isRental) {
    toast.add({
      severity: severity,
      summary: severity === 'success' ? 'Rental Created' : 'Failed',
      detail: severity === 'success' ? 'Successfully created new rental.' : 'Could not create new rental.',
      life: 3000
    });
    await fetchRentals()
    return
  }
  toast.add({
    severity: severity,
    summary: severity === 'success' ? 'Booking Created' : 'Failed',
    detail: severity === 'success' ? 'Successfully created new booking.' : 'Could not create new booking.',
    life: 3000
  });
  await fetchBookings()
}

onMounted(async () => {
  await fetchRentals()
  await fetchBookings()
})
</script>