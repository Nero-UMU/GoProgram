module example.com/hello

go 1.19

replace example.com/greetings => ../greetings

replace example.com/add => ../add

require (
	add v0.0.0-00010101000000-000000000000
	example.com/greetings v0.0.0-00010101000000-000000000000
)

replace add => ../add
