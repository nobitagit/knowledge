# Reactive Microsystems

Book by Jonas Boner

## Act autonomously

- Microservices allow team to move more independently
- The cost is **complexity**
- The advantage is the fact that you can roll out or replace services faster

**Authonomy** is the keyword, what we are striving for.
Autonomy is the foundation on which we can scale both the system and the development organization.

Systems should be moved around in isolation

## Single Responsiblity Principle

Similar to the Unix phylosophy, a MS should do one thing only.

## To each their own state

MS are for the most part stateful, they hold their own state.

This simple fact has huge implications. It means that data can be strongly consistent only within each service but never between services, for which we need to rely on eventual consistency and abandon transactional semantics. 

> You must give up on the idea of a single database for all your data, normalized data, and joins across services

![A MS owns its data](images/microservices-state.png)

## Monoliths: how not to break them down

Starting from a typical scenario like this:

![Alt text](images/monolith.png)

Many times an attempted move to micorservice architecture ends up with:

![Alt text](images/microliths.png)

This is **not** a microservice architecture.

> A microlith is defined as a single-instance service in which synchro‐ nous method calls have been turned into synchronous REST calls and blocking database access remains blocking. This creates an architecture that is maintaining the strong coupling we wanted to move away from but with higher latency added by interprocess communication (IPC).

Keywords here are **blocking** and **higher latency**.

The communication between MSs is the hardest part.

> What’s difficult in microservices design is not creating the individ‐ ual services themselves, but managing the space between the serv‐ ices. 

## Looking at the past

> [We need to] come to terms with the fact that information is always from the past, and always represents another present, another view of the world (you are, for example, always seeing the sun as it was 8 minutes and 20 seconds ago). “Now” is in the eye of the beholder, and in a way, we are always looking into the past.
> 
> It’s important to remember that reality is not strongly consistent, but eventually consistent. Everything is relative and there is no single “now.”

How to avoid microliths:

- Events-First Domain-Driven Design
- Reactive Programming and Reactive Systems
- Event-Based Persistence

## Events-First Domain-Driven Design

EF DDD is a set of principles help us to shift the focus from the nouns (the domain objects) to the verbs (the events) in the domain. 

- Think event-first
- Do not be biased by the current structure

> Object-Oriented Programming (OOP) and later Domain-Driven Design (DDD) taught us that we should begin our design sessions focusing on the things—the nouns—in the domain, as a way of finding the Domain Objects, and then work from there. It turns out that this approach has a major flaw: **it forces us to focus on structure too early.****
>
> Instead, we should turn our attention to the things that happen — the flow of events — in our domain. This forces us to understand how change propagates in the system—things like communication patterns, workflow, **figuring out who is talking to whom, who is responsible for what data**, and so on. We need to model the business domain from a data dependency and communication perspective.

> When you start modeling events, it forces you to think about the behavior of the system, as opposed to thinking about structure inside the system.
> Modeling events forces you to have a temporal focus on what’s going on in the system. Time becomes a crucial factor of the system.

Keyword here **"We need to model the business domain from a data dependency and communication perspective."**

## Events represent facts

Events are facts. A fact represents something that has happened in the past.

> Something that truly exists or happens: something that has actual existence, a true piece of information.

Facts are **immutable**. They can’t be changed or be retracted.

We need to focus on the flow of information, rather than on its representation at a point in time. Event Storiming is how we do that.

Domain events are suited to:

- describing the business
- implementing the supporting software

## Event storming

> It’s a design process in which you bring all of the stakeholders—the domain experts and the programmers—into a single room, where they brainstorm using Post-it notes, trying to find the domain lan‐ guage for the events and commands, exploring how they are causally related and the reactions they cause.

- Explore the domain from the perspective of what happens in the system.
- Explore what triggers the events. 
- Agree on the shared terminology, the "Ubiquitous Language"
- Different functions work together at the same time

Example: https://www.youtube.com/watch?v=mLXQIYEwK24

## Inter-service communication

> It is unfortunate that synchronous HTTP (most often using REST) is widely considered as the “state of the art” microservice communi‐ cation protocol. Its synchronous nature introduces strong coupling between the services making it a very bad default protocol for inter‐ service communication. Asynchronous messaging makes a much better default for communication between microservices (or any set of distributed components, for that matter).

- Always apply backpressure. If a service is degrading, it should ask the brokers to reduce their load.
- Prefer async messaging to REST for communication
- REST requires both services to be up and able to interact (ie. not overloaded/degraded)

## How to improve microliths

- Replace REST over HTTPS with async messaging
- Examples are Kafka and AWS Kinesis
- This helps to decouple the services by introducing temporal decoupling — **the services communicating do not need to be available at the same time**

MSs should embrace the fact that data is not always at rest, but it's more and more in motion.
Asynchronicity in cmmunication favours decoupling.

## Scaling persistence

CRUD is most often the wrong way to think about the design of MSs.

>  When bookkeeping was done with clay tablets or paper and ink, accountants developed some clear rules about good accounting practices.
> One never alters the books; if an error is made, it is annotated and a new compensating entry is made in the books. The books are thus a complete history of the transactions of the business.
> Update-in-place strikes many systems designers as a cardinal sin: it vio‐ lates traditional accounting practices that have been observed for hun‐ dreds of years.
>
> —Jim Gray, The Transaction Concept, 1981

Disk space used to be very expensive. This is one of the reasons why most SQL databases are using update-in-place—overwriting existing records with new data as it arrives.
Today disk space is cheap so there is little-to-no reason to use update-in-place for System of Record. 

> We can afford to store all data that has ever been created in a system, giving us the entire history of everything that has ever happened in it.

> The truth is the log. The database is a cache of a subset of the log.
>
> —Pat Helland

## Event sourcing

> Event Sourcing ensures that all changes to application state are stored as a sequence of events. Not just can we query these events, we can also use the event log to reconstruct past states, and as a foundation to automatically adjust the state to cope with retroactive changes.
> 
> [M. Fowler](https://martinfowler.com/eaaDev/EventSourcing.html)

- You store every event that change the state 
- You can rebuild the entire state anytime
- You can decide to persist the state at the end of the day, but keep the daily increments in memory, so we avoid the expensive I/O that comes from the interaction with the System of Record.

> Application states can be stored either in memory or on disk. Since an application state is purely derivable from the event log, you can cache it anywhere you like. A system in use during a working day could be started at the beginning of the day from an overnight snapshot and hold the current application state in memory. Should it crash it replays the events from the overnight store. At the end of the working day a new snapshot can be made of the state. New snapshots can be made at any time in parallel without bringing down the running application.

In a waym, we'd be storing 2 entities: 
- the state
- the event log

This series of events is in fact an event aggregate, and it's generally tied to an event stream for the outer world to subscribe to.

A simple way to help you name events is to remember that they should be a past-tense description of what happened. Unlike com‐ mands, which would be phrased in the imperative (CreateOrder, SubmitPayment, and ShipProducts).

Compare: `createOrder` to `orderCreated`, `shipProduct` to `productShipped`.

## CQRS (Command Query Responsibility Segregation)

This patterns separates the read from the write when dealing with data objects.

> "CQRS is simply the creation of two objects where there was previously only one. The separation occurs based upon whether the methods are a command or a query (the same definition that is used by Meyer in Command and Query Separation: a command is any method that mutates state and a query is any method that returns a value)."
>
> — [Greg Young, CQRS, Task Based UIs, Event Sourcing](https://learn.microsoft.com/en-us/previous-versions/msp-n-p/jj591573(v=pandp.10)?redirectedfrom=MSDN)

The nature of event logs means that we can aggregate their "view" (ie. their state representation) in a number of ways. In practice, we don't really have to choose, we can use many ways of aggregating the state, based on what we need it for.
This is the **read side**.
You could pick one, two or all of these:

- Traditional RDBMSs, like MySQL, Oracle, or Postgres
- Scalable distributed SQL, like Spanner or CockroachDB
- Key-value-based, like DynamoDB
- Hybrid key-value/column-oriented, like Cassandra
- Document-oriented, like MongoDB 
- Graph-oriented, like Neo4j
- Streaming products, like Flink or Kafka Streams
- Search products, like ElasticSearch

The advantage is that, based on our specific use case, we can decide to materialise the state in the best suited infrastructure for the job.

As an example, if you want to build a graph of friends, and run queries along the lines of “who’s my friend’s best friend?”, this query will be most efficiently answered by a graph-oriented database (such as Neo4j).

In other words, by having the event log as the single source of truth in your system, you easily can produce any kind of view on the data that works best for the specific use-case—so-called [Polyglot Persistence](https://martinfowler.com/bliki/PolyglotPersistence.html).

### Tradeoffs and disadvantages of CQRS with event sourcing

- complex: if you need to ship an MVP, this is surely too much overhead, and REST over HTTP is a better fit
- the write side and read side will be eventually consistent
- deletion of events (think GDPR) is not easy at all

## Short history of big data strategies

- First, big data was at rest. Hadoop was an example. Data was collected an processed offline overnight, with hours of latency
- Then data in motion appeared and a hybrid approach came to life. Lambda Architecture is an example of that, with 2 layers, one for speed (close to real time) and one for bigger, batch processed jobs.
- Lastly, data in motion has been embraced fully, and the intent is to remove batch processing altogether.


