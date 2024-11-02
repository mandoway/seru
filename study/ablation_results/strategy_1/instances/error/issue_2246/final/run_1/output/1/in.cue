#FormFoo: fooID: string
#FormBar: barID: string
#Form: {#FormFoo | #FormBar}
data: fooID: "123"
out1: #Form & data
#Form & out1
