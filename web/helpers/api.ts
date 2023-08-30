// Mixpay
// curl -i -X GET -G https://api.mixpay.me/v1/payments_estimated \
//   -d "paymentAssetId"="c6d0c728-2624-429b-8e0d-d9d19b6592fa" \
//   -d "settlementAssetId"="c6d0c728-2624-429b-8e0d-d9d19b6592fa" \
//   -d "quoteAssetId"="usd" \
//   -d "quoteAmount"="10" 

// 4swap
// curl -i -X POST -H "Content-Type: application/json" \ 
//   https://mtgswap-api.fox.one/api/orders/pre \ 
//   --data '{"pay_asset_id": "c6d0c728-2624-429b-8e0d-d9d19b6592fa", \
//           "fill_asset_id": "6cfe566e-4aad-470b-8c9a-2fd35b49c68d", \
//           "amount": "1"}'