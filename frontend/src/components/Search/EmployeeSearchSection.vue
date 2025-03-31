<template>
  <div>
    <div class="mb-5">
      <div class="flex flex-wrap justify-center gap-6 mb-5">
        <!-- Dates -->
        <div class="flex flex-col">
          <label class="font-medium">Dates</label>
          <DatePicker v-model="filters.dates" selectionMode="range" :manualInput="false" showIcon fluid iconDisplay="input" placeholder="Select Dates" />
          <Message v-if="error" severity="error" size="small" variant="simple">{{ error }}</Message>
        </div>
      </div>
  
    <!-- Sliders -->
      <div class="flex flex-wrap justify-center gap-6 mb-10">
        <div class="flex items-center gap-4">
          <div class="flex-1">
            <label class="font-medium">Max Room Price: {{filters.roomPrice ? `$${filters.roomPrice}`:'' }}</label>
            <Slider v-model="filters.roomPrice" :min="5" :max="500" :step="5" class="w-56 mt-3" />
          </div>
        </div>
        <div class="flex items-center gap-4">
          <div class="flex-1">
            <label class="font-medium">Minimum Room Capacity: {{ filters.roomCapacity }}</label>
            <Slider v-model="filters.roomCapacity" :min="1" :max="10" class="w-56 mt-3" />
          </div>
        </div>
      </div>
      <!-- Search Button -->
      <div class="flex justify-center gap-6">
        <Button label="Search" icon="pi pi-search" @click="submitSearch" />
        <Button label="Clear Filters" severity="secondary" icon="pi pi-times" @click="clearFilters" />
      </div>
    </div>
    
    <div v-if="searchResults?.length > 0 && showResults" class="grid grid-cols-1 md:grid-cols-2 gap-4">
      <template v-for="(room) in searchResults" :key="`${room.chain_name}-${room.hotel_ID}`">
        <SearchCard 
          :hotel_name="room.hotel_name"
          :chain_name="room.chain_name"
          :address="room.address"
          :view_type="room.view_type"
          :room_number="room.room_number"
          :capacity="room.capacity"
          :price="room.price"
          :category="room.category"
          :total_rooms="room.total_rooms"
          :extendable="room.extendable"
          :damaged="room.damaged"
          @show-booking-modal="showBookingModal(room)"
        />
      </template>
    </div>
    <div v-else-if="showResults">
      <hr class="mb-8">
      <p class="text-gray-600 text-center">No results found for your filters.</p>
    </div>
    <Dialog v-model:visible="bookingModalIsVisible" modal header="Search Results" class="w-full max-w-4xl" @hide="bookingModalIsVisible= false">
      <BookCard
        :room="selectedRoom"
        :start_date="filters.dates[0]"
        :end_date="filters.dates[1]"
        employee
        @close="bookingModalIsVisible = false"
        @bookingSubmitted="handleBookingSubmitted"
      />
    </Dialog>
  </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import SearchCard from './SearchCard.vue'
import Slider from 'primevue/slider';
import Button from 'primevue/button';
import type { SearchResult } from '../../types';
import type { ToastMessageOptions } from 'primevue';
import { getUserID } from '../../utils/auth';

const emit = defineEmits<{
  bookingSubmitted: [severity: ToastMessageOptions["severity"], isRental: boolean]
}>()

const filters = reactive<{
  dates: Date[],
  roomCapacity?: number,
  roomPrice?: number,
}>({
  dates: [],
});

function clearFilters() {
  filters.dates = []
  filters.roomCapacity = undefined
  filters.roomPrice = undefined

  searchResults.value = []
  showResults.value = false
}

const searchResults = ref<SearchResult[]>([]);
const bookingModalIsVisible = ref(false)

const error = ref('')
function validate() {
  error.value = ''
  if (!filters?.dates || !filters.dates[0] || !filters.dates[1]) {
    error.value  = 'Dates are required.'
    return false
  }
  filters.dates[0].setHours(0,0,0,0)
  filters.dates[1].setHours(0,0,0,0)
  const oneDay = 24 * 60 * 60 * 1000;
  const duration = (filters.dates[1].getTime() - filters.dates[0].getTime()) / oneDay;
  if (duration < 1) {
    error.value = 'Booking duration must be at least 1 night.'
    return false
  } 

  const today = new Date()
  today.setHours(0, 0, 0, 0)
  if (filters.dates[0] < today) {
    error.value = 'Dates cannot be in the past.'
    return false
  }
  
  return true
}

const employeeID = getUserID()
const showResults = ref(false) 
async function submitSearch() {
  if (validate()) {
      try {
        const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/search/employee`,
            {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                  "employee_ID": employeeID,
                  "start_date": filters.dates[0],
                  "end_date": filters.dates[1],
                  "capacity": filters.roomCapacity,
                  "max_price": filters.roomPrice,
                })
            }
        )
        if (res.ok) {
            searchResults.value = await res.json()
            showResults.value = true
        }
      } catch (error) {
          console.error('Error calling API:', error);
      }
  } else {
    showResults.value = false
  }
}

const selectedRoom = ref<SearchResult>()
function showBookingModal(room: SearchResult) {
  selectedRoom.value = room
  bookingModalIsVisible.value = true
}

async function handleBookingSubmitted(severity: ToastMessageOptions["severity"], isRental: boolean) {
  bookingModalIsVisible.value = false
  await submitSearch()
  emit('bookingSubmitted', severity, isRental)
}
</script>

<style scoped>
label {
  margin-bottom: 4px;
}
</style>
