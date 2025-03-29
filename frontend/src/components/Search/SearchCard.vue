<template>
    <div class="border p-4 rounded-xl shadow-md hover:bg-blue-50 transition">
        <p><strong>Hotel:</strong> {{ hotel_name }}</p>
        <p><strong>Room #:</strong> {{ room_number }}</p>
        <p><strong>Capacity:</strong> {{ capacity }}</p>
        <p><strong>Price:</strong> ${{ price }}</p>
        <p><strong>View:</strong> {{ view_type }}</p>
        <p><strong>Extendable:</strong> {{ extendable ? 'Yes' : 'No' }}</p>
        <p><strong>Damaged:</strong> {{ damaged ? 'Yes' : 'No' }}</p>
        <p><strong>Chain:</strong> {{ chain_name }}</p>
        <p><strong>Category:</strong> {{ category }}-star</p>
        <p><strong>Address:</strong> {{ address }}</p>
        <p><strong>Total Rooms:</strong> {{ total_rooms }}</p>

        <Button label="Book Now" icon="pi pi-book" class="mt-4 pl-2" @click="emit('showBookingModal')" />
    </div>
</template>
<script setup lang="ts">
defineProps<{
    hotel_name: string
    room_number: string
    view_type: string
    chain_name: string
    address: string
    price: number
    category: number
    capacity: number
    total_rooms: number
    damaged: boolean
    extendable: boolean
}>()

const emit = defineEmits<{
    showBookingModal: []
}>()
</script>