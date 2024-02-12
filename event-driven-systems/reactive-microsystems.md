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

We need to focus on the flow of information, rather than on its representation at a point in time.