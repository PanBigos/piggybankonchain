import { AppContext } from "@/contexts/AppContext";
import { FuseSDK } from "@fuseio/fusebox-web-sdk";
import { ethers } from "ethers";
import { ReactNode, useEffect, useState } from "react";
import { useAccount } from "wagmi";
import axios from "axios";
import { Token } from "@/types";
import { fuseToken } from "@/constants";

const apiKey = process.env.NEXT_PUBLIC_FUSE_PUBLIC_KEY;

interface AppProviderProps {
  children: ReactNode;
}

export const AppProvider = ({ children }: AppProviderProps) => {
  const account = useAccount();
  const [fuse, setFuse] = useState<FuseSDK | undefined>();
  const [userAddress, setUserAddress] = useState<string | undefined>();
  const [tokens, setTokens] = useState<Token[]>([]);

  useEffect(() => {
    axios
      .get(`https://api.fuse.io/api/v0/trade/tokens?apiKey=${apiKey}`)
      .then(({ data }) => {
        const tokens = data.data.tokens;
        const filteredTokens = tokens.filter(
          (token: Token) => token.type === "misc"
        );
        setTokens([fuseToken, ...filteredTokens]);
      });
  }, []);

  const initFuse = async () => {
    const provider = new ethers.BrowserProvider(window.ethereum);
    const signer = await provider.getSigner();

    const fuseObj = await FuseSDK.init(apiKey!, signer as any, {
      withPaymaster: true,
    });

    setFuse(fuseObj);

    console.log(`logged in, ${fuseObj.wallet.getSender()}`);
  };

  useEffect(() => {
    if (!userAddress && account.status === "connected" && !fuse) {
      setUserAddress(account.address);
      initFuse().catch(console.error);
    }
  }, [account]);

  useEffect(() => {
    if (account.isDisconnected) {
      setUserAddress(undefined);
    }
  }, [account.isDisconnected]);

  return (
    <AppContext.Provider value={{ userAddress, fuse, tokens }}>
      {children}
    </AppContext.Provider>
  );
};
