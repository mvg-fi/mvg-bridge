import { defineStore } from "pinia";

export const useSettingsStore = defineStore('settings', {
  state: () => ({
    useCustomAddress: getFromLocalStorage('useCustomAddress'),
    onlyUseDex: getFromLocalStorage('onlyUseDex'),
    enhancePrivacy: getFromLocalStorage('enhancePrivacy'),
  }),
  actions: {
    setUseCustomAddress(v: boolean) {
      this.useCustomAddress = v
    },
    setOnlyUseCex(v: boolean) {
      this.onlyUseDex = v
    },
    setEnhancePrivacy(v: boolean) {
      this.enhancePrivacy = v
    },
  }
})

const getFromLocalStorage = (key: string) => {
  const v = localStorage.getItem(key)
  return Boolean(v) || false
}
const changeLocalStorage = (key: string, value: boolean) => {
  const k = localStorage.getItem(key)
  if (!k) return
  localStorage.setItem(key, String(value))
}