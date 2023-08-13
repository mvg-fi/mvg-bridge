# MVG Bridge MTG

Here is the technical design of MVG Bridge backend (MTG)




## Technical Design

### Workers to have

- Syncer
In charge of updating UTXO list.

- OrderWorker
In charge of order creation, address creation, provides addresses query

- DepositWorker
In charge of monitoring deposit addresses, provides deposit states

- SwapWorker
In charge of asset exchange, provides swap states

- WithdrawalWorker
In charge of initiating withdrawals, provides withdrawal states


### Steps

1. Build an API for creating addresses

1.1 Binding of deposit address and task

Format: (prefix|address) to (chainid|assetid|toaddress|memo|recvamount) 

Prefix: PENDING|CS (Cross Chain Swap) / SS (Same Chain Swap) / BE (Bridge)

e.g. 
BTC to ERC-20 USDT 
Key: PENDING|CS|bc1qau86e05syn63tfuaz7gxj8w50mtefhegez2gp2|20230801090703x1
Value: 43d61dcde413450d80b8101d5e903357|4d8c508b91c5375b92b0ee702ed2dac5|0x5ddfBdCe220A4951a624ccAdb7a267c0877ff86c||1000


TRON to TRC-20 USDT
Key: PENDING|SS|TNq2ya5DhZeAvtpCoBHiV97stgfUXqtKWa|20230801090703x1
Value: 25dabac5056a48ffb9f9f67395dc407c|b91e18ffa9ae3dc78679e935d9a4b34b|TNq2ya5DhZeAvtpCoBHiV97stgfUXqtKWa|x|1000


Ethereum ETH to Arbitrum ETH
Key: PENDING|BE|0x1AE60D36412a6745fce4d4935FF5Bf2b8139a371|20230801090703x1
Value: 25dabac5056a48ffb9f9f67395dc407c|25dabac5056a48ffb9f9f67395dc407c|0x1AE60D36412a6745fce4d4935FF5Bf2b8139a371||1000


ERC20-USDT to ETH
handled by the frontend. Using 1inch, Uniswap, Curve, Dodo or Mixpay.



How to determine tx type:

CS: ChainID1 != ChainID2
SS: ChainID1 == ChainID2
BE: SAMECHAINMAP[AssetID]=SameChainAssetList AssetID1 in SameChainAssetList and AssetID2 in SameChainAssetList

1.1 Timeout worker

In charge of cleaning pending deposits. Get it out of the list.

Prefix: TIMEOUT|


1.2 Deposit worker

In charge of monitoring PENDING deposits. Once received, change the state to RECEIVED

Prefix: RECEIVED|
Key: RECEIVED|CS|bc1qau86e05syn63tfuaz7gxj8w50mtefhegez2gp2|20230801090703x1
Key: RECEIVED|BE|bc1qau86e05syn63tfuaz7gxj8w50mtefhegez2gp2|20230801090703x1
Value: 43d61dcde413450d80b8101d5e903357|4d8c508b91c5375b92b0ee702ed2dac5|0x5ddfBdCe220A4951a624ccAdb7a267c0877ff86c||1000

1.3 Swap worker

If Prefix is RECEIVED and Swap needed, Call Swap and change the state to SWAPINITED
If swap is not needed, Call withdrawal and change the state to WITHDRAWINITED

Prefix: SWAPINITED|
Key: SWAPINITED|CS|bc1qau86e05syn63tfuaz7gxj8w50mtefhegez2gp2|20230801090703x1
Value: 43d61dcde413450d80b8101d5e903357|4d8c508b91c5375b92b0ee702ed2dac5|0x5ddfBdCe220A4951a624ccAdb7a267c0877ff86c||1000
Key: SWAPINITED|BE|0x1AE60D36412a6745fce4d4935FF5Bf2b8139a371|20230801090703x1
Value: 43d61dcde413450d80b8101d5e903357|4d8c508b91c5375b92b0ee702ed2dac5|0x5ddfBdCe220A4951a624ccAdb7a267c0877ff86c||1000


1.4 Swap monitor worker
Constantly checking SWAPINITED transactions, if the broker API returns success, change the INITED state to WITHDRAWINITED, SWAPFAILED otherwise,

Prefix: WITHDRAWINITED|
Prefix: SWAPFAILED|
Key: WITHDRAWINITED|CS|bc1qau86e05syn63tfuaz7gxj8w50mtefhegez2gp2|20230801090703x1
Value: 43d61dcde413450d80b8101d5e903357|4d8c508b91c5375b92b0ee702ed2dac5|0x5ddfBdCe220A4951a624ccAdb7a267c0877ff86c||1000
Key: SWAPFAILED|BE|bc1qau86e05syn63tfuaz7gxj8w50mtefhegez2gp2|20230801090703x1
Value: 43d61dcde413450d80b8101d5e903357|4d8c508b91c5375b92b0ee702ed2dac5|0x5ddfBdCe220A4951a624ccAdb7a267c0877ff86c||1000


1.5 Withdrawal worker
Constantly checking WITHDRAWINITED transactions

Prefix: WITHDRAWSUCCESS|
Prefix: WITHDRAWFAILED|
Key: WITHDRAWSUCCESS|CS|bc1qau86e05syn63tfuaz7gxj8w50mtefhegez2gp2|20230801090703x1
Value: 43d61dcde413450d80b8101d5e903357|4d8c508b91c5375b92b0ee702ed2dac5|0x5ddfBdCe220A4951a624ccAdb7a267c0877ff86c||1000
Key: WITHDRAWFAILED|BE|bc1qau86e05syn63tfuaz7gxj8w50mtefhegez2gp2|20230801090703x1
Value: 43d61dcde413450d80b8101d5e903357|4d8c508b91c5375b92b0ee702ed2dac5|0x5ddfBdCe220A4951a624ccAdb7a267c0877ff86c||1000


1.6 Failure monitor worker

In charge of sending failed transaction in to Mixin Messenger group. For failed transactions, retry first, if fails again, send message.

1.7 Privacy worker (cleaner)

In charge of removing old transaction histories to protect privacy. 

Also, a mixer should be considered.


1.8 Status API

Expose an API to get the status of the order.

Pending | Timeout | Received | Initialed | Failed | Success |



2. Private transaction
















