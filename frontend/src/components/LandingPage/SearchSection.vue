<template>
    <Card class="w-full max-w-3xl mx-auto mt-6">
      <template #title>Search Filters</template>
      <template #content>
        <div class="flex flex-col gap-6">

            <!-- Cities -->
            <div class="flex flex-col gap-1">
                <label class="font-medium">City</label>
                <MultiSelect v-model="filters.city" :options="cities" optionLabel="name" filter placeholder="Select Cities" :maxSelectedLabels="3" class="w-full md:w-80" />
            </div>
    
            <!-- Date Range -->
            <div class="flex flex-col md:flex-row gap-4">
                <div class="flex flex-col">
                <label class="font-medium">Start Date</label>
                <Calendar v-model="filters.startDate" showIcon placeholder="Start Date" class="w-full md:w-56" />
                </div>
                <div class="flex flex-col">
                <label class="font-medium">End Date</label>
                <Calendar v-model="filters.endDate" :minDate="filters.startDate" showIcon placeholder="End Date" class="w-full md:w-56" />
                </div>
            </div>
    
            <!-- Hotel Chain -->
            <div class="flex flex-col gap-1">
                <label class="font-medium">Hotel Chain</label>
                <Dropdown v-model="filters.chain" :options="chains" optionLabel="name" placeholder="Select Hotel Chain" class="w-full md:w-80" />
            </div>
    
            <!-- Sliders -->
            <div class="grid grid-cols-1 md:grid-cols-2 gap-6">
                <div class="flex items-center gap-4">
                    <div class="flex-1">
                    <label class="font-medium">Room Capacity</label>
                    <Slider v-model="filters.roomCapacity" :min="1" :max="10" class="w-56" />
                    </div>
                    <span>{{ filters.roomCapacity }}</span>
                </div>
                <div class="flex items-center gap-4">
                    <div class="flex-1">
                    <label class="font-medium">Hotel Category</label>
                    <Slider v-model="filters.hotelCategory" :min="1" :max="5" class="w-56" />
                    </div>
                    <span>{{ filters.hotelCategory }}</span>
                </div>
                <div class="flex items-center gap-4">
                    <div class="flex-1">
                    <label class="font-medium">Total Rooms</label>
                    <Slider v-model="filters.totalRooms" :min="1" :max="10" class="w-56" />
                    </div>
                    <span>{{ filters.totalRooms }}</span>
                </div>
                <div class="flex items-center gap-4">
                    <div class="flex-1">
                    <label class="font-medium">Room Price</label>
                    <Slider v-model="filters.roomPrice" :min="50" :max="1000" class="w-56" />
                    </div>
                    <span>${{ filters.roomPrice }}</span>
                </div>
            </div>

            <!-- Search Button -->
            <div class="flex justify-end mt-4">
                <Button label="Search" icon="pi pi-search" @click="submitSearch" />
            </div>
        </div>
      </template>
    </Card>
</template>
  
<script setup lang="ts">
import { ref } from 'vue';
import MultiSelect from 'primevue/multiselect';
import Calendar from 'primevue/calendar';
import Dropdown from 'primevue/dropdown';
import Slider from 'primevue/slider';
import Button from 'primevue/button';
import Card from 'primevue/card';

const filters = ref({
    city: [],
    startDate: undefined as Date | undefined,
    endDate: undefined as Date | undefined,
    chain: null,
    roomCapacity: 1,
    hotelCategory: 1,
    totalRooms: 50,
    roomPrice: 200
});


const cities = ref([
    { name: 'New York' },
    { name: 'Los Angeles' },
    { name: 'Toronto' },
    { name: 'Chicago' },
]);

const chains = ref([
    { name: 'Hilton' },
    { name: 'Marriott' },
    { name: 'Holiday Inn' },
    { name: 'Best Western' },
]);

function submitSearch() {
    console.log('Search triggered with filters:', filters.value);
}
</script>

<style scoped>
label {
margin-bottom: 4px;
}
</style>
