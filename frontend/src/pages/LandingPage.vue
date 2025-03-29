<template>
    <div class="w-full m-auto">
      <NavBar role="customer" @toggleProfile="toggleProfileModal" />
  
      <div class="mt-16 text-center">
        <h1 class="text-2xl mt-4">Welcome to your one-stop shop for all things Hotel Room Rentings!</h1>
        <p class="mt-4">You are logged in!</p>
      </div>
  
      <LayoutSection title="Your Activity">
          <Accordion value="0">
            <AccordionPanel value="0">
                <AccordionHeader>Rentals</AccordionHeader>
                <AccordionContent>
                  <ActivityCard
                    v-if="customerRentals"
                    v-for="(rental, index) in customerRentals"
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
                  <p v-else class="text-center text-gray-500">You have no active rentals.</p>
                </AccordionContent>
            </AccordionPanel>
            <AccordionPanel value="1">
                <AccordionHeader>Bookings</AccordionHeader>
                <AccordionContent>
                  <ActivityCard
                    v-if="customerBookings"
                    v-for="(booking, index) in customerBookings"
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
                  <p v-else class="text-center text-gray-500">You have no upcoming bookings.</p>
                </AccordionContent>
            </AccordionPanel>
            <AccordionPanel value="2">
                <AccordionHeader>Archives</AccordionHeader>
                <AccordionContent>
                  <!-- <ActivityCard
                    v-for="(booking, index) in customerBookings"
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
                  /> -->
                </AccordionContent>
            </AccordionPanel>
        </Accordion>
      </LayoutSection>
      <Profile v-if="isProfileModalOpen" role="customer" :toggleProfileModal="toggleProfileModal" />
  
      <div class="mt-16 text-center">
        <button class="mt-4 px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition" @click="handleLogout"> Logout </button>
      </div>
    </div>
</template>
  
<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { getUserID, removeAuthCookie } from '../utils/auth';
import NavBar from '../components/LandingPage/NavBar.vue';
import Profile from '../components/LandingPage/Profile.vue';
import type { RentalItem, BookingItem } from '../types';

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

const customerRentals = ref<Array<RentalItem>>([])
const customerBookings = ref<Array<BookingItem>>([])
onMounted(async () => {
  const customerID = getUserID()
  
  try {
      const rentalResponse = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/rentings?customer_ID=${customerID}`)
      if (rentalResponse.ok) {
        customerRentals.value = await rentalResponse.json()
      }

      const bookingResponse = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/bookings?customer_ID=${customerID}`)
      if (bookingResponse.ok) {
        customerBookings.value = await bookingResponse.json()
      }

  } catch (error) {
      console.error('Error calling API:', error);
  }
})
</script>

<style lang="css" scoped>
.p-accordionpanel-active {
  margin: 0;
  border: 1px solid black;
}
</style>