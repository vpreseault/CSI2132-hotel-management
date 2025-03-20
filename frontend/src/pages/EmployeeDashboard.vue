<template>
    <div class="w-full m-auto">
        <div class="mt-16 text-center">
            <h1 class="text-2xl mt-4">Employee View</h1>
            <p class="mt-4">You are logged in!</p>
        </div>

        <NavBar role="employee" @toggleProfile="toggleProfileModal" @toggleHotel="toggleHotelModal" />

        <!-- Current Rentings -->
        <div id="rentings" class="mt-16 p-8 bg-green-100 rounded-lg shadow-md">
            <h2 class="text-center text-xl text-gray-600 font-semibold">Current Rentings</h2>
            <div class="mt-4 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-4 justify-center">
                <GenericCard 
                    v-for="(renting, index) in currentRentings" 
                    :key="index"
                    :title="`Room ${renting.roomNumber}`" 
                    :description="`Rented by ${renting.customerName} until ${renting.checkoutDate}`"
                    :section="'rentings'"
                    :index="index"
                    :expandedCard="expandedCard"
                    @toggle="toggleCard"
                />
            </div>
        </div>

        <!-- Make a Booking -->
        <div id="book" class="mt-16 p-8 bg-green-100 rounded-lg shadow-md">
            <h2 class="text-center text-xl text-gray-600 font-semibold">Make a Booking</h2>
            <form @submit.prevent="makeBooking" class="mt-4">
                <input v-model="newBooking.customerName" type="text" placeholder="Customer Name" class="border p-2 text-black rounded w-full mb-2" required />
                <input v-model.number="newBooking.roomNumber" type="number" placeholder="Room Number" class="border p-2 text-black rounded w-full mb-2" required />
                <input v-model="newBooking.checkoutDate" type="date" class="border p-2 text-black rounded w-full mb-2" required />
                <button type="submit" class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition">
                    Book Room
                </button>
            </form>
        </div>

        <!-- Profile Modal -->
        <div v-if="isProfileModalOpen" class="fixed inset-0 flex justify-center items-center">
            <div class="bg-green-200 p-6 rounded-lg shadow-lg w-1/3">
                <h2 class="text-center text-xl text-black font-semibold">Profile Details</h2>
                <div class="mt-4 space-y-4">
                    <p class="text-lg text-black"><strong>Name:</strong> John Doe</p>
                    <p class="text-lg text-black"><strong>Email:</strong> johndoe@example.com</p>
                    <p class="text-lg text-black"><strong>Phone:</strong> +1 234 567 890</p>
                    <p class="text-lg text-black"><strong>Address:</strong> 123 Main St, New York, NY</p>
                </div>
                <button class="mt-4 w-full px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition"
                        @click="toggleProfileModal">
                    Close
                </button>
            </div>
        </div>

        <!-- Hotel Modal -->
        <div v-if="isHotelModalOpen" class="fixed inset-0 flex justify-center items-center">
            <div class="bg-green-200 p-6 rounded-lg shadow-lg w-1/3">
                <h2 class="text-center text-xl text-black font-semibold">Hotel Details</h2>
                <div class="mt-4 space-y-4">
                    <p><strong>Name:</strong> Grand Palace Hotel</p>
                    <p><strong>Location:</strong> New York City, NY</p>
                    <p><strong>Rating:</strong> 4.5</p>
                    <p><strong>Available Rooms:</strong> 20</p>
                </div>
                <button class="mt-4 w-full px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition"
                        @click="toggleHotelModal">
                    Close
                </button>
            </div>
        </div>

        <!-- Logout Button -->
        <div class="mt-16 text-center">
            <button 
                class="mt-4 px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition" 
                @click="handleLogout">
                Logout
            </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { removeAuthCookie } from '../utils/auth';
import GenericCard from '../components/LandingPage/GenericCard.vue';
import NavBar from '../components/LandingPage/NavBar.vue';

const expandedCard = ref<{ section: string | null; index: number | null }>({ section: null, index: null });
const isProfileModalOpen = ref(false);
const isHotelModalOpen = ref(false);

function toggleProfileModal() {
    isProfileModalOpen.value = !isProfileModalOpen.value;
}

function toggleHotelModal() {
    isHotelModalOpen.value = !isHotelModalOpen.value;
}

const newBooking = ref({
    customerName: '',
    roomNumber: '',
    checkoutDate: ''
});

function makeBooking() {
    if (newBooking.value.customerName && newBooking.value.roomNumber && newBooking.value.checkoutDate) {
        currentRentings.value.push({
            id: currentRentings.value.length + 1,
            customerName: newBooking.value.customerName,
            roomNumber: Number(newBooking.value.roomNumber),
            checkoutDate: newBooking.value.checkoutDate
        });
        newBooking.value = { customerName: '', roomNumber: '', checkoutDate: '' };
    }
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

const currentRentings = ref([
    { id: 1, customerName: "John Doe", roomNumber: 101, checkoutDate: "2025-03-25" },
    { id: 2, customerName: "Jane Smith", roomNumber: 202, checkoutDate: "2025-03-30" }
]);

</script>

