import { checkStatus, parseJSON, API_ROOT } from "./utils"

function get(address) {
  return fetch(`${API_ROOT}/account/${address}`)
    .then(checkStatus)
    .then(parseJSON)
}

function getCoinTxs(address) {
  return fetch(`${API_ROOT}/txs/coin?address=${address}&page=1&size=1000`)
    .then(checkStatus)
    .then(parseJSON)
}

export default { get, getCoinTxs }
