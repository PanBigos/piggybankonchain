import { Chain } from "viem";

export type ConfigChain = {
  chain: Chain;
  chainLogo: string;
  piggyBankFactory: `0x${string}`;
  piggyBankPriceFactory: `0x${string}`;
  piggyBankRouter: `0x${string}`;
  api: string;
};

export type TokenBalance = {
  name: string;
  balance: string;
};

export type Token = {
  name: string;
  symbol: string;
  decimals: number;
  address: string;
  logoURI: string;
  type: string;
};

export type BankBalance = {
  bank: string;
  balances: Balance[];
};

export type Balance = {
  token: Token;
  balance: bigint;
};
export type MsgItem = {
  date: string;
  nick: string;
  content: string;
  token: string;
  amount: string;
};
