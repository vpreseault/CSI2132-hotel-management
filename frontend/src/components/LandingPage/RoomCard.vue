<template>
  <div class="bg-white p-4 rounded-lg shadow-md relative flex flex-col justify-between min-h-[320px]">
    <h3 class="text-lg font-semibold text-gray-800 mb-1">Room {{ room.room_number }}</h3>
    <p class="text-gray-600">Capacity: {{ room.capacity }}</p>
    <p class="text-gray-600">Price: ${{ room.price }}</p>
    <p class="text-gray-600">View: {{ room.view_type }}</p>
    <p class="text-gray-600">Extendable: {{ room.extendable ? 'Yes' : 'No' }}</p>
    <p class="text-gray-600">Damaged: {{ room.damaged ? 'Yes' : 'No' }}</p>
    <p class="text-gray-600">Amenities:
      <span v-if="room.amenities.length" v-for="(amenity, index) in room.amenities" :key="amenity.amenity_ID">
        {{ `${index !== 0 ? ', ' : ''}${amenity.amenity_name}` }}
      </span>
      <span v-else>N/A</span>
    </p>

    <div class="flex justify-between mt-4">
      <button class="bg-blue-500 text-white text-sm px-4 py-2 rounded hover:bg-blue-600 transition" @click="isModalOpen = true"> Edit Room </button>
      <button class="bg-red-500 text-white text-sm px-4 py-2 rounded hover:bg-red-600 transition" @click="confirmDelete"> Delete Room </button>
    </div>

    <RoomModal
      v-if="isModalOpen"
      :room="room"
      :allAmenities="allAmenities"
      :existingRooms="props.existingRooms"
      :onClose="() => isModalOpen = false"
      :onSave="handleRoomSave"
    />
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import RoomModal from './RoomModal.vue'
import type { Amenity, Room } from '../../types';

const props = defineProps<{
  room: Room
  allAmenities: Amenity[]
  existingRooms: Room[]
}>()


const emit = defineEmits(['delete', 'edit'])

const isModalOpen = ref(false)

function confirmDelete() {
  const confirmed = confirm(`Are you sure you want to delete Room ${props.room.room_number}?`)
  if (confirmed) {
    emit('delete', props.room.room_ID)
  }
}

function handleRoomSave(updatedRoom: Room) {
  emit('edit', updatedRoom)
}
</script>