import { observable, action, runInAction } from "mobx"
import { BlockAPI } from "../api"

export class BlockStore {
  @observable isLoading = false
  @observable error = undefined
  @observable blocksRegistry = observable.map()
  @observable recent = []

  getBlock(height) {
    return this.blocksRegistry.get(height)
  }
  @action
  async loadBlock(height, { acceptCached = false } = {}) {
    this.error = undefined
    if (acceptCached) {
      const block = this.getBlock(height)
      if (block) return Promise.resolve(block)
    }
    this.isLoading = true
    try {
      const block = await BlockAPI.get(height)
      runInAction(() => {
        this.blocksRegistry.set(height, block)
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
  async loadRecent(length = 20) {
    this.error = undefined
    this.isLoading = false
    this.recent = []
    try {
      const blocks = await BlockAPI.getRecent()
      runInAction(() => {
        this.recent = blocks.block_metas
          .map(b => {
            return {
              height: b.header.height,
              time: b.header.time,
              num_txs: b.header.num_txs
            }
          })
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

export default new BlockStore()
