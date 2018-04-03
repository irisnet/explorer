const API_ROOT = process.env.REACT_APP_API_ROOT

function checkStatus(response) {
  if (response.ok || response.status === 400) {
    return response.json()
  } else {
    throw new Error("Something went wrong ...")
  }
}

function parseJSON(json) {
  if (json.code != null) {
    throw new Error(json.error)
  } else {
    return json.data || json
  }
}

export { API_ROOT, checkStatus, parseJSON }
