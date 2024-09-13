module mainstack

replace example.com/simplestack => ../simplestack

replace example.com/nodestack => ../nodestack

go 1.22.2

require (
	example.com/nodestack v0.0.0-00010101000000-000000000000
	example.com/simplestack v0.0.0-00010101000000-000000000000
)
