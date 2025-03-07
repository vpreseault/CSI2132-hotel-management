<template>
<Form v-slot="$form" :resolver="resolver" @submit="onFormSubmit" class="flex flex-col gap-4 w-full sm:w-56">
    <div class="flex flex-col gap-1">
        <InputText name="fullname" type="text" placeholder="Full name" fluid />
        <Message v-if="$form.fullname?.invalid" severity="error" size="small" variant="simple">{{ $form.fullname.error?.message }}</Message>
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
    <Button type="submit" severity="secondary" label="Submit" />
</Form>
</template>
<script setup lang="ts">
import { ref } from 'vue';
import { type FormResolverOptions, type FormSubmitEvent } from '@primevue/forms';

const props = defineProps<{
    employee: boolean
}>()

const idTypes = ref([
    { type: 'SSN' },
    { type: 'SIN' },
]);
if (!props.employee) {
    idTypes.value.push({ type: 'Drivers License' })
}

interface Error {
    type: string;
    message: string;
}

const resolver = ({ values }: FormResolverOptions) => {
    const errors: Record<string, Error[]> = {
        fullname: [],
        address: [],
        idType: [],
        idNumber: [],
    }

    console.log(values);

    if (!values.fullname) {
        errors.fullname.push({ type: 'required', message: 'Full name is required.' });
    }

    if (!values.address) {
        errors.address.push({ type: 'required', message: 'Address is required.' });
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

    return {
        values,
        errors
    };
};

async function onFormSubmit(e: FormSubmitEvent) {
    if (e.valid) {
        await fetch(`${import.meta.env.VITE_BACKEND_HOST}/api/${props.employee ? 'employees' : 'customers'}`,
            {
                method: 'POST',
                body: JSON.stringify({
                    fullname: e.values.fullname,
                    address: e.values.address,
                    "ID_type": e.values.idType.type,
                    "ID_number": e.values.idNumber,
                })
            }
        )
    }
}
</script>
