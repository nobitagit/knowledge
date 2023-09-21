# Kafka Fundamentals

- [video course link](https://training.confluent.io/learningpath-viewer/3292/1/161265/16)
- [Kafka resources and videos](https://kafka.apache.org/videos#applications_use_cases)

## Use cases

In IT we are seeing a shift from dealing with snapshots of state to streams of events. Kafka is very suited for that.
Kafka's job is to manage and process those events.

## Fundamentals

Kafka is deployed on a cluster.

> In distributed computing, identical workloads or services run on different computers but present themselves to the developer as a single system. Why would that be useful? 
> 
> Well, say that you go to the library and want to check out a book, like I <3 Logs. Unfortunately, your local library is hosting a fundraising event, and is closed to patrons! Luckily, you can head over to the other side of town, and request a copy of that same resource. Distributed computing solves problems in a similar way: if one computer has a problem, other computers in the group can pick up the slack. A distributed computer system can also easily handle more load by transparently integrating more computers into the system. 
> 
> Apache Kafka® is more than a data streaming technology; it’s also an example of a distributed computing system. 
> 
> Within the context of Kafka, a cluster is a group of servers working together for three reasons: speed (low latency), durability, and scalability. Several data streams can be processed by separate servers, which decreases the latency of data delivery. Data is replicated across multiple servers, so if one fails, another server has the data backed up, ensuring stability - meaning data durability and availability. Kafka also balances the load across multiple servers to provide scalability.

[Source](https://www.confluent.io/blog/what-is-an-apache-kafka-cluster/)

