<template>
    <Form v-slot="$form" :resolver="resolver" @submit="onFormSubmit" class="flex flex-col gap-4 w-full sm:w-56">
        <div class="flex flex-col gap-1">
            <InputText name="fullname" type="text" placeholder="Full name" fluid />
            <Message v-if="$form.fullname?.invalid" severity="error" size="small" variant="simple">{{ $form.fullname.error?.message }}</Message>
        </div>
        <Button type="submit" severity="secondary" label="Submit" />
    </Form>
</template>
<script setup lang="ts">
import { type FormResolverOptions, type FormSubmitEvent } from '@primevue/forms';
import { setAuthCookie } from '../../utils/auth';

const props = defineProps<{
    employee: boolean
}>()

interface Error {
    type: string;
    message: string;
}

const resolver = ({ values }: FormResolverOptions) => {
    const errors: Error[] = []

    if (!values.fullname) {
        errors.push({ type: 'required', message: 'Full name is required.' });
    }

    return {
        values,
        errors
    };
};

async function onFormSubmit(e: FormSubmitEvent) {
    if (e.valid) {
        try {
            const res = await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/${props.employee ? `employees?employee_name=${e.values.fullname}` : `customers/${e.values.fullname}`}`)
            if (res.ok) {
                const data = await res.json()

                const user = {
                    ID: props.employee ? data.employee_ID: data.customer_ID,
                    name: data.full_name,
                    role: data?.role ? data.role : 'Customer',
                }

                setAuthCookie(user)
            }
        } catch (error) {
            console.error('Error calling API:', error);
        }
    }
}
</script>
    