import { Chain as WagmiChain } from "wagmi";
import type { Chain } from "../types";

export const chains: Record<string, Chain> = {
  31336: {
    id: 31336,
    name: "Mainnet",
    rpc: "http://34.132.67.34:8545",
    enabled: true,
  },
  167001: {
    id: 167001,
    name: "Taiko",
    rpc: "http://rpc.a1.testnet.taiko.xyz",
    enabled: true,
  },
};

export const mainnet: WagmiChain = {
  id: 31336,
  name: "Mainnet",
  network: "",
  nativeCurrency: { name: "Ether", symbol: "ETH", decimals: 18 },
  rpcUrls: {
    default: "http://34.132.67.34:8545",
  },
  blockExplorers: {
    default: {
      name: "Main",
      url: "https://34.132.67.34:4000",
    },
  },
  ens: {
    address: "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e",
  },
  multicall: {
    address: "0xca11bde05977b3631167028862be2a173976ca11",
    blockCreated: 14353601,
  },
};

export const taiko: WagmiChain = {
  id: 167001,
  name: "Taiko",
  network: "",
  nativeCurrency: { name: "Ether", symbol: "ETH", decimals: 18 },
  rpcUrls: {
    default: "http://rpc.a1.testnet.taiko.xyz",
  },
  blockExplorers: {
    default: {
      name: "Main",
      url: "https://a1.testnet.taiko.xyz",
    },
  },
  ens: {
    address: "0x00000000000C2E074eC69A0dFb2997BA6C7d2e1e",
  },
  multicall: {
    address: "0xca11bde05977b3631167028862be2a173976ca11",
    blockCreated: 14353601,
  },
};
