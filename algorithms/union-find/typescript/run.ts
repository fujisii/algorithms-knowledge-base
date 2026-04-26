import { UnionFind } from './union-find'

const uf = new UnionFind(5)
uf.union(0, 1)
uf.union(2, 3)
console.log('connected(0,1):', uf.connected(0, 1))
console.log('connected(0,2):', uf.connected(0, 2))
console.log('find(1):', uf.find(1))
