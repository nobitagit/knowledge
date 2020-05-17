# Algorithms and data structures

From [PS Algorithms and Data Structures - Part 1](https://app.pluralsight.com/player?course=ads-part1&author=robert-horvick&name=ads-linked-list&clip=0&mode=live).

## Node

Basic building block for many DSs (Data Structures).

```
-----      -----
| 1 | ---> | 3 |
-----      -----
```

Implementation:

```ts
class BaseNode {
  value: number;
  next?: BaseNode;
  constructor(value: number, next?: BaseNode) {
    this.value = value;
    this.next = next;
  }
}

// We create 2 independent Nodes
const head: BaseNode = new BaseNode(3);
const middle: BaseNode = new BaseNode(5);

// we now create the link between the first Node and the second
head.next = middle;

function printNodes(node: BaseNode) {
  let n: BaseNode | undefined = node;
  while (n !== undefined) {
    console.log(n.value);
    n = n.next;
  }
}

printNodes(head);
// 3
// 5
```

## Linked List

- a single chain of nodes
- has a reference to the head
- has a reference to the tail
- Supports a few basic operations:
  - add
  - remove
  - find
  - enumerate

### Linked list: adding an item to the front

```ts
// Same as above, but we now use a generic
class BaseNode<T> {
  value: T;
  next?: BaseNode<T>;
  constructor(value: T, next?: BaseNode<T>) {
    this.value = value;
    this.next = next;
  }
}

function printNodes<T>(node: BaseNode<T> | undefined) {
  let n: BaseNode<T> | undefined = node;
  while (n !== undefined) {
    console.log(n.value);
    n = n.next;
  }
}

class LinkedList<T> {
  public head: BaseNode<T> | undefined;
  public tail: BaseNode<T> | undefined;
  public count: number;
  constructor() {
    this.count = 0;
    this.head = undefined;
    this.tail = undefined;
  }

  addToFront(node: BaseNode<T>) {
    // we save a reference to the current head
    // so we don't lose it when reassigning
    // (see next step)
    const tempHead = this.head;
    // we mark the node that is passed to the method
    // as the new head
    this.head = node;

    // we then move the previous head to be the 2nd
    // node in the LL
    this.head.next = tempHead;

    // we keep track of the length
    this.count = this.count++;

    // if there's only one item in the list, the head is equal to the tail
    if (this.count === 1) {
      this.tail = this.head;
    }
  }

  print() {
    console.log(this);
    printNodes<T>(this.head);
  }
}

const ll = new LinkedList();
const head = new BaseNode(4);
ll.addToFront(head);
const middle = new BaseNode(5);
ll.addToFront(middle);
ll.print();
// 5
// 4
```

### Linked list: adding an item to the tail

As we've seen above, we have a reference to the tail in the LL implementation.
This makes adding to the tail much easier.

```ts
class LinkedList<T> {
  // this extends the implementation above
  addToTail(node: BaseNode<T> | undefined) {
    // const tempTail = this.tail;
    // this.tail = node;
    // if LL is empty tail and head reference the same node
    if (this.length === 0) {
      this.head = node;
    } else {
      this.tail.next = node;
    }
    this.tail = node;
    this.length = this.length + 1;
  }
}

const ll = new LinkedList();
const head = new BaseNode(4);
ll.addToTail(head);
const middle = new BaseNode(5);
ll.addToTail(middle);
ll.print();
```
