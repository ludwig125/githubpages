import { wasmBrowserInstantiate } from "../instantiateWasm.js";

const go = new Go(); // Defined in wasm_exec.js. Don't forget to add this in your index.html.

// TinyGoのバグを無視するため
// https://github.com/tinygo-org/tinygo/issues/1140#issuecomment-671261465
go.importObject.env["syscall/js.finalizeRef"] = () => {};

const runWasm = async () => {
	// Get the importObject from the go instance.
	const importObject = go.importObject;

	// wasm moduleのインスタンスを作成
	const wasmModule = await wasmBrowserInstantiate("./calc.wasm", importObject);

	go.run(wasmModule.instance);
};
runWasm();
