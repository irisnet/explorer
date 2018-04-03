import { observable, action, runInAction } from "mobx"
import { TxAPI } from "../api"

export class TxStore {
  @observable isLoading = false
  @observable error = undefined
  @observable txsRegistry = observable.map()
  @observable recentCoinTx = []
  @observable recentStakeTx = []

  getTx(txhash) {
    return this.txsRegistry.get(txhash)
  }

  @action
  async loadTx(txhash, { acceptCached = false } = {}) {
    this.error = undefined
    if (acceptCached) {
      const tx = this.getTx(txhash)
      if (tx) return Promise.resolve(tx)
    }
    this.isLoading = true
    try {
      const tx = await TxAPI.get(txhash)
      runInAction(() => {
        this.txsRegistry.set(txhash, tx)
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
  @action
  async loadRecentCoinTx(length = 20) {
    this.error = undefined
    this.isLoading = false
    this.recentCoinTx = []
    try {
      const txs = await TxAPI.getRecentCoinTx()
      runInAction(() => {
        this.recentCoinTx = txs
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
  @action
  async loadRecentStakeTx(length = 20) {
    this.error = undefined
    this.isLoading = false
    this.recentStakeTx = []
    try {
      const txs = await TxAPI.getRecentStakeTx()

      runInAction(() => {
        this.recentStakeTx = txs
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

export default new TxStore()
