import { observable, action, runInAction } from "mobx"
import { AccountAPI } from "../api"

export class CoinTxsStore {
  @observable isLoading = false
  @observable error = undefined
  @observable address = undefined
  @observable coinTxs = []

  @action
  async loadCoinTxs(address) {
    this.error = undefined
    this.isLoading = true
    this.address = address
    this.coinTxs.clear()
    try {
      const txs = await AccountAPI.getCoinTxs(address)
      runInAction(() => {
        if (address === this.address) this.coinTxs = txs
      })
    } catch (error) {
      runInAction(() => {
        if (address === this.address) this.error = error.message
      })
    } finally {
      runInAction(() => {
        if (address === this.address) this.isLoading = false
      })
    }
  }
}

export default new CoinTxsStore()
