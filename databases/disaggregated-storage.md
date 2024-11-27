# Disaggregated Storage

From [this article](https://avi.im/blag/2024/disaggregated-storage/)

- All the recent database offerings are built on top of storage disaggregation: Amazon Aurora (the most famous one), [Neon](https://neon.tech/blog/get-page-at-lsn), Snowflake, TiDB, etc.

## Anatomy of a DB

- Every DB can be split in 2 parts: the FE and the BE
- The FE is the query engine, it manages the connections, the parsing and plans the query
- The BE is the storage layer that fetches and writes data to/from disk
- The FE is **CPU heavy**
- The BE is **I/O heavy**

Traditionally, the DBs assumed that the storage would be colocated with the FE, ie. physically on the same machine.
This limits the size of storage to the storage of the machine the DB runs on.

You can't store 3Tb of data on a machine with 1Tb of storage.

But if we store data remotely, and substitue write I/O with network I/O, the storage is not limited anymore by the machine itself.
This is called the separation of compute and storage, also known as **storage disaggregation**.

## Advantages

The first advantage is that disk or CPU can be scaled independently.
You can scale up many compute machines as you like, and if they crash, you won't have to re-instantiate the storage.

## Disadvantages

Latency. Disk I/O is way faster than network I/O. Also, while disk I/O operations rarely fail, network I/O calls can fail more frequently.
