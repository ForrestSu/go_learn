module github.com/ForrestSu/go_learn/3rd-libs/37_protobuf_any

go 1.21

require (
	github.com/stretchr/testify v1.10.0
	google.golang.org/protobuf v1.36.6
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/golang/protobuf => github.com/golang/protobuf v1.4.3
	google.golang.org/protobuf => google.golang.org/protobuf v1.25.0
)
