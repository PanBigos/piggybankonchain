import Image from "next/image";

export default function Home() {
  return (
    <div className="container grid mt-20 text-center align-center gap-4">
      <Image
        src="/piggybank.png"
        width={300}
        height={300}
        alt="pig"
        className="ml-auto mr-auto"
      />
      <p className="text-4xl">PiggyBankOnChain</p>
      <p className="text-2xl">Secure, Gasless Saving & Wishes.</p>

      <p className="text-2xl">Unlock financial goals with PiggyBankOnChain! Create gasless, secure vaults—time-locked, price-locked or dividend-earning. Send wishes, save smarter.</p>
      
    </div>
  );
}
