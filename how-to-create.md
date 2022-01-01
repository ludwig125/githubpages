# githubpages の作成

https://docs.github.com/ja/pages/quickstart

を参考に進めます。

![image](https://user-images.githubusercontent.com/18366858/147394273-b0d583c6-fc88-4bb7-b64b-885baa698360.png)

![image](https://user-images.githubusercontent.com/18366858/147394279-aeec9a5d-92b9-4122-8489-2dfc6df0141c.png)

![image](https://user-images.githubusercontent.com/18366858/147394316-8f2ed349-76ee-402d-8d3a-97333c5f93ab.png)

![image](https://user-images.githubusercontent.com/18366858/147394321-444403ec-4c5d-4ba2-947d-bac8d5212281.png)

![image](https://user-images.githubusercontent.com/18366858/147394385-2ecddeb1-a570-4f14-924a-5485198fc364.png)

![image](https://user-images.githubusercontent.com/18366858/147394397-ffd289fd-e38a-465b-bd03-270f17020a9f.png)

この状態で
https://ludwig125.github.io/githubpages/
を見ると以下の通りです。

![image](https://user-images.githubusercontent.com/18366858/147394439-3b240c9b-0ab3-4f5b-a3f3-fae2d31af25a.png)

この時点では code は以下の通りです。

![image](https://user-images.githubusercontent.com/18366858/147394457-3aed0778-c13c-4d79-a5c5-80823e8b0a9c.png)

このコードをいじるために、ターミナルから操作してみます。

```
[~/go/src/github.com/ludwig125] $g clone git@github.com:ludwig125/githubpages.git
Cloning into 'githubpages'...
warning: You appear to have cloned an empty repository.
[~/go/src/github.com/ludwig125] $
[~/go/src/github.com/ludwig125] $cd githubpages
```

gh-pages ブランチに以下のファイルがあります。

```
[~/go/src/github.com/ludwig125/githubpages] $ls
_config.yml index.md
```

\_config.yml を以下のように書き直してみます。

```
theme: jekyll-theme-cayman

title: ludwig125's homepage
description: ludwig125's homepage by githubpages
```

これで commmit して git に push します。
すこし，待つと
https://ludwig125.github.io/githubpages/

以下のようにページに上の説明が加わわりました。（title はタブの上にカーソルを重ねると浮かび上がる）

![image](https://user-images.githubusercontent.com/18366858/147394610-fa6b9508-32ed-4bc9-8efb-d9a143a2d255.png)

ここまでで、基本的な github pages については理解できました。

# WebAssembly

以下では、WebAssembly を使った Web ページの作成方法を確認します。
この後で、github pages 上で、Go Wasm のページを公開することが目的です。

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

上のコードで重要なのは以下の２つです

- `<script src="wasm_exec.js"></script>`
- `WebAssembly.instantiateStreaming`
  - これは Javascript API で、wasm ファイルの読み込みを可能にします

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

# go wasm を github pages で動かす

2021/12/30 の時点で、github pages で Web ページを公開する方法は３通りしかないようです

- master ブランチ
- master ブランチ上の `docs/` フォルダ
- gh-pages ブランチ

- 参考：https://github.community/t/can-i-define-a-custom-source-or-folder-from-which-my-site-hosted-on-github-pages-can-load-from/10237

github リポジトリでは今後`master`ではなく`main`ブランチがデフォルトになったので、今回は`main`ブランチの`docs/`以下に wasm ファイルをおいてみます。

```
[~/go/src/github.com/ludwig125/githubpages] $ls docs
index.html  main.go  main.wasm*  wasm_exec.js
```

これで、以下の通り、`main`ブランチの`docs/`を選んで`Save`します。

![image](https://user-images.githubusercontent.com/18366858/147704500-860b4f75-3973-41fa-afe0-3b9373c5c7de.png)

30 秒ほど待つと、
`https://ludwig125.github.io/githubpages/`に更新が反映されて以下の通り、go wasm の結果が見られるようになりました。

![image](https://user-images.githubusercontent.com/18366858/147704625-c8fa5d48-bcec-46df-935a-81e27ba539af.png)

これで、githubpages で Go の wasm の Web ページを公開することができるようになりました。

以降、main ブランチを修正すれば、この Web ページも更新されるはずです。
毎回反映を待つのが嫌だったり、ローカルで確認したい場合は`goexec`を使えばいいわけです。

# wasm で計算機

もう少し複雑なケースを見てみます。
そこで、
https://github.com/golang/go/wiki/WebAssembly#getting-started
の下にあった
https://tutorialedge.net/golang/go-webassembly-tutorial/
を参考に足し算引き算だけの計算機を作ってみます。

ただ、このページは情報が古かったので、自分なりにかなり改変しました。
その結果が以下です。

## 計算機１（値は固定）

`wasm-calculator` ブランチを`main`から新しく切って修正をします。

### index.html

```html
<html>
	<head>
		<meta charset="utf-8" />
		<title>wasam-calculator</title>
		<link rel="shortcut icon" href="#" />
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
	<body>
		<button onClick="add(2,3);" id="addButton">Add</button>
		<button onClick="subtract(10,3);" id="subtractButton">Subtract</button>
	</body>
</html>
```

#### 説明

1. `<title>wasam-calculator</title>`

- Web ページのタイトルをつけてみました
- Chrome ではこれがタブに表示されます

2. <link rel="shortcut icon" href="#" />

`shortcut icon`の役割は、のように設定して任意の画像をタブに出すことです。

```
<link rel="shortcut icon" href="名前" type="＜画像のパス＞">
```

この設定がないと Console 上で以下のような`favicon.ico 404 (Not Found)`のエラーが出ます
![image](https://user-images.githubusercontent.com/18366858/147839432-73302827-80e2-486a-bae8-6b3d80b86739.png)

3. button

- `<button onClick="add(2,3);" id="addButton">Add</button>` のように、クリックされると`add`関数に２と３を引数に与えて実行します
- この`add`と`subtract`の処理内容は後述の Go プログラムで定義します

### main.go

```golang
package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	c := make(chan struct{})

	fmt.Println("Hello, WebAssembly!")
	registerCallbacks()
	<-c
}

func add(this js.Value, args []js.Value) interface{} {
	println(args[0].Int() + args[1].Int())
	return nil
}

func subtract(this js.Value, args []js.Value) interface{} {
	println(args[0].Int() - args[1].Int())
	return nil
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}
```

#### 説明

上のコードについて説明を書きます。

1. `"syscall/js"`

- Go で js の操作を行うためには syscall/js という標準パッケージを import する必要があります

2. `c := make(chan struct{})`と`<-c`

- ボタンを押すなどのイベント処理をするときにこれが必要になります
- イベント処理では、まず Web ページが表示されて、そのあとユーザがボタンを押して対応する処理が走るいう順番になりますが、Go のプログラムを普通に終わらせてしまうと、ボタンを押されても対応する処理ができずに以下のように`Uncaught Error: Go program has already exited`のエラーが発生します

![image](https://user-images.githubusercontent.com/18366858/147792725-59c3dbfd-9633-419a-a35b-157f07375356.png)

- channel を使うことで main 関数の実行が終了するのを防ぐことができます。
- channel を使う以外に `select {}` のように select で待ち続けることでプログラムの終了を防ぐやり方をしている人もいるようです

3. `registerCallbacks()`

- `js.Global().Set("property名", property)` で Javascript の property を登録することができます
  - https://pkg.go.dev/syscall/js#Value.Set
- ここで登録する`add`と`subtract`関数は前述の HTML に対応するものです

![image](https://user-images.githubusercontent.com/18366858/147708812-809133ba-cd13-4527-bc86-1b24ef3a68f6.png)

4. `js.FuncOf()`

- JavaScript の関数を返します
- この関数は以前は`js.NewCallback`という名前でしたが、Go1.12 で名前もインターフェースも大きく変わりました。そのため少し古い資料では`js.FuncOf()`ではなく`js.NewCallback`が多く使われていて、混乱の原因になっています
  - https://pkg.go.dev/syscall/js#FuncOf

5. `add`と`subtract`関数

- 上の`js.FuncOf()`の package の定義に沿って、`(this js.Value, args []js.Value)` を引数として取って、`interface{}` を返す関数です
- `args[0].Int()`のように引数２つをそれぞれ Int 型にしてから足しています。

### 実行

ここまでで保存して、以下の通り build してサーバを立ち上げます

```
$ GOOS=js GOARCH=wasm go build -o main.wasm
$ goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
```

ブラウザを見ると以下のように、`Add`ボタンや`Subtract`ボタンを押すと Console 上に結果が出力されます

![image](https://user-images.githubusercontent.com/18366858/147839642-c42a6f7d-bb36-4d04-a30d-4f931a01dac2.png)

## 計算機２（値は任意）

決まった数の足し算引き算では面白くないので、TextBox に数字を入力できるようにします。

`wasm-calculator`から新たに`wasm-calculator2`ブランチを切ります

### index.html

以下のように修正を加えます

```
        <body>
-               <button onClick="add(2,3);" id="addButton">Add</button>
-               <button onClick="subtract(10,3);" id="subtractButton">Subtract</button>
+               <input type="text" id="value1" />
+               <input type="text" id="value2" />
+
+               <button onClick="add('value1', 'value2');" id="addButton">Add</button>
+               <button onClick="subtract('value1', 'value2');" id="subtractButton">Subtract</button>
+
+               <div align="left">answer:</div>
+               <div id="answer"></div>
        </body>
```

#### 説明

- `add(2,3);`の代わりに、`text`入力値を`value1`,`value2`として受け取り、これを`add`や`subtract`に渡すようにしました
- 後述の Go プログラム側で、`<div id="answer"></div>`に計算結果を出力するようにします

### main.go

```go
package main

import (
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	c := make(chan struct{})

	fmt.Println("Hello, WebAssembly!")
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("add", js.FuncOf(add))
	js.Global().Set("subtract", js.FuncOf(subtract))
}

func add(this js.Value, args []js.Value) interface{} {
	value1 := textToStr(args[0])
	value2 := textToStr(args[1])

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)
	fmt.Println("int1:", int1, " int2:", int2)
	ans := int1 + int2

	printAnswer(ans)
	return nil
}

func subtract(this js.Value, args []js.Value) interface{} {
	value1 := textToStr(args[0])
	value2 := textToStr(args[1])

	int1, _ := strconv.Atoi(value1)
	int2, _ := strconv.Atoi(value2)
	fmt.Println("int1:", int1, " int2:", int2)
	ans := int1 - int2

	printAnswer(ans)
	return nil
}

func textToStr(v js.Value) string {
	return js.Global().Get("document").Call("getElementById", v.String()).Get("value").String()
}

func printAnswer(ans int) {
	println(ans)
	js.Global().Get("document").Call("getElementById", "answer").Set("innerHTML", ans)
}

```

#### 説明

1. `textToStr`

- HTML の一行 Text ボックスを`getElementById`で取得します
- この関数で、Javascript の世界の値を Go の文字列として変換しています

2. `printAnswer`

- 計算結果を Print して、そのあと HTML 側で用意した`answer`に値をセットします

### 実行結果

![image](https://user-images.githubusercontent.com/18366858/147861008-8b017e5d-8516-4fc7-9e04-8d20fa65820e.png)

左のテキスト入力欄と右のテキスト入力欄の値の和や差が answer としてブラウザ上にプリントされることが確認できました
