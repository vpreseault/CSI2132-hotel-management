<template>
    <div class="w-full m-auto">
      <div class="mt-16 text-center">
        <h1 class="text-2xl mt-4">Admin View</h1>
      </div>
  
      <NavBar role="admin" @exit-admin="exitAdminView" />
  
      <div class="my-8 p-8 bg-green-100 rounded-lg shadow-md xl:mx-24">
          <Accordion>
            <AccordionPanel value="0">
                <AccordionHeader>All Hotels</AccordionHeader>
                <AccordionContent>
                  <HotelList @hotelsUpdated="fetchViewsData" />
                </AccordionContent>
            </AccordionPanel>
            <AccordionPanel value="1">
                <AccordionHeader>Rooms Per Area</AccordionHeader>
                <AccordionContent>
                  <div class="grid grid-cols-2 gap-4">
                    <div v-for="area in viewsData.rooms_per_area" :key="area.area" class="mb-2">
                      <span class="font-medium">{{ area.area }}:</span> {{ area.total_rooms }} rooms
                    </div>
                  </div>
                </AccordionContent>
            </AccordionPanel>
            <AccordionPanel value="2">
                <AccordionHeader>Capacity Per Hotel</AccordionHeader>
                <AccordionContent>
                  <div class="grid grid-cols-2 gap-4">
                    <div v-for="hotel in viewsData.hotel_capacity" :key="hotel.hotel_name" class="mb-2">
                      <span class="font-medium">{{ hotel.hotel_name }}:</span> {{ hotel.total_capacity }}
                    </div>
                  </div>
                </AccordionContent>
            </AccordionPanel>
        </Accordion>
      </div>
  
      <Toast position="bottom-center" />
    </div>
</template>

<script setup lang="ts">
import { removeAuthCookie } from '../utils/auth'
import NavBar from '../components/LandingPage/NavBar.vue'
import HotelList from '../components/Admin/HotelList.vue'  
import { ref, onMounted } from 'vue'
import { useToast } from 'primevue/usetoast'

function exitAdminView() {
  handleLogout()
}

function handleLogout() {
  removeAuthCookie()
  window.location.href = '/'
}

interface ViewsData {
  rooms_per_area: {
    area: string
    total_rooms: number
  }[]
  hotel_capacity: {
    hotel_name: string
    total_capacity: number
  }[]
}

const viewsData = ref<ViewsData>({
  rooms_per_area: [],
  hotel_capacity: []
})

const toast = useToast()

async function fetchViewsData() {
  try {
    const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/admin/views`)
    if (!res.ok) {
      throw new Error('Failed to fetch views data')
    }
    viewsData.value = await res.json()
  } catch (error) {
    toast.add({
      severity: 'error',
      summary: 'Error',
      detail: 'Failed to fetch views data',
      life: 3000
    })
  }
}
</script>
