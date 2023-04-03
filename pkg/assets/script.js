WebAssembly.instantiateStreaming(fetch("main.wasm"), goWasm.importObject)
    .then((result) => {
        const goWasm = new Go()
        goWasm.run(result.instance)
    });