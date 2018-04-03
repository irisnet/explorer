import { checkStatus, parseJSON, API_ROOT } from "./utils"

function get(txhash) {
  return fetch(`${API_ROOT}/tx/${txhash}`)
    .then(checkStatus)
    .then(parseJSON)
    .then(tx => {
      return fetch(`${API_ROOT}/block/${tx.height}`)
        .then(checkStatus)
        .then(parseJSON)
        .then(block => {
          tx.time = block.block_meta.header.time
          console.log(tx)
          return tx
        })
    })
}

function getRecentCoinTx() {
  return fetch(`${API_ROOT}/txs/coin/1`)
    .then(checkStatus)
    .then(parseJSON)
}

function getRecentStakeTx() {
  return fetch(`${API_ROOT}/txs/stake/1`)
    .then(checkStatus)
    .then(parseJSON)
}

export default { get, getRecentCoinTx, getRecentStakeTx }
