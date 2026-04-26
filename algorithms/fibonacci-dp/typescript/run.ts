import { fibonacci } from './fibonacci-dp'

for (const n of [0, 1, 2, 5, 10]) {
  console.log(`fibonacci(${n}) = ${fibonacci(n)}`)
}
