import { DividendTab } from "@/components/tabs/DividendTab";
import { PriceLockTab } from "@/components/tabs/PriceLockTab";
import { TimeLockTab } from "@/components/tabs/TimeLockTab";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

const CreatePage = () => {
  return (
    <div className="container flex flex-col gap-4justify-center relative">
      <div className="flex flex-col mb-8 gap-2">
        <h1 className="text-2xl font-bold text-center">Create a goal</h1>
        <p className="text-center">You can create a new goal here.</p>

        <Tabs defaultValue="timeLock" className="text-center mt-8">
          <TabsList>
            <TabsTrigger value="timeLock">Time lock</TabsTrigger>
            <TabsTrigger value="priceLock">Price lock</TabsTrigger>
            <TabsTrigger value="dividend">Dividend</TabsTrigger>
          </TabsList>

          <TabsContent value="timeLock" className="text-left">
            <TimeLockTab />
          </TabsContent>

          <TabsContent value="priceLock" className="text-left">
            <PriceLockTab />
          </TabsContent>

          <TabsContent value="dividend" className="text-left">
            <DividendTab />
          </TabsContent>
        </Tabs>
      </div>
    </div>
  );
};

export default CreatePage;
