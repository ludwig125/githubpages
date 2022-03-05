import { wasmBrowserInstantiate } from "../instantiateWasm.js";

const go = new Go(); // Defined in wasm_exec.js. Don't forget to add this in your index.html.

// TinyGoのバグを無視するため
// https://github.com/tinygo-org/tinygo/issues/1140#issuecomment-671261465
go.importObject.env["syscall/js.finalizeRef"] = () => {};

function checkError(result) {
	if (result != null && "error" in result) {
		console.log("Go return value", result);
		answer.innerHTML = "";
		alert(result.error);
	}
}

const runWasm = async (w) => {
	// Get the importObject from the go instance.
	const importObject = go.importObject;

	// wasm moduleのインスタンスを作成
	const wasmModule = await wasmBrowserInstantiate("./calc.wasm", importObject);
	// wasmModule = await wasmBrowserInstantiate("./calc.wasm", importObject);

	go.run(wasmModule.instance);
	// wasmInstance = wasmModule.instance
	// go.run(wasmInstance);

	// Goの関数の実行
	// wasmModule.instance.exports.calcAdd("1", "3");

	// calcAddExport = wasmModule.instance.exports.calcAdd;

	document.getElementById("addButton4").onclick = function () {
		var value1 = document.getElementById("value1").value;
		var value2 = document.getElementById("value2").value;
		console.log("value:", value1, value2);
		// const res = wasmModule.instance.exports.calcAdd(value1, value2);
		// console.log(`calcAdd Result: ${res}`);

		// const res2 = wasmModule.instance.exports.calcAdd2(value1, value2);
		// console.log(`calcAdd2 Result: ${res2}`);

		// let addr1 = insertText(value1, wasmModule.instance);
		// // Pass the memory location to the relevant function.
		// wasmModule.instance.exports.calcAdd2(addr1, value1.length);

		console.log(JSON.stringify({ x: value1, y: value2 }));

		const params = JSON.stringify({ x: value1, y: value2 });
		let addr = insertText(params, wasmModule.instance);
		const result = wasmModule.instance.exports.calcAdd3(addr, params.length);
		console.log("result", result);
		checkError(result);

		// wasmModule.instance.exports.calcAdd3("1", "2");
		// const addResult = wasmModule.instance.exports.add(24, 24);
		// // Set the result onto the body
		// console.log(`Hello World! addResult: ${addResult}`);
	};
	// calcAddExport("1", "3");
	// function checkError(result) {
	// if (result != null && "error" in result) {
	//     console.log("Go return value", result);
	//     answer.innerHTML = "";
	//     alert(result.error);
	// }
	// }
	// addOrErr = function (value1, value2) {

	// https://stackoverflow.com/questions/49338193/how-to-use-code-from-script-with-type-module

	// addOrErr = function addOrErr(value1, value2) {
	//     var result = wasmModule.instance.exports.calcAdd(value1, value2);
	//     checkError(result);

	//     console.log("Go return value", result);
	// };

	// var subtractOrErr = function (value1, value2) {
	//     var result = calcSubtract(value1, value2);
	//     checkError(result);
	// };
	// wasmModule.instance.exports.setTimeZone();
	// setInterval(wasmModule.instance.exports.clock, 200);
	// document.getElementById("in").addEventListener("input", wasmModule.instance.exports.convTime);

	// 変換

	const mystring2 = "これはテスト2";
	// var array_to_pass = new Uint8Array([0, 9, 21, 32]);
	var array_to_pass = new TextEncoder().encode(mystring2);
	console.log("array_to_pass:", array_to_pass);
	PassUint8ArrayToGo(array_to_pass);
	console.log("array_to_pass2:", array_to_pass);

	var array_to_set = new Uint8Array(30);
	// const mystring = "これはテスト";
	// var array_to_set = new TextEncoder().encode(mystring);
	// console.log("array_to_set:", array_to_set);
	SetUint8ArrayInGo(array_to_set);
	console.log("array_to_set2:", array_to_set);
	console.log(array_to_set[0]); // 42
	console.log(array_to_set.length); // 2
	var str = new TextDecoder().decode(array_to_set);
	console.log(str);

	const mystring3 = "これはテスト3";
	var array_to_pass3 = new TextEncoder().encode(mystring3);
	wasmModule.instance.exports.passUint8ArrayToGo(array_to_pass3);
};
runWasm();

// calcAddExport("1", "3");

//  sayHi.js
export function sayHi(user) {
	console.log(`Hello, ${user}!`);
}

//export関数
export function foo(msg) {
	console.log("Hi", msg);
}

export function addOrErr3(value1, value2) {
	console.log("koko3", value1, value2);
	var result = wasmModule.instance.exports.calcAdd(value1, value2);
	console.log("koko3", result);
	// checkError(result);
}

function insertText(text, module) {
	// Get the address of the writable memory.
	let addr = module.exports.getBuffer();
	let buffer = module.exports.memory.buffer;

	let mem = new Int8Array(buffer);
	let view = mem.subarray(addr, addr + text.length);

	for (let i = 0; i < text.length; i++) {
		view[i] = text.charCodeAt(i);
	}

	// Return the address we started at.
	return addr;
}
