<template>
    <Form v-slot="$form" :resolver="resolver" @submit="onFormSubmit" class="flex flex-col gap-4 w-full sm:w-56 ">
        <div class="flex flex-col gap-1 ">
            <InputText name="fullname" type="text" placeholder="Full name" fluid />
            <Message v-if="$form.fullname?.invalid" severity="error" size="small" variant="simple"> {{ $form.fullname.error?.message }} </Message>
            <Message v-if="nameNotFound" severity="error" size="small" variant="simple"> Name not found in the database </Message>
        </div>
        <Button type="submit" severity="secondary" label="Submit" />
    </Form>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import { type FormResolverOptions, type FormSubmitEvent } from '@primevue/forms'
import { setAuthCookie } from '../../utils/auth'

const props = defineProps<{
    employee: boolean
}>()

interface Error {
    type: string;
    message: string;
}

const nameNotFound = ref(false)

const resolver = ({ values }: FormResolverOptions) => {
    const errors: Error[] = []

    if (!values.fullname) {
        errors.push({ type: 'required', message: 'Full name is required.' })
    }

    return {
        values,
        errors
    }
}

async function onFormSubmit(e: FormSubmitEvent) {
    nameNotFound.value = false  

    if (e.valid) {
        try {
            const url = props.employee
                ? `${import.meta.env.VITE_BACKEND_HOST}/api/employees?employee_name=${e.values.fullname}`
                : `${import.meta.env.VITE_BACKEND_HOST}/api/customers/${e.values.fullname}`

            const res = await fetch(url)

            if (res.ok) {
                const data = await res.json()

                const user = {
                    ID: props.employee ? data.employee_ID : data.customer_ID,
                    name: data.full_name,
                    role: data?.role ? data.role : 'Customer',
                }

                setAuthCookie(user)
            } else {
                nameNotFound.value = true
            }
        } catch (error) {
            console.error('Error calling API:', error)
            nameNotFound.value = true
        }
    }
}
</script>
