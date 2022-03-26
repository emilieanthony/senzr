import { useGetLatestCo2Value } from "../../utils/api/useGetLatestCo2Value";
import "./home.css";
import Value from "../../components/Value/Value";
import Card from "../../components/Card/Card";
import { useGetDailyAverageCo2Value } from "../../utils/api/useGetDailyAverageCo2Value";

const AVG_OUTDOOR_CO2_LEVEL_PPM = 400;
const AVG_ROOM_CO2_LEVEL_PPM = 1000;

const calculateCarbonDioxideThreshold = (co2) => {
  if (co2 < AVG_OUTDOOR_CO2_LEVEL_PPM) {
    return "good";
  }
  if (co2 < AVG_ROOM_CO2_LEVEL_PPM) {
    return "medium";
  }
  return "bad";
};

const Home = () => {
  const { state: cState, data: cData } = useGetLatestCo2Value();
  const { state: dailyAverageState, data: dailyAverage } =
    useGetDailyAverageCo2Value();
  return (
    <div className="container">
      <Card>
        <Value
          title="Live Carbon Dioxide"
          value={cData?.value}
          unit="ppm"
          state={cState}
          level={calculateCarbonDioxideThreshold(cData?.value)}
        />
      </Card>
      <Card>
        <Value
          title="Average / Day"
          value={dailyAverage?.value}
          unit="ppm"
          state={dailyAverageState}
          level={calculateCarbonDioxideThreshold(dailyAverage?.value)}
        />
      </Card>
      <Card className="full-width-card">
        <Value
          title="Week"
          /* value={hData?.value}
          unit="%"
          timestamp={hData?.createdAt}
          state={hState}
          level={calculateHumidityThreshold(hData?.value)} */
        />
      </Card>
    </div>
  );
};

export default Home;
