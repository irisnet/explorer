import { checkStatus, parseJSON, API_ROOT } from "./utils"

function get(address) {
  return fetch(`${API_ROOT}/account/${address}`)
    .then(checkStatus)
    .then(parseJSON)
}

function getCoinTxs(address) {
  return fetch(`${API_ROOT}/account/${address}/tx/coin`)
    .then(checkStatus)
    .then(parseJSON)
}

export default { get, getCoinTxs }
