import { binarySearch } from './binary-search'

const arr = [1, 3, 5, 7, 9]
const target = 5
const result = binarySearch(arr, target)
console.log(`binarySearch([${arr}], ${target}) = ${result}`)
