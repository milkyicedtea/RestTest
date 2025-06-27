import { Provider } from "./components/ui/provider.tsx";
import { GraphValueProvider } from "./context/graphContext.tsx";
import { StrictMode } from "react";
import { createRoot } from "react-dom/client";
import App from "./App.tsx";
import "./styles/Common.css";

createRoot(document.getElementById("root")!).render(
  <StrictMode>
    <Provider>
      <GraphValueProvider>
        <App />
      </GraphValueProvider>
    </Provider>
  </StrictMode>
);
