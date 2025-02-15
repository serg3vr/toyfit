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
            <!-- <q-btn
              dense
              icon="add"
              aria-label="Breakfast"
              color="primary"
              @click="'toggleLeftDrawer'"
              label="Breakfast"
            />
            <q-btn
              dense
              icon="add"
              aria-label="Lunch"
              color="primary"
              @click="'toggleLeftDrawer'"
              label="Lunch"
            />
            <q-btn
              dense
              icon="add"
              aria-label="Dinner"
              color="primary"
              @click="'toggleLeftDrawer'"
              label="Dinner"
            />
            <q-btn
              dense
              icon="add"
              aria-label="Snack"
              color="primary"
              @click="'toggleLeftDrawer'"
              label="Snack"
            /> -->
            <div class="q-gutter-sm">
              <q-radio v-model="fields.time" val="Breakfast" label="Breakfast" />
              <q-radio v-model="fields.time" val="Lunch" label="Lunch" />
              <q-radio v-model="fields.time" val="Dinner" label="Dinner" />
              <q-radio v-model="fields.time" val="Snack" label="Snack" />
            </div>
          </div>
          <div class="col-12 q-mt-sm">
            <div class="row q-col-gutter-sm">
              <div class="col-10">
                <q-input filled v-model="fields.meal" placeholder="Add meal" />
              </div>
              <div class="col-2">
                <q-input filled v-model="fields.kcal" placeholder="Kcal" />
              </div>
            </div>
          </div>
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

const kcalGoal = ref(0)

const fields = reactive({
  time: null,
  meal: null,
  kcal: null
})

// let data = [
//   {
//     mealType: 'Breakfast',
//     recipes: [
//       {
//         description: '2 huevos revueltos con espinacas y tomate'
//       },
//       {
//         description: '1 rebanada de pan integral'
//       }
//     ]
//   },
//   {
//     mealType: 'Lunch',
//     recipes: [
//       {
//         description: 'Ensalada de pollo (150 g) con verduras (lechuga, pepino, zanahoria)'
//       },
//       {
//         description: '1/2 taza de quinoa'
//       }
//     ]
//   },
//   {
//     mealType: 'Dinner',
//     recipes: [
//       {
//         description: 'Pescado a la plancha (150 g) con verduras al vapor (brócoli, zanahorias, calabacín)'
//       }
//     ]
//   },
//   {
//     mealType: 'Snack',
//     recipes: [
//       {
//         description: '1 puñado de almendras o nueces '
//       }
//     ]
//   }
// ]

const data = reactive([
{
    mealType: 'Dinner',
    recipes: [
      {
        description: 'Pescado a la plancha (150 g) con verduras al vapor (brócoli, zanahorias, calabacín)',
        kcal: 100
      }
    ]
  }
])

const doTheTest = () => {
  if (fields.time && fields.meal && fields.kcal) {
    data.push({
      mealType: fields.time,
      recipes: [
      {
        description: fields.meal,
        kcal: fields.kcal
      }
    ]
    })

    kcalGoal.value += parseInt(fields.kcal)

    fields.time = null
    fields.meal = null
    fields.kcal = null
  }
}
</script>
