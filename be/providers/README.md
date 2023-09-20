# Swap providers

|Name|Type|Endpoint|Website|
|--|--|--|--|
|Mixpay|CEX|https://api.mixpay.me|https://mixpay.me|
|Exinone|CEX|https://app.eiduwejdk.com/api/v2|https://exinone.com|
|PandoSwap|DEX|https://api.4swap.org|https://pando.im|

## Interface

### GetPrice
return amount and fee
- amount: The amount of asset will be received
- fee: The amount of asset will be charged as fee

### Swap
return a transfer input, and will always be a multisig transfer input.
