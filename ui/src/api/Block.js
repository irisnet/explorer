import { checkStatus, parseJSON, API_ROOT } from "./utils"

function get(height) {
  return fetch(`${API_ROOT}/block/${height}`)
    .then(checkStatus)
    .then(parseJSON)
}

function getRecent() {
  return fetch(`${API_ROOT}/blocks/recent`)
    .then(checkStatus)
    .then(parseJSON)
}

export default { get, getRecent }
