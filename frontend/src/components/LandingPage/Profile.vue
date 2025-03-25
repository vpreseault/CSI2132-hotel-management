<template>
    <div class="fixed inset-0 flex justify-center items-center">
      <div class="bg-green-200 p-6 rounded-lg shadow-lg w-1/3">
        <h2 class="text-center text-xl text-black font-semibold">Profile Details</h2>
        <div class="mt-4 space-y-4">
          <p class="text-lg text-black"><strong>Name:</strong> {{ profileData.name }}</p>
          <p class="text-lg text-black"><strong>Email:</strong> {{ profileData.email }}</p>
          <p class="text-lg text-black"><strong>Phone:</strong> {{ profileData.phone }}</p>
          <p class="text-lg text-black"><strong>Address:</strong> {{ profileData.address }}</p>
        </div>
        <button class="mt-4 w-full px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition" @click="toggleProfileModal"> Close </button>
      </div>
    </div>
</template>
  
<script setup lang="ts">
import { ref } from 'vue';

defineProps<{
toggleProfileModal: () => void;
}>();

const profileData = ref({
    name: "John Doe",
    email: "johndoe@example.com",
    phone: "+1 234 567 890",
    address: "123 Main St, New York, NY"
});
</script>
  