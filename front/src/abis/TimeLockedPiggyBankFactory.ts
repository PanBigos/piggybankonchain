export const TimeLockedPiggyBankFactory = [
  {
    type: "function",
    name: "createPiggyBank",
    inputs: [
      { name: "_owner", type: "address", internalType: "address" },
      { name: "_unlockDate", type: "uint256", internalType: "uint256" },
    ],
    outputs: [{ name: "piggyBank", type: "address", internalType: "address" }],
    stateMutability: "nonpayable",
  },
  {
    type: "function",
    name: "getCreatedPiggyBanks",
    inputs: [{ name: "_user", type: "address", internalType: "address" }],
    outputs: [{ name: "", type: "address[]", internalType: "address[]" }],
    stateMutability: "view",
  },
  {
    type: "function",
    name: "getPiggyBanks",
    inputs: [{ name: "_user", type: "address", internalType: "address" }],
    outputs: [{ name: "", type: "address[]", internalType: "address[]" }],
    stateMutability: "view",
  },
  {
    type: "event",
    name: "CreatedPiggyBank",
    inputs: [
      {
        name: "piggyBank",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "creator",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "owner",
        type: "address",
        indexed: false,
        internalType: "address",
      },
      {
        name: "createdAt",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
      {
        name: "unlockDate",
        type: "uint256",
        indexed: false,
        internalType: "uint256",
      },
    ],
    anonymous: false,
  },
];
