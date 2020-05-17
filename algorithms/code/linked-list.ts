// This is the complete code of the Linked List section of
// algorithms/data-structures.md

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
    this.count = this.count + 1;

    // if there's only one item in the list, the head is equal to the tail
    if (this.count === 1) {
      this.tail = this.head;
    }
  }

  addToTail(node: BaseNode<T> | undefined) {
    // const tempTail = this.tail;
    // this.tail = node;
    // if LL is empty tail and head reference the same node
    if (this.count === 0) {
      this.head = node;
    } else {
      this.tail!.next = node;
    }
    this.tail = node;
    this.count = this.count + 1;
  }

  removeFirst() {
    // empty LL -> do nothing
    if (this.count == 0) {
      return;
    }

    this.head = this.head!.next;
    this.count = this.count - 1;

    if (this.count == 0) {
      this.tail = undefined;
    }
  }

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

  print() {
    console.log(this);
    printNodes<T>(this.head);
  }
}

const ll = new LinkedList();
const head = new BaseNode(4);
ll.addToTail(head);
const middle = new BaseNode(5);
ll.addToTail(middle);
ll.removeFirst();
ll.print();
