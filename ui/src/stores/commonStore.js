import { observable, action, runInAction } from "mobx"
import { StatusAPI } from "../api"

class CommonStore {
  @observable appName = "Pangu Testnet Iris Block Explorer"
  @observable isLoading = false
  @observable status = {}
  @observable error = undefined

  @action
  async loadStatus() {
    this.isLoading = true
    try {
      const status = await StatusAPI.get()
      // after await, modifying state again, needs an actions:
      runInAction(() => {
        console.log(status,1111)
        this.status = status
      })
    } catch (error) {
      runInAction(() => {
        this.error = error.message
      })
    } finally {
      runInAction(() => {
        this.isLoading = false
      })
    }
  }
}

export default new CommonStore()
