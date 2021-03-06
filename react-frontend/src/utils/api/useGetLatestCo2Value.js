import { useEffect, useState } from 'react'
import { client } from './client'

export const useGetLatestCo2Value = () => {
  const [state, setState] = useState(''); // 'loading', 'error' or 'success'
  const [data, setData] = useState({value: 0, createdAt: '' });
  const [error, setError] = useState('');
  useEffect(() => {
    async function getInitialData(){
      try {
        const data = await client.getLatestCo2Value()
        if (data) {
          setData(data)
          setState('success')
        }
      } catch (error) {
        setError(error)
        setState('error')
      }
    }
    setState('loading')
    getInitialData()
  }, [])
  return {
    state, data, error
  }
}
