<template>
  <Card class="w-full max-w-2xl mx-auto shadow-none rounded-none border-x-1 border-[#b4b4b4]">
    <template #title>Book Room</template>
    <template #content>
      <div class="mb-6 space-y-1 text-gray-700">
        <p><strong>Hotel:</strong> {{ room.hotel_name }}</p>
        <p><strong>Address:</strong> {{ room.address }}</p>
        <p><strong>Room Number:</strong> {{ room.room_number }}</p>
        <p><strong>Capacity:</strong> {{ room.capacity }}</p>
        <p><strong>Booking Dates:</strong> from {{ props.start_date.toDateString() }} to {{ props.end_date.toDateString() }}</p>
        <p><strong>Room Price:</strong> ${{ room.price }}/night</p>
        <p><strong>Total Price:</strong> ${{ totalPrice }}</p>
      </div>

      <Form v-slot="$form" :initialValues :resolver="resolver" @submit="onFormSubmit" class="flex flex-col gap-4">
        <div class="flex flex-col gap-1">
          <label class="font-medium">Customer Name</label>
          <InputText name="customer_name" placeholder="Enter customer name" class="w-full" />
          <Message v-if="$form.customer_name?.invalid" severity="error">{{ $form.customer_name.error?.message }}</Message>
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
import { defineProps, defineEmits, ref, computed, reactive } from 'vue';
import { Form } from '@primevue/forms';
import InputText from 'primevue/inputtext';
import Message from 'primevue/message';
import Button from 'primevue/button';
import Card from 'primevue/card';
import type { FormResolverOptions, FormSubmitEvent } from '@primevue/forms';
import type { BookingPayload, SearchResult } from '../../types';
import { getUserName } from '../../utils/auth';
import type { ToastMessageOptions } from 'primevue';

type FormErrors = {
  customer_name: FormError[];
}
type FormError = { type: string; message: string }

const emit = defineEmits<{
  close: []
  bookingSubmitted: [ToastMessageOptions["severity"]]
}>()

const props = defineProps<{
  room: SearchResult
  start_date: Date,
  end_date: Date,
}>();

const customerName = ref(getUserName())
const initialValues = reactive({
  customer_name: customerName
})

const totalNights = computed(() => {
  const diff = (props.end_date.getTime() - props.start_date.getTime()) / (1000 * 60 * 60 * 24);
  return Math.max(0, Math.floor(diff));
});

const totalPrice = computed(() => {
  return totalNights.value * props.room.price;
});

const resolver = ({ values }: FormResolverOptions) => {
  const errors: FormErrors = {
    customer_name: [],
  };

  if (!values.customer_name) {
    errors.customer_name.push({ type: 'required', message: 'Customer name is required.' });
  }

  return {
    values,
    errors,
  };
};

async function onFormSubmit(e: FormSubmitEvent) {
  if (e.valid) {
    try {
      const newBooking: BookingPayload = {
        "customer_name": e.values.customer_name,
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
        emit('bookingSubmitted', 'success')
        return 
      }
      
      const errorMsg = await res.text()
      console.log(res.status, errorMsg, e.errors)
      if (res.status === 404 && errorMsg.trim() === "Customer not found") {
        const error = { type: 'required', message: 'Customer name not found.' } as FormError
        ((e.errors as unknown) as FormErrors).customer_name.push(error)
        e.states.customer_name.invalid = true
        e.states.customer_name.valid = false
        e.states.customer_name.error = error
        return
      }

      emit('bookingSubmitted', 'error')
    } catch (error) {
      console.error('Error calling API:', error);
    }
  }
}
</script>
