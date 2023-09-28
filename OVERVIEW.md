# MVG Bridge V2

A powerful cross chain bridge that connects all chains and all tokens.

## Motivation

Connect all the chains, all the tokens, and making cross chain experience flawless.

## Types

There will be two types of bridges:

1. The bridge for MVM Mainnet

2. The bridge that connects all chains

There will be a switch inside settings that allow people to switch between these two types of bridges. For the first type, people can deposit and withdraw to/from MVM mainnet

## Limitation

For now, the MTG doesn't support direct deposit and withdrawal. The current design will be updated once the MTG supports them.

## Use cases

### 1. Cross chain swap

- From Bitcoin to USDT-ERC20. 

User inputs the amount of Bitcoin, gets excepted amount of USDT-ERC20 to receive, and inputs the address for receiving the USDT-ERC20, gets a deposit address of bitcoin. This is an address owned by MVG Bridge MTG, which the MTG monitors all the deposits and initiates the swap and withdrawal processes. User can see the current state of the swap/withdrawal process at any time in MVG Bridge's front end. This should be implemented by MTG's API.

- From USDT-ERC20 to USDT-Matic

Both token are EVM-compatible, which means the user can make the transaction directly and see the result through the frontend. They will also be able to specify an address to receive the fund. But by default the recipient will be the account connected. The rest of the parts are quite same as above.

- From USDT-ERC20 to USDT-Tron

Tron is not EVM-compatible, the user will have to enter an address for receiving the fund. This will applies to all non-EVM chains. The rest of the parts are quite same as above.

### 2. Bridge into MVM Mainnet

User will connect to their EVM wallet to begin.

- Deposit 

There will be three types of deposit:
1. deposit from EVM compatible chains (e.g. Ethereum, Polygon, Arbitrum, ...)
2. deposit from non-EVM chains (e.g. Bitcoin, Monero, ...)
3. deposit from Mixin Wallet (e.g. Mixin Messenger, Fennec, MultiSig Wallet, bots)

For the 1st type, the user can deposit directly from their EVM Wallet within a few clicks.
For the 2nd type, the user will need to transfer to the address provided by the MVG Bridge front end. And the front end will have an active pending-deposit checking queue. Make sure the user can see the exact status of their asset at any moment.
For the 3rd type, the user will be able to scan QRCode (Mixin Messenger), sign transfer message (Fennec), get deposit address (Multisig or Bot). For updating the deposit process, trace the snapshot and tract id, and the transaction hash.

- Withdrawal

There will also be three types of withdraw:
1. Withdraw to EVM compatible chains (e.g. Ethereum, Polygon, Arbitrum, ...)
2. Withdraw to non-EVM chains (e.g. Bitcoin, Monero, ...)
3. Withdraw to Mixin Wallet (e.g. Mixin Messenger, Fennec, MultiSig Wallet, bots)

For the 1st type, the user can also withdraw directly within a few clicks.
For the 2nd type, the user will have to provide an address for receiving the fund.
For the 3rd type, the user can withdraw to Mixin Messenger through OAuth, to others through client id or a list of client id.

## Swap service

### In Mixin Ecosystem

Currently (2023-06-19), there're two exchange service provider, each represents CEX and DEX.

- MixPay (CEX)
- 4Swap (DEX)

The MVG Bridge MTG will compare and choose the provider that provides the best price & the lowest slippage. In the future, more services will be integrated if needed.

### In Other Ecosystem

The DEX with the best liquidity and able to provides the best price & lowest slippage will be integrated into the MVG Bridge.

## Privacy

The user will be able to disable centralized swap service provider in the settings for the potential privacy concerns.

Also, the MTG will only collects the information needed to accomplish the swap/withdrawal job. The swap/withdrawal record will be kept 7 days for possible incidents, then will be deleted permanently from the database.

## Mixin Safe integration

To increase security and avoid further possible attacks, the integration with Mixin Safe would be built in. The technical details will be settled as soon as Mixin Safe supports more blockchains. 

## License

This project is under GPL-3.0 License.
