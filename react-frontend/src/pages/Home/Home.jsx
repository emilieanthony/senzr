import { useGetLatestCo2Value } from "../../utils/api/useGetLatestCo2Value";
import { useGetDailyAverageCo2Value } from "../../utils/api/useGetDailyAverageCo2Value";
import { useWeeklyTimeSeries } from "../../utils/api/useWeeklyTimeSeries";
import "./home.css";
import Value from "../../components/Value/Value";
import Card from "../../components/Card/Card";
import {
  LineChart,
  Line,
  XAxis,
  YAxis,
  CartesianGrid,
  Tooltip,
  Legend,
  ResponsiveContainer,
} from "recharts";

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
  const { state: liveState, data: liveData } = useGetLatestCo2Value();
  const { state: dailyAverageState, data: dailyAverage } =
    useGetDailyAverageCo2Value();
  const { state: weeklyState, data: weeklyData } = useWeeklyTimeSeries();
  return (
    <div className="container">
      <Card>
        <Value
          title="Live Carbon Dioxide"
          value={liveData?.value}
          unit="ppm"
          state={liveState}
          level={calculateCarbonDioxideThreshold(liveData?.value)}
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
      <Card className="full-width-card chart-card">
        <h1 className="chart-title">Weekly Carbon Dioxide Levels</h1>
        {weeklyState === "loading" && <p className="loading">Loading...</p>}
        <ResponsiveContainer width="100%" height="90%">
          <LineChart
            width={500}
            height={300}
            data={weeklyData}
            margin={{
              top: 5,
              right: 30,
              left: 20,
              bottom: 5,
            }}
          >
            <CartesianGrid strokeDasharray="3 3" />
            <XAxis
              tickFormatter={(date) => new Date(date).toLocaleDateString()}
              dataKey="createdAt"
            />
            <YAxis />
            <Tooltip />
            <Legend />
            <Line
              type="monotone"
              dataKey="value"
              stroke="#8884d8"
              activeDot={{ r: 8 }}
            />
          </LineChart>
        </ResponsiveContainer>
      </Card>
    </div>
  );
};

export default Home;
