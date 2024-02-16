# Designing event driven systems

Concepts and Patterns for Streaming Services with Apache Kafka
Book by Ben Stopford

## Messaging systems vs replayable logs

Replayable logs, like Kafka, can play the role of an event store: a middle ground between a messaging system and a database. 
They decouple services from one another, much like a messaging system does, but they also provide a central point of storage that is fault-tolerant and scalable—a shared source of truth that any application can fall back to.

## Shared DBs or Integration DB

An integration DB is a DB that is shared across multiple (on paper) independent services. It is considered an antipattern as it creates strict coupling between services.

A replayable log instead provides the ability to materialise the state wherever we need, and not necessarily in one place only.

## What Kafka is, and what it is not

Kafka is a mechanism for programs to exchange information, but its home ground is event-based communication, where **events are business facts that have value to more than one service and are worth keeping around.**

The term queue has something ephemeral to it, but logs are durable and immutable.

Streaming encourages services to retain control, especially of their data, rather than providing orchestration from a single point, or from a single team or platform.

> Centralize an immutable stream of facts. Decentralize the freedom to act, adapt, and change.

or

> Decentralise the model, centralise the streaming of data

### Is Kafka a DB?

> Some people like to compare Kafka to a database. It certainly comes with similar features. It provides storage; production topics with hundreds of terabytes are not uncommon. It has a SQL interface that lets users define queries and execute them over the data held in the log. These can be piped into views that users can query directly. It also supports transactions. 

While all these elements are inherent to a DB, Kafka is a DB inside out, a tool for storing data, processing it in real time, and creating views. 

Streams are optimized for continual computation rather than batch processing.

Kafka is designed to move data, operating on that data as it does so. It’s about real-time processing first, long-term storage second.

Kafka is ultimately a streaming platorform.

## Quick overview of Kafka

- A **Kafka cluster** is essentially a collection of files, filled with messages, spanning many different machines