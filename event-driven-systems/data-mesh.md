# Data Mesh

Audiobook by Z. Dehghani

A journey in distibuted data ownership.
....

## Chapter 2

### Unique source of truth

- The search for one source of truth is costly and most of the time impossible. 
- What people want is not to have out of date data and dozens of different copies of the same attribute. 
- Data mesh doesn't guarantee one source of truth and doesn't aim to provide one. 
- DM describes the series of practices and infra that decrease the likelihood of that multiple copies of out of date data.

DM endorses reshaping data. Data can be loaded, transformed and stored from another domain.

### Data pipelines

- In traditional data architectures, data pipelines are first class architectural concerns
- In DM, data pipelines are implementation details of the data domain. They are managed within the domain and invisible from outside of the domain

With DM, the domains take ownership and responsibility of their own data, vs. the central team that has to manage the entire complexity.

## Chapter 3 - Principles of data as a product

For data to be a product it needs to abide to set of rules.
This phylosophy introduces new roles in the businees domain:

- Domain data product ownwer
- Domain data developer/engineer

### Apply product thinking to data

A domain data team can expose different pieces of data, in different formats based on the need of the (internal) customers that are interested in consuming the data.

- infinite strems of events (like "sessions started on a music player")
- aggregated, materialsed views of a slice of the data ("all playlist currently available")

### Baseline characteristics of a data product to be useful

Every data product needs to have these attributes need to adhere to these attributes to be 

- discoverable
- understandable
- trustworthy
- accessible
- interoperabile 
- usable

#### Discoverable
The first step in venturing into a system of data is to explore and discover. 
To achieve this, a **catalogue** or **a centrailsed registry** with some additional info like:

- owners
- location
- sample data
- use cases and applications enabled by this data

The data product must have an addressable aggreagate route where

#### Understandable

A data user should easily understand the pieces of data and the actions available.
Sample data sets and code examples for consuming help in that sense.

#### Trustworthy

Data users should trust that the data is reliable and correct.
Data owners should clearly represent SLOs:

- interval of change: how often is the data updated
- lead time: how much time between an event occurs and the time it is captured
- completeness
- lineage: if the data is transformed from the source to the consumption location

In traditional data pipelines/warehouse approach it is not uncommon to ingest data that have errors and fix it after ingestion. Cleansing data is common practice and part of the lingo.

In DM, products must adhere to quality standards and  be considered reliable. Data is cleansed at the **point of creation, not after ingestion**.

#### Accessible