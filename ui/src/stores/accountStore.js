import { observable, action, runInAction } from "mobx"
import { AccountAPI } from "../api"

export class AccountStore {
  @observable isLoading = false
  @observable error = undefined
  @observable accountsRegistry = observable.map()

  getAccount(address) {
    return this.accountsRegistry.get(address)
  }

  @action
  async loadAccount(address, { acceptCached = false } = {}) {
    this.error = undefined
    if (acceptCached) {
      const account = this.getAccount(address)
      if (account) return account
    }
    this.isLoading = true
    try {
      const account = await AccountAPI.get(address)
      runInAction(() => {
        this.accountsRegistry.set(address, account)
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

export default new AccountStore()
