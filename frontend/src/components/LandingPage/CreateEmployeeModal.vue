<template>
    <div class="fixed inset-0 flex justify-center items-center">
        <div class="bg-green-200 p-6 rounded-lg shadow-lg w-1/3">
            <h2 class="text-center text-xl text-black font-semibold mb-3">Create Employee Account</h2>
            <Form v-slot="$form" :resolver="resolver" @submit="onFormSubmit" class="flex flex-col gap-4 w-full">
                <div class="flex flex-col gap-1">
                    <InputText name="fullName" type="text" placeholder="Full name" fluid />
                    <Message v-if="$form.fullName?.invalid" severity="error" size="small" variant="simple">{{ $form.fullName.error?.message }}</Message>
                </div>
                <div class="flex flex-col gap-1">
                    <InputText name="address" type="text" placeholder="Address" fluid />
                    <Message v-if="$form.address?.invalid" severity="error" size="small" variant="simple">{{ $form.address.error?.message }}</Message>
                </div>
                <div class="flex flex-col gap-1">
                    <Select name="idType" :options="idTypes" optionLabel="type" placeholder="Select an ID type" fluid />
                    <Message v-if="$form.idType?.invalid" severity="error" size="small" variant="simple">{{ $form.idType.error?.message }}</Message>
                </div>
                <div class="flex flex-col gap-1">
                    <InputText name="idNumber" type="text" placeholder="ID number" fluid />
                    <Message v-if="$form.idNumber?.invalid" severity="error" size="small" variant="simple">{{ $form.idNumber.error?.message }}</Message>
                </div>
                <div class="flex flex-col gap-1">
                    <Select name="role" :options="roles" optionLabel="type" placeholder="Select a role" fluid />
                    <Message v-if="$form.role?.invalid" severity="error" size="small" variant="simple">{{ $form.role.error?.message }}</Message>
                </div>
                <Button type="submit" :severity="serverError ? 'error' : 'secondary'" label="Create Account" class="px-4 py-2 bg-blue-500 text-white rounded-md hover:bg-blue-600 transition" />
                <Message v-if="serverError" severity="error" size="small" variant="simple">{{ serverError }}</Message>
            </Form>
            <button class="mt-4 w-full px-4 py-2 bg-red-500 text-white rounded-md hover:bg-red-600 transition" @click="emit('close')"> Close </button>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref } from 'vue';
import { type FormResolverOptions, type FormSubmitEvent } from '@primevue/forms';
import type { FormError } from '../../types';
import { getUserID } from '../../utils/auth';

const emit = defineEmits<{
    close: []
    created: []
}>()

const idTypes = ref([
    { type: 'SSN' },
    { type: 'SIN' },
    { type: 'Drivers License' },
]);

const roles = ref([
    { type: 'Employee' },
    { type: 'Manager' },
]);

const serverError = ref('')
const success = ref(false)

const resolver = ({ values }: FormResolverOptions) => {
    const errors: Record<string, FormError[]> = {
        fullName: [],
        address: [],
        idType: [],
        idNumber: [],
        role: [],
    }

    if (!values.fullName) {
        errors.fullName.push({ type: 'required', message: 'Full name is required.' });
    } else if (!/^[A-Za-z]+(?:[\s'-][A-Za-z]+)+$/.test(values.fullName.trim())) {
        errors.fullName.push({
            type: 'format',
            message: 'Enter a valid full name (e.g., John Doe).',
        });
    }

    if (!values.address) {
        errors.address.push({ type: 'required', message: 'Address is required.' });
    } else if (!/^\d+\s+[A-Za-z]/.test(values.address)) {
        errors.address.push({ type: 'format', message: 'Enter a valid address (e.g., 123 Main Street, City, State).' });
    }

    if (!values.idType) {
        errors.idType.push({ type: 'required', message: 'ID type is required.' });
    }

    if (!values.idNumber) {
        errors.idNumber.push({ type: 'required', message: 'ID number is required.' });
    }

    if (values.idNumber?.length !== 9) {
        errors.idNumber.push({ type: 'length', message: 'ID number must be 9 characters long.' });
    }

    if (!values.role) {
        values.role = roles.value[0]
    }

    return {
        values,
        errors
    };
};

async function onFormSubmit(e: FormSubmitEvent) {
    serverError.value = ''
    success.value = false
    if (e.valid) {
        const managerID = getUserID()
        try {
            const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/employees`,
                {
                    method: 'POST',
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify({
                        "manager_ID": managerID,
                        "full_name": e.values.fullName,
                        address: e.values.address,
                        "ID_type": e.values.idType.type,
                        "ID_number": e.values.idNumber,
                        role: e.values.role.type,
                    })
                }
            )
            
            if (res.ok) {
                emit('created')
                return
            }

            serverError.value = await res.text()
        } catch (error) {
            console.error('Error calling API:', error);
            serverError.value = error as string
        }
    }
}
</script>
