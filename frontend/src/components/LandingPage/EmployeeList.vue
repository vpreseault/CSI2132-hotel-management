<template>
    <div id="employee-list" class="bg-green-100 p-6 rounded-lg shadow-md border border-gray-300 mb-10 max-w-5xl mx-auto">
      <h2 class="text-center text-xl text-gray-600 font-semibold">All Employees</h2>
      <div v-if="employees.length" class="mt-6 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 justify-center">
        <EmployeeCard 
          v-for="employee in employees" 
          :key="employee.employee_ID" 
          :name="employee.full_name"
          :role="employee.role"
          :address="employee.address"
          @delete="deleteEmployee(employee.employee_ID)"/>
      </div>
      <p v-else class="text-center text-gray-500 mt-8">No employees found.</p>
    </div>
</template>
  
<script setup lang="ts">
import { onMounted, ref } from 'vue';
import EmployeeCard from './EmployeeCard.vue';
import { getUserID, removeAuthCookie } from '../../utils/auth';
import type { ToastMessageOptions } from 'primevue';

const emit = defineEmits<{
  'delete': [severity: ToastMessageOptions["severity"]],
}>()

type Employee = {
  employee_ID: number;
  full_name: string;
  address: string;
  role: string;
};

const employees = ref<Employee[]>([]);

onMounted(async () => {
  await fetchEmployees()
})

const managerID = ref(getUserID())
async function fetchEmployees() {
  
  try {
      const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/employees?manager_ID=${managerID.value}`)
      if (res.ok) {
        employees.value = await res.json()
      }
  } catch (error) {
      console.error('Error calling API:', error);
  }
}

async function deleteEmployee(id: number) {
  try {
    const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/employees?employee_ID=${id}`,
        {
            method: 'DELETE',
        }
    )
    
    if (res.ok) {
        if (managerID.value === id) {
          removeAuthCookie()
          return
        }
        
        emit('delete', 'success')
        await fetchEmployees()
    } else {
      emit('delete', 'error')
    }
  } catch (error) {
      console.error('Error calling API:', error);
      emit('delete', 'error')
  }
}
</script>
  