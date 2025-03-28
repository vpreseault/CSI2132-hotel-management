<template>
    <div class="w-full m-auto">
      <div class="mt-16 text-center">
        <h1 class="text-2xl mt-4">Manager View</h1>
        <p class="mt-4">You are logged in as a manager!</p>
      </div>
  
      <NavBar role="manager" @toggleProfile="toggleProfileModal" @toggleHotel="toggleHotelModal" @toggleCreateEmployee="toggleCreateEmployeeModal"/>
      <Rental :expandedCard="expandedCard" :toggleCard="toggleCard" :isEmployee="true" :employeeRentals="allHotelRentings"/>
      <Booking :expandedCard="expandedCard" :toggleCard="toggleCard" :isEmployee="true" @createBooking="handleEmployeeBooking"/>
      <CreateEmployeeModal v-if="isCreateEmployeeModalOpen" @close="toggleCreateEmployeeModal" @created="employeeCreatedToast"/>
      <Profile v-if="isProfileModalOpen" role="employee" :toggleProfileModal="toggleProfileModal"/>
      <HotelModal v-if="isHotelModalOpen" @close="toggleHotelModal"/>
  
      <div class="mt-16 text-center">
        <button class="mt-4 px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition" @click="handleLogout"> Logout </button>
      </div>
      <Toast position="bottom-center" />
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { removeAuthCookie } from '../utils/auth';
import { useToast } from "primevue/usetoast";
import NavBar from '../components/LandingPage/NavBar.vue';
import Rental from '../components/LandingPage/Rental.vue';
import Booking from '../components/LandingPage/Booking.vue';
import Profile from '../components/LandingPage/Profile.vue';
import HotelModal from '../components/LandingPage/HotelModal.vue';
import CreateEmployeeModal from '../components/LandingPage/CreateEmployeeModal.vue';

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

const toast = useToast();
const expandedCard = ref<{ section: string | null; index: number | null }>({ section: null, index: null });
const isProfileModalOpen = ref(false);
const isHotelModalOpen = ref(false);
const isCreateEmployeeModalOpen = ref(false);

function toggleProfileModal() {
isProfileModalOpen.value = !isProfileModalOpen.value;
}

function toggleHotelModal() {
isHotelModalOpen.value = !isHotelModalOpen.value;
}

function toggleCreateEmployeeModal() {
isCreateEmployeeModalOpen.value = !isCreateEmployeeModalOpen.value;
}

function toggleCard(section: string, index: number) {
expandedCard.value = expandedCard.value.section === section && expandedCard.value.index === index 
    ? { section: null, index: null } 
    : { section, index };
}

function handleLogout() {
removeAuthCookie();
window.location.href = '/';
}

function handleEmployeeBooking(booking: {
customerName: string;
roomNumber: number;
checkoutDate: string;
email: string;
}) {
allHotelRentings.value.push({
    customer_name: booking.customerName,
    room_number: booking.roomNumber,
    start_date: new Date().toISOString().split('T')[0],
    end_date: booking.checkoutDate,
    employee_name: "Manager",
    cardType: "renting",
    total_price: 0,
    payment: false
});
}

const allHotelRentings = ref<RentalItem[]>([
{
    customer_name: "John Doe",
    room_number: 101,
    start_date: "2025-03-20",
    end_date: "2025-03-25",
    employee_name: "Manager",
    cardType: "renting",
    total_price: 500,
    payment: true
},
{
    customer_name: "Jane Smith",
    room_number: 202,
    start_date: "2025-03-27",
    end_date: "2025-03-30",
    employee_name: "Manager",
    cardType: "renting",
    total_price: 400,
    payment: false
},
{
    customer_name: "Bob Builder",
    room_number: 303,
    start_date: "2025-03-28",
    end_date: "2025-04-02",
    employee_name: "Manager",
    cardType: "renting",
    total_price: 450,
    payment: true
}
]);

function employeeCreatedToast() {
toggleCreateEmployeeModal();
toast.add({ severity: 'success', summary: 'Success', detail: 'Employee account created.', life: 3000 });
}
</script>
  