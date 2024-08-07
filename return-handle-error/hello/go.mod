module example.com/hello

go 1.22.5

// ! $ go mod edit -replace example.com/greetings=../greetings
replace example.com/greetings => ../greetings

// ! go mod tidy
// * go: found example.com/greetings in example.com/greetings v0.0.0-00010101000000-000000000000
require example.com/greetings v0.0.0-00010101000000-000000000000
