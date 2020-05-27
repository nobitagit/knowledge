# Algorithms and data structures

From [PS Algorithms and Data Structures - Part 1](https://app.pluralsight.com/player?course=ads-part1&author=robert-horvick&name=ads-linked-list&clip=0&mode=live).

## TOC

- [Node](#Node)
- [Linked List](#Linked-List)
- [Doubly Linked List](#Doubly-Linked-List)
- [Stack](#Stack)

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
    // if LL is empty tail and head reference the same node
    if (this.length === 0) {
      this.head = node;
      // otherwise we mutate the pointer of the current tail
      // to point to the newly appended node
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
// 4
// 5
```

### Linked lists: removing nodes

Removing from the front is very easy, since the head references the 2nd node in the LL.

```ts
class LinkedList<T> {
  removeFirst() {
    // empty LL -> do nothing
    if (this.count == 0) {
      return;
    }

    this.head = this.head.next;
    this.count = this.count - 1;

    if (this.count == 0) {
      this.tail = undefined;
    }
  }
}

const ll = new LinkedList();
const head = new BaseNode(4);
ll.addToTail(head);
const middle = new BaseNode(5);
ll.addToTail(middle);
ll.removeFirst();
ll.print();
// 5
```

Removing from the tail, is more complex, since the last node doesn't reference the previous one, so we have to traverse the LL from the head up to the last - 1 node.

```ts
class LinkedList<T> {
  removeLast() {
    // empty LL -> do nothing
    if (this.count == 0) {
      // do nothing
    } else if (this.count == 1) {
      this.tail = undefined;
      this.head = undefined;
    } else {
      let currentNode = this.head as BaseNode<T>;
      // we loop until we find the last - 1 node
      // (the one referencing the tail directly)
      while (currentNode.next !== this.tail) {
        currentNode = currentNode.next!;
      }

      currentNode!.next = undefined;
      this.tail = currentNode;
    }
  }
}

const ll = new LinkedList();
const head = new BaseNode(4);
ll.addToTail(head);
const middle = new BaseNode(5);
ll.addToTail(middle);
ll.removeLast();
ll.print();
// 4
```

This might result in a very expensive operation in case of very deep LLs.

## Doubly Linked List

This is a specialisation of the linked list. DLLs have 2 pointers, one to the next and one to the previous node.

## Stack

- LIFO container
- push, pop, peek
- can be implemented by using a Linked List or an Array

## Queue

- FIFO container
- enqueue, dequeue, peek,

## Priority queue

A specialisation of the queue.

## trees

Trees are similar to linked lists, in that they chain together nodes of data.
instead of doing it in a **linear fashion** though, they do it in an **hierarchical** one.

Each node in a tree can have many child nodes but **only one** parent.

```
                    |----|
                    |    |
                    |----|
                      |
                      |
      _________________________________
      |        |      |       |       |
    |---|    |---|  |---|   |---|   |---|
    |   |    |   |  |   |   |   |   |   |
    |---|    |---|  |---|   |---|   |---|

```

This structure implies that **there will always be 1 and only 1 path to go from node A to node b (where B is a child node of A)**.

This is an important limitation of the tree structure.

## Binary tree

A binary tree is composed of nodes that have 0, 1 or 2 child nodes.
A node with 2 child nodes will have a **left child** and a **right child**.

```
    |---|
    |   |
    |---|
      |
  _________
  |       |
|---|   |---|
|   |   |   |
|---|   |---|

LEFT   RIGHT
```

## Binary search tree

This is a specialisation of the binary tree, where:

- nodes are sorted
- the left child node will contain a value that is **smaller than the parent**.
- the right child node will contain a value that is **bigger than the parent**.

```
    |---|
    | 4 |
    |---|
      |
  _________
  |       |
|---|   |---|
| 2 |   | 6 |
|---|   |---|

LEFT   RIGHT
```

### Adding data to a binary search tree

Adding data happens though a recursive algo.

### Case 1: an empty tree

Whatever value we start with becomes the root node.

```
    |---|
    | 4 |
    |---|
```

### Case 2: adding a smaller value

```
    |---|
    | 4 |
    |---|
      |
  _____
  |
|---|
| 2 |
|---|
```

This happens recursively if more smaller values are added.

```
        |---|
        | 4 |
        |---|
          |
      _____
      |
    |---|
    | 2 |
    |---|
      |
  _____
  |
|---|
| 1 |
|---|
```

### Case 3: adding a higher value

```
        |---|
        | 4 |
        |---|
          |
          |
      _________
      |       |
    |---|   |---|
    | 2 |   | 6 |
    |---|   |---|
      |
  _____
  |
|---|
| 1 |
|---|
```

### Case 4: adding a duplicate

Duplicate will be sent to the right. In the case above we can add 4 and it will behave like this:

```
        |---|
        | 4 |
        |---|
          |
          |
      _________
      |       |
    |---|   |---|
    | 2 |   | 6 |
    |---|   |---|
      |       |
  _____   _____
  |       |
|---|   |---|
| 1 |   | 4 |
|---|   |---|
```
