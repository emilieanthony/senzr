const BASE_URL = "api/v1"

export const client = {
  getLatestCo2Value: async () => {
    const response = await fetch(`${BASE_URL}/co2/latest`)
    if (response.status >= 200 || response.status < 300) {
      const body = await response.json()
      return body
    }
    throw new Error(`${response.status}: ${response.statusText}`)
  },
  getLatestTemperatureValue: async () => {
    const response = await fetch(`${BASE_URL}/temperature/latest`)
    if (response.status >= 200 || response.status < 300) {
      const body = await response.json()
      return body
    }
    throw new Error(`${response.status}: ${response.statusText}`)
  },
  getLatestHumidityValue: async () => {
    const response = await fetch(`${BASE_URL}/humidity/latest`)
    if (response.status >= 200 || response.status < 300) {
      const body = await response.json()
      return body
    }
    throw new Error(`${response.status}: ${response.statusText}`)
  }
}
