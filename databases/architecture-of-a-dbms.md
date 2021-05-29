# Architecture of a Database System

From [this paper](https://dsf.berkeley.edu/papers/fntdb07-architecture.pdf), written by Joseph M. Hellerstein, Michael Stonebraker and James Hamilton

Other interesting articles:

- [How to write a DB](https://cstack.github.io/db_tutorial/)
- Writing the worst SQL DB on earth/This is the worst DB on earth and I wrote it

## Life of a query

### 1. The client (like a user's pc) issues a query over the network.

First, a connection needs to be established, so the DBMS needs to be able to communicate in various protocols with different clients. Like, the the PC directly (2-tier communication, or client-server), or a middle tier like a web server (3-tier, or more).
It also needs to be able to **remember** this connection.
Finally, it needs to remember the query itself and be able to transmit it internally.

### 2. Process management

The DBMS needs to assign a **thread of computation** to this query and decide if the operation can be performed now or later (for instance, if there's not enough memory to perform the operation now)

### 3. Relational Query Processor

Connection is established and thread is assigned, it's time to process the query.
The SQL query is compiled into a **query plan**.

```
SQL query --> get translated into --> Query plan
```

The query plan consists of a series of operations such as:

- selection
- join
- projection
- sorting
- etc

More about [query planning here](https://www.sqlite.org/queryplanner.html).

### 4. Transactional Storage Manager

The TSM is the part of the system that is in charge of actually retrieving and writing the data to disk.
It gets the locks from the lock manager to perform operations in safety in regards to other concurrent calls.
It also ensures atomicity of operation, so it is responsible for rolling back partial operations that can't be completed.

### 5. Unwinding of the operations

At this point data is accessed on disk and can be sent back to the client, basically walking backwards in the chain of events we've seen so far. Once the connection is closed the memory is freed and any volatile state is erased.

## Process models

Since a DBMS is normally a **multi-user system** dealing with **concurrency** is a fundamental piece of functionality and one that influences the design from early stages.
