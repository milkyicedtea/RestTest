import type { LatencyObjects } from "./LanguageCard";
import { ResponsiveLine } from "@nivo/line";

interface GraphData {
  data: LatencyObjects[];
  lang_name: string;
}

export default function LanguageGraph({ data, lang_name }: GraphData) {
  const transformedData = [
    {
      id: lang_name,
      data: data.map((item) => ({
        x: item.percentile,
        y: item.latency_in_microseconds,
      })),
    },
  ];

  return (
    <>
      <div className="graph-container">
        <ResponsiveLine
          data={transformedData}
          margin={{ top: 20, right: 20, bottom: 50, left: 60 }}
          yScale={{
            type: "linear",
            min: "auto",
            max: "auto",
            stacked: true,
            reverse: false,
          }}
          axisBottom={{ legend: "percentile", legendOffset: 36 }}
          axisLeft={{ legend: "latency (mus)", legendOffset: -50 }}
          pointSize={10}
          pointColor={{ theme: "background" }}
          pointBorderWidth={2}
          pointBorderColor={{ from: "seriesColor" }}
          pointLabelYOffset={-12}
          enableTouchCrosshair={true}
          useMesh={true}
        />
      </div>
    </>
  );
}
