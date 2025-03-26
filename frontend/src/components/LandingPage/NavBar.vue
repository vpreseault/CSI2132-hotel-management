<template>
  <div class="bg-green-100 py-3 fixed top-0 left-0 w-full z-10 shadow-md">
    <ul class="flex justify-center space-x-6">
      <li v-if="role === 'customer'" :class="navItemClass"> Book </li>
      <li v-if="role === 'customer'" :class="navItemClass"> Rentings </li>
  
      <li v-if="role === 'employee'" @click="emit('toggleHotel')" :class="navItemClass"> Hotel </li>
      <li v-if="role === 'manager'" @click="emit('toggleCreateEmployee')" :class="navItemClass" > Create Employee </li>
      <li v-if="role === 'manager'" @click="emit('toggleHotel')" :class="navItemClass"> Hotel </li>
  
      <li @click="emit('toggleProfile')" :class="navItemClass"> Profile </li>
    </ul>
  </div>
</template>
  
<script setup lang="ts">
import { computed } from 'vue'
import { defineProps, defineEmits } from 'vue'

const emit = defineEmits<{
  toggleProfile: [],
  toggleHotel: [],
  toggleCreateEmployee: []
}>()

defineProps<{ role: string }>()

const navItemClass = computed(() =>
  'text-lg text-black hover:text-green-300 cursor-pointer'
)
</script>
  