import { Network } from "./types"
import BitcoinIcon from '~/assets/images/networks/bitcoin.png'
import EthereumIcon from '~/assets/images/networks/ethereum.png'
import PolygonIcon from '~/assets/images/networks/polygon.svg'
import BinanceIcon from '~/assets/images/networks/binance.svg'
import OptimismIcon from '~/assets/images/networks/optimism.svg'
import ArbitrumIcon from '~/assets/images/networks/arbitrum.svg'
import MVMIcon from '~/assets/images/networks/mvm.png'

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