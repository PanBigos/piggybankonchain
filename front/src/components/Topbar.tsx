import Image from "next/image";
import { Button } from "./ui/button";
import { useRouter } from "next/router";
import Link from "next/link";
import { cn } from "@/lib/utils";
import { ConnectButton } from "./ConnectButton";
import { useApp } from "@/contexts/AppContext";
import NetworkInfo from "./NetworkInfo";
import { ConnectSheet } from "./ConnectSheet";

interface TopbarProps {}

export const Topbar = (props: TopbarProps) => {
  const {} = props;
  const router = useRouter();

  const { userAddress } = useApp();

  const isHome = router.pathname === "/";
  const isCreate = router.pathname === "/create";
  const isMyGoals = router.pathname === "/goals";
  const isAccount = router.pathname === "/[account]";

  return (
    <div className="container flex items-center p-4 justify-between relative z-10 mb-10">
      <Link href="/" className="flex items-center gap-2">
        <Image src="pig.svg" width={60} height={60} alt="logo" />
        <h1 className="text-xl">PiggyBankOnChain</h1>
      </Link>
      {!isAccount && (
        <div className="flex items-center gap-4 uppercase">
          <MenuItem isActive={isHome} link="/">
            Home
          </MenuItem>
          <MenuItem isActive={isCreate} link="/create">
            Create a goal
          </MenuItem>
          <MenuItem isActive={isMyGoals} link="/goals">
            My goals
          </MenuItem>
        </div>
      )}
      <div>
        {userAddress ? (
          <div className="flex items-center gap-4">
            <NetworkInfo />
            <ConnectButton address={userAddress} />
          </div>
        ) : (
          <ConnectSheet></ConnectSheet>
        )}
      </div>
    </div>
  );
};

interface MenuItemProps {
  isActive: boolean;
  link: string;
  children: React.ReactNode;
}

const MenuItem = (props: MenuItemProps) => {
  const { isActive, link, children } = props;
  return (
    <Link
      href={link}
      className={cn(
        isActive ? "text-white" : " text-slate-500",
        "hover:text-white transition-all py-1 px-2 rounded-xl"
      )}
    >
      {children}
    </Link>
  );
};
