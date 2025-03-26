<template>
    <div id="book" class="mt-16 p-8 bg-green-100 rounded-lg shadow-md">
      <h2 class="text-center text-xl text-gray-600 font-semibold">
        {{ isEmployee ? 'Make a Booking' : 'Book a Hotel' }}
      </h2>
  
      <div v-if="!isEmployee" class="mt-4 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-4 justify-center">
        <GenericCard 
          v-for="(hotel, index) in hotels" 
          :key="index"
          :title="hotel.title" 
          :description="hotel.description" 
          :details="formatDetails(hotel.details)"
          :section="'book'"
          :index="index"
          :expandedCard="expandedCard"
          @toggle="toggleCard"
        />
      </div>
  
      <form v-else @submit.prevent="submitBooking" class="mt-4" >
        <input v-model="employeeBooking.customerName" type="text" placeholder="Customer Name" class="border p-2 text-black rounded w-full mb-2" required />
        <input v-model.number="employeeBooking.roomNumber" type="number" placeholder="Room Number" class="border p-2 text-black rounded w-full mb-2" required />
        <input v-model="employeeBooking.checkoutDate" type="date" class="border p-2 text-black rounded w-full mb-2" required />
        <input v-model="employeeBooking.contactEmail" type="contactEmail" placeholder="Email" class="border p-2 text-black rounded w-full mb-2" required />
        <button type="submit" class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition" > Book Room </button>
      </form>
    </div>
</template>
  
<script setup lang="ts">
import { ref } from 'vue';
import GenericCard from './GenericCard.vue';

defineProps<{
expandedCard: { section: string | null; index: number | null };
toggleCard: (section: string, index: number) => void;
isEmployee?: boolean;
}>();

const emit = defineEmits(['createBooking']);
  
const hotels = [
  {
  title: "Hotel 1",
  description: "Description",
  details: {
  cost: "$250 per night",
  rating: "5 stars",
  contactEmail: "contact@hotel1.com",
  phone: "+1-555-1234"
  }
  },
  {
  title: "Hotel 2",
  description: "Description",
  details: {
  cost: "$120 per night",
  rating: "3 stars",
  contactEmail: "info@hotel2.com",
  phone: "+1-555-5678"
  }
  },
  {
  title: "Hotel 3",
  description: "Description",
  details: {
  cost: "$180 per night",
  rating: "4 stars",
  contactEmail: "reservations@hotel3.com",
  phone: "+1-555-9876"
  }
  },
  {
  title: "Hotel 4",
  description: "Description",
  details: {
  cost: "$220 per night",
  rating: "4.5 stars",
  contactEmail: "bookings@hotel4.com",
  phone: "+1-555-4321"
  }
  },
  {
  title: "Hotel 5",
  description: "Description",
  details: {
  cost: "$300 per night",
  rating: "5 stars",
  contactEmail: "luxury@hotel5.com",
  phone: "+1-555-2468"
  }
  }
];
  
const employeeBooking = ref({
customerName: '',
roomNumber: '',
checkoutDate: '',
contactEmail: ''
});
  
function submitBooking() {
if (
  employeeBooking.value.customerName &&
  employeeBooking.value.roomNumber &&
  employeeBooking.value.checkoutDate
) {
  emit('createBooking', { ...employeeBooking.value });
  employeeBooking.value = { customerName: '', roomNumber: '', checkoutDate: '', contactEmail: '' };
}
}
  
function formatDetails(details: any) {
return `
  Cost: ${details.cost}
  Rating: ${details.rating}
  Email: ${details.contactEmail}
  Phone: ${details.phone}
`;
}
</script>
  