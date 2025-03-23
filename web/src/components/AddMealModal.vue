<template>
  <q-dialog
    v-model="model"
    @hide="onHide"
    :maximized="$q.screen.xs"
    transition-show="slide-up"
    transition-hide="slide-down"
    @escape-key="onHide"
  >
    <q-card style="min-width:30%;max-height: 95vh; border-radius: 16px;" class="q-pa-md">
      <q-bar class="bg-white">
        <span>Add {{ mealTypeNames[mealTypeId] }}</span>
        <q-space />
        <q-btn rounded dense flat icon="close" color="grey" v-close-popup>
        </q-btn>
      </q-bar>
      <div class="bg-white">
        <!-- <div class="row q-px-xs q-col-gutter-sm"> -->
        <div class="row q-mt-md">
          <div class="col-12 q-mt-sm">
            <q-input
              v-if="isInput"
              v-model="foodName"
              filled
              label="Type a food"
            >
              <template v-slot:append>
                <!-- <q-icon name="close" @click.stop.prevent="model = ''" class="cursor-pointer" /> -->
                <q-btn round dense flat icon="list" @click.stop.prevent @click="toggleInput(false)"/>
              </template>
            </q-input>
            <q-select
              v-else
              filled
              v-model="fields.foodId"
              :options="foodsFilteredOptions"
              label="Select a food"
              @filter="filterFn"
              use-input
              emit-value
              map-options
            >
              <template v-slot:option="{ itemProps, opt, selected, toggleOption }">
                <q-item v-bind="itemProps">
                  <q-item-section>
                    <!-- <q-item-label v-html="opt.label" /> -->
                    {{ opt.label }}
                  </q-item-section>
                  <q-item-section side>
                    <!-- <q-toggle :model-value="selected" @update:model-value="toggleOption(opt)" /> -->
                      {{ opt.kcal }} kcal
                  </q-item-section>
                </q-item>
              </template>
              <!-- <template v-slot:no-option>
                <q-item>
                  <q-item-section class="text-grey">
                    No results
                  </q-item-section>
                </q-item>
              </template> -->
              <template v-slot:append>
                <!-- <q-icon name="close" @click.stop.prevent="model = ''" class="cursor-pointer" /> -->
                <q-btn round dense flat icon="edit" @click.stop.prevent @click="toggleInput(true)"/>
              </template>
            </q-select>
          </div>
          <div class="col-12 q-mt-lg text-right">
            <q-btn
              class="full-width"
              icon="add"
              aria-label="Add"
              color="primary"
              full-width
              @click="addMealFood"
              label="Add"
            />
          </div>
        </div>
      </div>
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import { ref, computed, reactive, defineModel } from 'vue'
import { api } from 'boot/axios'
import { useQuasar } from 'quasar'

const emit = defineEmits(['hide'])
const $q = useQuasar()

const props = defineProps({
  mealTypeId: {
    type: Number,
    // default: false,
    required: true
  }
})

// const mealTypeId = ref(props.mealTypeId)

const mealTypeId = computed(() => {
  return props.mealTypeId
})

const mealTypeNames = ['', 'Breakfast', 'Lunch', 'Dinner', 'Snack']

const fields = reactive({
  time: null,
  // meal: null,
  foodId: null
  // kcal: null
})

const foodsOptions = reactive([])
const foodsFilteredOptions = reactive([])
const isInput = ref(false)
const foodName = ref('')

// const show = ref(props.show)

const model = defineModel({ required: true })

const filterFn = (val, update, abort) => {
  update(() => {
    const needle = val.toLowerCase()
    foodsFilteredOptions.length = 0
    foodsFilteredOptions.push(...foodsOptions.filter(v => v.label.toLowerCase().indexOf(needle) > -1))
  })
}

const onHide = () => {
  emit('hide', true)
}

// interface IFoods {
//   id: number
//   name: string
//   kcal: number
//   carbs: number
//   proteins?: number
//   fats?: number
//   sodium?: number
// }

// const rows = reactive<IFoods[]>([])

const toggleInput = (toggle: boolean) => {
  isInput.value = toggle
  if (isInput.value) {
    foodName.value = ''
  }
}

const loadFoods = () => {
  api.get('foods').then(response => {
    const opts = (response.data || []).reduce((opt, val) => {
      opt.push({ label: val.name, value: val.id })
      return opt
    }, [])
    foodsOptions.length = 0
    foodsOptions.push(...opts)
    foodsFilteredOptions.length = 0
    foodsFilteredOptions.push(...opts)
  }).catch(error => error)
}

const addMealFood = async () => {
  const params = {
    meal_type_id: mealTypeId.value,
    date: "2025-03-19T00:15:43.890-06:00",
    food_id : fields.foodId,
    custom_food_id : null,
    qty : 1
  }
  const { data } = await api.post('meal-foods', params).catch(error => error)
  if (data) {
    $q.notify({
      message: 'Meal food added.',
      position: 'top-right',
      color: 'primary',
      icon: 'check'
    })
  }
}

loadFoods()
</script>

<style>
</style>
