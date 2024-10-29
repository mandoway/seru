_configs: {
 CHILD1: {}
 CHILD2: test: true
}
for k, v in _configs {
 let tmp = v.test
 "\(k)": tmp
}
