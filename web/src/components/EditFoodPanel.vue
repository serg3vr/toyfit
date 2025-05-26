<template>
  <q-drawer
    side="right"
    v-model="model"
    show-if-above
    bordered
    overlay
    :width="$q.screen.width / ($q.screen.xs ? 1 : 3.33)"
    :breakpoint="500"
    behavior="mobile"
    no-swipe-open
    no-swipe-close
    no-swipe-backdrop
    :class="$q.dark.isActive ? 'bg-grey-9' : 'bg-grey-3'"
    @hide="close"
  >
    <div class="row" style="height: 90vh;">
      <div class="col-xs-12">
        <q-scroll-area class="fit">
          <div class="row">
            <div class="col-xs-12">
              <div class="row q-ma-sm q-col-gutter-sm">
                <div class="col-xs-10">
                  <span class="text-h6 text-weight-regular">Edit food</span>
                </div>
                <div class="col-xs-2 text-right">
                  <q-input
                    v-model="fields.id"
                    filled
                    label="Id"
                    :maxlength="100"
                    disable
                    input-class="text-right"
                  >
                  </q-input>
                </div>
                <div class="col-xs-12">
                  <!-- :error="v$.fields.name.$error"
                  :rules="r$.fields.name" -->
                  <q-input
                    v-model="fields.name"
                    filled
                    label="Name"
                    :maxlength="100"
                    :error="v$.fields.name.$error"
                    :rules="r$.fields.name"
                  >
                    <template v-slot:label>
                      <span>Name </span><span class="text-red">*</span>
                    </template>
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
                    :error="v$.fields.kcal.$error"
                    :rules="r$.fields.kcal"
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
    <div class="row q-ma-sm q-col-gutter-sm" style="height: 10vh;">
      <div class="col-xs-12 text-right">
        <q-btn
          icon="close"
          aria-label="Close"
          color="black"
          label="Close"
          flat
          @click="close"
        />
        &nbsp;
        <q-btn
          icon="edit"
          aria-label="Update"
          color="primary"
          label="Update"
          @click="createFood"
          :loading="loading"
        />
      </div>
    </div>
  </q-drawer>
</template>

<script setup lang="ts">
import { ref, reactive, computed, defineModel, watch } from 'vue'
import { api } from 'boot/axios'
import { useQuasar } from 'quasar'
// import moment from 'moment'
import { handleRequestError, lockDecimals } from 'src/commons/utils'
import { useVuelidate } from '@vuelidate/core'
import { required } from '@vuelidate/validators'

// const emit = defineEmits(['hide', 'loadDailyMealFoods'])
const $q = useQuasar()

const props = defineProps({
  fields: {
    type: Object,
    default: {},
    required: false
  }
})

const model = defineModel({ required: true })
const defaultFields = {
  id: null,
  name: null,
	description: null,
	kcal: null,
	carbs: null,
	proteins: null,
	fats: null,
	sodium: null
}
const fields = reactive({ ...defaultFields })

const loading = ref(false)

const validations = {
  fields: {
    name: { required },
    kcal: { required },
  }
}
const v$ = useVuelidate(validations, { fields })
const ecer = 'El campo es requerido.'
const r$ = computed(() => {
  return {
    fields: {
      name: [() => (!v$.value.fields.name.$error) || ecer],
      kcal: [() => (!v$.value.fields.kcal.$error) || ecer],
    }
  }
})

const close = () => {
  cleanFields()
  model.value = false
}

const cleanFields = () => {
  Object.assign(fields, defaultFields)
  v$.value.fields.$reset()
}

const createFood = async () => {
  const fieldsAreCorrect = await v$.value.fields.$validate()
  if (!fieldsAreCorrect) return

  loading.value = true

  const params = {
    ...fields,
    kcal: Number(fields.kcal)
  }
  const { data, response } = await api.put(`foods/${fields.id}`, params).catch(error => error)
  handleRequestError(response)

  if (data) {
    $q.notify({
      message: 'Food added.',
      position: 'bottom-left',
      color: 'primary',
      icon: 'check'
    })
    close()
  }
  loading.value = false
}

watch(() => props.fields, () => {
  Object.assign(fields, props.fields)
})
</script>

<style>
</style>
