<template>
  <q-drawer
    side="right"
    v-model="model"
    show-if-above
    bordered
    overlay
    :width="$q.screen.width / 3 - 1"
    :breakpoint="500"
    behavior="mobile"
    no-swipe-open
    no-swipe-close
    no-swipe-backdrop
    :class="$q.dark.isActive ? 'bg-grey-9' : 'bg-grey-3'"
    @hide="model = false"
  >
    <div class="row" style="height: 90vh;">
      <div class="col-xs-12">
        <q-scroll-area class="fit">
          <div class="row">
            <div class="col-xs-12">
              <div class="row q-ma-sm q-col-gutter-sm">
                <div class="col-xs-12">
                  <span class="text-h6 text-weight-regular">Add food</span>
                </div>
                <div class="col-xs-12">
                  <q-input
                    v-model="fields.name"
                    filled
                    label="Name"
                    :maxlength="100"
                  >
                  </q-input>
                </div>
                <div class="col-xs-12">
                  <q-input
                    v-model="fields.description"
                    filled
                    label="Description"
                    :maxlength="200"
                  >
                  </q-input>
                </div>
                <div class="col-xs-12">
                  <q-input
                    v-model="fields.kcal"
                    filled
                    label="Kcal"
                    :maxlength="6"
                    @keypress="lockDecimals"
                  >
                  </q-input>
                </div>
                <div class="col-xs-12">
                  <q-input
                    v-model="fields.carbs"
                    filled
                    label="Carbs"
                    :maxlength="6"
                    @keypress="lockDecimals"
                  >
                  </q-input>
                </div>
                <div class="col-xs-12">
                  <q-input
                    v-model="fields.proteins"
                    filled
                    label="Proteins"
                    :maxlength="6"
                    @keypress="lockDecimals"
                  >
                  </q-input>
                </div>
                <div class="col-xs-12">
                  <q-input
                    v-model="fields.fats"
                    filled
                    label="Fats"
                    :maxlength="6"
                    @keypress="lockDecimals"
                  >
                  </q-input>
                </div>
                <div class="col-xs-12">
                  <q-input
                    v-model="fields.sodium"
                    filled
                    label="Sodium"
                    :maxlength="6"
                    @keypress="lockDecimals"
                  >
                  </q-input>
                </div>
              </div>
            </div>
          </div>
        </q-scroll-area>
      </div>
    </div>
    <div class="row" style="height: 10vh;">
      <div class="col-xs-12 text-right">
        <q-btn
          icon="close"
          aria-label="Close"
          color="black"
          label="Close"
          flat
          @click="model = false"
        />
        &nbsp;
        <q-btn
          icon="add"
          aria-label="Add"
          color="primary"
          label="Add"
          @click="createFood"
          :loading="loading"
        />
      </div>
    </div>
  </q-drawer>
</template>

<script setup lang="ts">
import { ref, reactive, defineModel } from 'vue'
import { api } from 'boot/axios'
import { useQuasar } from 'quasar'
// import moment from 'moment'
import { lockDecimals } from 'src/commons/utils'

// const emit = defineEmits(['hide', 'loadDailyMealFoods'])
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

const model = defineModel({ required: true })

const fields = reactive({
  name: null,
	description: null,
	kcal: null,
	carbs: null,
	proteins: null,
	fats: null,
	sodium: null
})

const loading = ref(false)

const createFood = () => {
  loading.value = true
  const params = { ...fields }
  api.post('foods', params).then(({ data }) => {
    if (data) {
      $q.notify({
        message: 'Food added.',
        position: 'bottom-left',
        color: 'primary',
        icon: 'check'
      })
    }
  }).catch(error => error)
  loading.value = false
}
</script>

<style>
</style>
