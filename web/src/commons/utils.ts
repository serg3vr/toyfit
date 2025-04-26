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

export {
  lockIntegers
}
