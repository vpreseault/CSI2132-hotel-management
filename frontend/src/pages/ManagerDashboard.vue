<template>
    <div class="w-full m-auto">
        <div class="mt-16 text-center">
            <h1 class="text-2xl mt-4">Manager View</h1>
            <p class="mt-4">You are logged in as a manager!</p>
        </div>

        <NavBar role="manager" @toggleProfile="toggleProfileModal" @toggleHotel="toggleHotelModal" @toggleCreateEmployee="toggleCreateEmployeeModal" />

        <!-- Current Rentings -->
        <div id="rentings" class="mt-16 p-8 bg-green-100 rounded-lg shadow-md">
            <h2 class="text-center text-xl text-gray-600 font-semibold">Current Rentings</h2>
            <div class="mt-4 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-4 justify-center">
                <div v-for="(renting, index) in currentRentings" :key="index" 
                     class="p-4 rounded-lg shadow-md cursor-pointer hover:bg-gray-100 transition grid place-items-center"
                     @click="toggleCard('rentings', index)">
                    <h3 class="text-lg text-gray-600 font-semibold">Room {{ renting.roomNumber }}</h3>
                    <p class="text-sm text-gray-600">Rented by {{ renting.customerName }} until {{ renting.checkoutDate }}</p>
                    <div v-if="expandedCard.section === 'rentings' && expandedCard.index === index" 
                         class="mt-4 p-3 bg-green-100 rounded-md">
                        <p class="text-sm text-gray-600">More details coming soon...</p>
                    </div>
                </div>
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

        <!-- Create Employee Account -->
        <div v-if="isCreateEmployeeModalOpen" class="fixed inset-0 flex justify-center items-center">
            <div class="bg-green-200 p-6 rounded-lg shadow-lg w-1/3">
                <h2 class="text-center text-xl text-black font-semibold">Create Employee Account</h2>
                <form @submit.prevent="createEmployee" class="mt-4">
                    <input v-model="newEmployee.name" type="text" placeholder="Employee Name" class="border p-2 text-black rounded w-full mb-2" required />
                    <input v-model="newEmployee.email" type="email" placeholder="Employee Email" class="border p-2 text-black rounded w-full mb-2" required />
                    <input v-model="newEmployee.contact" type="text" placeholder="Contact Info" class="border p-2 text-black rounded w-full mb-2" required />
                    <input v-model="newEmployee.ssn" type="text" placeholder="SSN/SIN" class="border p-2 text-black rounded w-full mb-2" required />
                    <button type="submit" class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition">
                        Create Account
                    </button>
                </form>
                <button class="mt-4 w-full px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition"
                        @click="toggleCreateEmployeeModal">
                    Close
                </button>
            </div>
        </div>

        <!-- Hotel -->
        <div v-if="isHotelModalOpen" class="fixed inset-0 flex justify-center items-center">
            <div class="bg-green-200 p-6 rounded-lg shadow-lg w-1/3">
                <h2 class="text-center text-xl text-black font-semibold">Hotel Details</h2>
                <div class="mt-4 space-y-4">
                    <p><strong>Name:</strong> {{ hotelData.name }}</p>
                    <p><strong>Location:</strong> {{ hotelData.location }}</p>
                    <p><strong>Rating:</strong> {{ hotelData.rating }} </p>
                    <p><strong>Available Rooms:</strong> {{ hotelData.availableRooms }}</p>
                </div>
                <button class="mt-4 w-full px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition"
                        @click="toggleHotelModal">
                    Close
                </button>
            </div>
        </div>

        <!-- Profile -->
        <div v-if="isProfileModalOpen" class="fixed inset-0 flex justify-center items-center">
            <div class="bg-green-200 p-6 rounded-lg shadow-lg w-1/3">
                <h2 class="text-center text-xl text-black font-semibold">Profile Details</h2>
                <div class="mt-4 space-y-4">
                    <div>
                        <label class="block text-black text-sm font-medium">Name</label>
                        <input type="text" v-model="profileData.name" class="w-full p-2 text-black border rounded-md" disabled>
                    </div>
                    <div>
                        <label class="block text-sm text-black font-medium">Email</label>
                        <input type="email" v-model="profileData.email" class="w-full p-2 text-black border rounded-md" disabled>
                    </div>
                    <div>
                        <label class="block text-sm text-black font-medium">Phone</label>
                        <input type="text" v-model="profileData.phone" class="w-full text-black p-2 border rounded-md" disabled>
                    </div>
                    <div>
                        <label class="block text-sm text-black font-medium">Address</label>
                        <input type="text" v-model="profileData.address" class="w-full p-2 text-black border rounded-md" disabled>
                    </div>
                </div>
                <button class="mt-4 w-full px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition"
                        @click="toggleProfileModal">
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

const expandedCard = ref<{ section: string | null; index: number | null }>({ section: null, index: null });
const isCreateEmployeeModalOpen = ref(false);
const isProfileModalOpen = ref(false);
const isHotelModalOpen = ref(false);

const profileData = ref({
    name: "John Doe",
    email: "johndoe@example.com",
    phone: "+1 234 567 890",
    address: "123 Main St, New York, NY"
});

const hotelData = ref({
    name: "Grand Palace Hotel",
    location: "New York City, NY",
    rating: 4.5,
    availableRooms: 20
});

const newEmployee = ref({
    name: '',
    email: '',
    contact: '',
    ssn: ''
});

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

function toggleCreateEmployeeModal() {
    isCreateEmployeeModalOpen.value = !isCreateEmployeeModalOpen.value;
}

function toggleProfileModal() {
    isProfileModalOpen.value = !isProfileModalOpen.value;
}

function toggleHotelModal() {
    isHotelModalOpen.value = !isHotelModalOpen.value;
}

function createEmployee() {
    console.log("New employee data:", newEmployee.value);
    newEmployee.value = { name: '', email: '', contact: '', ssn: '' };
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

