<template>
    <div class="bg-white p-4 rounded-lg shadow-md relative">
      <h3 class="text-lg font-semibold text-gray-800 mb-1">{{ hotel.hotel_name }}</h3>
      <p class="text-gray-600">Address: {{ hotel.address }}</p>
      <p class="text-gray-600">Phone: {{ hotel.phone_number }}</p>
      <p class="text-gray-600">Email: {{ hotel.email }}</p>
      <p class="text-gray-600">Category: {{ hotel.category }}</p>
  
      <div class="flex justify-end mt-4">
        <button class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600 transition" @click="confirmDelete"> Delete Hotel </button>
      </div>
    </div>
</template>

<script setup lang="ts">
import type { HotelDisplay } from '../../types'

const props = defineProps<{
hotel: HotelDisplay
}>()

const emit = defineEmits(['delete'])

function confirmDelete() {
const confirmed = confirm(`Are you sure you want to delete ${props.hotel.hotel_name}?`)
if (confirmed) {
    emit('delete', props.hotel.hotel_ID)
}
}
</script>
  