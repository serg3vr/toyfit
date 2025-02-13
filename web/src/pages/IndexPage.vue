<template>
  <q-page class="q-pa-md">
    <div class="row">
      <div class="col-sm-12">
        Goal: 2000 kcal
      </div>
      <div class="col-sm-12">
        Carbs: 88/120g
      </div>
      <div class="col-sm-12">
        Protein: 88/120g
      </div>
      <div class="col-sm-12">
        Fat: 88/120g
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
          <div class="col-12">
            <q-list bordered separator>
              <q-item clickable v-ripple v-for="(r, idx2) in d.recipes" :key="idx2">
                <q-item-section>{{ r.description }}</q-item-section>
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
// import { ref } from 'vue';
// import type { Meal } from 'components/models';
// import ExampleComponent from 'components/ExampleComponent.vue';

const data = [
  {
    mealType: 'Breakfast',
    recipes: [
      {
        description: '2 huevos revueltos con espinacas y tomate'
      },
      {
        description: '1 rebanada de pan integral'
      }
    ]
  },
  {
    mealType: 'Lunch',
    recipes: [
      {
        description: 'Ensalada de pollo (150 g) con verduras (lechuga, pepino, zanahoria)'
      },
      {
        description: '1/2 taza de quinoa'
      }
    ]
  },
  {
    mealType: 'Dinner',
    recipes: [
      {
        description: 'Pescado a la plancha (150 g) con verduras al vapor (brócoli, zanahorias, calabacín)'
      }
    ]
  },
  {
    mealType: 'Snack',
    recipes: [
      {
        description: '1 puñado de almendras o nueces '
      }
    ]
  }
]
</script>
