# Reactive vs Passive programming

## Sources

- [Reactive programming vs Passive programming](https://vaibhavgupta.me/2017/12/31/reactive-programming-vs-passive-programming/)
- [Reactive programming - why it matters](https://www.youtube.com/watch?v=49dMGC1hM1o&feature=youtu.be)

Consider two modules, **Cart** and **Invoice** such that if an item is added to the Cart, Invoice should update.

```

  |--------|      |---------|
  |  Cart  |      | Invoice |
  |--------|      |---------|

```

## Passive

One way to achieve this is to add dependency of Invoice on Cart, such that when Cart updates itself, it updates Invoice as well. Something like the following pseudo code in Cart:

```ts
class Cart {
  addItemInCart(item: Item) {
    Invoice.update(item.price);
  }
}
```

We can represent this relationship like so:

```

  |--------|      |---------|
  |  Cart  |-->   | Invoice |
  |--------|      |---------|

```

## Reactive

Another way to achieve this is to invert the relationship so that Invoice depends on Cart, updating itself when Cart changes.

```ts
class Invoice {
  addItemInCart(item: Item) {
    Invoice.update(item.price);
  }
}
```

This would be represented as:

```

  |--------|      |---------|
  |  Cart  |   -->| Invoice |
  |--------|      |---------|

```
