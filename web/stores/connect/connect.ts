import { defineStore } from "pinia";

export const useConnectStore = defineStore('connect', {
  state: () => ({
    connectDialog: false,
  }),
  getters: {
  },
  actions: {
    mutateDialog(open: boolean) {
      this.connectDialog = open
    },
  }
})