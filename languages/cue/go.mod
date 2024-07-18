module github.com/mandoway/seru/languages/cue

go 1.22.5

replace github.com/mandoway/seru => ../../.

require (
	cuelang.org/go v0.9.2
	github.com/mandoway/seru v0.0.0-00010101000000-000000000000
)

require github.com/cockroachdb/apd/v3 v3.2.1 // indirect
