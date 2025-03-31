<template>
    <div class="w-full m-auto">
      <NavBar role="customer" @toggleProfile="toggleProfileModal" />
  
      <div class="mt-16 text-center">
        <h1 class="text-2xl mt-4">Welcome to your one-stop shop for all things Hotel Room Rentings!</h1>
        <p class="mt-4">You are logged in!</p>
      </div>
      
      <LayoutSection title="Search For Room">
        <div class="mt-4">
          <SearchSection />
        </div>
      </LayoutSection>
      
      <LayoutSection title="Your Activity">
          <Accordion value="0">
            <AccordionPanel value="0">
                <AccordionHeader>Rentals</AccordionHeader>
                <AccordionContent>
                  <ActivityCard
                    v-if="customerRentals?.length"
                    v-for="(rental, index) in customerRentals"
                    :key="`rental-${index}`"
                    :customerName="rental.customer_name"
                    :hotelName="rental.hotel_name"
                    :startDate="new Date(rental.check_in_date)"
                    :endDate="new Date(rental.check_out_date)"
                    :employeeName="rental.employee_name"
                    :roomNumber="rental.room_number"
                    :price="rental.total_price"
                    :payment="rental.payment"
                    :section="'rental'"
                    :index="index"
                    :expandedCard="expandedCard"
                    @toggle="toggleCard"
                  />
                  <NoResultsLabel v-else>You have no active rentals.</NoResultsLabel>
                </AccordionContent>
            </AccordionPanel>
            <AccordionPanel value="1">
                <AccordionHeader>Bookings</AccordionHeader>
                <AccordionContent>
                  <ActivityCard
                    v-if="customerBookings?.length"
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
                  <NoResultsLabel v-else>You have no upcoming bookings.</NoResultsLabel>
                </AccordionContent>
            </AccordionPanel>
            <AccordionPanel value="2">
                <AccordionHeader>Archives</AccordionHeader>
                <AccordionContent>
                  <ArchiveCard
                    v-if="customerArchives?.length"
                    v-for="(archive, index) in customerArchives"
                    :key="`archive-${index}`"
                    :id="archive.archive_ID"
                    :startDate="new Date(archive.start_date)"
                    :endDate="new Date(archive.end_date)"
                    :price="archive.total_price"
                    :section="'archive'"
                    :index="index"
                    :expandedCard="expandedCard"
                    @toggle="toggleCard"
                  />
                  <NoResultsLabel v-else>You have no archives.</NoResultsLabel>
                </AccordionContent>
            </AccordionPanel>
        </Accordion>
      </LayoutSection>
      
      <Profile v-if="isProfileModalOpen" role="customer" :toggleProfileModal="toggleProfileModal" />
    </div>
</template>
  
<script setup lang="ts">
import { onMounted, ref } from 'vue';
import { getUserID } from '../utils/auth';
import NavBar from '../components/LandingPage/NavBar.vue';
import Profile from '../components/LandingPage/Profile.vue';
import type { RentalItem, BookingItem, ArchiveItem } from '../types';
import NoResultsLabel from '../components/Layout/NoResultsLabel.vue';
import ArchiveCard from '../components/LandingPage/ArchiveCard.vue';

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

const customerRentals = ref<Array<RentalItem>>([])
const customerBookings = ref<Array<BookingItem>>([])
const customerArchives = ref<Array<ArchiveItem>>([])
onMounted(async () => {
  const customerID = getUserID()
  
  try {
      const activityResponse = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/activity?customer_ID=${customerID}`)
      if (activityResponse.ok) {
        const {bookings, rentings, archives} = await activityResponse.json()
        customerBookings.value = bookings
        customerRentals.value = rentings
        customerArchives.value = archives
      }
  } catch (error) {
      console.error('Error calling API:', error);
  }
})
</script>