# MVG Bridge MTG

- Syncer
In charge of updating UTXO list.
- 

- OrderWorker
In charge of order creation, address creation, provides addresses query

- DepositWorker
In charge of monitoring deposit addresses, provides deposit states

- SwapWorker
In charge of asset exchange, provides swap states

- WithdrawalWorker
In charge of initiating withdrawals, provides withdrawal states




## TODO

1. 写充值API，创建充值地址，关联订单信息，监控充值
2. 写Swap worker，创建swap，监控swap，设置提现状态
3. 写提现worker，监控提现订单，完成提现


1. 匿名创建订单，有效期10分钟，包含地址 
