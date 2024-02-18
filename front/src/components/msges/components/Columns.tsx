"use client";

import { ColumnDef } from "@tanstack/react-table";
import { DEFAULT_CHAIN, fuseToken } from "@/constants";
import { MsgItem } from "@/types";
import moment from "moment";
import Image from "next/image";
import { DataTableColumnHeader } from "./ColumnHeader";

export const columns: ColumnDef<MsgItem>[] = [
  {
    accessorKey: "date",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Date" />
    ),
    cell: ({ row }) => {
      return <span>{moment(row.getValue("date")).fromNow()}</span>;
    },
    filterFn: (row, id, value) => {
      return value.includes(row.getValue(id));
    },
  },
  {
    accessorKey: "chainId",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Chain" />
    ),
    cell: ({ row }) => {
      return (
        <div>
          <Image
            src={DEFAULT_CHAIN.chainLogo}
            width={20}
            height={20}
            alt="FUSE"
            className="rounded-full"
          />
        </div>
      );
    },
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "nick",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Name" />
    ),
    cell: ({ row }) => <span>{row.getValue("nick")}</span>,
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "content",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Message" />
    ),
    cell: ({ row }) => {
      // const label = labels.find((label) => label.value === row.original.label);

      return (
        <div className="flex">
          <span className="max-w-[500px] font-medium">
            {row.getValue("content")}
          </span>
        </div>
      );
    },
  },

  {
    accessorKey: "token",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Token" />
    ),
    cell: ({ row }) => {
      const tokenData = fuseToken;
      return (
        <div className="w-[60px] flex items-center justify-center">
          <div className="flex gap-2">
            <Image
              src={tokenData.logoURI}
              alt={tokenData.symbol}
              width={20}
              height={20}
              className="rounded-full"
            />
            <span className="flex items-center justify-center">
              {tokenData.symbol}
            </span>
          </div>
        </div>
      );
    },
    enableSorting: false,
    enableHiding: false,
  },
  {
    accessorKey: "amount",
    header: ({ column }) => (
      <DataTableColumnHeader column={column} title="Amount" />
    ),
    cell: ({ row }) => <div className="w-[80px]">{row.getValue("amount")}</div>,
    enableSorting: false,
    enableHiding: false,
  },
];
