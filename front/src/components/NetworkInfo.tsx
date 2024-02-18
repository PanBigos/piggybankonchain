import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { DEFAULT_CHAIN } from "@/constants";
import Image from "next/image";
import { useEffect } from "react";
import { useAccount, useSwitchChain } from "wagmi";
import { Button } from "./ui/button";
import { useToast } from "./ui/use-toast";

function NetworkInfo() {
  const { chain } = useAccount();
  const { chains, error, switchChain } = useSwitchChain();

  const defaultChainId = chain?.id?.toString();

  const { toast } = useToast();

  useEffect(() => {
    if (error?.message) {
      toast({
        title: "Error",
        description: error.message,
        variant: "destructive",
      });
    }
  }, [error]);

  if (chain === undefined || !chains.map((item) => item.id).includes(chain.id))
    return (
      <Button
        variant="destructive"
        onClick={() => switchChain?.({ chainId: DEFAULT_CHAIN.chain.id })}
      >
        Switch to supported chain
      </Button>
    );

  return (
    <>
      <Select
        onValueChange={(value: string) => {
          switchChain?.({ chainId: Number(value) });
        }}
        defaultValue={defaultChainId}
        value={chain.id.toString()}
      >
        <SelectTrigger>
          <SelectValue />
        </SelectTrigger>
        <SelectContent>
          {chains.map((x) => (
            <SelectItem value={x.id.toString()} key={x.id}>
              <div className="flex gap-2">
                <Image
                  src={DEFAULT_CHAIN.chainLogo}
                  width={20}
                  height={20}
                  alt="fuse"
                  className="rounded-xl"
                />
                {x.name}
              </div>
            </SelectItem>
          ))}
        </SelectContent>
      </Select>
    </>
  );
}
export default NetworkInfo;
