# API

Endpoint: https://api.mvg.fi/

## PriceOracle

POST /price

### parameters

- input_asset: The uuid of input asset
- output_asset: The uuid of output asset
- amount: The amount of output_asset (optional, use without except)
- except: The amount of input_asset  (optional, use without amount)

### response

```json
{
    "output_amount": "",        // optional
    "input_amount": "",        // optional
}
```

### examples
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

### Ts interfaces
interface PriceInput {
    input_amount: string
    output_amout: stirng    
    amount?: string
    except?: string
}

interface PriceOutput {
    input_amount?: string
    ouput_amount?: string
}


## Order

POST /order

### parameters

- input_asset: The uuid of input asset
- output_asset: The uuid of output asset
- amount: The amount of output_asset

### response
```json
{
    "trace_id": "",
    "deposit_address": "",
}
```

### examples

### Ts interfaces



## Status

GET /status

### parameters

- trace_id: The uuid returned by /order 

### response
```json
{
    "status": "",
}
```

### examples

### Ts interfaces
