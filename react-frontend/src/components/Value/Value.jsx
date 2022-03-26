import React, { useMemo } from "react";
import "./value.css";

const Value = ({ title, value, unit, state, level }) => {
  const renderState = useMemo(() => {
    if (state === "error") {
      return "error";
    } else if (state === "loading" || !value) {
      return "loading";
    } else {
      return "success";
    }
  }, [state, value]);

  return (
    <div className="metric">
      <h1 className="title">{title}</h1>
      {renderState === "error" && <p className="error">An error occured.</p>}
      {renderState === "loading" && <p className="center-text">Loading...</p>}
      {renderState === "success" && (
        <>
          <h2 className={`value ${level}`}>
            {Math.floor(value)} {unit}
          </h2>
        </>
      )}
    </div>
  );
};

export default Value;
