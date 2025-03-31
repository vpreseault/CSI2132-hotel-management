<template>
    <div class="mt-16 p-8 bg-green-100 rounded-lg shadow-md mx-75">
      <h2 class="text-center text-xl text-gray-600 font-semibold">All Hotels</h2>
  
      <div class="text-center mt-4">
        <Button label="Create Hotel" icon="pi pi-plus" class="bg-blue-500 border-none hover:bg-blue-600" @click="openCreateModal" />
      </div>
  
      <div v-if="hotels.length" class="mt-6 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 justify-center">
        <HotelCard
          v-for="hotel in hotels"
          :key="hotel.hotel_ID"
          :hotel="hotel"
          @delete="deleteHotel"
        />
      </div>
      <p v-else class="text-center text-gray-500 mt-8">No hotels found.</p>
  
      <CreateHotels
        v-if="isCreateModalOpen"
        :onClose="closeCreateModal"
        :onCreate="handleHotelCreated"
      />
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import Button from 'primevue/button'
import CreateHotels from './CreateHotels.vue'
import HotelCard from './HotelCard.vue'
import type { HotelDisplay } from '../../types'

const hotels = ref<HotelDisplay[]>([])

const isCreateModalOpen = ref(false)

function openCreateModal() {
isCreateModalOpen.value = true
}

function closeCreateModal() {
isCreateModalOpen.value = false
}

function handleHotelCreated(hotel: HotelDisplay) {
  hotels.value.push(hotel)
}

function deleteHotel(id: number) {
hotels.value = hotels.value.filter(h => h.hotel_ID !== id)
}
</script>
  