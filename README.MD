To run Goophy on your browser, you need to run the following commands:
To compile for Windows Powershell
`cd server 
$Env:GOOS = "js"; $Env:GOARCH = "wasm"; go build -o ../pkg/assets/main.wasm ../pkg/wasm/main.go`
To compile on other terminals
`cd server 
GOOS=js GOARCH=wasm go build -o ../pkg/assets/main.wasm ../pkg/wasm/main.go`
