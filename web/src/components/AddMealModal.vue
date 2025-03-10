<template>
  <q-dialog
    v-model="model"
    @hide="onHide"
    :maximized="$q.screen.xs"
    transition-show="slide-up"
    transition-hide="slide-down"
    @escape-key="onHide"
  >
    <q-card style="min-width:30%;max-height: 95vh; border-radius: 16px;">
      <q-bar class="bg-white">
        <q-space />
        <q-btn rounded dense flat icon="close" color="grey" v-close-popup>
        </q-btn>
      </q-bar>
      <div class="bg-white">
        <!-- <div class="row q-px-xs q-col-gutter-sm"> -->
        <div class="row q-mt-md">
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
    </q-card>
  </q-dialog>
</template>

<script setup lang="ts">
import { ref, reactive, computed, defineModel } from 'vue'

const emit = defineEmits(['hide'])

const fields = reactive({
  time: null,
  // meal: null,
  food: null
  // kcal: null
})

const foodsOptions = reactive([])
const foodsFilteredOptions = reactive([])

const props = defineProps({
  // show: {
  //   type: Boolean,
  //   default: false,
  //   required: true
  // }
})
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
</script>

<style>
</style>
