<template>
  <div id="rentings" class="mt-16 p-8 bg-green-100 rounded-lg shadow-md">
    <h2 class="text-center text-xl text-gray-600 font-semibold">
      {{ isEmployee ? 'All Hotel Rentings' : 'Your Rentings' }}
    </h2>

    <div v-if="currentRentals.length" class="mt-6" >
      <h3 class="text-lg font-semibold text-gray-700 mb-4 text-center">Current Rentals</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-4 justify-center" >
        <ActivityCard 
          v-for="(rental, index) in currentRentals" 
          :key="`current-${index}`"
          :title="getTitle(rental)"
          :description="getDescription(rental)"
          :data="rental"
          :cardType="rental.cardType"
          :section="'current'"   
          :index="index"
          :expandedCard="expandedCard"
          @toggle="toggleCard"
        />
      </div>
    </div>

    <div v-if="futureRentals.length" class="mt-10">
      <h3 class="text-lg font-semibold text-gray-700 mb-4 text-center">Future Rentals</h3>
      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-4 justify-center">
        <ActivityCard 
          v-for="(rental, index) in futureRentals" 
          :key="`future-${index}`"
          :title="getTitle(rental)"
          :description="getDescription(rental)"
          :data="rental"
          :cardType="rental.cardType"
          :section="'future'"  
          :index="index"
          :expandedCard="expandedCard"
          @toggle="toggleCard"
        />
      </div>
    </div>

    <div v-if="archivedRentals.length" class="mt-10">
      <h3 class="text-lg font-semibold text-gray-700 mb-4 text-center">Past Rentals</h3>
      <div class="flex flex-col gap-4">
        <ActivityCard
          v-for="(rental, index) in archivedRentals"
          :key="`archive-${index}`"
          :title="getTitle(rental)"
          :description="getDescription(rental)"
          :data="rental"
          :cardType="'archive'"
          :section="'archive'"
          :index="index"
          :expandedCard="expandedCard"
          @toggle="toggleCard"
        />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed, onMounted } from 'vue';
import { getUserID } from '../../utils/auth';
import ActivityCard from './ActivityCard.vue';

type RentalItem = {
  cardType: 'booking' | 'renting' | 'archive';
  customer_name: string;
  employee_name?: string;
  room_number?: number;
  start_date: string;
  end_date: string;
  total_price?: number;
  payment?: boolean;
};

const props = defineProps<{
  expandedCard: { section: string | null; index: number | null };
  toggleCard: (section: string, index: number) => void;
  isEmployee?: boolean;
  employeeRentals?: RentalItem[];
}>();

onMounted(async () => {
  const customerID = getUserID()
  
  try {
      const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/activity/${customerID}`)
      if (res.ok) {
        // TODO: use data with RentalCard
        // const {bookings, rentings, archives} = await res.json()
      }
  } catch (error) {
      console.error('Error calling API:', error);
  }
})

const customerRentals: RentalItem[] = [
  {
    customer_name: "You",
    room_number: 12,
    start_date: "2025-03-24",
    end_date: "2025-03-30",
    cardType: "booking"
  },
  {
    customer_name: "You",
    room_number: 7,
    start_date: "2025-04-01",
    end_date: "2025-04-05",
    cardType: "renting",
    total_price: 300,
    payment: false
  },
  {
    customer_name: "You",
    start_date: "2025-03-10",
    end_date: "2025-03-15",
    cardType: "archive"
  },
  {
    customer_name: "You",
    start_date: "2025-03-10",
    end_date: "2025-03-15",
    cardType: "archive"
  },
  {
    customer_name: "You",
    start_date: "2025-03-10",
    end_date: "2025-03-15",
    cardType: "archive"
  }
];

const defaultEmployeeRentals: RentalItem[] = [
  {
    customer_name: "Alice Johnson",
    room_number: 101,
    start_date: "2025-03-20",
    end_date: "2025-03-26",
    total_price: 600.0,
    payment: true,
    employee_name: "John Doe",
    cardType: "renting"
  },
  {
    customer_name: "Bob Smith",
    room_number: 102,
    start_date: "2025-03-15",
    end_date: "2025-03-20",
    total_price: 500.0,
    payment: true,
    employee_name: "John Doe",
    cardType: "renting"
  },
  {
    customer_name: "Bob Smith",
    room_number: 102,
    start_date: "2025-03-15",
    end_date: "2025-03-20",
    total_price: 500.0,
    payment: true,
    employee_name: "John Doe",
    cardType: "renting"
  }
];

const allRentals = computed<RentalItem[]>(() => {
  return props.isEmployee
    ? props.employeeRentals ?? defaultEmployeeRentals
    : customerRentals;
});

const today = new Date();

const currentRentals = computed(() => {
  return allRentals.value.filter(r => {
    const start = new Date(r.start_date);
    const end = new Date(r.end_date);
    return r.cardType !== 'archive' && start <= today && end >= today;
  });
});

const futureRentals = computed(() => {
  return allRentals.value.filter(r => {
    const start = new Date(r.start_date);
    return r.cardType !== 'archive' && start > today;
  });
});

const archivedRentals = computed(() => {
  return allRentals.value.filter(r => r.cardType === 'archive');
});


function getTitle(rental: RentalItem) {
  return rental.room_number ? `Room ${rental.room_number}` : `Archived Rental`;
}

function getDescription(rental: RentalItem) {
  return `Rented by ${rental.customer_name} until ${rental.end_date}`;
}
</script>
