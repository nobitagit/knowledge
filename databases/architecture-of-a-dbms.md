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

Since a DBMS is normally a **multi-user system**, dealing with **concurrency** is a fundamental piece of functionality and one that influences the design from early stages.

Some terminology first:

- **OS process**: a program execution unit to which the kernel allocates a space in memory. The memory is **private** to that process. A process is an executing instance of a program. In its simple form, 1 program:1 process, but nevertheless, 1 program can spin up more than 1 process. On a multiprocessor system, **multiple processes** can be executed in parallel. On a **uni-processor system**, though true parallelism is not achieved, a process scheduling algorithm is applied and the processor is scheduled to execute each process one at a time **yielding an illusion of concurrency**.
- **OS thread**: a program execution unit that "lives" inside a process. In fact, it's a sub-unit of a process. It has no truly private space and **shares the resources with all the other OS threads in that process**. They are called also **Kernel threads**.
- **Lightweight Thread Package**: an **application level** thread (as opposed to OS level threads, see above). It comprises multiple OS threads in a single process and the Kernel is unaware of it.
- **DBMS Client**: a software that is the link between programs and the DBMS itself, via its API.
- **DBMS Worker**: a thread of execution in the DBMS that is in charge of attending the requests of the **DBMS client**. It's a 1:1 mapping, 1 client, 1 worker.

Read [more](http://web.archive.org/web/20100807021758/http://kquest.co.cc/2010/03/program-process-task-thread)

There are various ways one can map 1 worker to OS threads of processes.
Starting from the simplest scenario of a uni-processor machine there are 3 ways one can design the process model. Simplest to hardest:

1. 1 process per DBMS worker
2. 1 thread per DBMS worker
3. Process pool

![1 process x 1 worker](./images/1processx1worker.png)

![1 thread x 1 worker](./images/1threadx1worker.png)