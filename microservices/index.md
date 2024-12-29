# Microservices

## Service to service communication patterns

### Request/Response

- Description: Services interact in real-time using protocols like HTTP/REST or gRPC.
- Use Cases: When immediate feedback is required, e.g., fetching user details from an authentication service.
- Challenges: Coupling, higher latency, risk of cascading failures if dependent services are unavailable.

Note: while req/res might imply synchronous communication, it's not really a given that it's a synchronous pattern. A service could still send a request and not wait for a response, hence being non blocking.

### Message passing

- Description: Services communicate via message brokers (e.g., RabbitMQ, Kafka, Amazon SQS) without waiting for a response.
- Use Cases: Event-driven systems, order processing, notifications.
- Benefits: Looser coupling, higher fault tolerance, and scalability.

## Sync vs Async

As soon as you have multiple receivers, the communication pattern has to be async, otherwise the likelihood of bad performance or chains of failures increases drastically.

## Communication protocols

- REST: Simple and widely supported but less efficient for high-performance use cases.
- gRPC: Binary protocol for low-latency and high-performance communication.
- GraphQL Federation: Enables querying multiple microservices via a single schema.

Direct communication between services should be avoided, if possible.
Direct communciation creates coupling. A change needed in one system might need to be replicated in all the other systems that are aware of it.

## Events vs. Messages

The terms events and messages are often used interchangeably in distributed systems, but they have distinct meanings depending on the context and intent of communication. Here’s a breakdown of the differences:

_Messages:_

- Purpose: A message is a directed communication between two parties, typically a sender and a receiver.
- Intent: It often conveys a command, query, or request for an action.
- Example: “Process this order.”

Messages are often used in request-response or direct communication patterns.

_Events:_

- Purpose: An event is a notification that something has happened, without directing the receiver to take specific action.
- Intent: It signals a state change or occurrence, leaving it to the subscribers to decide how to respond.
- Example: “Order has been placed.”

Events are used in pub-sub or event-driven architectures.
Events **represents facts**. Ex. "this has happened".

Events favours decoupling.
