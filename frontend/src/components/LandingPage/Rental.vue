<template>
  <div id="rentings" class="mt-16 p-8 bg-green-100 rounded-lg shadow-md">
    <h2 class="text-center text-xl text-gray-600 font-semibold"> {{ isEmployee ? 'All Hotel Rentings' : 'Your Rentings' }} </h2>
    <div class="mt-4 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-4 justify-center">
      <GenericCard 
        v-for="(rental, index) in rentalsToShow" 
        :key="index"
        :title="getTitle(rental)"
        :description="getDescription(rental)"
        :details="getDetails(rental)"
        :section="'rentings'"
        :index="index"
        :expandedCard="expandedCard"
        @toggle="toggleCard"
      />
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import GenericCard from './GenericCard.vue';

type CustomerRental = {
  title: string;
  description: string;
  details: string;
};

type EmployeeRental = {
  customerName: string;
  roomNumber: number;
  checkoutDate: string;
  email: string;
};

type RentalItem = CustomerRental | EmployeeRental;

const props = defineProps<{
  expandedCard: { section: string | null; index: number | null };
  toggleCard: (section: string, index: number) => void;
  isEmployee?: boolean;
  employeeRentals?: EmployeeRental[];
}>();

const customerRentals: CustomerRental[] = [
  { title: "Current Rental", description: "You have an active rental", details: "Check-in: March 24, 2025" },
  { title: "Upcoming Rental", description: "You have an upcoming stay", details: "Check-in: April 1, 2025" },
  { title: "Past Rental", description: "You stayed at Hotel 3", details: "Check-out: March 20, 2025" }
];

const defaultEmployeeRentals: EmployeeRental[] = [
  { customerName: "Alice Johnson", roomNumber: 101, checkoutDate: "2025-03-26", email: "alice@example.com" },
  { customerName: "Bob Smith", roomNumber: 202, checkoutDate: "2025-03-28", email: "bob@example.com" },
  { customerName: "Charlie Brown", roomNumber: 303, checkoutDate: "2025-03-30", email: "charlie@example.com" }
];

const rentalsToShow = computed<RentalItem[]>(() => {
  return props.isEmployee
    ? props.employeeRentals ?? defaultEmployeeRentals
    : customerRentals;
});

function isEmployeeRental(rental: RentalItem): rental is EmployeeRental {
  return 'roomNumber' in rental;
}

function getTitle(rental: RentalItem) {
  return isEmployeeRental(rental) ? `Room ${rental.roomNumber}` : rental.title;
}

function getDescription(rental: RentalItem) {
  return isEmployeeRental(rental)
    ? `Rented by ${rental.customerName} until ${rental.checkoutDate}`
    : rental.description;
}

function getDetails(rental: RentalItem) {
  return isEmployeeRental(rental)
    ? `Email: ${rental.email}`
    : rental.details;
}
</script>
