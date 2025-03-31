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

          <div class="flex flex-col gap-1">
            <label class="text-black font-medium">Phone Number</label>
            <InputText v-model="hotelData.phone_number" :disabled="!isEditing" placeholder="Hotel Phone" class="w-full"/>
          </div>
          <div class="flex flex-col gap-1">
            <label class="text-black font-medium">Email</label>
            <InputText v-model="hotelData.email" :disabled="!isEditing" placeholder="Hotel Email" class="w-full"/>
          </div>
        </div>
  
        <div class="flex justify-between mt-6">
          <Button :label="isEditing ? 'Save' : 'Edit'" :icon="isEditing ? 'pi pi-check' : 'pi pi-pencil'" severity="success" class="pl-2" @click="toggleEdit" />
          <Button label="Close" icon="pi pi-times" severity="info" class="pl-2" @click="closeButtonClicked" />
        </div>
        <Message v-if="message.text" :severity="message.severity" variant="simple" class="flex justify-center mt-8">{{ message.text }}</Message>
    </div>
  </div>
</template>

<script setup lang="ts">
import { onMounted, reactive, ref } from 'vue';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import Slider from 'primevue/slider';
import type { Hotel } from '../../types';
import { getUserID } from '../../utils/auth';
import type { ToastMessageOptions } from 'primevue';

const emit = defineEmits(['close']);

const isEditing = ref(false)
const editsMade = ref(false)

const hotelData = ref<Hotel>({
  hotel_ID: 0,
  hotel_name: '',
  address: '',
  phone_number: '',
  email: '',
  category: 0,
});

onMounted(async () => {
  const managerID = getUserID()
  
  try {
      const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/hotels?manager_ID=${managerID}`)
      if (res.ok) {
        hotelData.value = await res.json()
      }
  } catch (error) {
      console.error('Error calling API:', error);
  }
})

const message = reactive<{
  severity?: ToastMessageOptions["severity"],
  text?: string,
}>({})

function isValidAddress(address: string): boolean {
  const addressRegex = /^\d+\s+[A-Za-z\s]+,\s*[A-Za-z]+,\s*[A-Z]{2}$/;
  return addressRegex.test(address.trim());
}

function isValidPhoneNumber(phone: string): boolean {
  const phoneRegex = /^\d{3}-\d{3}-\d{4}$/;
  return phoneRegex.test(phone.trim());
}

function isValidEmail(email: string): boolean {
  const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
  return emailRegex.test(email.trim());
}

async function toggleEdit() {
  if (isEditing.value) {
    const success = await updateHotel();
    if (!success) return;  
  }
  isEditing.value = !isEditing.value;
}

async function updateHotel(): Promise<boolean> {
  hotelData.value.address = hotelData.value.address.trim();
  hotelData.value.phone_number = hotelData.value.phone_number.trim();
  hotelData.value.email = hotelData.value.email.trim();

  if (!hotelData.value.address || !hotelData.value.phone_number || !hotelData.value.email) {
    message.severity = 'warn';
    message.text = 'All fields must be filled out.';
    return false;
  }

  if (!isValidAddress(hotelData.value.address)) {
    message.severity = 'warn';
    message.text = 'Address must follow the format: "123 Main Street, City, State".';
    return false;
  }

  if (!isValidPhoneNumber(hotelData.value.phone_number)) {
    message.severity = 'warn';
    message.text = 'Phone number must be 10 digits with optional country code. (E.x. 123-456-7890)';
    return false;
  }

  if (!isValidEmail(hotelData.value.email)) {
    message.severity = 'warn';
    message.text = 'Please enter a valid email address. (E.x. example@gmail.com)';
    return false;
  }

  try {
    const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/hotels/${hotelData.value.hotel_ID}`, {
      method: 'PATCH',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        hotel_name: hotelData.value.hotel_name,
        address: hotelData.value.address,
        phone_number: hotelData.value.phone_number,
        email: hotelData.value.email,
        category: hotelData.value.category,
      })
    });

    if (res.ok) {
      editsMade.value = true;
      message.severity = 'success';
      message.text = 'Hotel information updated successfully.';
      return true;
    }

    message.severity = 'error';
    message.text = 'Failed to update hotel information.';
    return false;
  } catch (error) {
    console.error('Error calling API:', error);
    message.severity = 'error';
    message.text = 'Failed to update hotel information.';
    return false;
  }
}

function closeButtonClicked() {
  if (editsMade.value) {
    editsMade.value = false
    window.location.reload()
  }
  emit('close')
}
</script>
