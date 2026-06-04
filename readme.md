### tejascoin

- create project structure.
- create blockchain with genesis.
- now move to mining and wallate building.

 Bitcoin miners repeatedly try different nonces until the hash satisfies a difficulty rule.

```
Example difficulty:
0000xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
The hash must start with four zeros.
```

## what we have now.

✓ Genesis Block
✓ Transactions
✓ Blockchain
✓ SHA256 Hashing
✓ Mining
✓ Proof of Work

we have build till now :
```
Bitcoin Block
    +
SHA256
    +
Proof of Work
```

starting creating wallet ->

in transaction we have from(bob) and to(jon) properties, but bitcoin does not know who is bob and jon.

- trnasaction use addresses not name.

each wallet link to perticular entity(person). wallet has own private key, public key and address.

flow of Address creation :

Generate Private Key
          ↓
Generate Public Key
          ↓
Hash Public Key
          ↓
Address


flow of signed transaction varification.

```
Alice Wallet
    │
    ├── Create transaction
    │
    ├── Sign transaction using private key
    │
    ▼
Signed Transaction
    │
    ▼
Blockchain verifies signature
    │
    ▼
Accepted
```

### Balance

- how many coins does address own?
    - this is foundation of cryptocurrency.


## Minnig Rewards 
- first we need coins to exist.
- bitcoin creats coins through mining.

```
Mine Block
    ↓
Reward Miner
    ↓
+50 BTC
```

in cryptocurrencies, a special transaction creates coins.
coins are created my protocol rules.

## Validate Transaction

current flow :
```
Create Transaction
      ↓
Mine Block
      ↓
Update Balances
```


New flow :

```
Create Transaction
      ↓
Validate Transaction
      ↓
Mine Block
      ↓
Update Balances
```


AddBlock
    │
    ├── Validate Tx 1
    ├── Validate Tx 2
    ├── Validate Tx 3
    │
    ▼
Mine Block
    │
    ▼
Append Block


Next Step: Mempool (Pending Transactions)

the flow will becomes:

```
Create Transaction
      ↓
Mempool
      ↓
Miner Selects Transactions
      ↓
Mine Block
```

- a mempool is just waiting area.

Example :- 
'''
Alice -> Bob      10
Charlie -> David   5
Bob -> Eve         3
'''

all sit in memory area.
- PendingTransactions []Transaction

Then a miner picks them up and creates a block.


till now blockchain lifecycle :

```
SYSTEM creates coins
       ↓
Alice owns coins
       ↓
Alice creates transaction
       ↓
Transaction enters mempool
       ↓
Miner mines block
       ↓
Miner gets reward
       ↓
Balances updated
```

a real cryptocurrency transaction contians :

```
From
To
Amount
Public Key
Signature
```

the node varifies :

```
Public Key
      ↓
Produces Address?
      ↓
YES

Signature Valid?
      ↓
YES

Accept Transaction
```

we need block chain to varify :

```
Transaction
     ↓
Public Key
     ↓
Address Match?
     ↓
YES

Signature Valid?
     ↓
YES

Accept
```

current state :

```
✓ Blockchain
✓ Blocks
✓ Mining
✓ Proof of Work
✓ Wallets
✓ Addresses
✓ Balances
✓ Mempool
✓ Mining Rewards
✓ Transaction Signing
✓ Signature Verification
✓ Tamper Detection
```

- now validation becomes :

Public Key
      ↓
create Address from pub key
      ↓
Matches tx.From ?
      ↓
YES

Verify Signature
      ↓
YES

Accept
