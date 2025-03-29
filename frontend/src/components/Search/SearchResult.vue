<template>
    <div class="border p-4 rounded-xl shadow-md hover:bg-blue-50 transition">
        <p><strong>Hotel:</strong> {{ hotelName }}</p>
        <p><strong>Room #:</strong> {{ roomNumber }}</p>
        <p><strong>Capacity:</strong> {{ capacity }}</p>
        <p><strong>Price:</strong> ${{ price }}</p>
        <p><strong>View:</strong> {{ viewType }}</p>
        <p><strong>Extendable:</strong> {{ extendable ? 'Yes' : 'No' }}</p>
        <p><strong>Damaged:</strong> {{ damaged ? 'Yes' : 'No' }}</p>
        <p><strong>Chain:</strong> {{ chainName }}</p>
        <p><strong>Category:</strong> {{ category }}-star</p>
        <p><strong>Address:</strong> {{ address }}</p>
        <p><strong>Total Rooms:</strong> {{ totalRooms }}</p>

        <Button label="Book Now" icon="pi pi-book" class="mt-4 pl-2" @click="" />
    </div>
</template>
<script setup lang="ts">
defineProps<{
    hotelName: number
    chainName: string
    address: string
    viewType: string
    roomNumber: number
    capacity: number
    price: number
    category: number
    totalRooms: number
    extendable: boolean
    damaged: boolean
}>()
</script>