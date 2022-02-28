import { wasmBrowserInstantiate } from "../instantiateWasm.js";

const go = new Go(); // Defined in wasm_exec.js. Don't forget to add this in your index.html.

// export let addOrErr = function () {};
// export function calcAddExport() {
// 	console.log("上書き前");
// }

// export function addOrErr2() {
// 	console.log("上書き前2");
// }

// var wasmModule = WebAssembly.WebAssemblyInstantiatedSource;

// var wasmInstance;

// TinyGoのバグを無視するため
// https://github.com/tinygo-org/tinygo/issues/1140#issuecomment-671261465
go.importObject.env["syscall/js.finalizeRef"] = () => {};

const runWasm = async () => {
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

	function checkError(result) {
		if (result != null && "error" in result) {
			console.log("Go return value", result);
			answer.innerHTML = "";
			alert(result.error);
		}
	}

	// w.addOrErr2 = function (value1, value2) {
	// 	console.log("koko", value1, value2);
	// 	var result = wasmModule.instance.exports.calcAdd(value1, value2);
	// 	console.log("koko", result);
	// 	checkError(result);
	// };

	// w.calcAddExport = function (value1, value2) {
	// 	return wasmModule.instance.exports.calcAdd(value1, value2);
	// };

	document.getElementById("addButton4").onclick = function () {
		var result = calcAdd2(value1, value2);
		checkError(result);

		// 	// var value1 = document.getElementById("value1").value;
		// 	// var value2 = document.getElementById("value2").value;
		// 	// console.log("value:", value1, value2);
		// 	// var result = wasmModule.instance.exports.calcAdd();
		// 	// console.log("Go return value", result);
		// 	// checkError(result);
		// 	var result = calcAdd2(value1, value2);
		// 	checkError(result);
		// 	// const res = wasmModule.instance.exports.calcAdd(value1, value2);
		// 	// console.log(`calcAdd Result: ${res}`);
		// 	// const res2 = wasmModule.instance.exports.calcAdd2(value1, value2);
		// 	// console.log(`calcAdd2 Result: ${res2}`);
		// 	// let addr1 = insertText(value1, wasmModule.instance);
		// 	// // Pass the memory location to the relevant function.
		// 	// wasmModule.instance.exports.calcAdd2(addr1, value1.length);
		// 	// console.log(JSON.stringify({ x: value1, y: value2 }));
		// 	// const params = JSON.stringify({ x: value1, y: value2 });
		// 	// let addr = insertText(params, wasmModule.instance);
		// 	// wasmModule.instance.exports.calcAdd3(addr, params.length);
		// 	// wasmModule.instance.exports.calcAdd3("1", "2");
		// 	// const addResult = wasmModule.instance.exports.add(24, 24);
		// 	// // Set the result onto the body
		// 	// console.log(`Hello World! addResult: ${addResult}`);
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
};
runWasm();

// calcAddExport("1", "3");

// // 📁 sayHi.js
// export function sayHi(user) {
// 	console.log(`Hello, ${user}!`);
// }

// //export関数
// export function foo(msg) {
// 	console.log("Hi", msg);
// }

// export function addOrErr3(value1, value2) {
// 	console.log("koko3", value1, value2);
// 	var result = wasmModule.instance.exports.calcAdd(value1, value2);
// 	console.log("koko3", result);
// 	// checkError(result);
// }
