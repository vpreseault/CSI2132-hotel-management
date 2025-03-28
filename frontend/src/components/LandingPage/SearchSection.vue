<template>
  <div>
    <!-- <Card class="w-full max-w-3xl mx-auto mt-6"> -->
      <!-- <template #content> -->
        <div class="flex flex-wrap gap-6 mb-3">
          <!-- Dates -->
          <div class="flex flex-col">
            <label class="font-medium">Dates</label>
            <DatePicker v-model="filters.dates" selectionMode="range" :manualInput="false" showIcon fluid iconDisplay="input" placeholder="Select Dates" />
          </div>
          
          <!-- City -->
          <div class="flex flex-col">
            <label class="font-medium">City</label>
            <InputText v-model="filters.city" placeholder="Enter City Name" class="w-full" />
          </div>

          <!-- Hotel Chain -->
          <div class="flex flex-col">
            <label class="font-medium">Hotel Chain</label>
            <Dropdown v-model="filters.chain" :options="chains" optionLabel="name" placeholder="Select Hotel Chain" class="w-full"/>
          </div>
        </div>

          <!-- Sliders -->
        <div class="flex flex-wrap gap-6 mb-3">
          <div class="flex items-center gap-4">
            <div class="flex-1">
              <label class="font-medium">Room Capacity: {{ filters.roomCapacity }}</label>
              <Slider v-model="filters.roomCapacity" :min="1" :max="10" class="w-56 mt-3" />
            </div>
          </div>
          <div class="flex items-center gap-4">
            <div class="flex-1">
              <label class="font-medium">Hotel Category: {{ filters.hotelCategory }}</label>
              <Slider v-model="filters.hotelCategory" :min="1" :max="5" class="w-56 mt-3" />
            </div>
          </div>
          <div class="flex items-center gap-4">
            <div class="flex-1">
              <label class="font-medium">Total Hotel Rooms: {{ filters.totalRooms }}</label>
              <Slider v-model="filters.totalRooms" :min="1" :max="25" class="w-56 mt-3" />
            </div>
          </div>
          <div class="flex items-center gap-4">
            <div class="flex-1">
              <label class="font-medium">Max Room Price: ${{ filters.roomPrice }}</label>
              <Slider v-model="filters.roomPrice" :min="50" :max="1000" class="w-56 mt-3" />
            </div>
          </div>
        </div>
        <!-- Search Button -->
        <div class="flex justify-center mb-3">
          <Button label="Search" icon="pi pi-search" @click="submitSearch" />
        </div>
      <!-- </template> -->
    <!-- </Card> -->

    <SearchCard :rooms="searchResults" :show="showDialog" @close="showDialog = false" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import Dropdown from 'primevue/dropdown';
import Slider from 'primevue/slider';
import Button from 'primevue/button';
import InputText from 'primevue/inputtext';
import SearchCard from './SearchCard.vue';

interface Room {
  hotel_name: string;
  room_number: number;
  capacity: number;
  price: number;
  view_type: string;
  extendable: boolean;
  damaged: boolean;
  chain_name: string;
  category: number;
  address: string;
  total_rooms: number;
}

const filters = ref({
  city: '',
  dates: [],
  chain: '',
  roomCapacity: 1,
  hotelCategory: 1,
  totalRooms: 5,
  roomPrice: 200
});

const chains = ref([
  { name: 'Hilton' },
  { name: 'Marriott' },
  { name: 'Holiday Inn' },
  { name: 'Best Western' },
]);

const searchResults = ref<Room[]>([]);
const showDialog = ref(false);

function submitSearch() {
  const { startDate, endDate } = filters.value;

  if (!startDate || !endDate) {
    alert('Please select both start and end dates.');
    return;
  }

  const oneDay = 24 * 60 * 60 * 1000;
  const duration = (endDate.getTime() - startDate.getTime()) / oneDay;

  if (duration < 1) {
    alert('Bookings must be at least 2 days long.');
    return;
  }

  searchResults.value = [
    {
      hotel_name: 'Marriott',
      room_number: 101,
      capacity: 3,
      price: 250,
      view_type: 'Sea',
      extendable: true,
      damaged: false,
      chain_name: 'Marriott Group',
      category: 4,
      address: '123 Ocean View Blvd, Miami, FL',
      total_rooms: 200
    },
    {
      hotel_name: 'Holiday Inn',
      room_number: 202,
      capacity: 2,
      price: 180,
      view_type: 'City',
      extendable: false,
      damaged: false,
      chain_name: 'IHG',
      category: 3,
      address: '789 City Center Ave, Orlando, FL',
      total_rooms: 150
    }
  ];

  showDialog.value = true;
}
</script>

<style scoped>
label {
  margin-bottom: 4px;
}
</style>
