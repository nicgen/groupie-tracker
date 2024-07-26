# Golang web app starter template

This project is made to accelerate a web app creation by
- creating a correct structure
- handling the server
- html templates layout
- handling errors

## Usage

```go
// with air
air
// manualy
go run cmd/groupie_tracker/main.go
```

Note: for air to work you need to modify a line in `.air.toml`

```.air.toml
 cmd = "go build -o ./tmp/main ./cmd/<project_name>/main.go"
```

## Architecture

### Single go.mod

**Simplicity**: A single go.mod file at the root of your project keeps things simple and straightforward. It reduces complexity by centralizing dependency management.

**Consistency**: Ensures that all parts of the project use the same versions of dependencies, avoiding conflicts and inconsistencies.

**Standard Practice**: This follows the convention used in most Go projects, making it easier for other developers to understand and navigate your project.

### Reproducibility

```bash
mkdir -p {cmd,handlers,models,static/{css,img,js},templates}
mkdir cmd/<project-name>
touch cmd/<project-name>/main.go handlers/{index,about,error}.go README.md static/{css/styles.css,img/about.txt} templates/{about,error,index,layout}.html
go mod init <project-name>
air init
sed 's/  cmd = "go build -o .\/tmp\/main ."/  cmd = "go build -o .\/tmp\/main .\/cmd\/<project-name>\/main.go"/g' .air.toml
air -c .air.toml
# after that launch it with `air`
```

## error handling

<!-- middleware refers to a function that wraps an HTTP handler to add additional behavior before
or after the handler processes an HTTP request


The HandleError function will set the appropriate HTTP status code and render an error page with the provided message.
call HandleError directly within your HTTP handlers when you encounter an error.

The WithErrorHandling middleware uses HandleError within a defer statement to handle any panics that occur during the request processing. It logs the panic and sends an appropriate error response using HandleError. -->

## testing

<!-- not done yet -->

## Attribution

This favicon was generated on [favicon.io](https://favicon.io/) using the following graphics from Twitter Twemoji:

- Graphics Title: 2620.svg
- Graphics Author: Copyright 2020 Twitter, Inc and other contributors (https://github.com/twitter/twemoji)
- Graphics Source: https://github.com/twitter/twemoji/blob/master/assets/svg/2620.svg
- Graphics License: CC-BY 4.0 (https://creativecommons.org/licenses/by/4.0/)
