<template>
  <div class="fixed inset-0 flex justify-center items-center">
    <div class="bg-green-200 p-6 rounded-lg shadow-lg w-full max-w-xl">
      <h2 class="text-center text-xl text-black font-semibold mb-4">{{ userCookie.role }} Profile Details</h2>

      <div class="flex flex-col gap-4">
        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Name</label>
          <InputText v-model="profileData.name" :disabled="!isEditing" placeholder="Your Full Name" class="w-full"/>
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Address</label>
          <InputText v-model="profileData.address" :disabled="!isEditing" placeholder="Your Address" class="w-full"/>
        </div>
      </div>

      <div class="flex justify-between mt-6">
        <Button :label="isEditing ? 'Save' : 'Edit'" :icon="isEditing ? 'pi pi-check' : 'pi pi-pencil'" class="bg-green-500 border-none hover:bg-green-600 pl-2" @click="toggleEdit" />
        <Button label="Delete Account" icon="pi pi-trash" class="bg-red-600 border-none hover:bg-red-700 px-4 pl-2" @click="deleteAccount"/>
        <Button label="Close" icon="pi pi-times" class="pl-2" severity="info" @click="closeButtonClicked"/>
      </div>
      <Message v-if="message.text" :severity="message.severity" variant="simple" class="flex justify-center mt-8">{{ message.text }}</Message>
    </div>
  </div>
</template>

<script setup lang="ts">
import { reactive, ref } from 'vue';
import InputText from 'primevue/inputtext';
import Button from 'primevue/button';
import { getAuthCookie, setAuthCookie } from '../../utils/auth';
import type { ToastMessageOptions } from 'primevue';

const props = defineProps<{
  toggleProfileModal: () => void,
}>();

const isEditing = ref(false);
const editsMade = ref(false)

const userCookie = getAuthCookie()
const profileData = ref({
  name: userCookie.name,
  address: userCookie.address
});

const message = reactive<{
  severity?: ToastMessageOptions["severity"],
  text?: string,
}>({})

function toggleEdit() {
  if (isEditing.value) {
    updateAccount()
  }
  isEditing.value = !isEditing.value;
}

async function updateAccount() {
  try {
      const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/${userCookie.role === "Employee" ? 'employees' : 'customers'}/${userCookie.ID}`,
          {
              method: 'PATCH',
              body: JSON.stringify({
                  "full_name": profileData.value.name,
                  address: profileData.value.address,
              })
          }
      )
      
      if (res.ok) {
          editsMade.value = true

          message.severity = 'success'
          message.text = 'Account updated successfully.'
          
          userCookie.name = profileData.value.name
          userCookie.address = profileData.value.address
          setAuthCookie(userCookie)
          return
      }
      message.severity = 'error'
      message.text = 'Account update Failed.'
  } catch (error) {
      console.error('Error calling API:', error);
      message.severity = 'error'
      message.text = 'Account update Failed.'
  }
}

function closeButtonClicked() {
  if (editsMade.value) {
    editsMade.value = false
    window.location.reload()
  }
  props.toggleProfileModal()
}

function deleteAccount() {
  const confirmDelete = confirm("Are you sure you want to delete your account?");
  if (confirmDelete) {
    console.log(`${props.role} account deleted.`);
    alert("Your account has been deleted.");
  }
}
</script>
