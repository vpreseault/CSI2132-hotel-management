<template>
  <div class="fixed inset-0 flex justify-center items-center">
    <div class="bg-green-200 p-6 rounded-lg shadow-lg w-full max-w-xl">
      <h2 class="text-center text-xl text-black font-semibold mb-4">{{ role === 'employee' ? 'Employee' : 'Customer' }} Profile Details</h2>

      <div class="flex flex-col gap-4">
        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Name</label>
          <InputText v-model="profileData.name" :disabled="!isEditing" placeholder="Your Full Name" class="w-full"/>
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Email</label>
          <InputText v-model="profileData.email" :disabled="!isEditing" placeholder="Your Email Address" class="w-full"/>
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Phone</label>
          <InputText v-model="profileData.phone" :disabled="!isEditing" placeholder="Your Phone Number" class="w-full"/>
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Address</label>
          <InputText v-model="profileData.address" :disabled="!isEditing" placeholder="Your Address" class="w-full"/>
        </div>
      </div>

      <div class="flex justify-between mt-6">
        <Button :label="isEditing ? 'Save' : 'Edit'" :icon="isEditing ? 'pi pi-check' : 'pi pi-pencil'" class="bg-green-500 border-none hover:bg-green-600 pl-2" @click="toggleEdit" />
        <Button label="Close" icon="pi pi-times" class="bg-red-500 border-none hover:bg-red-600 pl-2" @click="toggleProfileModal"/>
      </div>

      <div class="mt-4 text-center">
        <Button label="Delete Account" icon="pi pi-trash" class="bg-red-600 border-none hover:bg-red-700 px-4 pl-2" @click="deleteAccount"/>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, watch } from 'vue';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';

const props = defineProps<{
  role: 'customer' | 'employee';
  toggleProfileModal: () => void;
}>();

const isEditing = ref(false);

const profileData = ref({
  name: '',
  email: '',
  phone: '',
  address: ''
});

watch(() => props.role, (newRole) => {
  if (newRole === 'employee') {
    profileData.value = {
      name: 'Alice Employee',
      email: 'alice@hotel.com',
      phone: '+1 555 234 5678',
      address: '456 Staff Rd, Hotel HQ'
    };
  } else {
    profileData.value = {
      name: 'John Doe',
      email: 'johndoe@example.com',
      phone: '+1 234 567 890',
      address: '123 Main St, New York, NY'
    };
  }
}, { immediate: true });

function toggleEdit() {
  if (isEditing.value) {
    console.log(`Saving ${props.role} profile:`, profileData.value);
    alert("Profile saved!");
  }
  isEditing.value = !isEditing.value;
}

function deleteAccount() {
  const confirmDelete = confirm("Are you sure you want to delete your account?");
  if (confirmDelete) {
    console.log(`${props.role} account deleted.`);
    alert("Your account has been deleted.");
  }
}
</script>
