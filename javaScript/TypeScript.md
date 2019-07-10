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
