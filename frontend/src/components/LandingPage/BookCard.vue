<template>
  <Card class="w-full max-w-2xl mx-auto shadow-none rounded-none border-x-1 border-[#b4b4b4]">
    <template #title>Create Room {{isRental ? 'Rental': 'Booking'}} For Customer</template>
    <template #content>
      <div class="mb-6 space-y-1 text-gray-700">
        <p><strong>Hotel:</strong> {{ room.hotel_name }}</p>
        <p><strong>Address:</strong> {{ room.address }}</p>
        <p><strong>Room Number:</strong> {{ room.room_number }}</p>
        <p><strong>Capacity:</strong> {{ room.capacity }}</p>
        <p><strong>{{isRental ? 'Rental': 'Booking'}} Dates:</strong> from {{ props.start_date.toDateString() }} to {{ props.end_date.toDateString() }}</p>
        <p><strong>Room Price:</strong> ${{ room.price }}/night</p>
        <p><strong>Total Price:</strong> ${{ totalPrice }}</p>
      </div>

      <Form v-slot="$form" :initialValues :resolver="resolver" @submit="onFormSubmit" class="flex flex-col gap-4">
        <div class="flex flex-col gap-1">
          <label class="font-medium">Customer Name</label>
          <InputText name="customer_name" placeholder="Enter customer name" class="w-full" />
          <Message v-if="$form.customer_name?.invalid" severity="error">{{ $form.customer_name.error?.message }}</Message>
        </div>

        <template v-if="employee && isRental">
          <div class="flex flex-col gap-1">
            <label class="font-medium">Card Type</label>
            <Select name="card_type" :options="cardTypes" optionLabel="label" placeholder="Select card type" class="w-full" />
            <Message v-if="$form.card_type?.invalid" severity="error">{{ $form.card_type.error?.message }}</Message>
          </div>
  
          <div class="flex flex-col gap-1">
            <label class="font-medium">Card Number</label>
            <InputText name="card_number" placeholder="1234 5678 9012 3456" class="w-full" />
            <Message v-if="$form.card_number?.invalid" severity="error">{{ $form.card_number.error?.message }}</Message>
          </div>
        </template>

        <div class="flex justify-between mt-6">
          <Button type="button" label="Back to Results" icon="pi pi-arrow-left" class="p-button-secondary" @click="emit('close')" />
          <Button type="submit" :label="`Confirm ${isRental ? 'Rental': 'Booking'}`" icon="pi pi-check" />
        </div>
      </Form>
    </template>
  </Card>
</template>

<script setup lang="ts">
import { defineProps, defineEmits, computed, reactive } from 'vue';
import { Form } from '@primevue/forms';
import InputText from 'primevue/inputtext';
import Message from 'primevue/message';
import Button from 'primevue/button';
import Card from 'primevue/card';
import type { FormResolverOptions, FormSubmitEvent } from '@primevue/forms';
import type { BookingPayload, RentalPayload, SearchResult } from '../../types';
import { getUserID, getUserName } from '../../utils/auth';
import type { ToastMessageOptions } from 'primevue';

type FormErrors = {
  customer_name: FormError[];
  card_type: FormError[];
  card_number: FormError[];
}
type FormError = { type: string; message: string }

const emit = defineEmits<{
  close: []
  bookingSubmitted: [severity: ToastMessageOptions["severity"], isRental: boolean]
}>()

const props = defineProps<{
  room: SearchResult
  start_date: Date,
  end_date: Date,
  employee: boolean
}>();

const cardTypes = [
  { label: 'Visa', value: 'Visa' },
  { label: 'Mastercard', value: 'Mastercard' },
  { label: 'Amex', value: 'Amex' },
  { label: 'Discover', value: 'Discover' }
];

const customerName = computed(() => props.employee ? '': getUserName())
const initialValues = reactive({
  customer_name: customerName.value
})

const totalNights = computed(() => {
  const diff = (props.end_date.getTime() - props.start_date.getTime()) / (1000 * 60 * 60 * 24);
  return Math.max(0, Math.floor(diff));
});

const isRental = computed(() => {
  const today = new Date();
  return props.start_date.toDateString() === today.toDateString();
});

const totalPrice = computed(() => {
  return totalNights.value * props.room.price;
});

const resolver = ({ values }: FormResolverOptions) => {
  const errors: FormErrors = {
    customer_name: [],
    card_type: [],
    card_number: [],
  };

  if (!values.customer_name) {
    errors.customer_name.push({ type: 'required', message: 'Customer name is required.' });
  }

  if (props.employee && !values.card_type) {
    errors.card_type.push({ type: 'required', message: 'Card type is required.' });
  }

  // || !/^\d{4} \d{4} \d{4} \d{4}$/.test(values.card_number)
  if (props.employee && !values.card_number) {
    errors.card_number.push({ type: 'invalid', message: 'Card number is required.' });
    // errors.card_number.push({ type: 'invalid', message: 'Enter a valid card number (e.g., 1234 5678 9012 3456).' });
  }

  return {
    values,
    errors,
  };
};

async function onFormSubmit(e: FormSubmitEvent) {
  if (e.valid) {
    try {
      const customerNameResponse = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/customers/${e.values.customer_name}`)

      if (customerNameResponse.status === 404) {
        const error = { type: 'required', message: 'Customer name not found.' } as FormError
        ((e.errors as unknown) as FormErrors).customer_name.push(error)
        e.states.customer_name.invalid = true
        e.states.customer_name.valid = false
        e.states.customer_name.error = error
        return
      }

      if (!customerNameResponse.ok) {
        emit('bookingSubmitted', 'error', isRental.value)
        return
      }

      const { customer_ID } = await customerNameResponse.json()

      if (isRental.value) {
        await createRental(customer_ID)
      } else {
        await createBooking(customer_ID)
      }
    } catch (error) {
      console.error('Error calling API:', error);
    }
  }
}

async function createBooking(customerID: number) {
  try {
    const newBooking: BookingPayload = {
      "customer_ID": customerID,
      "room_ID": props.room.room_ID,
      "start_date": props.start_date.toISOString().split('T')[0],
      "end_date": props.end_date.toISOString().split('T')[0],
      "total_price": totalPrice.value,
    }
  
    const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/bookings`,
      {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(newBooking)
      }
    )
    if (res.ok) {
      emit('bookingSubmitted', 'success', isRental.value)
      return 
    }
    emit('bookingSubmitted', 'error', isRental.value)
  } catch (error) {
    console.error('Error calling API:', error);
  }
}

const employeeID = getUserID()
async function createRental(customerID: number) {
  try {
    const newRental: RentalPayload = {
      "customer_ID": customerID,
      "employee_ID": employeeID,
      "room_ID": props.room.room_ID,
      "check_in_date": props.start_date.toISOString().split('T')[0],
      "check_out_date": props.end_date.toISOString().split('T')[0],
      "total_price": totalPrice.value,
    }
  
    const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/rentings`,
      {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(newRental)
      }
    )
    if (res.ok) {
      emit('bookingSubmitted', 'success', isRental.value)
      return 
    }
    emit('bookingSubmitted', 'error', isRental.value)
  } catch (error) {
    console.error('Error calling API:', error);
  }
}
</script>
