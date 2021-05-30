# Architecture of a Database System (pt. 2)

Part 1 was about the overall architecture of the DB and its relationship with the host OS. Here we see in depth the various components of the DB system.

## Relational Query Processor

A relational query processor job is to:

- take a declarative SQL statement
- validate it
- optimise it into a procedural data flow (query planner)
- execute this plan for the client (worker) who asked for it

### Query Parsing and Authorization

> Given an SQL query, the parser first considers each of the table references in the FROM clause. It canonicalizes table names into a fully qualified name of the form `server.database.schema.table`. This is also called a **four part name**.
>
> Systems that do not support queries spanning multiple servers need only canonicalize to `database.schema.table`, and systems that support only one database per DBMS can canonicalize to just `schema.table`

If the query syntax ("is it valid SQL?") and semantics ("do the tables queried for exist?"), the next step is to ensure the user is **authorised** to perform the `SELECT/DELETE/INSERT/UPDATE` action.

If all is good we go to **query rewrite**

### Query rewrite

The QR, or **rewriter**, simplifies the query without changing its semantics.
The name rewrite is misleading since the query stays the same, what changes is that the internal representation of the query is optimised.

An example:

> The expression `Emp.salary < 75000 AND Emp.salary > 1000000`, for example, can be replaced with `FALSE`. This might allow the system to return an empty query result without accessing the database.
