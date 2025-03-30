<template>
  <div class="fixed inset-0 flex justify-center items-center z-50">
    <div class="bg-green-200 p-6 rounded-lg shadow-lg w-full max-w-xl">
      <h2 class="text-center text-xl text-black font-semibold mb-4">Edit Room {{ form.room_number }}</h2>

      <div class="flex flex-col gap-4">
        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Room Number</label>
          <InputText v-model="form.room_number" class="w-full" />
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Capacity: {{ form.capacity }}</label>
          <Slider v-model="form.capacity" :min="1" :max="10" class="w-full" />
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Price ($)</label>
          <InputNumber v-model="form.price" class="w-full" inputClass="w-full" />
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">View Type</label>
          <InputText v-model="form.view_type" class="w-full" />
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Extendable</label>
          <InputSwitch v-model="form.extendable" />
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Damaged</label>
          <InputSwitch v-model="form.damaged" />
        </div>

        <div class="flex flex-col gap-1">
          <label class="text-black font-medium">Amenities</label>
          <div class="flex flex-wrap gap-2">
            <div v-for="amenity in allAmenities" :key="amenity.amenity_ID" class="flex items-center gap-1">
              <Checkbox v-model="selectedAmenities" :inputId="String(amenity.amenity_ID)" name="amenity"
                :value="amenity.amenity_name" />
              <label :for="String(amenity.amenity_ID)">{{ amenity.amenity_name }}</label>
            </div>
          </div>
        </div>
      </div>

      <div class="flex justify-between mt-6">
        <Button label="Save" icon="pi pi-check" class="bg-green-500 border-none hover:bg-green-600" @click="save" />
        <Button label="Cancel" icon="pi pi-times" class="bg-red-500 border-none hover:bg-red-600" @click="close" />
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from 'vue'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import Slider from 'primevue/slider'
import InputSwitch from 'primevue/inputswitch'
import Checkbox from 'primevue/checkbox'
import type { Amenity, Room } from '../../types'

const props = defineProps<{
  room: Room;
  allAmenities: Amenity[]
  onClose: () => void;
  onSave: (updatedRoom: Room) => void;
}>()

const form = ref<Room>({ ...props.room })
const selectedAmenities = ref(props.room.amenities.map((amenity) => amenity.amenity_name))

function save() {
  form.value.amenities = props.allAmenities.filter(
    (amenity) => selectedAmenities.value.find(
      (selectedAmenity: string) =>
        selectedAmenity === amenity.amenity_name
    )
  )

  props.onSave(form.value)
  props.onClose()
}

function close() {
  props.onClose()
}
</script>
