### **Wallet Service Documentation**

---

#### **Overview**

The Wallet Service provides API for managing digital wallets. It allows businesses to perform core wallet operations, including wallet creation, fund management, and transaction history retrieval. Built using Hexagonal Architecture, DDD, CQRS, and Golang with Echo, the service is modular and scalable.

> **Note:** This is an **experimental project** created as an example of how to implement **CQRS** with **Hexagonal Architecture** using Golang. It demonstrates best practices for structuring a domain-driven application in a microservices context.

---

### **Core Features**

1. **Wallet Management**
    - Create wallets with unique identifiers.
    - Store and manage balance information.
    - Handle wallet statuses (e.g., Active, Blocked).

2. **Fund Management**
    - Add funds to a wallet.
    - Deduct funds from a wallet.
    - Ensure transactional consistency and validations.

3. **Transaction History**
    - Record all wallet transactions (credits and debits).
    - Retrieve paginated transaction histories.

---

### **API Endpoints**

| **HTTP Method** | **Endpoint**                          | **Description**               |
|------------------|---------------------------------------|-------------------------------|
| `POST`           | `/api/v1/wallets`                   | Create a new wallet.          |
| `POST`           | `/api/v1/wallets/:walletID/funds/add` | Add funds to the wallet.      |
| `POST`           | `/api/v1/wallets/:walletID/funds/deduct` | Deduct funds from the wallet. |
| `GET`            | `/api/v1/wallets/:walletID`          | Retrieve wallet details.      |
| `GET`            | `/api/v1/wallets/:walletID/transactions` | Get wallet transaction history.|

---

### **Wallet Lifecycle**

1. **Create Wallet**: Initialize a wallet with an owner, currency, and starting balance.
2. **Add or Deduct Funds**: Modify the wallet balance while maintaining a transaction log.
3. **View Wallet Details**: Fetch wallet metadata and balance.
4. **Transaction History**: Access all transactions associated with a wallet for reporting.

---

### **Architecture Highlights**

- **Hexagonal Architecture**: Clean separation between business logic and external systems.
- **CQRS**: Commands handle state changes, while queries handle data retrieval.
- **DDD**: Focus on aggregates like Wallet and Transaction for domain consistency.
- **Echo Framework**: Lightweight and performant web framework for Golang.

---

### **Key Entities**

1. **Wallet**:
    - **Attributes**: ID, OwnerID, Balance, Currency, Status, Timestamps.
    - **Responsibilities**: Balance updates, status management.

2. **Transaction**:
    - **Attributes**: ID, WalletID, Amount, Type (Credit/Debit), Status, Timestamp.
    - **Responsibilities**: Maintain transaction details and history.

---

### **Example Use Case**

1. A user creates a wallet for their account.
2. The wallet receives funds through the `Add Funds` endpoint.
3. The user deducts funds to make a purchase.
4. The user views their wallet balance and transaction history.

---

This service is an **experimental project** designed as a learning example of implementing **CQRS** with **Hexagonal Architecture** in Golang. It demonstrates best practices in applying domain-driven design patterns and is suitable for educational purposes or as a starting point for similar production-grade systems.
