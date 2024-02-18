import { TimeLockedPiggyBankFactory } from "@/abis/TimeLockedPiggyBankFactory";
import { Button } from "@/components/ui/button";
import { DEFAULT_CHAIN } from "@/constants";
import { useApp } from "@/contexts/AppContext";
import Link from "next/link";
import { useEffect, useState } from "react";
import { createPublicClient, http } from "viem";
import { columns } from "../msges/components/Columns";
import { DataTable } from "../msges/components/Table";
import moment from "moment";

const mockedMessages = [
  {
    date: moment().subtract(2, "minute").format("YYYY-MM-DD HH:mm:ss"),
    address: "mockedaddress d",
    token: "FUSE",
    amount: "123.31",
    fee: "123.2",
    content: "Happy Birthday, I wish you every happiness!",
    nick: "Januszek",
  },
  {
    date: moment().subtract(2, "year").format("YYYY-MM-DD HH:mm:ss"),
    address: "mockedaddress d",
    token: "FUSE",
    amount: "6123.31",
    fee: "1123.2",
    content: "Best wishes on your birthday – may you have many, many more.",
    nick: "Dad",
  },
];

interface MessagesTabProps {}

export const MessagesTab = (props: MessagesTabProps) => {
  const {} = props;
  const [banks, setBanks] = useState<`0x${string}`[]>([]);

  const { userAddress } = useApp();

  const publicClient = createPublicClient({
    batch: {
      multicall: true,
    },
    chain: DEFAULT_CHAIN.chain,
    transport: http(),
  });

  const getBanks = async () => {
    const piggyBanks = (await publicClient.readContract({
      address: DEFAULT_CHAIN.piggyBankFactory,
      abi: TimeLockedPiggyBankFactory,
      functionName: "getPiggyBanks",
      args: [userAddress],
    })) as `0x${string}`[];

    setBanks(piggyBanks);
  };

  useEffect(() => {
    if (userAddress) {
      getBanks();
    }
  }, [userAddress]);

  return (
    <div className="space-y-4">
      <div className="text-right">
        <Button onClick={() => {}}>Refresh</Button>
      </div>

      <div className="flex flex-col gap-7">
        {banks.map((bank) => (
          <div key={bank} className="border rounded-xl px-6 py-10">
            <p className="text-xl text-center">
                Goal: Annie 18th birthday
              </p>
            <p className="text-xl text-center">
              <Link href={`/${`bank`}`}>{bank}</Link>
            </p>
            <DataTable data={mockedMessages} columns={columns} />
          </div>
        ))}
      </div>
    </div>
  );
};
