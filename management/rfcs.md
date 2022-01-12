# RFCs

- https://buriti.ca/6-lessons-i-learned-while-implementing-technical-rfcs-as-a-management-tool-34687dbf46cb

## Why

- enable individual contributors to make decisions for systems they’re responsible for
- allow domain experts to have input in decisions when they’re not directly involved in building a particular system
- manage the risk of decisions made
- include team members without it becoming design by committee
- have a snapshot of context for the future
- be asynchronous
- work on multiple projects in parallel

## When

You should write an RFC if you:

- are building something from scratch. New endpoint, component, system, library, application, etc.
- the need rewrite has crossed your mind
- will impact more than one system or other team members.
- would like to define a contract or interface between clients or systems.
- are adding a new dependency.
- are adding or replacing languages or tools to the stack
- are in doubt of whether you should write one

> I’ve found no better way yet to generate the sense of belonging in teams than by including people in decision making. Having an impact at work is important, and some of this impact happens through decision making. If we’re involved in important decisions, our work has the potential to be more impactful and this gives it a sense of purpose. By giving team members the opportunity to comment on decisions proposed by others, RFCs become excellent tools for inclusion and enable participation that can result in impact at work.

## How long

From 2 days to 7 days.

> engineers are not required to participate, but if they don’t do it in time they lose the opportunity of being included.

> RFCs have reduced the number of times managers find themselves with the dreaded “Why wasn’t I consulted about X” question. In addition, when there is a document they can refer to, individual contributors (ICs) can be given concrete feedback on performance if important decisions which are part of their responsibilities were not on their radar.

## RFCs as a KPI

Something worth measuring is the rate of participation on RFCs. A low participation rate (participants/total team size) may be the indicator for a few problems lurking in the background.

## RFCs as an inclusion strategy

> Using RFCs has allowed us to create spaces for team members who wouldn’t normally take the lead in technical decisions.

> Managers and senior ICs are encouraged to appoint less dominant members as authors of RFCs. Being an author makes it clear that you are the person who is responsible for the proposal. This is an explicit way of being empowered to take the lead without the need to be the loudest or most dominant member of a team.

## Examples

- [Rust RFC template](https://github.com/rust-lang/rfcs/blob/master/0000-template.md)
