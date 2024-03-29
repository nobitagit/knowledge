# URL shortener

## 1. Understand exactly what the system does.

Visit http://tinyurl.com/ and try it.
Check input and output:

- IN: https://docs.microsoft.com/en-us/learn/modules/rust-get-started/1-introduction
- OUT: http://tinyurl.com/jMg8z44d

## 2. Clarify requirements

### Functional Requirements

- Given a URL, the system should generate a shorter, unique URL
- Using the url should redirect the user to the original website
- User can pick a custom short URL
- Links expire. User can decide when.

### Non-Functional Requirements

- Highly available
- Minimal latency when redirecting
- URLs should be unpredictable

### 3. **Estimate.** Starting from some ballpark numbers write your expectations.

- Lots of reads compared to writes. Estimate: 100:1
- **Traffic estimate**. 500 millions of new URLs x month. This means:

```
100 (reads ratio) * 500 million = 50 billion reads x month
```

- How many queries x second?

```
Number of seconds in a month:
seconds x minutes in 1 hour x hours in 1 day x days in 1 month
60      X 60                x 24             x 30
= 2,592,000

500 million /
```

4. **Define the APIs**:
   - Endpoints that we expose
   - Parameters
   - Returns
5. Design the DB schema
6. Pick a DB. Relational vs non-relational.
7. **Dig into the problem** to solve the principal challenges.
8. How to scale the DB?
9. Caching
10. Load Balancing
11. Telemetry and analytics
12. Security
