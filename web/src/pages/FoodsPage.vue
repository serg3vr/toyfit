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
      <div class="col-12 q-mt-sm text-right">
        <q-btn icon="add" color="primary" size="12px" label="Add" @click="openDrawer"></q-btn>
      </div>
      <div class="col-12 q-mt-sm">
        <q-table
          style="height: 75vh"
          flat
          :rows="rows"
          :columns="columns"
          :pagination="pagination"
          row-key="id"
          :filter="filter"
        >
        <template v-slot:top-left>
          <q-input borderless dense debounce="300" v-model="filter" placeholder="Search">
            <template v-slot:append>
              <q-icon name="search" />
            </template>
          </q-input>
        </template>
          <template v-slot:body="props">
            <q-tr :props="props" @click="onRowClick(props.row)" style="cursor: pointer;">
              <q-td key="name" :props="props">
                {{ props.row.name }}
              </q-td>
              <q-td key="description" :props="props">
                {{ props.row.description }}
              </q-td>
              <q-td key="kcal" :props="props">
                {{ props.row.kcal }}
              </q-td>
              <q-td key="carbs" :props="props">
                {{ props.row.carbs }}
              </q-td>
              <q-td key="proteins" :props="props">{{ props.row.proteins }}</q-td>
              <q-td key="fats" :props="props">{{ props.row.fats }}</q-td>
              <q-td key="sodium" :props="props">{{ props.row.sodium }}</q-td>
            </q-tr>
          </template>
        </q-table>
      </div>
    </div>
    <!-- <FoodPanel
      v-model="showModal"
    /> -->
    <FoodPanel
      v-model="showFoodPanel"
      :fields="editFields"
      @reload="loadFoods"
    />
  </q-page>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue';
import { api } from 'boot/axios'
import FoodPanel from 'src/components/FoodPanel.vue';

// const showModal = ref(false)
const showFoodPanel = ref(false)
let editFields = {
  id: null,
  name: null,
	description: null,
	kcal: null,
	carbs: null,
	proteins: null,
	fats: null,
	sodium: null
}

interface IFoods {
  id: number
  name: string
  description?: string
  kcal: number
  carbs?: number
  proteins?: number
  fats?: number
  sodium?: number
}

const rows = reactive<IFoods[]>([])

const columns = [
  { name: 'name', required: true, label: 'Food', align: 'left', field: 'name', sortable: true },
  { name: 'description', required: true, label: 'Description', align: 'left', field: 'description', sortable: true },
  { name: 'kcal', label: 'Calories', field: 'kcal', align: 'right', sortable: true },
  { name: 'fats', label: 'Fat (g)', field: 'fats', align: 'right', sortable: true },
  { name: 'carbs', label: 'Carbs (g)', field: 'carbs', align: 'right', },
  { name: 'proteins', label: 'Protein (g)', field: 'proteins', align: 'right' },
  { name: 'sodium', label: 'Sodium (mg)', field: 'sodium', align: 'right' }
]

const pagination = {
  sortBy: 'desc',
  descending: false,
  page: 1,
  rowsPerPage: 25
  // rowsNumber: xx if getting data from a server
}

const filter = ref(null)

const loadFoods = () => {
  api.get('foods').then(response => {
    rows.length = 0
    rows.push(...response.data || [])
  }).catch(error => error)
}

const openDrawer = () => {
  showFoodPanel.value = true
}

const onRowClick = (row) => {
  showFoodPanel.value = true
  editFields = { ...row }
}

loadFoods()
</script>

<style scoped>
.email-cell {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.email-options {
  display: flex;
  gap: 4px;
}

.q-tr {
  transition: background-color 0.2s;
}

.q-tr:hover {
  background-color: #f0f0f0;
}
</style>
