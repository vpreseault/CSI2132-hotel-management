<template>
  <div class="fixed inset-0 flex justify-center items-center z-50">
    <div class="bg-green-200 p-6 rounded-lg shadow-lg w-full max-w-xl">
      <h2 class="text-center text-xl text-black font-semibold mb-4">Create Hotel</h2>
      <div class="flex flex-col gap-4">
        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Chain Name</label>
          <Select v-model="form.chain_ID" :options="chains" optionLabel="chain_name" optionValue="chain_ID" class="w-full"/>
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Hotel Name</label>
          <InputText v-model="form.hotel_name" class="w-full" />
          <p v-if="validationErrors.hotel_name" class="text-sm text-red-600">{{ validationErrors.hotel_name }}</p>
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Address</label>
          <InputText v-model="form.address" class="w-full" />
          <p v-if="validationErrors.address" class="text-sm text-red-600">{{ validationErrors.address }}</p>
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Phone Number</label>
          <InputText v-model="form.phone_number" class="w-full" />
          <p v-if="validationErrors.phone_number" class="text-sm text-red-600">{{ validationErrors.phone_number }}</p>
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Email</label>
          <InputText v-model="form.email" class="w-full" />
          <p v-if="validationErrors.email" class="text-sm text-red-600">{{ validationErrors.email }}</p>
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Category: {{ form.category }}</label>
          <Slider v-model="form.category" :min="1" :max="5" class="w-full" />
        </div>
      </div>

      <div class="flex justify-between mt-6">
        <Button label="Create" icon="pi pi-check" class="bg-green-500 border-none hover:bg-green-600" @click="save"/>
        <Button label="Cancel" icon="pi pi-times" class="bg-red-500 border-none hover:bg-red-600" @click="close"/>
      </div>
    </div>
  </div>
</template>


<script setup lang="ts">
import { ref, reactive, } from 'vue'
import InputText from 'primevue/inputtext'
import Slider from 'primevue/slider'
import Select from 'primevue/select'
import Button from 'primevue/button'
import type { HotelPayload, Chain } from '../../types'
import type { ToastMessageOptions } from 'primevue'

defineProps<{
  chains: Chain[]
}>()

const emit = defineEmits<{
  close: []
  create: [severity: ToastMessageOptions["severity"]]
}>()

const form = ref<HotelPayload>({
  chain_ID: 1,
  hotel_name: '',
  manager_ID: 1,
  address: '',
  phone_number: '',
  email: '',
  category: 1
})

const managers = [
  { id: 1, name: 'Lebron James' },
  { id: 2, name: 'Gerard Butler' },
  { id: 3, name: 'Luffy Monkey' },
  { id: 4, name: 'Ariana Grande' },
  { id: 5, name: 'Troye Sivan' },
]

const validationErrors = reactive<Record<keyof HotelPayload, string>>({
  chain_ID: '',
  hotel_name: '',
  manager_ID: '',
  address: '',
  phone_number: '',
  email: '',
  category: ''
})

function validateField(field: keyof HotelPayload, value: any) {
  switch (field) {
    case 'chain_ID':
      validationErrors.chain_ID = value < 1 || value > 5 ? 'Chain ID must be between 1 and 5.' : ''
      break
    case 'hotel_name':
      validationErrors.hotel_name = value.trim().length < 3 ? 'Hotel name must be at least 3 characters.' : ''
      break
    case 'address':
      validationErrors.address = !/^\d+\s[\w\s]+\s*,\s*[\w\s]+\s*,\s*[A-Z]{2}$/.test(value.trim())
        ? 'Address must be in format: 123 Main St, Ottawa, ON'
        : ''
      break
    case 'phone_number':
      validationErrors.phone_number = !/^[0-9\-+\s()]{7,15}$/.test(value.trim())
        ? 'Enter a valid phone number (e.g., 123 456 7890)'
        : ''
      break
    case 'email':
      validationErrors.email = !/^[^\s@]+@[^\s@]+\.[^\s@]+$/.test(value.trim())
        ? 'Enter a valid email address (e.g., admin@email.com)'
        : ''
      break
  }
}

function validateForm(): boolean {
  for (const field in form.value) {
    validateField(field as keyof HotelPayload, form.value[field as keyof HotelPayload])
  }

  return Object.values(validationErrors).every(error => error === '')
}

async function save() {
  if (!validateForm()) return

  try {
    const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/hotels`,
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(form.value)
        }
    )
    
    if (res.ok) {      
        emit('create', 'success')
        return
    }
    emit('create', 'error')
  } catch (error) {
      console.error('Error calling API:', error);
      emit('create', 'error')
  }
}

function close() {
  emit('close')
}
</script>

