import { useEffect, useState } from "react";
import { client } from "./client";

export const useWeeklyTimeSeries = () => {
  const [state, setState] = useState("");
  const [data, setData] = useState([]);
  const [error, setError] = useState("");
  useEffect(() => {
    async function getInitialData() {
      try {
        const data = await client.getWeeklyTimeSeries();
        if (data) {
          setData(
            data.map((d) => ({
              ...d,
              createdAt: new Date(d.createdAt).toLocaleString(),
            }))
          );
          setState("success");
        }
      } catch (error) {
        setError(error);
        setState("error");
      }
    }
    setState("loading");
    getInitialData();
  }, []);
  return {
    state,
    data,
    error,
  };
};
