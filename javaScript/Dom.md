# DOM

## How to change the text content of a node

```js
document.getElementById("div-1").textContent = "This text is different!";
```

## How to work with fragments

```html
<ul id="list"></ul>
```

```js
var list = document.querySelector("#list");
var fruits = ["Apple", "Orange", "Banana", "Melon"];

var fragment = new DocumentFragment();

fruits.forEach(function (fruit) {
  var li = document.createElement("li");
  li.textContent(fruit);
  fragment.appendChild(li);
});

list.appendChild(fragment);
```

## How does JSX/Hyperhtml work under the hood

```js
// from https://codeburst.io/taming-huge-collections-of-dom-nodes-bebafdba332
const h = (tag, attrs, ...children) => {
  const elm = document.createElement(tag);
  for (let key in attrs) {
    if (key.slice(0, 2) == "on") {
      const evtName = key.slice(2);
      const cb = attrs[key];
      if (cb == null) continue; // we can use null or undefnied to suppress
      elm.addEventListener(evtName, cb);
    } else if (
      ["disabled", "autocomplete", "selected", "checked"].indexOf(key) > -1
    ) {
      if (attrs[key]) {
        elm.setAttribute(key, key);
      }
    } else {
      if (attrs[key] == null) continue; // Don't set undefined or null attributes
      elm.setAttribute(key, attrs[key]);
    }
  }
  if (children.length === 0) {
    return elm;
  }
  forEach((child) => {
    if (child instanceof Node) {
      elm.appendChild(child);
    } else {
      elm.appendChild(document.createTextNode(child));
    }
  }, flatten(children));
  return elm;
};

h("div", { class: "product" }, [
  h("span", { class: "id" }, product.id),
  h("span", { class: "type" }, product.type),
]);
```

If you are using Babel with the react preset, you can use the @jsx h pragma to use JSX with the h() function:

```jsx
/** @jsx h **/
const renderProduct = (product) => (
  <div class="product">
    <span class="id">{product.id}</span>
    <span class="type">{product.type}</span>
  </div>
);
```
