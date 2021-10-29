# Algolia

## How to structure and index your data set

- Searchable data
- Filterable data
- Display data
- Business metrics

The mapping is 1 to 1: 1 index per data type.
So products, help, promos, all should go into a separate index each.

- For each DT, map out where this data is located in your system.
- What is the frequency of updates for each DT? Can we leverage any existing event system? In short, **what should trigger a re-index?** For instance:
  - when to update the stock level?
  - when to update the product list?
  - all other changes to the underlying data

## Once index per store/language vs one index for all

If we go for one index per dimension (be it store or language):

- PRO: clear isolation per dimension
- PRO: flexible index settings
- PRO: optimal configuration and analytics
- CONS: slower indexing
- CONS: messier handling of settings, easier to make mistakes/not copy setting across indices

If we go for one index for all:

- PRO: fewer indices
- PRO: multilingual search is possible
- CONS: less fine grained analytics

## Going live

Go through [this checklist](https://www.algolia.com/doc/guides/going-to-production/implementation-checklist/)

## Re-indexing

Algolia suggests to do a full reindex regularly but to do batched updates by default.

## Open questions

- Why is sandbox useful?





TL;DR: I don't always know what I want, but I know what I like.
