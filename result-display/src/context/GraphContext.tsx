import { createContext, useContext, useState, type ReactNode } from "react";

interface GraphContextType {
  maxYValue: number;
  setMaxYValue: (value: number) => void;
}

const ValueContext = createContext<GraphContextType | undefined>(undefined);

export const useGraphValue = () => {
  const context = useContext(ValueContext);
  if (!context) {
    throw new Error("useGraphValue must be used within a GraphValueProvider");
  }
  return context;
};

export const GraphValueProvider = ({ children }: { children: ReactNode }) => {
  const [maxYValue, setMaxYValue] = useState<number>(0);

  return (
    <ValueContext.Provider value={{ maxYValue, setMaxYValue }}>
      {children}
    </ValueContext.Provider>
  );
};
