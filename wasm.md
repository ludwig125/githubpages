# WebAssembly

https://webassembly.org/

公式の説明

> WebAssembly (abbreviated Wasm) is a binary instruction format for a stack-based virtual machine. Wasm is designed as a portable compilation target for programming languages, enabling deployment on the web for client and server applications.

```
WebAssembly（略称：Wasm）は、stack-baseの仮想マシン用のバイナリ命令形式です。
Wasmは、プログラミング言語用のポータブルなコンパイルターゲットとして設計されており、クライアントおよびサーバーアプリケーションのWeb上でのdeployを可能にします。
```

> The Wasm stack machine is designed to be encoded in a size- and load-time-efficient binary format. WebAssembly aims to execute at native speed by taking advantage of common hardware capabilities available on a wide range of platforms.

```
Wasmスタックマシンは、サイズとロード時間の効率的なバイナリ形式でエンコードされるように設計されています。
WebAssemblyは、幅広いプラットフォームで利用可能な一般的なハードウェア機能を活用することで、ネイティブスピードで実行することを目指しています。
```

## 補足説明

JavaScript より速度を上回ることを目標にしているらしいです。
Go でも、Go1.11 から標準の機能として Go のコードを WebAssembly にコンパイルする機能が追加されました。

今の自分の Go のバージョンは以下の通りでした。

```bash
[~/go/src/github.com/ludwig125/githubpages] $go version
go version go1.17 linux/amd64
```

# Go WebAssembly

https://github.com/golang/go/wiki/WebAssembly#getting-started

を参考に進めます。

## Getting Started

まずは簡単なプログラムを作成します。

main.go

```
package main

import "fmt"

func main() {
	fmt.Println("Hello, WebAssembly!")
}
```

このコードを WebAssembly 形式で、build するには以下のようにします。

Go にはクロスコンパイルという機能で、別のアーキテクチャや別の OS 向けのバイナリをビルドすることができます。
ここでは、 `GOOS`を`js`に、`GOARCH` を `wasm`にすることで、wasm 用のファイルにしています。

また、`-o` で`main.wasm`を指定したので、この名前の実行可能な WebAssembly ファイルが作られることになります。

```
$ GOOS=js GOARCH=wasm go build -o main.wasm
```

この `main.wasm`をブラウザ上で実行するために、Javascript と HTML が必要になります。

Go の最近のバージョンにはデフォルトで wasm 用の javascript(js)が同封されているので、それを以下のように手元に持ってきます。

```
$ cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" .
```

また、以下の通り、HTML ファイルを作成します。

```html
<html>
	<head>
		<meta charset="utf-8" />
		<script src="wasm_exec.js"></script>
		<script>
			const go = new Go();
			WebAssembly.instantiateStreaming(
				fetch("main.wasm"),
				go.importObject
			).then((result) => {
				go.run(result.instance);
			});
		</script>
	</head>
	<body></body>
</html>
```

https://github.com/golang/go/wiki/WebAssembly#getting-started
には、ブラウザが`WebAssembly.instantiateStreaming`に対応していない場合は `polyfill`を使うようにと書かれていますが、私の環境では普通に実行できたのでここではこのまま使用しました。

ここまでの段階で以下のファイルが存在します。

```
[~/go/src/github.com/ludwig125/githubpages] $ls
index.html  main.go  main.wasm*  wasm_exec.js
```

これを Web サーバ上で実行するために、 `goexec` を使います。
もちろん、別途 Go でサーバプログラムを作ってもいいです( 例：https://go.dev/play/p/pZ1f5pICVbV )が、ここでは公式ドキュメントに従って以下のように goexec でサーバを立てます。

goexec の install(初回のみ)

```
$ go get -u github.com/shurcooL/goexec
```

goexec でサーバ起動（ここでは Port 8080 でサーバを立ち上げています）

```
$ goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
```

注意：うまく動かないときは以下の通り Go の環境設定をする必要があります

- https://go.dev/doc/install

サーバ起動した状態でブラウザで`http://localhost:8080/`にアクセスします。
ちなみに、公式ドキュメントには`http://localhost:8080/index.html` となっていますが、普通の Web サーバでは `http://localhost:8080/`のようにスラッシュで終わる URL にアクセスすると自動で`index.html`を探すようになっているので同じことです。

この Web ページ上で、JavaScript のデバッグコンソールを開きます。
Chrome では、`F12` で開けます。

![image](https://user-images.githubusercontent.com/18366858/147610530-04f66edc-d75b-4aef-b222-98e871f98aa7.png)
