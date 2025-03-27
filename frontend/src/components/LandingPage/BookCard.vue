<template>
    <Card class="w-full max-w-2xl mx-auto">
      <template #title>Book Room</template>
      <template #content>
        <div class="mb-6 space-y-1 text-gray-700 text-sm">
          <p><strong>Hotel:</strong> {{ room.hotel_name }}</p>
          <p><strong>Address:</strong> {{ room.address }}</p>
          <p><strong>Capacity:</strong> {{ room.capacity }}</p>
          <p><strong>Price per Night:</strong> ${{ room.price }}</p>
        </div>
  
        <Form v-slot="$form" :resolver="resolver" @submit="onFormSubmit" class="flex flex-col gap-4">
          <!-- Customer Name -->
          <div class="flex flex-col gap-1">
            <label class="font-medium">Customer Name</label>
            <InputText name="customer_name" placeholder="Enter customer name" class="w-full" />
            <Message v-if="$form.customer_name?.invalid" severity="error">{{ $form.customer_name.error?.message }}</Message>
          </div>
  
          <!-- Room Number -->
          <div class="flex flex-col gap-1">
            <label class="font-medium">Room Number</label>
            <InputText :modelValue="String(room.room_number)" readonly class="w-full bg-gray-100" />
          </div>
  
          <!-- Dates -->
          <div class="grid grid-cols-1 md:grid-cols-2 gap-4">
            <div class="flex flex-col gap-1">
              <label class="font-medium">Start Date</label>
              <Calendar name="start_date" showIcon placeholder="Select start date" class="w-full" />
              <Message v-if="$form.start_date?.invalid" severity="error">{{ $form.start_date.error?.message }}</Message>
            </div>
  
            <div class="flex flex-col gap-1">
              <label class="font-medium">End Date</label>
              <Calendar name="end_date" showIcon placeholder="Select end date" class="w-full" :minDate="minEndDate"/>
              <Message v-if="$form.end_date?.invalid" severity="error">{{ $form.end_date.error?.message }}</Message>
            </div>
          </div>
  
          <!-- Total Price (readonly) -->
          <div class="flex flex-col gap-1">
            <label class="font-medium">Total Price</label>
            <InputText :modelValue="`$${room.price}`" readonly class="w-full bg-gray-100" />
          </div>
  
          <!-- Card Type -->
          <div class="flex flex-col gap-1">
            <label class="font-medium">Card Type</label>
            <Dropdown name="card_type" :options="cardTypes" optionLabel="label" placeholder="Select card type" class="w-full"/>
            <Message v-if="$form.card_type?.invalid" severity="error">{{ $form.card_type.error?.message }}</Message>
          </div>
  
          <!-- Card Number -->
          <div class="flex flex-col gap-1">
            <label class="font-medium">Card Number</label>
            <InputText name="card_number" placeholder="1234 5678 9012 3456" class="w-full" />
            <Message v-if="$form.card_number?.invalid" severity="error">{{ $form.card_number.error?.message }}</Message>
          </div>
  
          <!-- Action Buttons -->
          <div class="flex justify-between mt-6">
            <Button type="button" label="Back to Results" icon="pi pi-arrow-left" class="p-button-secondary" @click="emit('close')" />
            <Button type="submit" label="Confirm Booking" icon="pi pi-check" />
          </div>
        </Form>
      </template>
    </Card>
  </template>
  
<script setup lang="ts">
import { defineProps, defineEmits, ref, computed } from 'vue';
import { Form } from '@primevue/forms';
import InputText from 'primevue/inputtext';
import Dropdown from 'primevue/dropdown';
import Calendar from 'primevue/calendar';
import Message from 'primevue/message';
import Button from 'primevue/button';
import Card from 'primevue/card';
import type { FormResolverOptions, FormSubmitEvent } from '@primevue/forms';

const emit = defineEmits(['close', 'createBooking']);

const props = defineProps<{
  room: {
    hotel_name: string;
    room_number: number;
    capacity: number;
    price: number;
    view_type: string;
    extendable: boolean;
    damaged: boolean;
    chain_name: string;
    category: number;
    address: string;
    total_rooms: number;
  };
}>();

const cardTypes = [
  { label: 'Visa', value: 'Visa' },
  { label: 'Mastercard', value: 'Mastercard' },
  { label: 'Amex', value: 'Amex' },
  { label: 'Discover', value: 'Discover' }
];

const startDate = ref<Date | null>(null);

const minEndDate = computed(() => {
  if (!startDate.value) return undefined;
  const min = new Date(startDate.value);
  min.setDate(min.getDate() + 1);
  return min;
});

const resolver = ({ values }: FormResolverOptions) => {
  const errors: Record<string, { type: string; message: string }[]> = {
    customer_name: [],
    start_date: [],
    end_date: [],
    total_price: [],
    payment: [],
    card_type: [],
    card_number: [],
  };

  if (!values.customer_name) {
    errors.customer_name.push({ type: 'required', message: 'Customer name is required.' });
  }

  if (!values.start_date) {
    errors.start_date.push({ type: 'required', message: 'Start date is required.' });
  } else {
    startDate.value = new Date(values.start_date);
  }

  if (!values.end_date) {
    errors.end_date.push({ type: 'required', message: 'End date is required.' });
  }

  if (values.start_date && values.end_date) {
    const start = new Date(values.start_date);
    const end = new Date(values.end_date);
    const diffDays = (end.getTime() - start.getTime()) / (1000 * 60 * 60 * 24);
    if (diffDays < 1) {
      errors.start_date.push({ type: 'invalid', message: 'End date must be at least 1 day after start date.' });
      errors.end_date.push({ type: 'invalid', message: 'End date must be at least 1 day after start date.' });
    }
  }

  if (!values.payment) {
    errors.payment.push({ type: 'required', message: 'Payment status is required.' });
  }

  if (!values.card_type) {
    errors.card_type.push({ type: 'required', message: 'Card type is required.' });
  }

  if (!values.card_number || !/^\d{4} \d{4} \d{4} \d{4}$/.test(values.card_number)) {
    errors.card_number.push({ type: 'invalid', message: 'Enter a valid card number (e.g., 1234 5678 9012 3456).' });
  }

  return {
    values: {
      ...values,
      room_number: props.room.room_number,
      total_price: props.room.price,
      cardType: 'booking',
    },
    errors,
  };
};

function onFormSubmit(e: FormSubmitEvent) {
  if (e.valid) {
    emit('createBooking', e.values);
    emit('close');
  }
}
</script>
