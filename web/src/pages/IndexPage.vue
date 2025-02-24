<template>
  <q-page class="q-pa-md">
    <div class="row">
      <div class="col-xs-12 col-sm-6">
        <div class="row">
          <div class="col-sm-12">
            Goal: {{ kcalGoal }} / 2040 kcal
          </div>
          <div class="col-sm-12">
            Carbs: 88 / 120g
          </div>
          <div class="col-sm-12">
            Protein: 88 / 120g
          </div>
          <div class="col-sm-12">
            Fat: 88 / 120g
          </div>
        </div>
      </div>
      <div class="col-xs-12 col-sm-6">
        <div class="row">
          <div class="col-12">
            <div class="q-gutter-sm">
              <q-radio v-model="fields.time" val="Breakfast" label="Breakfast" />
              <q-radio v-model="fields.time" val="Lunch" label="Lunch" />
              <q-radio v-model="fields.time" val="Dinner" label="Dinner" />
              <q-radio v-model="fields.time" val="Snack" label="Snack" />
            </div>
          </div>
          <div class="col-12 q-mt-sm">
            <q-select
              filled
              v-model="fields.food"
              :options="foodsFilteredOptions"
              label="Food"
              @filter="filterFn"
              use-input
              emit-value
              map-options
            >
            <template v-slot:option="{ itemProps, opt, selected, toggleOption }">
              <q-item v-bind="itemProps">
                <q-item-section>
                  <q-item-label v-html="opt.label" />
                </q-item-section>
                <q-item-section side>
                  <!-- <q-toggle :model-value="selected" @update:model-value="toggleOption(opt)" /> -->
                   {{ opt.kcal }} kcal
                </q-item-section>
              </q-item>
            </template>
              <template v-slot:no-option>
                <q-item>
                  <q-item-section class="text-grey">
                    No results
                  </q-item-section>
                </q-item>
              </template>
            </q-select>
          </div>
          <!-- <div class="col-12 q-mt-sm">
            <div class="row q-col-gutter-sm">
              <div class="col-10">
                <q-input filled v-model="fields.meal" placeholder="Add meal" />
              </div>
              <div class="col-2">
                <q-input filled v-model="fields.kcal" placeholder="Kcal" />
              </div>
            </div>
          </div> -->
          <div class="col-12 q-mt-sm text-right">
            <q-btn
              dense
              icon="add"
              aria-label="Add"
              color="grey"
              full-width
              @click="'toggleLeftDrawer'"
            />
          </div>
          <div class="col-12 q-mt-lg text-right">
            <q-btn
              class="full-width"
              icon="add"
              aria-label="Add"
              color="primary"
              full-width
              @click="doTheTest"
              label="Add"
            />
          </div>
        </div>

      </div>
    </div>
    <!-- <div class="row q-mt-md">
      <div class="col-12">
        Breakfast: 2 huevos revueltos con espinacas y tomate, 1 rebanada de pan integral
      </div>
      <div class="col-12">
        Lunch: Ensalada de pollo (150 g) con verduras (lechuga, pepino, zanahoria) y 1/2 taza de quinoa
      </div>
      <div class="col-12">
        Dinner: Pescado a la plancha (150 g) con verduras al vapor (brócoli, zanahorias, calabacín)
      </div>
      <div class="col-12">
        Snack: 1 puñado de almendras o nueces
      </div>
    </div> -->
    <div class="row q-mt-md">
      <div class="col-12" v-for="(d, idx) in data" :key="idx">
        <div class="row">
          <div class="col-12">
            {{ d.mealType }}:
          </div>
          <div class="col-6">
            <q-list bordered separator>
              <q-item clickable v-ripple v-for="(r, idx2) in d.recipes" :key="idx2">
                <q-item-section>{{ r.description }}</q-item-section>
                <q-item-section side top>
                  <q-item-label caption>{{ r.kcal }} kcal</q-item-label>
                  <!-- <q-icon name="star" color="yellow" /> -->
                </q-item-section>
              </q-item>

              <!-- <q-item clickable v-ripple>
                <q-item-section>
                  <q-item-label>Item with caption</q-item-label>
                  <q-item-label caption>Caption</q-item-label>
                </q-item-section>
              </q-item>

              <q-item clickable v-ripple>
                <q-item-section>
                  <q-item-label overline>OVERLINE</q-item-label>
                  <q-item-label>Item with overline</q-item-label>
                </q-item-section>
              </q-item> -->
            </q-list>
          </div>
        </div>
      </div>
    </div>
  </q-page>
  <!-- <q-page class="row items-center justify-evenly">
    <example-component
      title="Example component"
      active
      :todos="todos"
      :meta="meta"
    ></example-component>
  </q-page> -->
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
// import type { Meal } from 'components/models';
// import ExampleComponent from 'components/ExampleComponent.vue';
import { api } from 'boot/axios'

const kcalGoal = ref(300)

const fields = reactive({
  time: null,
  // meal: null,
  food: null
  // kcal: null
})

const foodsOptions = reactive([])
const foodsFilteredOptions = reactive([])

const data = reactive([
{
    mealType: 'Breakfast',
    recipes: [
      {
        description: 'Pescado a la plancha (150 g) con verduras al vapor (brócoli, zanahorias, calabacín)',
        kcal: 100
      },
      {
        description: 'Pan tostado con mermelada',
        kcal: 200
      }
    ]
  }
])

const doTheTest = () => {
  if (fields.time && fields.food) {
    const founded = foodsOptions.find(v => v.value === fields.food)

    const foundedData = data.find(v => v.mealType === fields.time)
    if (foundedData) {
      foundedData.recipes.push({ description: founded.label, kcal: founded.kcal })
    } else {
      data.push({
        mealType: fields.time,
        recipes: [{ description: founded.label, kcal: founded.kcal }]
      })
    }

    kcalGoal.value += parseInt(founded.kcal)

    fields.time = null
    fields.food = null
  }
}

const loadFoods = () => {
  api.get('foods').then(({ data }) => {
    const fo = (data || []).reduce((opt, val) => {
      opt.push({ label: val.name, value: val.id, kcal: val.kcal })
      return opt
    }, [])
    foodsOptions.length = 0
    foodsOptions.push(...fo)
    foodsFilteredOptions.length = 0
    foodsFilteredOptions.push(...fo)
  }).catch(error => error)
}

loadFoods()

const filterFn = (val, update, abort) => {
  update(() => {
    const needle = val.toLowerCase()
    foodsFilteredOptions.length = 0
    foodsFilteredOptions.push(...foodsOptions.filter(v => v.label.toLowerCase().indexOf(needle) > -1))
  })
}
</script>
