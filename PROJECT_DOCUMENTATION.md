# TejasCoin - Project Documentation

**Project Name:** TejasCoin  
**Repository:** github.com/tejas/tejascoin  
**Language:** Go 1.26.2  
**Last Updated:** June 2, 2026  
**Status:** Active Development - Phase 1: Core Blockchain Implementation

---

## 1. Project Overview

### 1.1 Purpose
TejasCoin is a blockchain learning project that implements the fundamental concepts of cryptocurrency and distributed ledger technology. It demonstrates core blockchain mechanics including blocks, mining, proof-of-work, and transaction handling.

### 1.2 Goals
- ✓ Build a functional blockchain with genesis block
- ✓ Implement transaction support
- ✓ Create SHA256 hashing mechanism
- ✓ Develop Proof of Work (PoW) mining algorithm
- ✓ Implement wallet functionality with key generation (In Progress - Phase 1 Complete)
- ⏳ Build REST API endpoints (Planned)

### 1.3 Scope
This is an educational blockchain implementation focused on core concepts rather than production-grade security and scalability.

---

## 2. Technical Architecture

### 2.1 System Architecture

```
TejasCoin Architecture
├── Core Components
│   ├── Block Structure
│   ├── Blockchain Chain
│   ├── Transaction Management
│   ├── Hashing (SHA256)
│   └── Proof of Work (Mining)
├── API Layer (Planned)
├── Wallet Layer (In Progress)
└── CLI Interface
```

### 2.2 Blockchain Flow

```
User Creates Transaction
        ↓
Transaction Added to Block
        ↓
Mining Process Started (PoW)
        ↓
Block Hash Calculated with Difficulty Target
        ↓
Nonce Incremented Until Hash Matches Target
        ↓
Block Added to Blockchain
        ↓
Next Block Can Reference Previous Hash
```

---

## 3. Project Structure

```
tejascoin/
├── cmd/
│   └── main.go              # Application entry point
├── internal/
│   ├── api/                 # REST API endpoints (Planned)
│   ├── blockchain/
│   │   ├── block.go         # Block data structure
│   │   ├── blockchain.go    # Blockchain management
│   │   ├── hash.go          # SHA256 hashing
│   │   ├── mining.go        # Proof of Work implementation
│   │   └── transaction.go   # Transaction structure
│   └── wallet/
│       ├── wallet.go        # Wallet and key management
│       └── sign.go          # Transaction signing and verification
├── go.mod                   # Go module definition
├── go.sum                   # Dependency checksums
├── readme.md                # Project overview
└── PROJECT_DOCUMENTATION.md # This file
```

---

## 4. Core Components

### 4.1 Block Structure (`internal/blockchain/block.go`)

**Purpose:** Defines the data structure for a single block in the blockchain.

| Field | Type | Description |
|-------|------|-------------|
| `Index` | int | Block position in the chain |
| `Timestamp` | string | When the block was created |
| `Transactions` | []Transaction | Array of transactions in the block |
| `PrevHash` | string | SHA256 hash of the previous block |
| `Hash` | string | SHA256 hash of the current block |
| `Nonce` | int | Number used once in PoW algorithm |

**Key Points:**
- Nonce is incremented during mining to find valid hash
- Each block references the previous block's hash (immutability)
- JSON tags for serialization support

### 4.2 Blockchain Management (`internal/blockchain/blockchain.go`)

**Purpose:** Manages the chain of blocks and block operations.

**Key Functions:**
- `NewBlockchain()` - Initializes blockchain with genesis block
- `AddBlock(transactions []Transaction)` - Adds new block with transactions

**Genesis Block:**
- Index: 0
- Timestamp: "GENESIS"
- PrevHash: "" (empty, no previous block)
- Nonce: 0
- Hash: Calculated at initialization

### 4.3 Transaction Structure (`internal/blockchain/transaction.go`)

**Purpose:** Defines the transaction data structure.

```go
type Transaction struct {
    From   string  // Sender identifier
    To     string  // Receiver identifier
    Amount float64 // Transaction amount
}
```

**Properties:**
- Simple transfer model with sender, receiver, and amount
- Support for JSON serialization
- Foundation for wallet integration

### 4.4 SHA256 Hashing (`internal/blockchain/hash.go`)

**Purpose:** Generates consistent cryptographic hashes for blocks.

**Algorithm:** SHA256
- Input: Block data (index, timestamp, transactions, previous hash, nonce)
- Output: 64-character hexadecimal string
- Used for: Block identification and chain integrity verification

### 4.5 Proof of Work Mining (`internal/blockchain/mining.go`)

**Purpose:** Implements the mining algorithm with adjustable difficulty.

**Algorithm:**
1. Set difficulty target (leading zeros required in hash)
2. Loop until valid hash found:
   - Calculate SHA256 hash of block
   - Check if hash starts with required zeros
   - Increment nonce if not valid
   - Return block when valid hash found

**Current Difficulty:** 4 leading zeros

**Example Difficulty Target:**
```

### 4.6 Wallet System (`internal/wallet/wallet.go`)

**Purpose:** Manages cryptographic keys and wallet addresses for users.

**Key Components:**
- `PrivateKey` - ECDSA private key for signing transactions
- `PublicKey` - ECDSA public key for verification
- `Address` - SHA256 hash of public key (wallet identifier)

**Wallet Generation Process:**
1. Generate ECDSA key pair using P256 elliptic curve
2. Extract public key bytes
3. Hash public key with SHA256
4. Encode hash as hex string to create address
5. Return Wallet with all components

**Cryptography:**
- Algorithm: ECDSA (Elliptic Curve Digital Signature Algorithm)
- Curve: P256 (secp256r1)
- Address Format: 64-character hexadecimal string
- Security Level: 256-bit

**Example Usage:**
```go
wallet, err := wallet.NewWallet()
if err != nil {
    log.Fatal(err)
}
fmt.Println(wallet.Address) // Output: hex-encoded address
```

### 4.7 Transaction Signing & Verification (`internal/wallet/sign.go`)

**Purpose:** Enables transaction authentication through digital signatures.

**Functions:**

**Sign()**
- Input: Private key, hash of transaction data
- Output: Digital signature (ASN1 encoded)
- Purpose: Sign transactions to prove ownership

**Verify()**
- Input: Public key, transaction hash, signature
- Output: Boolean (true if signature valid)
- Purpose: Verify transaction was signed by owner

**Implementation Details:**
- Uses ECDSA ASN1 format for signature encoding/decoding
- Cryptographically secure random number generation
- Compatible with blockchain signature standards

**Example Flow:**
```
Transaction Created
    ↓
Calculate Hash of Transaction
    ↓
Sign Hash with Private Key
    ↓
Add Signature to Transaction
    ↓
Receiver Verifies with Public Key
    ↓
Transaction Authenticated
```
0000xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

---

## 5. Technical Stack

### 5.1 Core Dependencies

| Package | Version | Purpose |
|---------|---------|---------|
| Go | 1.26.2 | Programming language |
| Fiber | v3.3.0 | Web framework (for API) |
| google/uuid | v1.6.0 | UUID generation |
| golang/crypto | v0.51.0 | Cryptographic functions |

### 5.2 Dependency Tree

```
github.com/tejas/tejascoin
├── github.com/gofiber/fiber/v3
│   ├── github.com/valyala/fasthttp
│   └── github.com/gofiber/schema
├── golang.org/x/crypto
└── github.com/google/uuid
```

---

## 6. Current Implementation Status

### 6.1 Completed Features

| Feature | Status | Description |
|---------|--------|-------------|
| Project Structure | ✓ | Organized with cmd, internal packages |
| Genesis Block | ✓ | Initial block creation |
| Block Structure | ✓ | Defined with all required fields |
| Transaction Structure | ✓ | From, To, Amount |
| SHA256 Hashing | ✓ | Cryptographic hash implementation |
| Proof of Work | ✓ | Mining with adjustable difficulty |
| Block Addition | ✓ | Adding blocks to chain with PoW |
| ECDSA Key Generation | ✓ | Wallet address generation using P256 |
| Transaction Signing | ✓ | ECDSA ASN1 signature creation |
| Transaction Verification | ✓ | ECDSA signature verification |
| Chain Validation | Partial | Basic structure in place |

### 6.2 In Progress

- **Transaction Pool** - Mempool for pending transactions
- **Transaction Integration** - Connect wallet signing to transactions
- **REST API** - Endpoints for blockchain interaction

### 6.3 Not Started

- Consensus mechanisms (beyond basic PoW)
- Network synchronization
- Smart contracts
- Advanced difficulty adjustment
- Full transaction signing integration
- Transaction pool/mempool

---

## 7. Build & Execution

### 7.1 Prerequisites
- Go 1.26.2 or higher
- Linux/macOS/Windows environment

### 7.2 Build Instructions

```bash
# Navigate to project directory
cd /home/tntra/Documents/blockchain_practice/tejascoin

# Download dependencies
go mod download

# Build executable
go build -o bin/tejascoin ./cmd/main.go
```

### 7.3 Run Instructions

```bash
# Direct run (development)
go run ./cmd/main.go

# Run from built binary
./bin/tejascoin
```

### 7.4 Expected Output

```
Blockchain created
{Index:0 Timestamp:GENESIS Transactions:[] PrevHash: Hash:xxxx... Nonce:0}

{Index:1 Timestamp:2026-06-02... Transactions:[{From:Alice To:Bob Amount:10}] PrevHash:xxxx... Hash:0000xxxx... Nonce:xxxx}
```

---

## 8. Development Workflow

### 8.1 Current Testing Approach

**Main Test Case** (`cmd/main.go`):
1. Create new blockchain (genesis block created)
2. Create transaction (Alice → Bob, 10 units)
3. Add block with transaction (mining occurs)
4. Display all blocks

### 8.2 Code Organization

**Modular Structure:**
- Each blockchain component in separate file
- `package blockchain` for core logic
- `package main` for entry point
- Separation of concerns: hashing, mining, transactions

---

## 9. Performance Metrics

### 9.1 Mining Performance

| Difficulty | Target | Avg. Iterations | Time Impact |
|------------|--------|-----------------|-------------|
| 1 | `0` | ~16 | Minimal |
| 2 | `00` | ~256 | <1ms |
| 3 | `000` | ~4,096 | ~5ms |
| 4 | `0000` | ~65,536 | ~100-200ms |

**Current Setting:** Difficulty 4 (0000 prefix requirement)

---

## 10. API Endpoints (Planned)

### 10.1 Proposed REST API

```
GET  /api/blockchain       - Get entire blockchain
GET  /api/blocks           - Get all blocks
GET  /api/blocks/:index    - Get specific block
POST /api/transactions     - Submit new transaction
POST /api/mine             - Mine new block
GET  /api/chain/validate   - Validate chain integrity
```

---

## 11. Development Progress Tracker

### 11.1 Phase 1: Core Blockchain (CURRENT)
- [x] Project structure setup
- [x] Block structure definition
- [x] Genesis block creation
- [x] Transaction support
- [x] SHA256 hashing
- [x] Proof of Work implementation
- [x] Block addition with mining
- [ ] Block validation
- [ ] Chain integrity verification
- [ ] Difficulty adjustment

### 11.2 Phase 2: Wallet System (NEXT)
- [ ] Address generation (public key cryptography)
- [ ] Private key management
- [ ] Transaction signing
- [ ] Signature verification
- [ ] Balance tracking

### 11.3 Phase 3: API Layer
- [ ] REST API setup with Fiber
- [ ] Endpoint implementation
- [ ] Request/response handling
- [ ] Error handling
- [ ] Input validation

### 11.4 Phase 4: Advanced Features (FUTURE)
- [ ] Mempool for transaction pooling
- [ ] Network synchronization
- [ ] Consensus mechanism refinement
- [ ] Smart contract foundation
- [ ] Database persistence

---

## 12. Code Quality Guidelines

### 12.1 Naming Conventions
- **Packages:** lowercase, concise (blockchain, wallet, api)
- **Functions:** CamelCase, PascalCase for exports (NewBlockchain, AddBlock)
- **Variables:** camelCase for unexported, meaningful names
- **Constants:** UPPER_SNAKE_CASE

### 12.2 Documentation Standards
- Package-level comments for each file
- Function comments for public functions
- Inline comments for complex logic
- Type documentation with purpose

### 12.3 Code Structure
- Single responsibility per file
- Clean separation of concerns
- Modular design for reusability
- Consistent formatting (go fmt)

---

## 13. Common Issues & Solutions

| Issue | Cause | Solution |
|-------|-------|----------|
| High mining time | High difficulty | Reduce difficulty value in mining.go |
| Import errors | Missing go mod | Run `go mod download` |
| Version conflicts | Outdated Go | Update to Go 1.26.2+ |
| Nonce overflow | Extended mining | Consider bitmask operations |

---

## 14. Resources & References

### 14.1 Blockchain Concepts
- **Proof of Work:** Mining algorithm that requires computational work
- **SHA256:** Cryptographic hash function producing 256-bit output
- **Nonce:** "Number used once" - varied to find valid hash
- **Genesis Block:** First block in blockchain, has no parent

### 14.2 Bitcoin References
The mining mechanism is inspired by Bitcoin's proof-of-work system where miners compete to find hashes meeting difficulty requirements.

---

## 15. Future Roadmap

### Q3 2026
- Complete wallet implementation
- Implement transaction signing and verification

### Q4 2026
- REST API fully operational
- CLI interface improvements
- Database persistence

### Q1 2027
- Smart contracts foundation
- Network protocol design
- Consensus refinement

---

## 16. Contact & Contribution

**Project Author:** Tejas  
**Learning Purpose:** Educational blockchain implementation  
**Status:** Active Development

**Development Environment:**
- OS: Linux
- IDE: VS Code
- Version Control: Git
- Language: Go 1.26.2

---

## 17. Document History

| Version | Date | Changes |
|---------|------|---------|
| 1.0 | 2026-06-02 | Initial project documentation |

---

**End of Document**

Generated for tracking and reference purposes following industry documentation standards.
