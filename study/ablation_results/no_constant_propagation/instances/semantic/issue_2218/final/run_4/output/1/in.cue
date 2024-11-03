_input: {
 child1: {}
 child2: test: true
}
for k, v in _input {
 let tmp = v.test
 (k): tmp
}
