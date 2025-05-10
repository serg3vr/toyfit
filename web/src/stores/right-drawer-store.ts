import { defineStore, acceptHMRUpdate } from 'pinia';

export const useRightDrawerStore = defineStore('rightDrawer', {
  state: () => ({
    show: false,
  }),

  getters: {
    getShow: (state) => state.show,
  },

  actions: {
    toggle (show: boolean = false) {
      this.show = show;
    },
  },
});

if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useRightDrawerStore, import.meta.hot));
}
