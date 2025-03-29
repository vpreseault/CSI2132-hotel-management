<template>
  <div class="fixed inset-0 flex justify-center items-center">
    <div class="bg-green-200 p-6 rounded-lg shadow-lg w-full max-w-xl">
      <h2 class="text-center text-xl text-black font-semibold mb-4">Hotel Details</h2>

      <div class="flex flex-col gap-4">
        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Hotel Name</label>
          <InputText v-model="hotelData.hotel_name" :disabled="!isEditing" placeholder="Hotel Name" class="w-full"/>
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Address</label>
          <InputText v-model="hotelData.address" :disabled="!isEditing" placeholder="Hotel Address" class="w-full"/>
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Category (1 to 5)</label>
          <Slider v-model="hotelData.category" :min="1" :max="5" :step="1" :disabled="!isEditing" class="w-full"/>
          <div class="text-sm text-black mt-1 text-center">Selected: {{ hotelData.category }}</div>
        </div>
      </div>

      <div class="flex justify-between mt-6">
        <Button :label="isEditing ? 'Save' : 'Edit'" :icon="isEditing ? 'pi pi-check' : 'pi pi-pencil'" class="bg-green-500 border-none hover:bg-green-600 pl-2" @click="toggleEdit" />
        <Button label="Close" icon="pi pi-times" class="bg-red-500 border-none hover:bg-red-600 pl-2" @click="$emit('close')" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import Slider from 'primevue/slider';

const emit = defineEmits(['close']);

const isEditing = ref(false);

const hotelData = ref({
  hotel_name: 'Grand Palace Hotel',
  address: '123 Main St, New York, NY',
  category: 4
});

function toggleEdit() {
  if (isEditing.value) {
    console.log("Saving hotel data:", hotelData.value);
    alert("Hotel details saved!");
  }
  isEditing.value = !isEditing.value;
}
</script>
