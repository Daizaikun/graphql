@ECHO OFF  
go get github.com/99designs/gqlgen/internal/imports
go get github.com/99designs/gqlgen/codegen/config
go get github.com/99designs/gqlgen
go get github.com/99designs/gqlgen/codegen
@ECHO ON
go run github.com/99designs/gqlgen generate ./...
@ECHO OFF  
EXIT 