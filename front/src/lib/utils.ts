import { type ClassValue, clsx } from "clsx";
import { twMerge } from "tailwind-merge";
import { toHex } from "viem";

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs));
}

export const shortenAddr = (addr: string): string => {
  if (addr && addr.length >= 10) {
    const pre = addr.slice(0, 5);
    const post = addr.slice(addr.length - 3, addr.length);
    return pre.concat("...").concat(post);
  }
  return addr;
};

export const getEncoded = (nickname: string, msg: string) => {
  const json = JSON.stringify({
    nick: nickname,
    msg: msg,
  });
  const encoded = toHex(json);

  return encoded;
};
