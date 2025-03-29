<template>
    <div id="employee-list" class="mt-16 p-8 bg-green-100 rounded-lg shadow-md mx-75">
      <h2 class="text-center text-xl text-gray-600 font-semibold">All Employees</h2>
      <div v-if="employees.length" class="mt-6 grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4 justify-center">
        <EmployeeCard v-for="employee in employees" :key="employee.id" :employee="employee" @delete="deleteEmployee(employee.id)"/>
      </div>
      <p v-else class="text-center text-gray-500 mt-8">No employees found.</p>
    </div>
</template>
  
<script setup lang="ts">
import { ref } from 'vue';
import EmployeeCard from './EmployeeCard.vue';

type Employee = {
id: number;
name: string;
role: string;
email: string;
};

const employees = ref<Employee[]>([
{ id: 1, name: 'Alice Johnson', role: 'Employee', email: 'alice@hotel.com' },
{ id: 2, name: 'Bob Smith', role: 'Employee', email: 'bob@hotel.com' },
{ id: 3, name: 'Charlie Rose', role: 'Employee', email: 'charlie@hotel.com' },
]);

function deleteEmployee(id: number) {
employees.value = employees.value.filter(emp => emp.id !== id);
}
</script>
  