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
    <div class="row">
      <div class="col-xs-12 q-mb-sm" v-for="(m, idx) in meals" :key="idx">
        <div class="row">
          <div class="col-xs-12 debug">
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
                  <div class="col-xs-12">{{ m.currentKcal }} / {{ m.targetKcal }} kcal</div>
                </div>
              </div>
              <div class="col-xs-1">
                <q-btn icon="add" color="primary" round size="12px"></q-btn>
              </div>
            </div>
          </div>
          <div class="col-xs-12 debug">
            <div class="row" v-if="m?.foods">
              <div class="offset-xs-1 col-xs-11" v-for="(f, fIdx) in m.foods" :key="fIdx">{{ f.name }}</div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <AddMealModal v-model="addMealModal"></AddMealModal>
    </div>

  </q-page>
</template>

<script setup lang="ts">
import { ref, reactive, computed } from 'vue';
// import type { Meal } from 'components/models';
import AddMealModal from 'components/AddMealModal.vue';
// import { api } from 'boot/axios'
const carbsProgress = ref(80 * 100 / 120)

const addMealModal = ref(true)

const kcalProgress = computed(() => {
  return 0
})

const meals = reactive([
{
    name: 'Breakfast',
    targetKcal: 900,
    currentKcal: 100,
    foodSummarize: 'Huevo, tortilla',
    foods: [
      {
        name: 'Huevo frito'
      },
      {
        name: 'Café'
      }
    ]
},
{
    name: 'Lunch',
    targetKcal: 900,
    currentKcal: 100,
},
{
    name: 'Dinner',
    targetKcal: 900,
    currentKcal: 100,
},
{
    name: 'Snacks',
    targetKcal: 900,
    currentKcal: 100,
}
])
</script>
