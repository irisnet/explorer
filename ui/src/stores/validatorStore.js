import { observable, action, runInAction } from "mobx"
import { ValidatorAPI } from "../api"

export class ValidatorStore {
  @observable isLoading = false
  @observable error = undefined
  @observable height = undefined
  @observable validators = []

  @action
  async loadValidators(height, { acceptCached = false } = {}) {
    this.error = undefined
    this.isLoading = true
    this.height = height
    this.validators.clear()
    try {
      const ret = await ValidatorAPI.get(height)
      runInAction(() => {
        if (height === this.height) this.validators = ret.validators
      })
    } catch (error) {
      runInAction(() => {
        if (height === this.height) this.error = error.message
      })
    } finally {
      runInAction(() => {
        if (height === this.height) this.isLoading = false
      })
    }
  }
}

export default new ValidatorStore()
