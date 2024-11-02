Foo: #Abstract & {spec: foo: {}}
Bar: #Abstract & {spec: bar: {}}

#Abstract: X={
	spec: _#Spec

	resource: _Thing & {_X: spec: X.spec}
}

_#Spec: *_#SpecFoo | _#SpecBar

_#SpecFoo: {
	foo: {
		min: int | *10
		max: int | *20
	}
}

_#SpecBar: {
	bar: {
		min: int | *30
		max: int | *40
	}
}

_Thing: #Constrained & {
	_X: _

	spec: {
		if _X.spec.foo != _|_ {
			minFoo: _X.spec.foo.min
			maxFoo: _X.spec.foo.max
		}

		if _X.spec.bar != _|_ {
			minBar: _X.spec.bar.min
			maxBar: _X.spec.bar.max
		}
	}
}

#Constrained: #Base & {
	spec: {
		minFoo:  int | *10
		maxFoo:  int | *20
		minBar?: null
		maxBar?: null
	} | {
		minBar:  int | *30
		maxBar:  int | *40
		minFoo?: null
		maxFoo?: null
	}

	spec: *{
		fuga?: null
	} | {
		hoge?: null
	}
}

#Base: {
	spec: {
		minFoo?: null | int
		maxFoo?: null | int
		minBar?: null | int
		maxBar?: null | int

		hoge?: null | bool
		fuga?: null | bool
	}
}
