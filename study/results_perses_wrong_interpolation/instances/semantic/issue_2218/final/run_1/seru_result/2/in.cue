{}
for k, v in {
 child1: {}
 child2: test: true
} {
 let tmp = v.test
 (k): tmp
}
