import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useApp } from "@/contexts/AppContext";

import Image from "next/image";

interface TokenSelectProps {
  value: string;
  onChange: (value: string) => void;
}

export const TokenSelect = (props: TokenSelectProps) => {
  const { value, onChange } = props;

  const { tokens } = useApp();

  return (
    <Select value={value} onValueChange={onChange}>
      <SelectTrigger className="w-full">
        <SelectValue />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          {tokens.map((token) => (
            <SelectItem key={token.address} value={token.symbol}>
              <div className="flex gap-2">
                <Image
                  src={token.logoURI}
                  width={20}
                  height={20}
                  alt="fuse"
                  className="rounded-xl"
                />
                {token.name}
              </div>
            </SelectItem>
          ))}
        </SelectGroup>
      </SelectContent>
    </Select>
  );
};
