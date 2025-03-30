<template>
    <div class="bg-white border rounded-lg p-4 shadow-md cursor-pointer transition hover:shadow-lg mb-3" @click="emit('toggle', section, index)">
      <h3 class="text-lg font-semibold text-gray-800">Archive ID: ARCH-{{ index + 1 }}</h3>
      <p class="text-gray-600 mt-1"> From {{ startDate.toLocaleDateString() }} to {{ endDate.toLocaleDateString() }} </p>
      <p class="text-gray-600 mt-1">Click for more info.</p>
      <div v-if="isExpanded" class="mt-4 space-y-1 text-sm text-gray-700 whitespace-pre-line">
        <p><strong>Start Date:</strong> {{ startDate.toDateString() }}</p>
        <p><strong>End Date:</strong> {{ endDate.toDateString() }}</p>
        <p v-if="price"><strong>Total Price:</strong> ${{ price?.toFixed(2) }}</p>
      </div>
    </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';

const props = defineProps<{
startDate: Date;
endDate: Date;
price?: number;
section: string;
index: number;
expandedCard: { section: string | null; index: number | null };
}>();

const emit = defineEmits<{
toggle: [section: string, index: number];
}>();

const isExpanded = computed(() => {
return props.expandedCard.section === props.section && props.expandedCard.index === props.index;
});
</script>
  