<template>
    <Card class="p-4">
      <template #title>Book Room</template>
      <template #content>
        <p>This is where booking details will go for:</p>
        <ul class="mt-2 text-sm text-gray-700">
          <li><strong>Hotel:</strong> {{ room.hotel_name }}</li>
          <li><strong>Address:</strong> {{ room.address }}</li>
          <li><strong>Capacity:</strong> {{ room.capacity }}</li>
          <li><strong>Price:</strong> ${{ room.price }}</li>
        </ul>
  
        <Button label="Back to Results" icon="pi pi-arrow-left" class="mt-4 pl-2" @click="emit('close')" />
      </template>
    </Card>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from 'vue';
import Card from 'primevue/card';
import Button from 'primevue/button';

defineProps<{
room: {
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
};
}>();

defineProps

const emit = defineEmits(['close']);
</script>
  