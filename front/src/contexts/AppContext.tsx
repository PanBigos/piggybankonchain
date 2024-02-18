import { Token } from "@/types";
import { FuseSDK } from "@fuseio/fusebox-web-sdk";
import { createContext, useContext } from "react";

type AppContextType = {
  userAddress: string | undefined;
  fuse: FuseSDK | undefined;
  tokens: Token[];
};

export const AppContext = createContext<AppContextType>({} as AppContextType);

export const useApp = () => {
  const context = useContext(AppContext);
  if (!context) {
    throw new Error("useApp must be used within a AppProvider");
  }
  return context;
};
