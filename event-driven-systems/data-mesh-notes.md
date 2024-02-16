# DM notes

A collection of notes regarding data mesh.

> The data mesh does not need to be constructed in one fell swoop. Many companies attain positive results by taking serial steps. A biotech company began by providing data from an operational data warehouse through a data mesh to feed into operational reporting of its production performance (monitoring production variables). The data product team worked closely with business users to understand their needs, improve data quality and velocity, and standardize data into a harmonized format. Business users were able to explore and develop new applications more quickly at the proof-of-concept stage and then scale them to full production.
>
> Executive and nontechnical business users will all need a basic level of data literacy for data mesh success. Coaching, hackathons, online programs, and analytics academies can all work well.

-- Demistifying Data Mesh, [McKinsey](https://www.mckinsey.com/capabilities/quantumblack/our-insights/demystifying-data-mesh)

- DM is the corresponding paradigm of "shift left approach" applied to data. It is much more effective to curate data at source and fix problems as soon as possible in the pipeline, than fixing them at a later stage

- In traditional data pipelines/warehouse approach it is not uncommon to ingest data that have errors and fix it after ingestion. Cleansing data is common practice and part of the lingo.

- In DM, products must adhere to quality standards and  be considered reliable. Data is cleansed at the **point of creation, not after ingestion**.

> **Ubiquitous Language** is the term Eric Evans uses in Domain Driven Design for the practice of building up a common, rigorous language between developers and users. This language should be based on the Domain Model used in the software - hence the need for it to be rigorous, since software doesn't cope well with ambiguity.
>
> [Martin Fowler](https://martinfowler.com/bliki/UbiquitousLanguage.html)

- We need to focus on the flow of information, rather than on its representation at a point in time.

- Read the definition of Single Source of Truth https://en.wikipedia.org/wiki/Single_source_of_truth especially Master Data Management, Event Store etc.

> It's not the domain experts knowledge that goes to production, it's the assumption of the developers that goes to production
> 
> Alberto Brandolini

- A shared DB across multiple services is called [integration Database](https://martinfowler.com/bliki/IntegrationDatabase.html). An integration database needs a schema that takes all its client applications into account. The resulting schema is either more general, more complex or both - because it has to unify what should be separate BoundedContexts. 
> On the whole integration databases lead to serious problems becaue the database becomes a point of coupling between the applications that access it. This is usually a deep coupling that significantly increases the risk involved in changing those applications and making it harder to evolve them. As a result most software architects that I respect take the view that **integration databases should be avoided.**

-  Turning attention to that 5 to 20 percent of the equation, technology is still how we deliver change. But what technology best suits Data-as-a-Product?
Data mesh is a scalable way of applying order to the data we have today vs data lakes where data is centralized to one place and difficult to move, scale, and to get the agility we need. Imagine, for example, if Amazon in its early days had decided to stockpile every book into one distribution center, and then tried to distribute those books everywhere in the world from once single location. Sounds a little bit ludicrous, right? You can almost imagine the scale of that and how it would grind to a halt very quickly. It would become very difficult to manage, and very disorganized.
Companies like Amazon have always acknowledged the fact that the best way to scale is to distribute these effort through its network, hubs, and resellers. And data is a similar thing which is why data mesh is a better solution when preparing ourselves for a product thinking approach. From [this post](https://www.kinandcarta.com/en-us/insights/2021/10/5-steps-to-enable-data-as-a-product-thinking/)