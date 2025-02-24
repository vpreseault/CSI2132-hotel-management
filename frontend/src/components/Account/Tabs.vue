<template>

<template>
    <div class="card">
        <Tabs value="0">
            <TabList>
                <Tab value="0">Customer</Tab>
                <Tab value="1">Employee</Tab>
            </TabList>
            <TabPanels>
                <TabPanel value="0">
                    <slot name="customer" />
                </TabPanel>
                <TabPanel value="1">
                    <slot name="employee" />
                </TabPanel>
            </TabPanels>
        </Tabs>
    </div>
</template>

<script setup>
</script>
    
</template>
<script setup lang="ts">
import Tabs from 'primevue/tabs';
import TabList from 'primevue/tablist';
import Tab from 'primevue/tab';
import TabPanels from 'primevue/tabpanels';
import TabPanel from 'primevue/tabpanel';
</script>