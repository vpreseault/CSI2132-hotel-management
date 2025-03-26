<template>
    <div class="fixed inset-0 flex justify-center items-center">
        <div class="bg-green-200 p-6 rounded-lg shadow-lg w-1/3">
            <h2 class="text-center text-xl text-black font-semibold">Create Employee Account</h2>
            <form @submit.prevent="createEmployee" class="mt-4">
                <input v-model="newEmployee.name" type="text" placeholder="Employee Name" class="border p-2 text-black rounded w-full mb-2" required />
                <input v-model="newEmployee.email" type="email" placeholder="Employee Email" class="border p-2 text-black rounded w-full mb-2" required />
                <input v-model="newEmployee.contact" type="text" placeholder="Contact Info" class="border p-2 text-black rounded w-full mb-2" required />
                <input v-model="newEmployee.ssn" type="text" placeholder="SSN/SIN" class="border p-2 text-black rounded w-full mb-2" required />
                <button type="submit" class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition"> Create Account </button>
            </form>
            <button class="mt-4 w-full px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition" @click="$emit('close')"> Close </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';

const newEmployee = ref({
    name: '',
    email: '',
    contact: '',
    ssn: ''
});

function createEmployee() {
    console.log("New employee data:", newEmployee.value);
    newEmployee.value = { name: '', email: '', contact: '', ssn: '' };
}
</script>
