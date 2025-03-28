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
        <div class="flex flex-col gap-1">
          <label class="font-medium">Customer Name</label>
          <InputText name="customer_name" placeholder="Enter customer name" class="w-full" />
          <Message v-if="$form.customer_name?.invalid" severity="error">{{ $form.customer_name.error?.message }}</Message>
        </div>

        <div class="flex flex-col gap-1">
          <label class="font-medium">Room Number</label>
          <InputText :modelValue="String(room.room_number)" readonly class="w-full bg-gray-100" />
        </div>

        <div class="flex flex-col gap-1">
          <label class="font-medium">Booking Dates</label>
          <Calendar name="dates" selectionMode="range" v-model="dates" showIcon :manualInput="false" class="w-full" />
          <Message v-if="$form.dates?.invalid" severity="error">{{ $form.dates.error?.message }}</Message>
        </div>

        <div class="flex flex-col gap-1">
          <label class="font-medium">Total Price</label>
          <InputText :modelValue="`$${totalPrice}`" readonly class="w-full bg-gray-100" />
        </div>

        <div class="flex flex-col gap-1">
          <label class="font-medium">Card Type</label>
          <Dropdown name="card_type" :options="cardTypes" optionLabel="label" placeholder="Select card type" class="w-full" />
          <Message v-if="$form.card_type?.invalid" severity="error">{{ $form.card_type.error?.message }}</Message>
        </div>

        <div class="flex flex-col gap-1">
          <label class="font-medium">Card Number</label>
          <InputText name="card_number" placeholder="1234 5678 9012 3456" class="w-full" />
          <Message v-if="$form.card_number?.invalid" severity="error">{{ $form.card_number.error?.message }}</Message>
        </div>

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

const dates = ref<[Date, Date] | null>(null);

const totalNights = computed(() => {
  if (!dates.value || dates.value.length !== 2) return 0;
  const [start, end] = dates.value;
  const diff = (end.getTime() - start.getTime()) / (1000 * 60 * 60 * 24);
  return Math.max(0, Math.floor(diff));
});

const totalPrice = computed(() => {
  return totalNights.value * props.room.price;
});

const resolver = ({ values }: FormResolverOptions) => {
  const errors: Record<string, { type: string; message: string }[]> = {
    customer_name: [],
    dates: [],
    card_type: [],
    card_number: [],
  };

  if (!values.customer_name) {
    errors.customer_name.push({ type: 'required', message: 'Customer name is required.' });
  }

  if (!dates.value || dates.value.length !== 2) {
    errors.dates.push({ type: 'required', message: 'Select a valid start and end date.' });
  } else {
    const [start, end] = dates.value;
    if (start >= end) {
      errors.dates.push({ type: 'invalid', message: 'End date must be after start date.' });
    }
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
      start_date: dates.value?.[0],
      end_date: dates.value?.[1],
      total_price: totalPrice.value,
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
