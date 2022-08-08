const BASE_URL = "api/v1";

export const client = {
  getLatestCo2Value: async () => {
    const response = await fetch(`${BASE_URL}/co2/latest`);
    if (response.status >= 200 || response.status < 300) {
      const body = await response.json();
      return body;
    }
    throw new Error(`${response.status}: ${response.statusText}`);
  },
  getDailyAverageCo2Value: async () => {
    const response = await fetch(`${BASE_URL}/co2/daily-average`);
    if (response.status >= 200 || response.status < 300) {
      const body = await response.json();
      return body;
    }
    throw new Error(`${response.status}: ${response.statusText}`);
  },
  getWeeklyTimeSeries: async () => {
    const response = await fetch(`${BASE_URL}/co2/duration?seconds=604800`);
    if (response.status >= 200 || response.status < 300) {
      const body = await response.json();
      return body;
    }
    throw new Error(`${response.status}: ${response.statusText}`);
  },
};
