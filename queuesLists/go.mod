module main

replace example.com/slicequeue => ./slicequeue

replace example.com/nodequeue => ./nodequeue

go 1.22.2

require (
	example.com/nodequeue v0.0.0-00010101000000-000000000000
	example.com/slicequeue v0.0.0-00010101000000-000000000000
)
