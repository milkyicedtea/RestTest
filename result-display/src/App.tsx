import { LanguageDataFiles } from "./constants/LanguageDataFiles";
import LanguageCard from "./components/LanguageCard";
import Header from "./components/Header";
import { useGraphValue } from "./context/graphContext";
import { useEffect } from "react";

function App() {
  const languageResults = LanguageDataFiles;
  const { setMaxYValue } = useGraphValue();

  useEffect(() => {
    const maxLatency = Math.max(
      ...languageResults.flatMap((result) =>
        result.latency_distribution.map(
          (latency) => latency.latency_in_microseconds
        )
      )
    );

    setMaxYValue(maxLatency);
  }, [languageResults, setMaxYValue]);

  return (
    <>
      <Header />
      <div className="cards">
        {languageResults.map((result, index) => (
          <LanguageCard key={result.lang_name || index} {...result} />
        ))}
      </div>
    </>
  );
}

export default App;
