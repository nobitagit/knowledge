# Writing a DB

- Do not focus on algorithms, or speed, or scalability. Just the components.

- Support 1 server, 1 DB (or more?), multiple schemas and multiple tables.

- Start from the Relational Query Processor
  - implement query parsing using a 3rd party module
  - skip authorisation? or do dummy authorisation with 1 role at first
  - skip query rewrite (no optimisation of the query)
- Read/write operations are free from the Halloween problem

## Reading

- https://www.sqlite.org/queryplanner.html

## Resources

- SQL to AST https://github.com/taozhi8833998/node-sql-parser or https://github.com/godmodelabs/flora-sql-parser
