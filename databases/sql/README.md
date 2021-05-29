## SQL

### How to paginate results

```sql
SELECT * FROM this_table LIMIT 0, 20
```

## How to look for multiple values for the same column

```sql
SELECT * FROM stock_items WHERE client_id IN (4221, 4354, 4356, 4359, 4362, 4365, 4368, 4371, 4373, 4375, 4377, 4379, 4381, 4383, 4385)
```
