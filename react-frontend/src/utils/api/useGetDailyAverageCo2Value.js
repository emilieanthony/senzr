import { useEffect, useState } from "react";
import { client } from "./client";

export const useGetDailyAverageCo2Value = () => {
  const [state, setState] = useState(""); // 'loading', 'error' or 'success'
  const [data, setData] = useState({ value: 0, createdAt: "" });
  const [error, setError] = useState("");
  useEffect(() => {
    async function getInitialData() {
      try {
        const data = await client.getDailyAverageCo2Value();
        if (data) {
          setData(data);
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
