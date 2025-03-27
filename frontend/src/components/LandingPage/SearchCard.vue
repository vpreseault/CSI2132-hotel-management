<template>
    <Dialog v-model:visible="visible" modal header="Search Results" class="w-full max-w-4xl" @hide="handleClose">
      <BookCard v-if="selectedRoom" :room="selectedRoom" @close="selectedRoom = null" @createBooking="handleBooking"/>
  
      <Card v-else>
        <template #title>Rooms</template>
        <template #content>
          <div v-if="rooms.length > 0" class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div v-for="(room, index) in rooms" :key="index" class="border p-4 rounded-xl shadow-md hover:bg-blue-50 transition">
              <p><strong>Hotel:</strong> {{ room.hotel_name }}</p>
              <p><strong>Room #:</strong> {{ room.room_number }}</p>
              <p><strong>Capacity:</strong> {{ room.capacity }}</p>
              <p><strong>Price:</strong> ${{ room.price }}</p>
              <p><strong>View:</strong> {{ room.view_type }}</p>
              <p><strong>Extendable:</strong> {{ room.extendable ? 'Yes' : 'No' }}</p>
              <p><strong>Damaged:</strong> {{ room.damaged ? 'Yes' : 'No' }}</p>
              <p><strong>Chain:</strong> {{ room.chain_name }}</p>
              <p><strong>Category:</strong> {{ room.category }}-star</p>
              <p><strong>Address:</strong> {{ room.address }}</p>
              <p><strong>Total Rooms:</strong> {{ room.total_rooms }}</p>
  
              <Button label="Book Now" icon="pi pi-book" class="mt-4 pl-2" @click="selectedRoom = room"/>
            </div>
          </div>
          <div v-else class="text-gray-600">No results found for your filters.</div>
        </template>
      </Card>
    </Dialog>
  </template>
  
  <script setup lang="ts">
  import { ref, watch, defineProps, defineEmits } from 'vue';
  import Dialog from 'primevue/dialog';
  import Card from 'primevue/card';
  import Button from 'primevue/button';
  import BookCard from './BookCard.vue';
  
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
  
  const props = defineProps<{
    rooms: Room[];
    show: boolean;
  }>();
  
  const emit = defineEmits<{
    (e: 'close'): void;
  }>();
  
  const visible = ref(props.show);
  const selectedRoom = ref<Room | null>(null);
  
  watch(() => props.show, (val) => {
    visible.value = val;
  });
  
  function handleClose() {
    emit('close');
  }
  
  function handleBooking(bookingData: any) {
    console.log('Booking confirmed:', bookingData);
    selectedRoom.value = null;
    visible.value = false;
  }
  </script>
  