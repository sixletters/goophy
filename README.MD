# Goophy Language

![Goophy](https://raw.githubusercontent.com/sixletters/goophy/main/pkg/assets/images/Goofy.png)

Goophy is a programming language written in Go that aims to be a simple language for building concurrent programs.

## Installation

To install Goophy, you first need to have Go installed on your system. Once you have Go installed, you can clone this repo into your system.

## Running Goophy natively

To run Goophy locally, you can update the test.txt file with your code and run the following command:

```
go run main.go test.txt
```

## Running Goophy in your browser

To run Goophy on your browser, you need to run the following commands:

### To compile for Windows Powershell

```
cd server
$Env:GOOS = "js"; $Env:GOARCH = "wasm"; go build -o ../pkg/assets/main.wasm ../pkg/wasm/main.go
```

### To compile on other terminals

```
cd server
GOOS=js GOARCH=wasm go build -o ../pkg/assets/main.wasm ../pkg/wasm/main.go
```

This will compile the Goophy interpreter to WebAssembly, which can be run in a web browser.

## Usage

Goophy supports basic concurrency through a go keyword, similar to Go. You can create a new Goroutine using the go keyword:

```
var foo = func(){
// your code here
}
go foo()
```

## Contributing

We welcome contributions to Goophy! If you find a bug or have an idea for a new feature, please create an issue in the GitHub repository. If you want to contribute code, please fork the repository and submit a pull request with your changes.
