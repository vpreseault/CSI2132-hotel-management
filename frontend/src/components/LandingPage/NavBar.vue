<template>
  <div class="bg-green-100 py-4 fixed top-0 left-0 w-full z-10 shadow-md">
    <div class="flex justify-between items-center mx-8">
      <h1 class="text-2xl font-bold text-black">Ω Hotel Management</h1>
      <ul class="flex items-center space-x-6">
        <li v-if="role === 'customer'" :class="navItemClass"> Book </li>
        <li v-if="role === 'customer'" :class="navItemClass"> Rentings </li>

        <li v-if="role === 'manager'" @click="emit('toggleHotel')" :class="navItemClass">Edit Hotel</li>
        <li v-if="role === 'manager'" @click="emit('toggleCreateEmployee')" :class="navItemClass">Create Employee</li>

        <li @click="emit('toggleProfile')" :class="navItemClass"> Profile </li>
        <button class="ml-8 px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition" @click="handleLogout">
          Logout
        </button>
      </ul>
    </div>
  </div>
</template>

  
<script setup lang="ts">
import { computed } from 'vue'
import { defineProps, defineEmits } from 'vue'
import { removeAuthCookie } from '../../utils/auth';

const emit = defineEmits<{
  toggleProfile: [],
  toggleHotel: [],
  toggleCreateEmployee: [],
  exitAdmin: []
}>()

defineProps<{ role: string }>()

const navItemClass = computed(() =>
  'text-lg text-black hover:text-green-300 cursor-pointer'
)

function handleLogout() {
  removeAuthCookie();
}
</script>
  