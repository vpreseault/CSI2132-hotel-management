<template>
  <div class="fixed inset-0 flex justify-center items-center z-50">
    <div class="bg-green-200 p-6 rounded-lg shadow-lg w-full max-w-xl">
      <h2 class="text-center text-xl text-black font-semibold mb-4">Confirm Booking Payment</h2>

      <div class="flex flex-col gap-4">
        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Customer Name</label>
          <InputText :modelValue="booking?.customer_name || ''" disabled class="w-full" />
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Hotel</label>
          <InputText :modelValue="booking?.hotel_name || ''" disabled class="w-full" />
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Room Number</label>
          <InputText :modelValue="booking?.room_number || ''" disabled class="w-full" />
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Stay Dates</label>
          <Calendar :model-value="dateRange" disabled selectionMode="range" :manualInput="false" dateFormat="yy-mm-dd" class="w-full" />
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Total Price</label>
          <InputText :modelValue="`$${booking?.total_price}`" disabled class="w-full" />
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Card Type</label>
          <Dropdown v-model="cardType" :options="cardOptions" placeholder="Select Card Type" class="w-full" />
          <small v-if="showCardError && !cardType" class="text-red-600">Card type is required.</small>
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Card Number</label>
          <InputText v-model="cardNumber" placeholder="XXXX-XXXX-XXXX-XXXX" class="w-full" />
          <small v-if="showCardError && !cardNumber" class="text-red-600">Card number is required.</small>
        </div>
      </div>

      <div class="flex justify-between mt-6">
        <Button label="Confirm Payment" icon="pi pi-check" class="bg-green-500 border-none hover:bg-green-600 pl-2"
          :disabled="!isCardValid" @click="emitConfirm" />
        <Button label="Cancel" icon="pi pi-times" class="bg-red-500 border-none hover:bg-red-600 pl-2"
          @click="emitClose" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, computed } from 'vue';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import type { BookingItem } from '../../types';

const props = defineProps<{
  booking: BookingItem | null;
}>();

const emit = defineEmits(['confirm', 'close']);


const dateRange = computed(() => {
  if (props.booking?.start_date && props.booking.end_date) {
    return [new Date(props.booking.start_date), new Date(props.booking.end_date)]
  }
})
const cardType = ref('');
const cardNumber = ref('');
const showCardError = ref(false);
const cardOptions = ['Visa', 'MasterCard', 'Amex', 'Discover'];

const isCardValid = computed(() => cardType.value && cardNumber.value.trim() !== '');


function emitConfirm() {
  if (!isCardValid.value) {
    showCardError.value = true;
    return;
  }

  emit('confirm', props.booking);
}

function emitClose() {
  emit('close');
}
</script>