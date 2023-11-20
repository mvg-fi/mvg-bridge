import { Network } from "~/types/network"
import BitcoinIcon from '~/assets/images/networks/bitcoin.png'
import EthereumIcon from '~/assets/images/networks/ethereum.png'
import PolygonIcon from '~/assets/images/networks/polygon.svg'
import BinanceIcon from '~/assets/images/networks/binance.svg'
import OptimismIcon from '~/assets/images/networks/optimism.svg'
import ArbitrumIcon from '~/assets/images/networks/arbitrum.svg'
import MVMIcon from '~/assets/images/networks/mvm.png'
import type { Asset } from "~/types/asset"

const BitcoinExplorerURL = 'https://bitcoinexplorer.org/'
const EthereumExplorerURL = 'https://etherscan.io'
const PolygonExplorerURL = 'https://polygonscan.com'
const BinanceExplorerURL = 'https://bscscan.com'
const ArbitrumExplorerURL = 'https://arbiscan.io'
const OptimismExplorerURL = 'https://optimistic.etherscan.io'
const MVMExplorerURL = 'https://scan.mvm.dev'

export const Networks: Network[] = [
  {
    name: 'Bitcoin',
    evmId: -1,
    icon: BitcoinIcon,
    explorerURL: BitcoinExplorerURL,
  },
  {
    name: 'Ethereum',
    evmId: 1,
    icon: EthereumIcon,
    explorerURL: EthereumExplorerURL,
  },
  {
    name: 'Binance',
    evmId: 56,
    icon: BinanceIcon,
    explorerURL: BinanceExplorerURL
  },
  {
    name: 'Polygon',
    evmId: 137,
    icon: PolygonIcon,
    explorerURL: PolygonExplorerURL,
  },
  {
    name: 'Optimism',
    evmId: 10,
    icon: OptimismIcon,
    explorerURL: OptimismExplorerURL,
  },
  {
    name: 'Arbitrum',
    evmId: 42161,
    icon: ArbitrumIcon,
    explorerURL: ArbitrumExplorerURL,
  },
  {
    name: 'MVM',
    evmId: 73927,
    icon: MVMIcon,
    explorerURL: MVMExplorerURL,
  },
]

export const isEVMAsset = (asset: Asset) => {
  // determine if chain id is evm
  switch(asset.chain_id){
  // ETH
  case '43d61dcd-e413-450d-80b8-101d5e903357':
    return true;
  // BNB
  case '1949e683-6a08-49e2-b087-d6b72398588f':
    return true;
  // Polygon
  case 'b7938396-3f94-4e0a-9179-d3440718156f':
    return true;
  default:
    return false;
  }
}