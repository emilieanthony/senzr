import { useGetLatestCo2Value } from "../../utils/api/useGetLatestCo2Value";
import { useGetLatestTemperatureValue } from "../../utils/api/useGetLatestTemperatureValue";
import { useGetLatestHumidityValue } from "../../utils/api/useGetLatestHumidityValue";
import './home.css';
import Value from "../../components/Value/Value";

const AVG_OUTDOOR_CO2_LEVEL_PPM = 400;
const AVG_ROOM_CO2_LEVEL_PPM = 1000;

const COLD_INDOOR_TEMPERATURE_DEGREES = 18;
const LOW_INDOOR_TEMPERATURE_DEGREES = 20;
const HIGH_INDOOR_TEMPERATURE_DEGREES = 23;
const WARM_INDOOR_TEMPERATURE_DEGREES = 25;

const VERY_LOW_INDOOR_HUMIDITY_PERCENT = 20;
const LOW_INDOOR_HUMIDITY_PERCENT = 30;
const HIGH_INDOOR_HUMIDITY_PERCENT = 50;
const VERY_HIGH_INDOOR_HUMIDITY_PERCENT = 60;

const calculateCarbonDioxideThreshold = (co2) => {
  if (co2 < AVG_OUTDOOR_CO2_LEVEL_PPM) {
    return 'good'
  }
  if (co2 < AVG_ROOM_CO2_LEVEL_PPM) {
    return 'medium'
  }
  return 'bad'
}

const calculateTemperatureThreshold = (temperature) => {
  if (temperature >= COLD_INDOOR_TEMPERATURE_DEGREES && temperature <= LOW_INDOOR_TEMPERATURE_DEGREES) {
    return 'medium'
  }
  if (temperature >= HIGH_INDOOR_TEMPERATURE_DEGREES && temperature <= WARM_INDOOR_TEMPERATURE_DEGREES) {
    return 'medium'
  }
  if (temperature < COLD_INDOOR_TEMPERATURE_DEGREES) {
    return 'bad'
  }
  if (temperature > WARM_INDOOR_TEMPERATURE_DEGREES) {
    return 'bad'
  }
  return 'good'
}

const calculateHumidityThreshold = (humidity) => {
  if (humidity > VERY_LOW_INDOOR_HUMIDITY_PERCENT && humidity <= LOW_INDOOR_HUMIDITY_PERCENT) {
    return 'medium'
  }
  if (humidity > HIGH_INDOOR_HUMIDITY_PERCENT && humidity <= VERY_HIGH_INDOOR_HUMIDITY_PERCENT) {
    return 'medium'
  }
  if (humidity < VERY_LOW_INDOOR_HUMIDITY_PERCENT) {
    return 'bad'
  }
  if (humidity > VERY_HIGH_INDOOR_HUMIDITY_PERCENT) {
    return 'bad'
  }
  return 'good'
}

const Home = () => {
  const { state: tState, data: tData } = useGetLatestTemperatureValue()
  const { state: cState, data: cData } = useGetLatestCo2Value()
  const { state: hState, data: hData } = useGetLatestHumidityValue()
  return (
    <div className="container">
      <Value title="Carbon Dioxide" value={cData?.value} unit="ppm" timestamp={cData?.createdAt} state={cState} level={calculateCarbonDioxideThreshold(cData?.value)} />
      <Value title="Temperature" value={tData?.value} unit="C" timestamp={tData?.createdAt} state={tState} level={calculateTemperatureThreshold(tData?.value)} />
      <Value title="Humidity" value={hData?.value} unit="%" timestamp={hData?.createdAt} state={hState} level={calculateHumidityThreshold(hData?.value)} />
    </div>
  )
}

export default Home
