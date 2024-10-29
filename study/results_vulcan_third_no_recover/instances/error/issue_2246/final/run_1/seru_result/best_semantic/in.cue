#FormFoo: fooID: fooID
data: #Form
#Form: {
	{
		fooID: fooID
	} | {
		fooID: fooID
	}
}
data: fooID: "123"
data: data
{#FormFoo | #FormFoo} & data
