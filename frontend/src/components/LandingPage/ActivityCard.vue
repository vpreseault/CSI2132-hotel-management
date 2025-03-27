<template>
    <div class="bg-white border rounded-lg p-4 shadow-md cursor-pointer transition hover:shadow-lg" @click="$emit('toggle', section, index)">
      <h3 class="text-lg font-semibold text-gray-800">{{ title }}</h3>
      <p class="text-gray-600 mt-1">{{ description }}</p>
  
      <div v-if="isExpanded" class="mt-4 space-y-1 text-sm text-gray-700 whitespace-pre-line">
        <template v-if="cardType === 'booking'">
          <p><strong>Customer Name:</strong> {{ data.customer_name }}</p>
          <p><strong>Room Number:</strong> {{ data.room_number ?? 'N/A' }}</p>
          <p><strong>Start Date:</strong> {{ data.start_date }}</p>
          <p><strong>End Date:</strong> {{ data.end_date }}</p>
        </template>
  
        <template v-else-if="cardType === 'renting'">
          <p><strong>Customer Name:</strong> {{ data.customer_name }}</p>
          <p><strong>Employee Name:</strong> {{ data.employee_name ?? 'N/A' }}</p>
          <p><strong>Room Number:</strong> {{ data.room_number ?? 'N/A' }}</p>
          <p><strong>Start Date:</strong> {{ data.start_date }}</p>
          <p><strong>End Date:</strong> {{ data.end_date }}</p>
          <p><strong>Total Price:</strong> ${{ data.total_price?.toFixed(2) ?? 'N/A' }}</p>
          <p><strong>Payment:</strong> {{ data.payment ? 'Paid' : 'Pending' }}</p>
        </template>
  
        <template v-else-if="cardType === 'archive'">
          <p><strong>Customer Name:</strong> {{ data.customer_name }}</p>
          <p><strong>Start Date:</strong> {{ data.start_date }}</p>
          <p><strong>End Date:</strong> {{ data.end_date }}</p>
        </template>
      </div>
    </div>
</template>
  
<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{
data: any;
title: string;
description: string;
cardType: 'booking' | 'renting' | 'archive';
section: string;
index: number;
expandedCard: { section: string | null; index: number | null };
}>();

const isExpanded = computed(() => {
return props.expandedCard.section === props.section && props.expandedCard.index === props.index;
});
</script>
  