<template>
    <div class="p-4 rounded-lg shadow-md cursor-pointer transition grid place-items-center" :class="cardStyles" @click="toggleCard" >
        <h3 class="text-lg text-gray-600 font-semibold">{{ title }}</h3>
        <p class="text-sm text-gray-600">{{ description }}</p>
        <div v-if="isExpanded" class="mt-4 p-3 rounded-md" :class="expandedStyles">
            <p class="text-sm text-gray-600">{{ details }}</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import {defineProps, defineEmits, computed} from 'vue';

const props = defineProps<{
    title: string;
    description: string;
    details?: string;
    section: string;
    index: number;
    expandedCard: { section: string | null; index: number | null };
}>();

const emit = defineEmits(['toggle']);

const isExpanded = computed(() => 
    props.expandedCard.section === props.section && props.expandedCard.index === props.index
);

const cardStyles = computed(() => isExpanded.value ? 'bg-gray-200' : 'bg-green-100');

function toggleCard() {
    emit('toggle', props.section, props.index);
}
</script>
