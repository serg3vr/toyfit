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
        <span>Add <template v-if="isInput"> New</template> {{ mealTypeNames[mealTypeId] }}</span>
        <q-space />
        <q-btn rounded dense flat icon="close" color="grey" v-close-popup>
        </q-btn>
      </q-bar>
      <div class="bg-white q-mt-sm">
        <!-- <div class="row q-px-xs q-col-gutter-sm"> -->
        <div class="row q-mt-xs">
          <div class="col-xs-12 row justify-between q-gutter-sm">
            <q-radio v-model="fields.mealTypeId" val="1" label="Breakfast" />
            <q-radio v-model="fields.mealTypeId" val="2" label="Lunch" />
            <q-radio v-model="fields.mealTypeId" val="3" label="Dinner" />
            <q-radio v-model="fields.mealTypeId" val="4" label="Snacks" />
          </div>
          <div class="col-xs-12 q-mt-xs">
            <q-input filled v-model="fields.date" mask="date" :rules="['date']">
              <template v-slot:append>
                <q-icon name="event" class="cursor-pointer">
                  <q-popup-proxy cover transition-show="scale" transition-hide="scale">
                    <q-date v-model="fields.date">
                      <div class="row items-center justify-end">
                        <q-btn v-close-popup label="Close" color="primary" flat />
                      </div>
                    </q-date>
                  </q-popup-proxy>
                </q-icon>
              </template>
            </q-input>
          </div>
          <div class="col-xs-12">
            <q-input
              v-if="isInput"
              v-model="fields.foodName"
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
          <div class="col-xs-12 col-sm-10 q-mt-md q-pr-md" v-if="isInput">
            <q-input
              v-model="fields.foodDescription"
              filled
              label="Description"
            >
            </q-input>
          </div>
          <div class="col-xs-12 col-sm-2 q-mt-md" v-if="isInput">
            <q-input
              v-model="fields.foodKcal"
              filled
              label="Kcal"
              :maxlength="4"
              @keypress="lockIntegers"
            >
            </q-input>
          </div>
          <div class="col-xs-12 q-mt-lg text-right">
            <q-btn
              icon="add"
              aria-label="Add"
              color="primary"
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
import { ref, reactive, defineModel } from 'vue'
import { api } from 'boot/axios'
import { useQuasar } from 'quasar'
import moment from 'moment'
import { lockIntegers } from 'src/commons/utils'

const emit = defineEmits(['hide', 'loadDailyMealFoods'])
const $q = useQuasar()

// const props = defineProps({
//   mealTypeId: {
//     type: Number,
//     // default: false,
//     required: false
//   }
// })

// const mealTypeId = ref(1)

// const mealTypeId = computed(() => {
//   return props.mealTypeId
// })



const mealTypeNames = ['', 'Breakfast', 'Lunch', 'Dinner', 'Snack']

const fields = reactive({
  mealTypeId: '1',
  date: moment().format('YYYY-MM-DD'),
  foodId: null,
  foodName: '',
  foodDescription: '',
  foodKcal: ''
})

const foodsOptions = reactive([])
const foodsFilteredOptions = reactive([])
const isInput = ref(false)

const model = defineModel({ required: true })

const filterFn = (val, update, abort) => {
  update(() => {
    const needle = val.toLowerCase()
    foodsFilteredOptions.length = 0
    foodsFilteredOptions.push(...foodsOptions.filter(v => v.label.toLowerCase().indexOf(needle) > -1))
  })
}

const onHide = () => {
  cleanFields()
  emit('hide', true)
}

const toggleInput = (toggle: boolean) => {
  isInput.value = toggle
  cleanFields()
}

const cleanFields = () => {
  fields.mealTypeId = '1'
  fields.date = moment().format('YYYY-MM-DD')
  fields.foodId = null
  fields.foodName = ''
  fields.foodDescription = ''
  fields.foodKcal = ''
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
    meal_type_id: fields.mealTypeId * 1,
    date: moment(fields.date, 'YYYY/MM/DD').format('YYYY-MM-DDTHH:mm:ss.SSSZ'),
    food_id : fields.foodId,
    custom_food_id : null,
    qty: 1,
    food_name: fields.foodName,
    food_description: fields.foodDescription,
    food_kcal: fields.foodKcal ? 1 * fields.foodKcal: 0
  }
  const { data } = await api.post('meal-foods', params).catch(error => error)
  if (data) {
    emit('loadDailyMealFoods')
    onHide()
    $q.notify({
      message: 'Meal food added.',
      position: 'bottom-left',
      color: 'primary',
      icon: 'check'
    })
  }
}

loadFoods()
</script>

<style>
</style>
