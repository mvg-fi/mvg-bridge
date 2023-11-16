export interface Network {
  name: string,
  icon: string,
  evmId: number,
  rpcURL?: string,
  explorerURL: string,
}

export interface EVMNetworkParam {
	chainId: string;
	rpcUrls: string[];
	chainName: string;
	nativeCurrency: { name: string; decimals: number; symbol: string };
	blockExplorerUrls: string[];
	iconUrls?: string[];
}