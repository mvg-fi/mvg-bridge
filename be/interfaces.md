# API

Endpoint: https://api.mvg.fi

| Path                             | Method | Usage                        |
| -------------------------------- | ------ | ---------------------------- |
| [`/price/simple`](#price-simple) | POST   | Get the best price           |
| [`/price/all`](#price-all)       | POST   | Get all prices with provider |
| [`/order/new`](#order-new)       | POST   | Create new order             |
| [`/status/{id}`](#order-status)  | GET    | Get order status             |

## Price

### Price simple

POST /price/simple

#### Parameters

- input_asset: The uuid of input asset
- output_asset: The uuid of output asset
- amount: The amount of output_asset (optional, use without except)
- except: The amount of input_asset  (optional, use without amount)

```json
{
    "input_asset": "",
    "output_asset": "",
    "amount": "",                   // optional
    "except": "",                   // optional
}
```

#### Response

- output_amount: This occurs when `amount` is specified in parameters, and represents the amount user will finally receive.
- input_amount: This occurs when `except` is specified in parameters, and represents the amount user to pay to receive excepted amount.

```json
{
    "output_amount": "",            // optional
    "input_amount": "",             // optional
}
```


#### Examples
From bitcoin to ethereum, pay 1.35 BTC, get output_amount of ETH

Parameter:
```json
{
    "input_asset": "c6d0c728-2624-429b-8e0d-d9d19b6592fa",
    "output_asset": "43d61dcd-e413-450d-80b8-101d5e903357",
    "amount": "1.35",
}
```

Response:
```json
{
    "output_amount": "20.556"
}
```

From bitcoin to ethereum, except to have 35.1 ETH, get input_amount of BTC
```json
{
    "input_asset": "c6d0c728-2624-429b-8e0d-d9d19b6592fa",
    "output_asset": "43d61dcd-e413-450d-80b8-101d5e903357",
    "except": "35.1",
}
```

Response:
```json
{
    "input_amount": "1.52348"
}
```

#### Ts interfaces
```ts
interface PriceSimpleReq {
    input_asset: string
    output_asset: string    
    amount?: string
    except?: string
}

interface PriceSimpleResp {
    input_amount?: string
    output_amount?: string
}
```

### Price all

POST /price/all

#### Parameters

- input_asset: The uuid of input asset
- output_asset: The uuid of output asset
- amount: The amount of output_asset (optional, use without except)
- except: The amount of input_asset  (optional, use without amount)
- cex: Enable cex or not

```json
{
    "input_asset": "",
    "output_asset": "",
    "amount": "",                   // optional
    "except": "",                   // optional
    "cex": "",
}
```

#### Response

The response is an array. And will be sorted by the highest price. (if return output_amount, sort the highest, otherwise lowest)

```json
[
    {
        "output_amount": "",        // optional
        "input_amount": "",         // optional
        "provider": "Mixpay",
        "type": "cex",
    },
    {
        "output_amount": "",        // optional
        "input_amount": "",         // optional
        "provider": "Pando swap",
        "type": "dex",
    }
]
```




#### Examples
From bitcoin to ethereum, pay 1.35 BTC, enable CEX

Parameter:
```json
{
    "input_asset": "c6d0c728-2624-429b-8e0d-d9d19b6592fa",
    "output_asset": "43d61dcd-e413-450d-80b8-101d5e903357",
    "amount": "1.35",
    "cex": true,
}
```

Response:
```json
[
    {
        "output_amount": "10.4321",
        "provider": "Mixpay",
        "type": "cex",
    },
    {
        "output_amount": "10.3321",
        "provider": "Pando Swap",
        "type": "dex",
    }
]
```
From bitcoin to ethereum, except to have 35.1 ETH, disable CEX

Parameter:
```json
{
    "input_asset": "c6d0c728-2624-429b-8e0d-d9d19b6592fa",
    "output_asset": "43d61dcd-e413-450d-80b8-101d5e903357",
    "except": "35.1",
    "cex": false,
}
```

Response:
```json
[
    {
        "input_amount": "2.1231",
        "provider": "Pando Swap",
        "type": "dex",
    }
]
```

#### Ts interfaces
```ts
interface PriceAllReq {
    input_asset: string
    output_asset: string    
    amount?: string
    except?: string
    cex: boolean
}

interface PriceAllResp {
    input_amount?: string
    output_amount?: string
    provider: string
    type: string
}
```

## Order

### Order new

POST /order/new

#### Parameters

- input_asset: The uuid of input asset
- output_asset: The uuid of output asset
- amount: The amount of output_asset
- cex: Use CEX or not

#### Response
```json
{
    "trace_id": "",
    "address": "",
    "memo": "",
}
```

#### Examples

#### Ts interfaces
```ts
```


## Status

### Order status

GET /status/{id}

#### Parameters

- trace_id: The uuid returned by /order 

#### Response
```json
{
    "status": "",
}
```

#### Examples

#### Ts interfaces
```ts
```
