<template>
    <div class="w-full m-auto">
        <NavBar role="customer" @toggleProfile="toggleProfileModal" />

        <div class="mt-16 text-center">
            <h1 class="text-2xl mt-4">Welcome to your one-stop shop for all things Hotel Room Rentings!</h1>
            <p class="mt-4">You are logged in!</p>
        </div>

        <!-- Book -->
        <div id="book" class="mt-16 p-8 bg-green-100 rounded-lg shadow-md">
            <h2 class="text-center text-xl text-gray-600 font-semibold">Book a Hotel</h2>
            <div class="mt-4 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-4 justify-center">
                <GenericCard 
                    v-for="(hotel, index) in bookings" 
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
        </div>

        <!-- Rentings -->
        <div id="rentings" class="mt-16 p-8 bg-green-100 rounded-lg shadow-md">
            <h2 class="text-center text-xl text-gray-600 font-semibold">Rentings Section</h2>
            <div class="mt-4 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-2 gap-4 justify-center">
                <GenericCard 
                    v-for="(rental, index) in rentals" 
                    :key="index"
                    :title="rental.title" 
                    :description="rental.description" 
                    :details="rental.details"
                    :section="'rentings'"
                    :index="index"
                    :expandedCard="expandedCard"
                    @toggle="toggleCard"
                />
            </div>
        </div>

        <!-- Profile -->
        <div v-if="isProfileModalOpen" class="fixed inset-0 flex justify-center items-center">
            <div class="bg-green-200 p-6 rounded-lg shadow-lg w-1/3">
                <h2 class="text-center text-xl text-black font-semibold">Profile Details</h2>
                <div class="mt-4 space-y-4">
                    <p class="text-lg text-black"><strong>Name:</strong> {{ profileData.name }}</p>
                    <p class="text-lg text-black"><strong>Email:</strong> {{ profileData.email }}</p>
                    <p class="text-lg text-black"><strong>Phone:</strong> {{ profileData.phone }}</p>
                    <p class="text-lg text-black"><strong>Address:</strong> {{ profileData.address }}</p>
                </div>
                <button class="mt-4 w-full px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition"
                        @click="toggleProfileModal">
                    Close
                </button>
            </div>
        </div>

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

function toggleCard(section: string, index: number) {
    expandedCard.value = expandedCard.value.section === section && expandedCard.value.index === index 
        ? { section: null, index: null } 
        : { section, index };
}

function toggleProfileModal() {
    isProfileModalOpen.value = !isProfileModalOpen.value;
}

function handleLogout() {
    removeAuthCookie();
    window.location.href = '/';
}

function formatDetails(details: any) {
    return `
        Cost: ${details.cost}
        Rating: ${details.rating}
        Email: ${details.contactEmail}
        Phone: ${details.phone}
    `;
}

const profileData = ref({
    name: "John Doe",
    email: "johndoe@example.com",
    phone: "+1 234 567 890",
    address: "123 Main St, New York, NY"
});

const bookings = ref([
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
]);

const rentals = ref([
    {title: "Current Rental", description: "You have an active rental", details: "Description" },
    {title: "Upcoming Rentals", description: "Your next booking", details: "Description"},
    {title: "Past Rentals", description: "Your previous stay at Hotel 3", details: "Description" }
]);

</script>


