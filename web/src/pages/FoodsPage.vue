<template>
  <q-page class="q-pa-md">
    <div class="row">
      <div class="col-12">
        <q-breadcrumbs>
          <q-breadcrumbs-el icon="home" to="/" />
          <q-breadcrumbs-el label="Foods" />
        </q-breadcrumbs>
      </div>
    </div>
    <div class="row">
      <div class="col-12 q-mt-sm">
        <q-table
          style="height: 400px"
          flat
          :rows="rows"
          :columns="columns"
          row-key="id"
        />
      </div>
    </div>
  </q-page>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
// import type { Meal } from 'components/models';
// import ExampleComponent from 'components/ExampleComponent.vue';
import { api } from 'boot/axios'

interface IFoods {
  id: number
  name: string
  kcal: number
  carbs: number
  proteins?: number
  fats?: number
  sodium?: number
}

const rows = reactive<IFoods[]>([])

const columns = [
  { name: 'name', required: true, label: 'Food', align: 'left', field: 'name', sortable: true },
  { name: 'kcal', align: 'center', label: 'Calories', field: 'kcal', sortable: true },
  { name: 'fats', label: 'Fat (g)', field: 'fats', sortable: true },
  { name: 'carbs', label: 'Carbs (g)', field: 'carbs' },
  { name: 'proteins', label: 'Protein (g)', field: 'proteins' },
  { name: 'sodium', label: 'Sodium (mg)', field: 'sodium' }
]

const loadFoods = () => {
  api.get('foods').then(response => {
    console.log(response)
    rows.length = 0
    rows.push(...response.data || [])
  })
}

loadFoods()
</script>
