import React, { useMemo } from 'react'
import { format } from "date-fns";
import './value.css';

const Value = ({title, value, unit, timestamp, state, level}) => {

  const renderState = useMemo(() => {
    if(state === 'error') {
    return "error"
  } else if (state === 'loading' || !value || !timestamp) {
    return "loading"
  } else {
    return "success"
  }
  }, [state, value, timestamp])

  return (
    <div className="metric">
      <h1 className="title">{title}</h1>
      {renderState === "error" && <p className="error">An error occured.</p>}
      {renderState === "loading" && <p>Loading...</p>}
      {renderState === "success" && (
        <>
        <h2 className={`value ${level}`}>{Math.floor(value)} {unit}</h2>
        <h6 className="timestamp">{format(Date.parse(timestamp), 'PP kk:mm:ss')}</h6>
      </>
      )}
    </div>
  )
}

export default Value
