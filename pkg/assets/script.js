const goWasm = new Go()

WebAssembly.instantiateStreaming(fetch("main.wasm"), goWasm.importObject)
    .then((result) => {
        goWasm.run(result.instance)
        
        const input = "some code to compile";
        const encoded = new TextEncoder().encode(input);
        const len = encoded.length;
        const ptr = instance.exports.malloc(len); // Use WebAssembly memory management to allocate memory
        new Uint8Array(memory.buffer, ptr, len).set(encoded);
    
        const resPtr = instance.exports.WasmRunGo(ptr, len); // Call the Wasm function
        const output = new TextDecoder().decode(new Uint8Array(memory.buffer, resPtr));
        console.log(output); // Output result from the Go program
    });