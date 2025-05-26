import { Notify } from 'quasar'

// Restringir un <q-input> de vue3 a solo utilizar enteros
const lockIntegers = (evt: Event) => {
  if (!evt) {
    evt = window.event
  }

  const charCode = evt.which ? evt.which : evt.keyCode
  if (charCode > 31 && (charCode < 48 || charCode > 57)) {
    evt.preventDefault()
  }
  // else {
  //   return true
  // }
}

// Restringir un <q-input> de vue3 a poder utilizar decimales
const lockDecimals = (evt: Event) => {
  if (!evt) {
    evt = window.event
  }

  const charCode = evt.which ? evt.which : evt.keyCode
  if (
    charCode > 31 &&
    (charCode < 48 || charCode > 57) &&
    charCode !== 46 &&
    charCode !== 45
  ) {
    evt.preventDefault()
  }
}

// Mostrar notificación si se detecta un error en un request
const handleRequestError = (response) => {
  if (response) {
    Notify.create({
      message: response.data,
      position: 'bottom-left',
      color: 'red-6',
      icon: 'error',
      actions: [
        { icon: 'close', color: 'white', round: true, handler: () => { /* ... */ } }
      ]
    })
  }
}

export {
  lockIntegers,
  lockDecimals,
  handleRequestError
}
