<template>
  <Combobox v-model="value">
    <div class="relative">
      <div
          class="relative cursor-pointer overflow-hidden text-left focus:outline-none text-sm rounded-l-md border border-r-0 border-slate-800"
      >
        <ComboboxInput
            class="w-full border-0 bg-transparent dark:text-gray-400 py-2 pl-3 pr-10 text-sm leading-5 text-gray-900 focus:ring-0"
            :displayValue="(i) => i.name"
            @change="query = $event.target.value"
        />
        <ComboboxButton
            class="absolute inset-y-0 right-0 flex items-center pr-2"
        >
          <SelectorIcon class="h-5 w-5 text-gray-400"/>
        </ComboboxButton>
      </div>
      <TransitionRoot
          leave="transition ease-in duration-100"
          leaveFrom="opacity-100"
          leaveTo="opacity-0"
          @after-leave="query = ''"
      >
        <ComboboxOptions
            class="absolute mt-1 max-h-60 w-full overflow-auto rounded-md bg-slate-900 py-1 text-base ring-1 ring-black ring-opacity-5 focus:outline-none text-sm"
        >
          <ComboboxOption
              v-for="item in filteredItems"
              as="template"
              :key="item.id"
              :value="item"
              v-slot="{ selected, active }"
          >
            <li
                class="relative cursor-pointer select-none p-2 hover:text-green-500">
                <span
                    class="block truncate"
                >
                  {{ item.name }}
                </span>
            </li>
          </ComboboxOption>
        </ComboboxOptions>
      </TransitionRoot>
    </div>
  </Combobox>
</template>

<script setup lang="ts">
import {
  Combobox,
  ComboboxButton,
  ComboboxInput,
  ComboboxOption,
  ComboboxOptions,
  TransitionRoot,
} from '@headlessui/vue'
import {SelectorIcon} from '@heroicons/vue/solid'
import {computed, ref} from 'vue'

const props = defineProps({
  items: {type: Array as () => Array<Item>, required: true},
  modelValue: {type: Object as () => Item, required: true},
})

const emits = defineEmits(['update:modelValue'])

interface Item {
  id: string
  name: string
}

let query = ref('')
let value = computed({
  get() {
    return props.modelValue
  },
  set(value) {
    emits('update:modelValue', value)
  }
})

const filteredItems = computed(() => {
  return query.value === ''
      ? props.items
      : props.items.filter((item) =>
          item.name
              .toLowerCase()
              .replace(/\s+/g, '')
              .includes(query.value.toLowerCase().replace(/\s+/g, ''))
      )
})

</script>