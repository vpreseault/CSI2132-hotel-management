<template>
    <div class="w-100 m-auto">
        <div class="card">
            <Tabs v-model:value="tabValue">
                <TabList>
                    <Tab value="customer">Customer</Tab>
                    <Tab v-if="mode === 'login'" value="employee">Employee</Tab>
                </TabList>
                <TabPanels class="flex flex-col items-center">
                    <h1 class="text-xl mb-5 my-2">{{ headerText }}</h1>
                    <LoginUser v-if="mode === 'login'" :employee="tabValue === 'employee'" />
                    <RegisterUser />
                </TabPanels>
            </Tabs>
            <p class="my-3">{{ modeToggleText }} <Button :label="modeToggleButtonText" variant="link" @click="handleModeToggle" class="!p-0" /> instead.</p>
        </div>
    </div>
</template>

<script setup lang="ts">
import { computed, ref } from 'vue'
import RegisterUser from "../components/Account/RegisterUser.vue";
import LoginUser from "../components/Account/LoginUser.vue";
import { Tabs } from 'primevue';

type Tabs = 'employee'|'customer'
const tabValue = ref<Tabs>('customer')

type AuthModes = 'login'|'register'

const mode = ref<AuthModes>('login')
const modeToggleText = computed(() => `${mode.value === 'login' ? 'Don\'t':'Already'} have an account?`)
const modeToggleButtonText = computed(() => mode.value === 'login' ? 'Register':'Login')

function handleModeToggle() {
    if (mode.value === 'login') {
        mode.value = 'register'
    } else {
        mode.value = 'login'
    }
}

const headerText = computed(() => `${tabValue.value ==='employee' ? "Employee":"Customer"} ${mode.value === 'login' ? "login":"registration"}`)
</script>
