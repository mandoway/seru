{}
for k, v in {
 CHILD1: {}
 CHILD2: test: true
} {
 let tmp = v.test
 "\(k)": tmp
}
