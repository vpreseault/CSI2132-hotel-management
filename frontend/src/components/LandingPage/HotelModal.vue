<template>
  <div class="fixed inset-0 flex justify-center items-center">
    <div class="bg-green-200 p-6 rounded-lg shadow-lg w-1/3">
      <h2 class="text-center text-xl text-black font-semibold">Hotel Details</h2>
      <div class="mt-4 space-y-4">
        <p><strong>Name:</strong> Grand Palace Hotel</p>
        <p><strong>Location:</strong> New York City, NY</p>
        <p><strong>Rating:</strong> 4.5</p>
        <p><strong>Available Rooms:</strong> 20</p>
      </div>
      <button class="mt-4 w-full px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition" @click="$emit('close')"> Close </button>
    </div>
  </div>
</template>
  
<script setup lang="ts">
  defineEmits(['close']);
</script>
  