<template>
    <div id="room-list" class="mt-16 p-8 bg-green-100 rounded-lg shadow-md mx-75">
      <h2 class="text-center text-xl text-gray-600 font-semibold">All Rooms</h2>
  
      <div class="text-center mt-4">
        <Button label="Create Room" icon="pi pi-plus" class="bg-blue-500 border-none hover:bg-blue-600" @click="openCreateModal" />
      </div>
  
      <div v-if="rooms.length" class="mt-6 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 justify-center">
        <RoomCard v-for="room in rooms" :key="room.room_number" :room="room" @delete="deleteRoom" @edit="updateRoom"/>
      </div>
      <p v-else class="text-center text-gray-500 mt-8">No rooms found.</p>

      <RoomModal v-if="isCreateModalOpen" :room="newRoom" :onClose="closeCreateModal" :onSave="handleCreateRoom"/>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { ToastMessageOptions } from 'primevue'
import Button from 'primevue/button'
import RoomCard from './RoomCard.vue'
import RoomModal from './RoomModal.vue'
import { getUserID } from '../../utils/auth'

type Room = {
room_number: string;
capacity: number;
price: number;
view_type: string;
extendable: boolean;
damaged: boolean;
amenities: string[];
};

const emit = defineEmits<{
delete: [severity: ToastMessageOptions["severity"]]
}>()

const rooms = ref<Room[]>([])
const isCreateModalOpen = ref(false)

const newRoom = ref<Room>({
room_number: '',
capacity: 1,
price: 100,
view_type: '',
extendable: false,
damaged: false,
amenities: []
})

function openCreateModal() {
isCreateModalOpen.value = true
}

function closeCreateModal() {
isCreateModalOpen.value = false
}

function handleCreateRoom(room: Room) {
rooms.value.push({ ...room })
isCreateModalOpen.value = false
}

function updateRoom(updatedRoom: Room) {
const index = rooms.value.findIndex(r => r.room_number === updatedRoom.room_number)
if (index !== -1) {
    rooms.value[index] = { ...updatedRoom }
}
}

const managerID = ref(getUserID())
async function deleteRoom(roomNumber: string) {
try {
    const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/rooms?room_number=${roomNumber}`, {
    method: 'DELETE'
    })

    if (res.ok) {
    rooms.value = rooms.value.filter(room => room.room_number !== roomNumber)
    emit('delete', 'success')
    } else {
    emit('delete', 'error')
    }
} catch (error) {
    console.error('Error deleting room:', error)
    emit('delete', 'error')
}
}

onMounted(() => {
rooms.value = [
    {
    room_number: "101A",
    capacity: 2,
    price: 120,
    view_type: "Ocean",
    extendable: true,
    damaged: false,
    amenities: ["WiFi", "TV", "Mini Fridge"]
    },
    {
    room_number: "202B",
    capacity: 4,
    price: 200,
    view_type: "Mountain",
    extendable: false,
    damaged: false,
    amenities: ["WiFi", "Balcony"]
    },
    {
    room_number: "303C",
    capacity: 1,
    price: 90,
    view_type: "City",
    extendable: false,
    damaged: true,
    amenities: ["WiFi"]
    }
]

fetchRooms()
})

async function fetchRooms() {
try {
    const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/rooms?manager_ID=${managerID.value}`)
    if (res.ok) {
    rooms.value = await res.json()
    }
} catch (error) {
    console.error('Error fetching rooms:', error)
}
}
</script>
  