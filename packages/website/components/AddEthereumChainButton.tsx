type Props = {
  buttonText: string;
};

async function addEthereumChain() {
  interface AddEthereumChainParameter {
    chainId: string; // A 0x-prefixed hexadecimal string
    chainName: string;
    nativeCurrency: {
      name: string;
      symbol: string; // 2-6 characters long
      decimals: 18;
    };
    rpcUrls: string[];
    blockExplorerUrls?: string[];
    iconUrls?: string[]; // Currently ignored.
  }

  const taikoParams: AddEthereumChainParameter = {
    chainId: "0x28C5A",
    chainName: "Taiko Hackathon Testnet",
    nativeCurrency: {
      name: "ETH",
      symbol: "eth",
      decimals: 18,
    },
    rpcUrls: ["https://l2rpc.hackathon.taiko.xyz"],
    blockExplorerUrls: ["https://l2explorer.hackathon.taiko.xyz/"],
    iconUrls: [],
  };

  await (window as any).ethereum.request({
    method: "wallet_addEthereumChain",
    params: [taikoParams],
  });
}

export default function AddEthereumChainButton(props: Props) {
  return (
    <div
      onClick={() => addEthereumChain()}
      className="hover:cursor-pointer text-neutral-900 bg-white hover:bg-neutral-100 border-solid border-neutral-200 focus:ring-4 focus:outline-none focus:ring-neutral-100 font-medium rounded-lg text-sm px-3 py-2 text-center inline-flex items-center dark:focus:ring-neutral-600 dark:bg-neutral-800 dark:border-neutral-700 dark:text-white dark:hover:bg-neutral-700"
    >
      {props.buttonText}
    </div>
  );
}
