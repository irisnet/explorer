import { checkStatus, parseJSON, API_ROOT } from "./utils"

function get(height) {
  return fetch(`${API_ROOT}/validators/${height}`)
    .then(checkStatus)
    .then(parseJSON)
}

export default { get }
