<script setup lang="ts">
import { ref } from 'vue';
interface Item {
  id: number;
  name: string;
  value: string;
}

const items = ref<Item[]>([]); // items is now an array of Item

const callApi = async () => {
    try {
        const response = await fetch('http://localhost:8080/api/hello');
        const data: Item[] = await response.json(); // Type the response as Item[]
        items.value = data;
    } catch (error) {
        console.error('Error calling API:', error);
    }
};
</script>

<template>
    <div>
        <button @click="callApi">Call API</button>
        <div v-if="items.length > 0">
            <ul>
                <li v-for="item in items" :key="item.id">
                    {{ item.name }}: {{ item.value }} (ID: {{item.id}})
                </li>
            </ul>
        </div>
        <p v-else>No items found.</p>
    </div>
</template>