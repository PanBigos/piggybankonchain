import { TimeLockedPiggyBankFactory } from "@/abis/TimeLockedPiggyBankFactory";
import { erc20 } from "@/abis/erc20";
import { Button } from "@/components/ui/button";
import { DEFAULT_CHAIN, fuseToken } from "@/constants";
import { useApp } from "@/contexts/AppContext";
import { Balance, BankBalance } from "@/types";
import { ethers } from "ethers";
import { sortBy } from "lodash";
import moment from "moment";
import Image from "next/image";
import Link from "next/link";
import { useEffect, useState } from "react";
import { createPublicClient, http } from "viem";
import { Separator } from "../ui/separator";

interface BalancesTabProps {}

export const BalancesTab = (props: BalancesTabProps) => {
  const {} = props;

  const [balances, setBalances] = useState<BankBalance[]>([]);
  const {} = props;
  const { fuse, userAddress, tokens } = useApp();

  const getBalances = async () => {
    const publicClient = createPublicClient({
      batch: {
        multicall: true,
      },
      chain: DEFAULT_CHAIN.chain,
      transport: http(),
    });

    const piggyBanks = (await publicClient.readContract({
      address: DEFAULT_CHAIN.piggyBankFactory,
      abi: TimeLockedPiggyBankFactory,
      functionName: "getPiggyBanks",
      args: [userAddress],
    })) as `0x${string}`[];

    const fetchBalancesForBank = async (bank: `0x${string}`) => {
      const tokensWithoutEther = tokens.filter(
        (token) => token.symbol !== "FUSE"
      );

      let balanceList = await publicClient.multicall({
        contracts: tokensWithoutEther.map((token) => ({
          abi: erc20.abi,
          address: token.address as `0x${string}`,
          functionName: "balanceOf",
          args: [bank],
        })),
      });

      const balances = balanceList.map((value, idx) => ({
        token: tokensWithoutEther[idx],
        balance: value.result ? BigInt(value.result) : BigInt(0),
      }));

      const userEtherBalance = await publicClient.getBalance({
        address: bank,
      });

      const etherBalance = {
        token: fuseToken,
        balance: BigInt(userEtherBalance),
      };

      return { bank, balances: [...balances, etherBalance] };
    };

    const totalBalances = Promise.all(piggyBanks.map(fetchBalancesForBank));

    totalBalances.then((balances) => {
      setBalances(balances);
    });
  };

  useEffect(() => {
    if (fuse) {
      getBalances();

      setInterval(() => {
        getBalances();
      }, 120000);
    }
  }, [fuse]);

  if (!fuse) {
    return (
      <div className="flex items-center justify-center h-full">
        <p className="text-2xl">Loading...</p>
      </div>
    );
  }

  return (
    <div className="space-y-4">
      <div className="text-right">
        <Button onClick={getBalances}>Refresh</Button>
      </div>

      <div className="flex flex-col gap-7">
        {balances.map((balance) => {
          const sortedBalances: Balance[] = sortBy(
            balance.balances,
            (balance: Balance) => balance.token.name.toLowerCase()
          ).filter((balance) => balance.balance.toString() !== "0");

          return (
            <div key={balance.bank} className="border rounded-xl px-6 py-10">
              <p className="text-lg">
                Goal: Annie 18th birthday
              </p>
              <p className="text-lg">
                Bank: <Link href={`/${balance.bank}`}>{balance.bank}</Link>
              </p>
              <p className="text-lg">Type: time lock</p>
              <p className="text-lg">
                Time left: {moment().add(1, "hour").fromNow()}
              </p>
              <Separator className="my-4" />
              <p className="text-lg">Balances:</p>
              {sortedBalances.length === 0 && (
                <p className="italic p-2">No balances</p>
              )}

              {sortedBalances.map((balance) => (
                <div key={balance.token.address}>
                  <div
                    key={balance.token.address}
                    className="flex items-center gap-2 p-2"
                  >
                    <Image
                      src={balance.token.logoURI}
                      width={20}
                      height={20}
                      alt="token"
                      className="w-5 h-5 rounded-xl"
                    />
                    <p>{balance.token.name.toUpperCase()}:</p>
                    <p>
                      {ethers.formatUnits(
                        balance.balance.toString(),
                        balance.token.decimals
                      )}
                    </p>
                  </div>
                </div>
              ))}
            </div>
          );
        })}
      </div>
    </div>
  );
};
