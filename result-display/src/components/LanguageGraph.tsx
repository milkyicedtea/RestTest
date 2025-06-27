import type { LatencyObjects } from "./LanguageCard";
import { ResponsiveLine } from "@nivo/line";
import { useGraphValue } from "@/context/graphContext";

interface GraphData {
  data: LatencyObjects[];
  lang_name: string;
}

export default function LanguageGraph({ data, lang_name }: GraphData) {
  const { maxYValue } = useGraphValue();
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
          margin={{ top: 20, right: 20, bottom: 50, left: 70 }}
          yScale={{
            type: "linear",
            min: "auto",
            max: maxYValue,
            stacked: true,
            reverse: false,
          }}
          axisBottom={{ legend: "percentile", legendOffset: 36 }}
          axisLeft={{ legend: "latency (mus)", legendOffset: -60 }}
          pointSize={10}
          pointColor={{ theme: "background" }}
          pointBorderWidth={2}
          pointBorderColor={{ from: "seriesColor" }}
          pointLabelYOffset={-12}
          enableTouchCrosshair={true}
          useMesh={true}
          tooltip={({ point }) => (
            <div
              style={{
                background: "white",
                padding: "9px 12px",
                border: "1px solid #ccc",
                borderRadius: "4px",
                color: "black",
                fontSize: "10px",
                whiteSpace: "nowrap",
              }}
            >
              <div>{point.data.yFormatted} μs</div>
            </div>
          )}
        />
      </div>
    </>
  );
}
