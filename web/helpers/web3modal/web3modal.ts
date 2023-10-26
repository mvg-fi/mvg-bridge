import { createWeb3Modal, defaultConfig } from '@web3modal/ethers5/vue'
import { AppDescription, AppFullName, AppURL } from '../constants'

export const initWeb3Modal = () => {
  const projectId = "43b980fb7399c6267103bf2d0aa1ea69"
  const chains = [1, 42161]
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
    themeMode: 'light',
    themeVariables: {
      '--w3m-font-family': "Satoshi-Medium",
      '--w3m-z-index':2500,
    },
    chains,
    projectId
  })
}