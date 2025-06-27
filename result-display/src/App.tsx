import { LanguageDataFiles } from "./constants/LanguageDataFiles";
import LanguageCard from "./components/LanguageCard";
import Header from "./components/Header";

function App() {
  const languageResults = LanguageDataFiles;

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
