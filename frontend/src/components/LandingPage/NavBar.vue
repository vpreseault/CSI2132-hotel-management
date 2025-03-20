<template>
    <div class="bg-green-100 py-3 fixed top-0 left-0 w-full z-10 shadow-md">
        <ul class="flex justify-center space-x-6">
            <li v-if="role === 'customer'" class="text-lg text-black hover:text-green-300">Book</li>
            <li v-if="role === 'customer'" class="text-lg text-black hover:text-green-300">Rentings</li>
            
            <li v-if="role === 'employee' || role === 'manager'" @click="$emit('toggleHotel')" 
                class="text-lg text-black hover:text-green-300 cursor-pointer">Hotel</li>

            <li v-if="role === 'manager'" @click="$emit('toggleCreateEmployee')" 
                class="text-lg text-black hover:text-green-300 cursor-pointer">Create Employee</li>

            <li @click="$emit('toggleProfile')" class="text-lg text-black hover:text-green-300 cursor-pointer">Profile</li>
        </ul>
    </div>
</template>

<script setup lang="ts">
import { defineProps, defineEmits } from 'vue';

defineProps<{ role: string }>();

defineEmits(['toggleProfile', 'toggleHotel', 'toggleCreateEmployee']);
</script>
