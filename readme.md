# Go Programming for Beginners 
    * Static Typed language - not a dynamically typed language like php, python, ruby,javascript

# Basic CLI commands
    * go run hello.go - compile & executes file
    * go build hello.go - compiles the file & creates binary
    * gofmt - Format the GO program. Align it properly but won't write permanently 
    * gofmt -d main.go - Shows the difference between aligned one with previous one
    * gofmt -w main.go - Rewrite the file with proper alignment [ go fmt main.go does the same thing]
    * go install - 
    * go get - go get golang.org/x/tools/cmd/goimports [ Install the package in $GOPATH [/home/akilan/go ]]
    * goimports main.go - removes/adds import packages
    * goimports -d main.go - shows the difference removes/adds import packages
    * goimports -w main.go - Rewrite the file with proper with removes/adds import packages
    * go test -
    * go list -f "{{ .Name }}: {{ .Doc }}" - List main package with Documentation
    * go list -f "{{ .Imports }}" - Lists all the import packages
    * go list -f "{{ .Imports }}" fmt - Lists all the dependent packages for "fmt"
    * go list -f '{{ join .Imports"\n"  }}' fmt - Formatted o/p
    * go doc - Documentaion for main package
    * go doc fmt - Documentation for fmt package
    * go doc fmt Println - Documentation for Println under fmt package
    * godoc -http :6060 -  Access it in browser 

# Build for multiple platforms

    * env GOOS=windows go build hello.go - creates binary for windows [ hello.exe ]
    * env GOOS=darwin go build hello.go - creates binary for macos [ hello ]
    * file has "package main" - Executable one. Any other name for package will be dependency to other files & reusable code
    * Array - Fixed length
    * Slice - We may increase & Shrink Array length

