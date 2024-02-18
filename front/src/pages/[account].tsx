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
import { useForm } from "react-hook-form";
import { z } from "zod";

import { TimeLockedPiggyBankRouter } from "@/abis/TimeLockedPiggyBankRouter";
import { TokenSelect } from "@/components/TokenSelect";
import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { Textarea } from "@/components/ui/textarea";
import { useToast } from "@/components/ui/use-toast";
import { DEFAULT_CHAIN } from "@/constants";
import { getEncoded } from "@/lib/utils";
import Image from "next/image";
import { useParams } from "next/navigation";
import { useEffect, useState } from "react";
import { parseUnits } from "viem";
import { useWriteContract } from "wagmi";

const formSchema = z.object({
  name: z.string().min(2, {
    message: "Goal name must be at least 2 characters.",
  }),
  message: z.string(),
  chain: z.string(),
  token: z.string(),
  amount: z.preprocess(
    (value) => {
      if (value === "") return undefined;
      const parsed = parseFloat(value as string);
      return isNaN(parsed) ? undefined : parsed;
    },
    z
      .number()
      .min(0, "Minimum must be 0 or greater")
      .refine((value) => value !== -0, {
        message: "Zero is not allowed",
      })
  ),
});

const DonatePage = () => {
  const [showPiggy, setShowPiggy] = useState(false);
  const params = useParams();
  const account = params?.account;

  const { toast } = useToast();

  const {
    data: hash,
    error,
    isError,
    isPending,
    writeContract,
    isSuccess,
    status,
  } = useWriteContract();

  const form = useForm<z.infer<typeof formSchema>>({
    resolver: zodResolver(formSchema),
    defaultValues: {
      name: undefined,
      message: undefined,
      chain: "fuse",
      token: "FUSE",
      amount: 0,
    },
  });

  useEffect(() => {
    if (error) {
      toast({
        title: "Error",
        description: error.message,
        variant: "destructive",
      });
    }
  }, [error]);

  useEffect(() => {
    if (status === "success") {
      setShowPiggy(true);

      setTimeout(() => {
        setShowPiggy(false);
      }, 10000);
    }
  }, [status]);

  const onSubmit = async (values: z.infer<typeof formSchema>) => {
    const transactionData = {
      address: DEFAULT_CHAIN.piggyBankRouter,
      abi: TimeLockedPiggyBankRouter,
      functionName: "transferWithFee",
      args: [account],
      dataSuffix: getEncoded(values.name, values.message),
      value: BigInt(parseUnits(values.amount.toString(), 18)),
    };

    writeContract({ ...transactionData });
  };

  return (
    <>
      <div className="container flex flex-col justify-center relative">
        {showPiggy && (
          <div className="absolute top-[50%] left-[50%] -translate-x-1/2 -translate-y-1/2">
            <Image
              src="https://cdn.dribbble.com/users/2261302/screenshots/7982163/media/fd1b4b3b26b542119cb49a21eb1236e7.gif"
              width={500}
              height={500}
              alt="loading"
              className="rounded-xl"
            />
          </div>
        )}
        <div className="flex flex-col mb-8 gap-2">
          <p className="text-2xl text-center">Send a gift</p>
          <p className="text-center">Sending a gift to {account}</p>
        </div>
        <div className="flex items-center justify-center">
          <Card className="p-4 w-[500px] my-8">
            <div className="space-y-6">
              <Form {...form}>
                <form
                  onSubmit={form.handleSubmit(onSubmit)}
                  className="space-y-4"
                >
                  <FormField
                    control={form.control}
                    name="name"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Name</FormLabel>
                        <FormControl>
                          <Input {...field} />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />
                  <FormField
                    control={form.control}
                    name="message"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Message</FormLabel>
                        <FormControl>
                          <Textarea
                            placeholder="I want to support you!"
                            {...field}
                          />
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
                    name="token"
                    render={({ field }) => (
                      <FormItem>
                        <FormLabel>Token</FormLabel>
                        <FormControl>
                          <TokenSelect
                            value={field.value}
                            onChange={field.onChange}
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />

                  <FormField
                    control={form.control}
                    name="amount"
                    render={({ field }) => (
                      <FormItem className="flex flex-col flex-1">
                        <FormLabel>Amount</FormLabel>
                        <FormControl>
                          <Input
                            {...field}
                            type="number"
                            min={0}
                            step="0.00000000001"
                          />
                        </FormControl>
                        <FormMessage />
                      </FormItem>
                    )}
                  />

                  <div className="text-center">
                    <Button
                      type="submit"
                      className="ml-auto mr-auto"
                      disabled={isPending}
                    >
                      {isPending ? "Submitting..." : "Send"}
                    </Button>
                  </div>
                </form>
              </Form>
            </div>
          </Card>
        </div>
      </div>
    </>
  );
};

export default DonatePage;
