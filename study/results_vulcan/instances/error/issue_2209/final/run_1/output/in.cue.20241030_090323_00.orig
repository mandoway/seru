Foo: #Obj & {spec: foo: {}}
Bar: #Obj & {spec: bar: {}}

#Obj: X={
	spec: *#SpecFoo | #SpecBar

	out: #Out & {
		_Xspec: X.spec
		if _Xspec.foo != _|_ {
			minFoo: _Xspec.foo.min
		}
		if _Xspec.bar != _|_ {
			minBar: _Xspec.bar.min
		}
	}
}

#SpecFoo: foo: min: int | *10
#SpecBar: bar: min: int | *20

#Out: {
	{minFoo: int} | {minBar: int}

	*{nullFoo: null} | {nullBar: null}
}
