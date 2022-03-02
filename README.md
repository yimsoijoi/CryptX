# cryptx
cryptx stands between future transactions, hopefully facilitating your untrusted trades.

# Overview
![CryptxDiagram](https://github.com/yimsoijoi/cryptx/blob/main/cryptx.png?raw=true)

## Why?
*Just for fun, obviously*. I wanted to build a crypto version of PayPal, as one of my very first programming projects. Although cryptx in itself has no actual use, it helped me learn software architecture and implementation, e.g. API and database designs, and most importantly, the hyped Ethereum blockchain.

The rationale behind it was that PayPal was a jerk for charging insane amount of fees just to connect buyer and seller accounts. Recently in Thailand, PayPal Thailand made absurd changes to its policies regarding overseas transaction, greatly hindering the laymen's ability to receive money from aboard. On blockchains, such insufficient intermediaries are almost nonexistent.

This is why I wanted to some how design something like PayPal, but with no fees charged, and no nonesense restrictions. However, I'm just a single developer just starting learning how to code, so the code is not *polished* or made for normal end users - it's just meant to be a proof of concept that would also help me learn Go.

## How cryptx works
1. Buyer creates an order by hitting one of the APIs. The order information includes token data, destination wallet data, the pay deadline, and the amount. The default deadline for an order is 21 days.

2. Buyer deposits token into cryptx central wallet.

3. Seller performs trade obligations.

4. Buyer, satisfied with the delivery of goods or services, triggers the payment from the central wallet to the seller wallet.

5. If the buyer does not trigger the pay, the money remains in cryptx central wallet, and if the buyer fails to give reasonable notice as to why the pay is not triggered, cryptx automatically honors the seller after the deadline.

## cryptx API
cryptx exposes the following API endpoints:

- `POST /orders` is for creating new orders

- `POST /orders/pay/:uuid` is for triggering ERC20 transfers

- `GET /orders/` is for getting orders

- `GET /orders/:uuid` is for getting a specific order

With these API endpoints, the seller can also verify that their future payment is safe with cryptx.

## WIP
- The endpoint `/order/pay/:uuid` is still unsafe - there's still no verification.

- cryptx does not yet verify if the buyer has indeed deposit their money into the central wallet.