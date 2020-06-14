# Reactive vs Passive programming

## Sources

- [Reactive programming - why it matters](https://www.youtube.com/watch?v=49dMGC1hM1o&feature=youtu.begit)

Consider two modules, **Cart** and **Invoice** such that if an item is added to the Cart, Invoice should update.

```

  |--------|      |---------|
  |  Cart  |      | Invoice |
  |--------|      |---------|

```

## Passive

One way to achieve this is to add dependency of Invoice on Cart, such that when Cart updates itself, it updates Invoice as well. Something like the following pseudo code in Cart:

```ts
import { Invoice } from "../invoice";

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
import { Cart } from "../cart";

class Invoice {
  setup() {
    Cart.on("product:added", this.updateItems.bind(this));
  }

  updateItems(product: Product) {
    this.items(product);
  }
}
```

This would be represented as:

```

  |--------|      |---------|
  |  Cart  |   -->| Invoice |
  |--------|      |---------|

```

## In short...

- Passive: remote setters and getters

- Reactive: Event observation and self updates
