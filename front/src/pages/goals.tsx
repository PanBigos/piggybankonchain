import { BalancesTab } from "@/components/tabs/BalancesTab";
import { MessagesTab } from "@/components/tabs/MessagesTab";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";

interface GoalsPageProps {}

const GoalsPage = (props: GoalsPageProps) => {
  return (
    <div className="container space-y-4 py-10">
      <div className="flex flex-col mb-8 gap-2">
        <p className="text-2xl text-center">My goals</p>
        <p className="text-center">
          Here you can see all of your created goals and balances.
        </p>
      </div>

      <Tabs defaultValue="messages" className="text-center mt-8">
        <TabsList>
          <TabsTrigger value="messages">Messages</TabsTrigger>
          <TabsTrigger value="balances">Balances</TabsTrigger>
        </TabsList>

        <TabsContent value="messages" className="text-left">
          <MessagesTab />
        </TabsContent>
        <TabsContent value="balances" className="text-left">
          <BalancesTab />
        </TabsContent>
      </Tabs>
    </div>
  );
};

export default GoalsPage;
