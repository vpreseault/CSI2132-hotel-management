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
            <InputText v-model="roomNumber" class="w-full" />
          </div>
  
          <div class="flex flex-col gap-1">
            <label class="text-black font-medium">Stay Dates</label>
            <DatePicker v-model="dateRange" selectionMode="range" :manualInput="false" dateFormat="yy-mm-dd" :minDate="today" class="w-full"/>
            <small v-if="!isDateRangeValid" class="text-red-600">End date must be at least one day after start date.</small>
          </div>
  
          <div class="flex flex-col gap-1">
            <label class="text-black font-medium">Total Price</label>
            <InputText :modelValue="`$${totalPrice}`" disabled class="w-full" />
          </div>
  
          <div class="flex flex-col gap-1">
            <label class="text-black font-medium">Card Type</label>
            <Select v-model="cardType" :options="cardOptions" placeholder="Select Card Type" class="w-full" />
            <small v-if="showCardError && !cardType" class="text-red-600">Card type is required.</small>
          </div>
  
          <div class="flex flex-col gap-1">
            <label class="text-black font-medium">Card Number</label>
            <InputText v-model="cardNumber" placeholder="XXXX-XXXX-XXXX-XXXX" class="w-full" />
            <small v-if="showCardError && !cardNumber" class="text-red-600">Card number is required.</small>
          </div>
        </div>
  
        <div class="flex justify-between mt-6">
          <Button label="Confirm Payment" icon="pi pi-check" class="bg-green-500 border-none hover:bg-green-600 pl-2" :disabled="!isFormValid" @click="emitConfirm" />
          <Button label="Cancel" icon="pi pi-times" class="bg-red-500 border-none hover:bg-red-600 pl-2" @click="emitClose"/>
        </div>
      </div>
    </div>
</template>

<script setup lang="ts">
import { ref, watch, computed } from 'vue';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import DatePicker from 'primevue/datepicker';
import Select from 'primevue/select';
import type { BookingItem } from '../../types';

const props = defineProps<{
booking: BookingItem | null;
}>();

const emit = defineEmits(['confirm', 'close']);

const roomNumber = ref('');
const dateRange = ref<[Date, Date] | null>(null);
const pricePerDay = ref(0);

const cardType = ref('');
const cardNumber = ref('');
const showCardError = ref(false);
const cardOptions = ['Visa', 'MasterCard', 'Amex', 'Discover'];

const today = new Date();
today.setHours(0, 0, 0, 0);

const isDateRangeValid = computed(() => {
if (!dateRange.value) return false;
const [start, end] = dateRange.value;
if (!start || !end) return false;
return end.getTime() > start.getTime() && start.getTime() >= today.getTime();
});

const isCardValid = computed(() => cardType.value && cardNumber.value.trim() !== '');

const isFormValid = computed(() => isDateRangeValid.value && isCardValid.value);

const totalPrice = computed(() => {
if (!dateRange.value) return 0;
const [start, end] = dateRange.value;
if (!start || !end || end <= start) return 0;
const diffDays = Math.ceil((end.getTime() - start.getTime()) / (1000 * 60 * 60 * 24));
return diffDays * pricePerDay.value;
});

watch(() => props.booking, (val) => {
if (val) {
    roomNumber.value = val.room_number.toString();
    const start = new Date(val.start_date);
    const end = new Date(val.end_date);
    if (!isNaN(start.getTime()) && !isNaN(end.getTime())) {
    dateRange.value = [start, end];
    } else {
    dateRange.value = null;
    }
    const stayLength = Math.ceil((end.getTime() - start.getTime()) / (1000 * 60 * 60 * 24));
    pricePerDay.value = stayLength > 0 ? val.total_price / stayLength : 0;
}
}, { immediate: true });

function emitConfirm() {
if (!isFormValid.value) {
    showCardError.value = true;
    return;
}

emit('confirm', {
    ...props.booking,
    room_number: parseInt(roomNumber.value),
    start_date: dateRange.value?.[0].toISOString().split('T')[0],
    end_date: dateRange.value?.[1].toISOString().split('T')[0],
    card_type: cardType.value,
    card_number: cardNumber.value,
    total_price: totalPrice.value
});
}

function emitClose() {
emit('close');
}
</script>
  