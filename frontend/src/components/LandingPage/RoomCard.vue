<template>
    <div class="bg-white p-4 rounded-lg shadow-md relative">
      <h3 class="text-lg font-semibold text-gray-800 mb-1">Room {{ room.room_number }}</h3>
      <p class="text-gray-600">Capacity: {{ room.capacity }}</p>
      <p class="text-gray-600">Price: ${{ room.price }}</p>
      <p class="text-gray-600">View: {{ room.view_type }}</p>
      <p class="text-gray-600">Extendable: {{ room.extendable ? 'Yes' : 'No' }}</p>
      <p class="text-gray-600">Damaged: {{ room.damaged ? 'Yes' : 'No' }}</p>
      <p class="text-gray-600">Amenities: 
        <span v-if="room.amenities.length">{{ room.amenities.join(', ') }}</span>
        <span v-else>None</span>
      </p>
  
      <div class="flex justify-between mt-4">
        <button class="bg-blue-500 text-white px-4 py-2 rounded hover:bg-blue-600 transition" @click="isModalOpen = true"> Edit Room </button>
        <button class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-600 transition" @click="confirmDelete"> Delete Room </button>
      </div>
  
      <RoomModal v-if="isModalOpen" :room="room" :onClose="() => isModalOpen = false" :onSave="handleRoomSave" />
    </div>
  </template>
  
  <script setup lang="ts">
  import { ref } from 'vue'
  import RoomModal from './RoomModal.vue'
  
  const props = defineProps<{
    room: {
      room_number: string;
      capacity: number;
      price: number;
      view_type: string;
      extendable: boolean;
      damaged: boolean;
      amenities: string[];
    };
  }>()
  
  const emit = defineEmits(['delete', 'edit'])
  
  const isModalOpen = ref(false)
  
  function confirmDelete() {
    const confirmed = confirm(`Are you sure you want to delete Room ${props.room.room_number}?`)
    if (confirmed) {
      emit('delete', props.room.room_number)
    }
  }
  
  function handleRoomSave(updatedRoom: typeof props.room) {
    emit('edit', updatedRoom)
  }
  </script>
  