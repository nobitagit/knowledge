# TypeScript

## How to type an arrow function with generics

```ts
export const identity = <T extends {}>(o: T): T => o;
// or 
const foo = <T>(o: T) => o
// or - if using JSX, which would be thrown off by the above
const foo = <T,>(o: T) => o
// see https://stackoverflow.com/questions/32308370/what-is-the-syntax-for-typescript-arrow-functions-with-generics#comment99104831_45576880
```

## How to suppress "Object is possibly 'null' or 'undefined'” error when strict null checks are enabled
In some cases we are sure that an expression is non null and we want to override the error Ts is throwing.
For example in a test I might have 

```ts
const svg = container.querySelector('svg');
expect(svg.style).toEqual('position: absolute'); // this will throw because svg might be null
      
// but since it's a test file we are 100% sure the svg is there so:
expect(svg!.style)... // error goes away
```

See [here](https://stackoverflow.com/a/40350534/1446845) and [here](https://stackoverflow.com/a/38875179/1446845)
