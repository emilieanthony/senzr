import { format } from "date-fns";
import { useGetLatestCo2Value } from "../../utils/api/useGetLatestCo2Value";
import './home.css';

const AVG_OUTDOOR_CO2_LEVEL = 400;
const AVG_ROOM_CO2_LEVEL = 1000;

const calculateCarbonDioxideThreshold = (co2Value) => {
  if (co2Value < AVG_OUTDOOR_CO2_LEVEL) {
    return 'low'
  }
  if (co2Value < AVG_ROOM_CO2_LEVEL) {
    return 'medium'
  }
  return 'high'
}

const Home = () => {
  const { state, data, error } = useGetLatestCo2Value()
  if (state === "loading" || data.createdAt === '') {
    return <div className="container"><h3>Loading...</h3></div>
  }
  if (state === "error") {
    return <p className="error">An error occured: ${error}</p>
  }
  return (
    <div className="container">
      <div className="metric">
        <h1 className="title">Carbon Dioxide</h1>
        <h2 className={`value co2-${calculateCarbonDioxideThreshold(data.value)}`}>{data.value} ppm</h2>
        <h6 className="timestamp">{format(Date.parse(data.createdAt), 'PP kk:mm:ss')}</h6>
      </div>
      <div className="metric">
        <h1 className="title">Humidity (TBD)</h1>
        <h2 className="value">50%</h2>
        <h6 className="timestamp">{format(Date.parse(data.createdAt), 'PP kk:mm:ss')}</h6>
      </div>
      <div className="metric">
        <h1 className="title">Temperature (TBD)</h1>
        <h2 className="value">20C</h2>
        <h6 className="timestamp">{format(Date.parse(data.createdAt), 'PP kk:mm:ss')}</h6>
      </div>
    </div>
  )
}

export default Home
