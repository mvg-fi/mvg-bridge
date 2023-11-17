import { createWeb3Modal, defaultConfig } from '@web3modal/ethers5/vue'
import { AppDescription, AppFullName, AppURL } from '../constants'

export const initWeb3Modal = () => {
  const projectId = "43b980fb7399c6267103bf2d0aa1ea69"

  const mainnet = {
    chainId: 1,
    name: 'Ethereum',
    currency: 'ETH',
    explorerUrl: 'https://etherscan.io',
    rpcUrl: 'https://eth.llamarpc.com'
  }
  const mvm = {
    chainId: 73927,
    name: 'MVM',
    currency: 'ETH',
    explorerUrl: 'https://scan.mvm.dev',
    rpcUrl: 'https://geth.mvm.dev'
  }

  const chains = [mainnet, mvm]
  const metadata = {
    name: AppFullName,
    description: AppDescription,
    url: AppURL,
    icons: []
  }

  createWeb3Modal({
    ethersConfig: defaultConfig({
      metadata,
      defaultChainId: 1,
      enableEIP6963: true,
      enableInjected: true,
      enableCoinbase: true,
      rpcUrl: ''
    }),
    featuredWalletIds: [
      'c57ca95b47569778a828d19178114f4db188b89b763c899ba0be274e97267d96', // Metamask
      '18388be9ac2d02726dbac9777c96efaac06d744b2f6d580fccdd4127a6d01fd1', // Rabby
      'fd20dc426fb37566d803205b19bbc1d4096b248ac04548e3cfb6b3a38bd033aa', // Coinbase
      '1ae92b26df02f0abca6304df07debccd18262fdf5fe82daa81593582dac9a369', // Rainbow
      'ef333840daf915aafdc4a004525502d6d49d77bd9c65e0642dbaefb3c2893bef', // ImToken
    ],
    themeMode: 'light',
    themeVariables: {
      '--w3m-font-family': "Satoshi-SemiBold",
      '--w3m-z-index': 2500,
    },
    chains,
    projectId
  })
}