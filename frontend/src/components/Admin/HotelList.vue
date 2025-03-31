<template>
  <div>
    <div class="text-center pt-4">
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
      :chains="chains"
      @close="closeCreateModal"
      @create="handleHotelCreated"
    />
  </div>
  <Toast position="bottom-center" />
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import Button from 'primevue/button'
import CreateHotels from './CreateHotels.vue'
import HotelCard from './HotelCard.vue'
import type { Chain, HotelDisplay } from '../../types'
import { useToast } from "primevue/usetoast";
import type { ToastMessageOptions } from 'primevue'
const toast = useToast();

const hotels = ref<HotelDisplay[]>([])

const chains = ref<Chain[]>([]);
onMounted(async () => {  
  try {
      const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/chains`)
      if (res.ok) {
        chains.value = await res.json()
      }
      await fetchHotels()
  } catch (error) {
      console.error('Error calling API:', error);
  }
})

async function fetchHotels() {
  const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/all-hotels`)
  if (res.ok) {
    hotels.value = await res.json()
  }
}

const isCreateModalOpen = ref(false)

function openCreateModal() {
isCreateModalOpen.value = true
}

function closeCreateModal() {
isCreateModalOpen.value = false
}

async function handleHotelCreated(severity: ToastMessageOptions["severity"]) {
  closeCreateModal()
  toast.add({ 
    severity, 
    summary: severity === 'success' ? 'Success' : 'Failed', 
    detail: severity === 'success' ? 'Created hotel successfully.' : 'Failed to create hotel.', 
    life: 3000 
  })
  await fetchHotels()
}

async function deleteHotel(id: number) {
  try {
    const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/hotels/${id}`,
        {
            method: 'DELETE',
        }
    )
    
    if (res.ok) {
        toast.add({
          severity: 'success',
          summary: 'Success',
          detail: 'Hotel was deleted successfully.',
          life: 3000
        });
        await fetchHotels()
        return
    }
    toast.add({
      severity: 'error',
      summary: 'Failed',
      detail: 'Failed to delete hotel.',
      life: 3000
    });
  } catch (error) {
      console.error('Error calling API:', error);
  }
}
</script>
  