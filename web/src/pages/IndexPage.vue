<template>
  <q-page class="q-pa-md">
    <div class="row">
      <div class="col-xs-12 col-sm-4">
        <q-knob
          v-model="kcalProgress"
          size="90px"
          color="primary"
          show-value
          :step="0"
          flat
        />
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
              <div class="col-xs-1">
                <q-knob
                  v-model="kcalProgress"
                  size="42px"
                  color="primary"
                  show-value
                  :step="0"
                  flat
                />
              </div>
              <div class="col-xs-10">
                <div class="row">
                  <div class="col-xs-12 text-weight-medium">{{ m.name }}</div>
                  <!-- <div class="col-xs-12">{{ m.currentKcal }} / {{ m.targetKcal }} kcal</div> -->
                </div>
              </div>
              <div class="col-xs-1">
                <q-btn icon="add" color="primary" round size="12px" @click="openAddMealFoodModal(m.id)"></q-btn>
              </div>
            </div>
          </div>
          <div class="col-xs-12">
            <div class="row" v-if="m?.foods">
              <div class="offset-xs-1 col-xs-10" v-for="(f, fIdx) in m.foods" :key="fIdx">
                {{ f.name }}
                <q-btn icon="delete" flat color="red-3" round size="12px" @click="deleteMealFood(f.id)"></q-btn>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <AddMealModal
        v-model="addMealModal"
        :meal-type-id="mealTypeId"
        @loadDailyMealFoods="getDailyMealFoods"
        @hide="addMealModal = false"></AddMealModal>
    </div>

  </q-page>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue';
// import type { Meal } from 'components/models';
import AddMealModal from 'components/AddMealModal.vue';
import { api } from 'boot/axios'
const carbsProgress = ref(80 * 100 / 120)

const addMealModal = ref(false)
const mealTypeId = ref(0)

const kcalProgress = computed(() => {
  return 0
})

const meals = reactive([
{
  id: 1,
  name: 'Breakfast',
  foods: []
  // foods: [
  //   {
  //     name: 'Huevo frito'
  //   },
  //   {
  //     name: 'Café'
  //   }
  // ]
},
{
    id: 2,
    name: 'Lunch',
    foods: []
},
{
    id: 3,
    name: 'Dinner',
    foods: []
},
{
    id: 4,
    name: 'Snacks',
    foods: []
    // foods: [
    //   {
    //     name: 'Doraditas'
    //   }
    // ]
}
])

// const breakfast = reactive({
//   id: 1,
//   name: 'Breakfast',
//   foods: []
// })

const openAddMealFoodModal = (mtId: number) => {
  if (mtId > 0) {
    mealTypeId.value = mtId
    addMealModal.value = true
  }
}

const getDailyMealFoods = async () => {
  const { data } = await api.get('meal-foods/daily').catch(error => error)
  if (data) {
    meals[0].foods.length = 0
    meals[1].foods.length = 0
    meals[2].foods.length = 0
    meals[3].foods.length = 0

    data.forEach(elm => {
      if (elm.meal_type_id === 1) {
        meals[0].foods.push({ ...elm })
      }
      if (elm.meal_type_id === 2) {
        meals[1].foods.push({ ...elm })
      }
      if (elm.meal_type_id === 3) {
        meals[2].foods.push({ ...elm })
      }
      if (elm.meal_type_id === 4) {
        meals[3].foods.push({ ...elm })
      }
    })
  }
}

const deleteMealFood = async (id: number) => {
  const { data } = await api.delete(`meal-foods/${id}`).catch(error => error)
  if (data) {
    getDailyMealFoods()
    // $q.notify({
    //   message: 'Meal food added.',
    //   position: 'bottom-left',
    //   color: 'primary',
    //   icon: 'check'
    // })
  }
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
