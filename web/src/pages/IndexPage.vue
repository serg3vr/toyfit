<template>
  <q-page class="q-pa-md">
    <div class="row">
      <div class="col-xs-4 col-sm-8 q-pl-sm text-h4 self-center">
        <span class="text-blue-grey-9">{{ formattedDate }}</span>
      </div>
      <div class="col-xs-4 col-sm-2 text-right self-center">
        <q-btn outline round color="primary" icon="arrow_back" size="sm" @click="modifyDay(-1)" />
        &nbsp;
        <q-btn outline round color="primary" icon="arrow_forward" size="sm" @click="modifyDay(1)" />
      </div>
      <div class="col-xs-4 col-sm-2 text-right self-center">
        <q-btn icon="add" color="primary" size="12px" label="Add" @click="openAddMealFoodModal"></q-btn>
      </div>
    </div>

    <div class="row q-mt-sm">
      <!-- <div class="col-xs-12 q-pl-sm text-h4">
        <span class="text-blue-grey-9">2025-04-28</span>
      </div> -->
      <div class="col-xs-12 col-sm-4 q-pl-sm text-h4">
        <!-- <q-knob
          v-model="kcalProgress"
          size="90px"
          color="primary"
          show-value
          :step="0"
          flat
        /> -->

        <span class="text-blue-grey-9">{{ kcalProgress }}</span>
      </div>
      <div class="col-xs-12 col-sm-8">
        <div class="row q-col-gutter-lg">
          <div class="col-xs-12 col-sm-4" :class="{'q-px-md': !$q.screen.xs}">
            <q-linear-progress rounded size="8px" :value="carbsProgress" />
            Carbs: 88 / 120g
          </div>
          <div class="col-xs-12 col-sm-4" :class="{'q-px-md': !$q.screen.xs}">
            <q-linear-progress rounded size="8px" :value="carbsProgress" />
            Protein: 88 / 120g
          </div>
          <div class="col-xs-12 col-sm-4" :class="{'q-px-md': !$q.screen.xs}">
            <q-linear-progress rounded size="8px" :value="carbsProgress" />
            Fat: 88 / 120g
          </div>
        </div>
      </div>

    </div>
    <div class="row meals">
      <div class="col-xs-12 q-mb-sm" :class="{ 'meal-top-border': idx !== 0 }" v-for="(m, idx) in meals" :key="idx">
        <div class="row">
          <div class="col-xs-12">
            <div class="row">
              <div class="col-xs-1 q-pl-sm">
                <!-- <q-knob
                  v-model="kcalProgress"
                  size="42px"
                  color="primary"
                  show-value
                  :step="0"
                  flat
                /> -->
                {{ m.kcal }}
              </div>
              <div class="col-xs-10">
                <div class="row">
                  <div class="col-xs-12 text-weight-medium">{{ m.name }}</div>
                  <!-- <div class="col-xs-12 text-weight-light">{{ m.description }}</div> -->
                  <!-- <div class="col-xs-12">{{ m.currentKcal }} / {{ m.targetKcal }} kcal</div> -->
                </div>
              </div>
              <div class="col-xs-1">
                <!-- <q-btn icon="add" color="primary" round size="12px" @click="openAddMealFoodModal(m.id)"></q-btn> -->
              </div>
            </div>
          </div>
          <div class="col-xs-12">
            <div class="row" v-if="m?.foods">
              <div class="offset-xs-1 col-xs-10" v-for="(f, fIdx) in m.foods" :key="fIdx">
                <div class="row">
                  <div class="col-1">
                      {{ f.qty }}
                  </div>
                  <div class="col-10">
                    <span>{{ f.name }} {{ f.kcal }} kcal</span> <span class="text-grey">{{ f.description }}</span>
                  </div>
                  <div class="col-1">
                    <q-btn icon="delete" flat color="red-3" round size="12px" @click="deleteMealFood(f.id)"></q-btn>
                  </div>
                  <!-- <div class="col-12 text-weight-light">{{ f.description }}</div> -->
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <!-- :meal-type-id="mealTypeId" -->
      <AddMealModal
        v-model="addMealModal"
        @loadDailyMealFoods="getDailyMealFoods"
        @hide="addMealModal = false">
      </AddMealModal>
    </div>

  </q-page>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue';
// import type { Meal } from 'components/models';
import AddMealModal from 'components/AddMealModal.vue';
import { api } from 'boot/axios'
import moment from 'moment'
import { useQuasar } from 'quasar'
const $q = useQuasar()

const carbsProgress = ref(80 * 100 / 120)
// const currentDate = ref(moment().format("YYYY-MM-DDTHH:mm:ss.SSSZ"))
const currentDate = ref(moment().startOf('day'))
const addMealModal = ref(false)
// const mealTypeId = ref(0)

const kcalProgress = computed(() => {
  return meals.reduce((acc, val) => {
    return acc + val.kcal
  }, 0)
})

const meals = reactive([
{
  id: 1,
  name: 'Breakfast',
  foods: [],
  kcal: 0
},
{
    id: 2,
    name: 'Lunch',
    foods: [],
    kcal: 0
},
{
    id: 3,
    name: 'Dinner',
    foods: [],
    kcal: 0
},
{
    id: 4,
    name: 'Snacks',
    foods: [],
    kcal: 0
}
])

const formattedDate = computed(() => {
  return currentDate.value.format('YYYY-MM-DD')
})

const modifyDay = (day: number) => {
  if (day > 0) {
    currentDate.value = moment(currentDate.value).add(1, 'day')
    getDailyMealFoods()
  }
  if (day < 0) {
    currentDate.value = moment(currentDate.value).subtract(1, 'day')
    getDailyMealFoods()
  }
}

const openAddMealFoodModal = () => {
  // if (mtId > 0) {
  //   // mealTypeId.value = mtId
  // }
  addMealModal.value = true
}

const getDailyMealFoods = async () => {
  const date = moment(currentDate.value).format("YYYY-MM-DDTHH:mm:ss.SSSZ")
  const { data } = await api.get('meal-foods/daily', { params: { date: date }}).catch(error => error)
  if (data) {
    meals.forEach(elm => {
      elm.foods.length = 0
      elm.kcal = 0
    })

    data.forEach(elm => {
      if (elm.meal_type_id === 1) {
        meals[0].foods.push({ ...elm })
        meals[0].kcal += elm.kcal * elm.qty
      }
      if (elm.meal_type_id === 2) {
        meals[1].foods.push({ ...elm })
        meals[1].kcal += elm.kcal * elm.qty
      }
      if (elm.meal_type_id === 3) {
        meals[2].foods.push({ ...elm })
        meals[2].kcal += elm.kcal * elm.qty
      }
      if (elm.meal_type_id === 4) {
        meals[3].foods.push({ ...elm })
        meals[3].kcal += elm.kcal * elm.qty
      }
    })
  }
}

const deleteMealFood = (id: number) => {
  $q.dialog({
    title: 'Confirm',
    message: 'Delete this food?',
    cancel: true,
    persistent: true
  }).onOk(async () => {
    const { data } = await api.delete(`meal-foods/${id}`).catch(error => error)
    if (data) {
      getDailyMealFoods()
    }
  })
}

getDailyMealFoods()
</script>

<style lang="sass">
.meals
  border: 1px solid #e1e1e1
  border-radius: 4px
.meal-top-border
  border-top: 1px solid #e1e1e1
</style>
