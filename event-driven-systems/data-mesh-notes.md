# DM notes

A collection of notes regarding data mesh.

> The data mesh does not need to be constructed in one fell swoop. Many companies attain positive results by taking serial steps. A biotech company began by providing data from an operational data warehouse through a data mesh to feed into operational reporting of its production performance (monitoring production variables). The data product team worked closely with business users to understand their needs, improve data quality and velocity, and standardize data into a harmonized format. Business users were able to explore and develop new applications more quickly at the proof-of-concept stage and then scale them to full production.
>
> Executive and nontechnical business users will all need a basic level of data literacy for data mesh success. Coaching, hackathons, online programs, and analytics academies can all work well.

-- Demistifying Data Mesh, [McKinsey](https://www.mckinsey.com/capabilities/quantumblack/our-insights/demystifying-data-mesh)

- DM is the corresponding paradigm of "shift left approach" applied to data. It is much more effective to curate data at source and fix problems as soon as possible in the pipeline, than fixing them at a later stage

- In traditional data pipelines/warehouse approach it is not uncommon to ingest data that have errors and fix it after ingestion. Cleansing data is common practice and part of the lingo.

- In DM, products must adhere to quality standards and  be considered reliable. Data is cleansed at the **point of creation, not after ingestion**.