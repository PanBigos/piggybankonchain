import { createConfig } from "wagmi";
import { ConfigChain, Token } from "./types";
import { Chain, http } from "viem";
import { injected } from "wagmi/connectors";
import { fuse } from "wagmi/chains";

export const piggyBankAddress = "0xff419C6fB39d37713CBD769Adad546e2eF09CC92";

export const fuseToken = {
  name: "Fuse",
  symbol: "FUSE",
  decimals: 18,
  address: "0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE",
  logoURI:
    "https://assets.coingecko.com/coins/images/10347/standard/fuse.png?1696510348",
  type: "misc",
};

export const foundryLocalDocker = {
  id: 31_337,
  name: "Test Fuse",
  nativeCurrency: {
    decimals: 18,
    name: "tFUSE",
    symbol: "tFUSE",
  },
  rpcUrls: {
    default: {
      http: ["http://localhost:5099/node/"],
      webSocket: ["ws://localhost:5099/node/"],
    },
  },
  blockExplorers: {
    default: {
      name: "Fuse explorer",
      url: "https://explorer.fuse.io/",
    },
  },
  contracts: {
    multicall3: {
      address: "0xcf7ed3acca5a467e9e704c703e8d87f634fb0fc9" as `0x${string}`,
      blockCreated: 1,
    },
  },
};

export const CHAIN_CONFIGS: ConfigChain[] = [
  {
    chain: foundryLocalDocker,
    chainLogo:
      "https://assets.coingecko.com/coins/images/10347/standard/fuse.png?1696510348",
    piggyBankFactory:
      "0x4332Df3c3938E9133432dAd0CBD3e087C39B1CcA" as `0x${string}`,
      piggyBankPriceFactory:
      "0xf039069Fd522b25a08abf958C1F631177b81b29E" as `0x${string}`,
    piggyBankRouter:
      "0xA0b025693de04bf0f4F4Ff99Ab4D98Bff0c822Fe" as `0x${string}`,
    api: "http://localhost:5051/",
  },
  {
    chain: fuse,
    chainLogo:
      "https://assets.coingecko.com/coins/images/10347/standard/fuse.png?1696510348",
    piggyBankFactory:
      "0x6678f99DE6bDee2b6074700c05F31C0571E1c703" as `0x${string}`,
    piggyBankPriceFactory:
      "0xf039069Fd522b25a08abf958C1F631177b81b29E" as `0x${string}`,
    piggyBankRouter:
      "0xA0b025693de04bf0f4F4Ff99Ab4D98Bff0c822Fe" as `0x${string}`,
    api: "http://localhost:5051/",
  },
];

export const CHAINS_LIST = CHAIN_CONFIGS.map((config) => config.chain) as [
  Chain,
  ...Chain[]
];

export const ConfigWagmi = createConfig({
  chains: CHAINS_LIST,
  transports: CHAINS_LIST.reduce((acc, chain) => {
    acc[chain.id] = http();
    return acc;
  }, {} as Record<number, ReturnType<typeof http>>),
  connectors: [injected()],
});

export const DEFAULT_CHAIN = CHAIN_CONFIGS[1];
