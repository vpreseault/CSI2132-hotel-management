<template>
  <div id="book" class="mt-16 p-8 bg-green-100 rounded-lg shadow-md">
    <h2 class="text-center text-xl text-gray-600 font-semibold">
      {{ isEmployee ? 'Make a Rental' : 'Book a Hotel' }}
    </h2>

    <template v-if="!isEmployee">
      <div class="text-center my-4">
        <button @click="isSearchCardVisible = !isSearchCardVisible" class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-indigo-600 transition">
          {{ isSearchCardVisible ? 'Close' : 'Search' }}
        </button>
      </div>

      <div v-if="isSearchCardVisible" class="mb-4">
        <SearchSection />
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-4 justify-center">
        <GenericCard 
          v-for="(hotel, index) in hotels" 
          :key="index"
          :title="hotel.title" 
          :description="hotel.description" 
          :details="formatDetails(hotel.details)"
          :section="'book'"
          :index="index"
          :expandedCard="expandedCard"
          @toggle="toggleCard"
        />
      </div>
    </template>

    <template v-else>
      <Form v-slot="$form" :resolver="resolver" @submit="onFormSubmit" class="flex flex-col gap-4 w-full lg:w-1/2 mx-auto mt-4">
        <div class="flex flex-col gap-1">
          <InputText name="customer_name" type="text" placeholder="Customer Name" fluid />
          <Message v-if="$form.customer_name?.invalid" severity="error" size="small" variant="simple">{{ $form.customer_name.error?.message }}</Message>
        </div>

        <div class="flex flex-col gap-1">
          <InputText name="employee_name" type="text" placeholder="Employee Name" fluid />
          <Message v-if="$form.employee_name?.invalid" severity="error" size="small" variant="simple">{{ $form.employee_name.error?.message }}</Message>
        </div>

        <div class="flex flex-col gap-1">
          <InputText name="room_number" type="number" placeholder="Room Number" fluid />
          <Message v-if="$form.room_number?.invalid" severity="error" size="small" variant="simple">{{ $form.room_number.error?.message }}</Message>
        </div>

        <div class="flex flex-col gap-1">
          <InputText name="start_date" type="date" placeholder="Start Date" fluid />
          <Message v-if="$form.start_date?.invalid" severity="error" size="small" variant="simple">{{ $form.start_date.error?.message }}</Message>
        </div>

        <div class="flex flex-col gap-1">
          <InputText name="end_date" type="date" placeholder="End Date" fluid />
          <Message v-if="$form.end_date?.invalid" severity="error" size="small" variant="simple">{{ $form.end_date.error?.message }}</Message>
        </div>

        <div class="flex flex-col gap-1">
          <InputText name="total_price" type="number" placeholder="Total Price" fluid />
          <Message v-if="$form.total_price?.invalid" severity="error" size="small" variant="simple">{{ $form.total_price.error?.message }}</Message>
        </div>

        <div class="flex flex-col gap-1">
          <Select name="payment" :options="paymentOptions" optionLabel="label" placeholder="Payment Status" fluid />
          <Message v-if="$form.payment?.invalid" severity="error" size="small" variant="simple">{{ $form.payment.error?.message }}</Message>
        </div>

        <Button type="submit" label="Book Room" />
      </Form>
    </template>
  </div>
</template>

  
<script setup lang="ts">
import { ref } from 'vue';
import GenericCard from './GenericCard.vue';
import SearchSection from './SearchSection.vue';
import { Form } from '@primevue/forms';
import InputText from 'primevue/inputtext';
import Message from 'primevue/message';
import Select from 'primevue/select';
import Button from 'primevue/button';
import type { FormResolverOptions, FormSubmitEvent } from '@primevue/forms';


defineProps<{
expandedCard: { section: string | null; index: number | null };
toggleCard: (section: string, index: number) => void;
isEmployee?: boolean;
}>();

const emit = defineEmits(['createBooking']);
  
const hotels = [
  {
  title: "Hotel 1",
  description: "Description",
  details: {
  cost: "$250 per night",
  rating: "5 stars",
  contactEmail: "contact@hotel1.com",
  phone: "+1-555-1234"
  }
  },
  {
  title: "Hotel 2",
  description: "Description",
  details: {
  cost: "$120 per night",
  rating: "3 stars",
  contactEmail: "info@hotel2.com",
  phone: "+1-555-5678"
  }
  },
  {
  title: "Hotel 3",
  description: "Description",
  details: {
  cost: "$180 per night",
  rating: "4 stars",
  contactEmail: "reservations@hotel3.com",
  phone: "+1-555-9876"
  }
  },
  {
  title: "Hotel 4",
  description: "Description",
  details: {
  cost: "$220 per night",
  rating: "4.5 stars",
  contactEmail: "bookings@hotel4.com",
  phone: "+1-555-4321"
  }
  },
  {
  title: "Hotel 5",
  description: "Description",
  details: {
  cost: "$300 per night",
  rating: "5 stars",
  contactEmail: "luxury@hotel5.com",
  phone: "+1-555-2468"
  }
  }
];

const isSearchCardVisible = ref(false);

function formatDetails(details: any) {
return `
  Cost: ${details.cost}
  Rating: ${details.rating}
  Email: ${details.contactEmail}
  Phone: ${details.phone}
`;
}

const paymentOptions = ref([
  { label: 'Paid', value: true },
  { label: 'Not Paid', value: false }
]);

const resolver = ({ values }: FormResolverOptions) => {
  const errors: Record<string, { type: string; message: string }[]> = {
    customer_name: [],
    employee_name: [],
    room_number: [],
    start_date: [],
    end_date: [],
    total_price: [],
    payment: [],
  };

  if (!values.customer_name) {
    errors.customer_name.push({ type: 'required', message: 'Customer name is required.' });
  }

  if (!values.employee_name) {
    errors.employee_name.push({ type: 'required', message: 'Employee name is required.' });
  }

  if (!values.room_number || values.room_number >= 0) {
    errors.room_number.push({ type: 'required', message: 'Valid room number is required.' });
  }

  if (!values.start_date) {
    errors.start_date.push({ type: 'required', message: 'Start date is required.' });
  }

  if (!values.end_date) {
    errors.end_date.push({ type: 'required', message: 'End date is required.' });
  }

  if (values.start_date && values.end_date) {
    const start = new Date(values.start_date);
    const end = new Date(values.end_date);
    const oneDay = 24 * 60 * 60 * 1000;
    const diffDays = (end.getTime() - start.getTime()) / oneDay;

    if (diffDays < 1) {
      errors.start_date.push({ type: 'invalid', message: 'Start date must be at least one day before end date.' });
      errors.end_date.push({ type: 'invalid', message: 'End date must be at least one day after start date.' });
    }
  }

  if (!values.total_price || values.total_price >= 0) {
    errors.total_price.push({ type: 'required', message: 'Total price is required.' });
  }

  if (values.payment === undefined || values.payment === null) {
    errors.payment.push({ type: 'required', message: 'Payment status is required.' });
  }

  return {
    values: {
      ...values,
      cardType: 'booking',
    },
    errors,
  };
};


function onFormSubmit(e: FormSubmitEvent) {
  if (e.valid) {
    emit('createBooking', e.values);
  }
}

</script>
  