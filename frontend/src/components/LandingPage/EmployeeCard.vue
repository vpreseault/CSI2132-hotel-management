<template>
    <div class="bg-white p-4 rounded-lg shadow-md relative">
      <h3 class="text-lg font-semibold text-gray-800 mb-1">{{ name }}</h3>
      <p class="text-gray-600">Role: {{ role }}</p>
      <p class="text-gray-600">Address: {{ address }}</p>
      <button class="bg-red-500 text-white text-sm px-4 py-1.5 rounded hover:bg-red-600 transition mt-4" @click="confirmDelete"> Delete Account </button>
    </div>
</template>

<script setup lang="ts">
const props = defineProps<{
    name: string;
    role: string;
    address: string;
}>();

const emit = defineEmits(['delete']);

function confirmDelete() {
    const confirmed = confirm(`Are you sure you want to delete ${props.name}'s account?`);
    if (confirmed) {
        emit('delete');
    }
}
</script>
  