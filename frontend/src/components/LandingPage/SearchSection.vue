<template>
  <div>
    <Card class="w-full max-w-3xl mx-auto mt-6">
      <template #title>Search Filters</template>
      <template #content>
        <div class="flex flex-col gap-6">
          <!-- City -->
          <div class="flex flex-col gap-1">
            <label class="font-medium">City</label>
            <InputText v-model="filters.city" placeholder="Enter City Name" class="w-full md:w-80" />
          </div>

          <!-- Dates -->
          <div class="flex flex-col md:flex-row gap-4">
            <div class="flex flex-col">
              <label class="font-medium">Start Date</label>
              <Calendar v-model="filters.startDate" showIcon placeholder="Start Date" class="w-full md:w-56" />
            </div>
            <div class="flex flex-col">
              <label class="font-medium">End Date</label>
              <Calendar
                v-model="filters.endDate"
                :minDate="minEndDate"
                showIcon
                placeholder="End Date"
                class="w-full md:w-56"
              />
            </div>
          </div>

          <!-- Hotel Chain -->
          <div class="flex flex-col gap-1">
            <label class="font-medium">Hotel Chain</label>
            <Dropdown
              v-model="filters.chain"
              :options="chains"
              optionLabel="name"
              placeholder="Select or Type Hotel Chain"
              editable
              class="w-full md:w-80"
            />
          </div>

          <!-- Sliders -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
            <div class="flex items-center gap-4">
              <div class="flex-1">
                <label class="font-medium">Room Capacity</label>
                <Slider v-model="filters.roomCapacity" :min="1" :max="10" class="w-56" />
              </div>
              <span>{{ filters.roomCapacity }}</span>
            </div>
            <div class="flex items-center gap-4">
              <div class="flex-1">
                <label class="font-medium">Hotel Category</label>
                <Slider v-model="filters.hotelCategory" :min="1" :max="5" class="w-56" />
              </div>
              <span>{{ filters.hotelCategory }}</span>
            </div>
            <div class="flex items-center gap-4">
              <div class="flex-1">
                <label class="font-medium">Total Rooms</label>
                <Slider v-model="filters.totalRooms" :min="1" :max="5" class="w-56" />
              </div>
              <span>{{ filters.totalRooms }}</span>
            </div>
            <div class="flex items-center gap-4">
              <div class="flex-1">
                <label class="font-medium">Room Price</label>
                <Slider v-model="filters.roomPrice" :min="50" :max="1000" class="w-56" />
              </div>
              <span>${{ filters.roomPrice }}</span>
            </div>
          </div>

          <!-- Search Button -->
          <div class="flex justify-end mt-4">
            <Button label="Search" icon="pi pi-search" @click="submitSearch" />
          </div>
        </div>
      </template>
    </Card>

    <SearchCard :rooms="searchResults" :show="showDialog" @close="showDialog = false" />
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import Calendar from 'primevue/calendar';
import Dropdown from 'primevue/dropdown';
import Slider from 'primevue/slider';
import Button from 'primevue/button';
import Card from 'primevue/card';
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
  startDate: undefined as Date | undefined,
  endDate: undefined as Date | undefined,
  chain: '',
  roomCapacity: 1,
  hotelCategory: 1,
  totalRooms: 50,
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

const minEndDate = computed(() => {
  if (!filters.value.startDate) return undefined;
  const min = new Date(filters.value.startDate);
  min.setDate(min.getDate() + 1);
  return min;
});

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
