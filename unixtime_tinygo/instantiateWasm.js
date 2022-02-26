export const wasmBrowserInstantiate = async (wasmModuleUrl, importObject) => {
	let response = undefined;

	if (!importObject) {
		importObject = {
			env: {
				abort: () => console.log("Abort!"),
			},
		};
	}

	response = await WebAssembly.instantiateStreaming(
		fetch(wasmModuleUrl),
		importObject
	);

	return response;
};
