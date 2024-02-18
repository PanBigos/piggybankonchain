import Head from "next/head";
import { ThemeProvider } from "../providers/ThemeProvider";
import { Topbar } from "./Topbar";
import { Toaster } from "./ui/toaster";

interface LayoutProps {
  children: React.ReactNode;
}

const Layout = (props: LayoutProps) => {
  const { children } = props;

  return (
    <div className="relative">
      <Head>
        <title>Piggy Bank Onchain</title>
      </Head>

      <Topbar />
      <div className="relative w-100 h-100">
        <div className="absolute top-[50%] translate-y-1/2 left-[-300px] rounded-full bg-pink-500 blur-[200px] w-[400px] h-[400px] z-0 opacity-80" />
        <div className="absolute top-[20%] translate-x-1/2 -translate-y-1/2 right-[0] rounded-full bg-pink-500 blur-[200px] w-[400px] h-[400px] z-0 opacity-80" />
      </div>

      <div className="z-10">{children}</div>
      <Toaster />
    </div>
  );
};

export default Layout;
