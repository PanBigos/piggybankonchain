import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import {
  Form,
  FormControl,
  FormField,
  FormItem,
  FormLabel,
  FormMessage,
} from "@/components/ui/form";
import { Input } from "@/components/ui/input";
import { zodResolver } from "@hookform/resolvers/zod";
import { CalendarIcon } from "@radix-ui/react-icons";
import { format } from "date-fns";
import { useForm } from "react-hook-form";
import { z } from "zod";

import { TimeLockedPiggyBankFactory } from "@/abis/TimeLockedPiggyBankFactory";
import { Calendar } from "@/components/ui/calendar";
import {
  Popover,
  PopoverContent,
  PopoverTrigger,
} from "@/components/ui/popover";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { DEFAULT_CHAIN } from "@/constants";
import { useApp } from "@/contexts/AppContext";
import { cn } from "@/lib/utils";
import { useState } from "react";
import { encodeFunctionData } from "viem";

import { useToast } from "@/components/ui/use-toast";
import Image from "next/image";
import { TokenSelect } from "../TokenSelect";

const formSchema = z.object({
  goalName: z.string().min(2, {
    message: "Goal name must be at least 2 characters.",
  }),
  chain: z.string(),
  lockDate: z.number(),
});

interface TimeLockTabProps {}

export const TimeLockTab = (props: TimeLockTabProps) => {
  const {} = props;

  const { fuse, userAddress } = useApp();
  const [isSubmitting, setIsSubmitting] = useState(false);

  const { toast } = useToast();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      goalName: "",
      chain: "fuse",
      lockDate: undefined,
    },
  });

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    setIsSubmitting(true);
    try {
      const txData = encodeFunctionData({
        abi: TimeLockedPiggyBankFactory,
        functionName: "createPiggyBank",
        args: [userAddress, values.lockDate],
      }) as string;

      const to = DEFAULT_CHAIN.piggyBankFactory;
      const value = 0;

      const data = Uint8Array.from(Buffer.from(txData.substring(2), "hex"));

      const res = await fuse!.callContract(to, value, data);

      console.log(`UserOpHash: ${res?.userOpHash}`);
      console.log("Waiting for transaction...");

      const receipt = await res?.wait();
      console.log("Transaction Hash:", receipt?.transactionHash);

      toast({
        title: "Success",
        description: "Goal created successfully",
      });
      setIsSubmitting(false);
    } catch (error) {
      console.log(error);
      setIsSubmitting(false);
    }
  };

  return (
    <div className="flex items-center justify-center">
      <Card className="p-4 w-[500px] my-4">
        <div className="space-y-6">
          <Form {...form}>
            <form onSubmit={form.handleSubmit(onSubmit)} className="space-y-4">
              <FormField
                control={form.control}
                name="goalName"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Goal name</FormLabel>
                    <FormControl>
                      <Input {...field} />
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <FormField
                control={form.control}
                name="chain"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Chain</FormLabel>
                    <FormControl>
                      <Select
                        value={field.value}
                        onValueChange={field.onChange}
                      >
                        <SelectTrigger className="w-full">
                          <SelectValue />
                        </SelectTrigger>
                        <SelectContent>
                          <SelectGroup>
                            <SelectItem value="fuse">
                              <div className="flex gap-2">
                                <Image
                                  src={DEFAULT_CHAIN.chainLogo}
                                  width={20}
                                  height={20}
                                  alt="fuse"
                                  className="rounded-xl"
                                />
                                Fuse
                              </div>
                            </SelectItem>
                          </SelectGroup>
                        </SelectContent>
                      </Select>
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />

              <FormField
                control={form.control}
                name="lockDate"
                render={({ field }) => (
                  <FormItem>
                    <FormLabel>Lock date</FormLabel>
                    <FormControl>
                      <Popover>
                        <PopoverTrigger asChild>
                          <Button
                            variant={"outline"}
                            className={cn(
                              "w-full justify-start text-left font-normal",
                              !field.value && "text-muted-foreground"
                            )}
                          >
                            <CalendarIcon className="mr-2 h-4 w-4" />
                            {field.value ? (
                              format(field.value, "PPP")
                            ) : (
                              <span>Pick a date</span>
                            )}
                          </Button>
                        </PopoverTrigger>
                        <PopoverContent className="w-auto p-0" align="start">
                          <Calendar
                            mode="single"
                            selected={new Date(field.value)}
                            onSelect={(val) => field.onChange(val?.getTime())}
                            initialFocus
                          />
                        </PopoverContent>
                      </Popover>
                    </FormControl>
                    <FormMessage />
                  </FormItem>
                )}
              />
              <div className="text-center">
                <Button
                  type="submit"
                  className="ml-auto mr-auto"
                  disabled={isSubmitting}
                >
                  {isSubmitting ? "Submitting..." : "Create"}
                </Button>
              </div>
            </form>
          </Form>
        </div>
      </Card>
    </div>
  );
};
