<template>
    <div class="card">
        <Tabs value="0">
            <TabList>
                <Tab value="0">Customer</Tab>
                <Tab value="1">Employee</Tab>
            </TabList>
            <TabPanels class="d-flex flex-align-center">
                <TabPanel value="0">
                    <LoginUser v-if="mode === 'login'" />
                    <RegisterCustomer v-else />
                </TabPanel>
                <TabPanel value="1">
                    <LoginUser v-if="mode === 'login'" />
                    <RegisterEmployee v-else />
                </TabPanel>
            </TabPanels>
        </Tabs>
        <p class="my-3">{{ modeToggleText }} <Button :label="modeToggleButtonText" variant="link" @click="handleModeToggle" class="!p-0" /> instead.</p>
    </div>
</template>

<script setup lang="ts">
import {computed, ref} from 'vue'
import RegisterCustomer from "../components/Account/RegisterCustomer.vue";

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
</script>
