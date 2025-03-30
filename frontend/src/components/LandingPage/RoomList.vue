<template>
    <div id="room-list" class="mt-16 p-8 bg-green-100 rounded-lg shadow-md mx-75">
        <h2 class="text-center text-xl text-gray-600 font-semibold">All Rooms</h2>

        <div class="text-center mt-4">
            <Button label="Create Room" icon="pi pi-plus" class="bg-blue-500 border-none hover:bg-blue-600"
                @click="openCreateModal" />
        </div>

        <div v-if="rooms.length" class="mt-6 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 justify-center">
            <RoomCard v-for="room in rooms" :key="room.room_number" :room="room" @delete="deleteRoom"
                @edit="updateRoom" />
        </div>
        <p v-else class="text-center text-gray-500 mt-8">No rooms found.</p>

        <RoomModal v-if="isCreateModalOpen" :room="newRoom" :onClose="closeCreateModal" :onSave="handleCreateRoom" />
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import type { ToastMessageOptions } from 'primevue'
import Button from 'primevue/button'
import RoomCard from './RoomCard.vue'
import RoomModal from './RoomModal.vue'
import { getUserID } from '../../utils/auth'
import type { Room } from '../../types'

const emit = defineEmits<{
    delete: [severity: ToastMessageOptions["severity"]]
    update: [severity: ToastMessageOptions["severity"]]
    create: [severity: ToastMessageOptions["severity"]]
}>()

const rooms = ref<Room[]>([])
const isCreateModalOpen = ref(false)
const managerID = ref(getUserID())

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

async function handleCreateRoom(room: Room) {
    isCreateModalOpen.value = false
    try {
        const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/rooms`,
            {
                method: 'POST',
                body: JSON.stringify({
                    ...room,
                    employee_ID: managerID.value,
                })
            }
        )
        
        if (res.ok) {
            rooms.value.push({ ...room })
            emit('create', 'success')
            return
        }
        emit('create', 'error')
    } catch (error) {
        console.error('Error calling API:', error);
        emit('create', 'error')
    }
}

async function updateRoom(updatedRoom: Room) {
    try {
        const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/rooms/${updatedRoom.room_ID}`,
            {
                method: 'PATCH',
                headers: {
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify(updatedRoom)
            }
        )
    
        if (res.ok) {
            emit('update', 'success')

            const index = rooms.value.findIndex(r => r.room_number === updatedRoom.room_number)
            if (index !== -1) {
                rooms.value[index] = { ...updatedRoom }
            }
            return
        }

        emit('update', 'error')
    } catch (error) {
        console.error('Error calling API:', error);
        emit('update', 'error')
    }
}

async function deleteRoom(id: number) {
    try {
        const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/rooms/${id}`, {
            method: 'DELETE'
        })

        if (res.ok) {
            rooms.value = rooms.value.filter(room => room.room_ID !== id)
            emit('delete', 'success')
        } else {
            emit('delete', 'error')
        }
    } catch (error) {
        console.error('Error deleting room:', error)
        emit('delete', 'error')
    }
}

onMounted(async () => {
    await fetchRooms()
})

async function fetchRooms() {
    try {
        const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/rooms?employee_ID=${managerID.value}`)
        if (res.ok) {
            rooms.value = await res.json()
        }
    } catch (error) {
        console.error('Error fetching rooms:', error)
    }
}
</script>