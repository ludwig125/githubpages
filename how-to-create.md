# 目的

WebAssembly 略称 WASM に興味があったので、Go で Web ツールを作成しました。
Web ページを無料で作れるところを探したところ、
github pages が良さそうだったのでこれを使ってみました。

## 環境と言語

私は Windows 上の WSL で Ubuntu20.04 を使っています。

```
$cat /etc/os-release
NAME="Ubuntu"
VERSION="20.04.3 LTS (Focal Fossa)"
ID=ubuntu
ID_LIKE=debian
PRETTY_NAME="Ubuntu 20.04.3 LTS"
VERSION_ID="20.04"
```

なお、この記事に登場する言語は Go、HTML、Javascript ですが、
私は Go は数年の開発経験があるものの、**HTML と Javascript はほぼ無知の素人** なので、手探りでの開発となりました。

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

WebAssembly を使用すると、JavaScript と同じように Rust、C、Go などの言語で Web ツールを作成できます。これにより、既存のライブラリを移植したり、JavaScript で利用できない機能を活用したりできます。

また、WebAssembly はバイナリ形式にコンパイルされるためコードの高速実行が可能になります。
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

polyfill

- https://github.com/golang/go/blob/b2fcfc1a50fbd46556f7075f7f1fbf600b5c9e5d/misc/wasm/wasm_exec.html#L17-L22

この辺の WASM を使う場合の説明は以下が詳しいです

- https://developer.mozilla.org/en-US/docs/WebAssembly/Loading_and_running

  > Fetch を使用する

- https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/WebAssembly/instantiateStreaming

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

また、goexec 実行時に以下のようなエラーが出た場合は、すでに同じ Port で goexec を起動していてバッティングしている可能性があります

```
(*net.OpError)(&net.OpError{
        Op:     (string)("listen"),
        Net:    (string)("tcp"),
        Source: (net.Addr)(nil),
        Addr: (*net.TCPAddr)(&net.TCPAddr{
                IP:   (net.IP)(nil),
                Port: (int)(8080),
                Zone: (string)(""),
        }),
        Err: (*os.SyscallError)(&os.SyscallError{
                Syscall: (string)("bind"),
                Err:     (syscall.Errno)(0x62),
        }),
})
```

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
- Go 側で関数を定義して、イベント発生時に javascript として実行されるものなのでいわゆる Callback 関数です

![image](https://user-images.githubusercontent.com/18366858/147708812-809133ba-cd13-4527-bc86-1b24ef3a68f6.png)

4. `js.FuncOf()`

- JavaScript の関数を返します
- この関数は以前は`js.NewCallback`という名前でしたが、Go1.12 で名前もインターフェースも大きく変わりました。そのため少し古い資料では`js.FuncOf()`ではなく`js.NewCallback`が多く使われていて、混乱の原因になっています
  - https://pkg.go.dev/syscall/js#FuncOf

5. `add`と`subtract`関数

- 上の`js.FuncOf()`の package の定義に沿って、`(this js.Value, args []js.Value)` を引数として取って、`interface{}` を返す関数です
- `args[0].Int()`のように引数２つをそれぞれ Int 型にしてから足しています。
- この引数のうち、`this`は JavaScript の [global object](https://developer.mozilla.org/en-US/docs/Glossary/Global_object) で、`args`は `add`（または`subtract`）関数に与えられる引数に相当します

#### `Value`について

この`Value`が曲者です。

これが Javascript の世界と Go の世界の橋渡しをするものですが、型が動的なので、
例えば Int に変換しようとしてできない、などの場合にいとも簡単に Panic します

どこで問題が起きたのか非常に分かりにくいです

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

## [脱線] Go の WASM はライブラリではなくアプリケーションである

`GoのWASMはライブラリではなくアプリケーションである` この言葉が最初が分かりませんでしたが、以下のような意味だと理解しています

- C/C++/Rust などの言語の WASM では、JavaScript に変換して「ライブラリ」として扱うことができる
- Go の WASM は、「アプリケーション」なので、HTML 側から`実行`しないといけない

そのため、イベント処理をするときは Go 側で終了させないようにチャネルで永久に待たせるとか、HTML 側で以下のように`go.run`で Go を実行させる処理が必要になります

```
const go = new Go();
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
	go.run(result.instance);
});
```

## 計算機３（エラーハンドリング）

前述までで、計算機としての最低限の機能は作れましたが、いくつか重要な欠点があります。

1. 数値のバリデーションチェックがない＆エラーハンドリングできていない

- テキスト欄に`a`や`あ`など、整数変換ができないものが入力された場合、`int1, err := strconv.Atoi("a")`の結果、int1 には `0`が設定されてしまいます
- このとき、`err`を適切にエラーハンドリングしたいです

2. Web ページ上でエラーが分かりにくい

- 上のエラーハンドリングができたら、Web ページにエラーメッセージを出して不正な入力値であることを分かりやすくしたいです

3. Panic を起こしやすい

```
js.Global().Get("document").Call("getElementById", v.String()).Get("value").String()
```

- 例えば`textToStr`関数のこの式ですが、`getElementById`で対象の ID が取得できない状態で`Get`メソッドを呼ぶと Panic を起こします
- 同様に、`Get("value")`の結果が空の時に`String`メソッドを呼んでも Panic となります
- 可能な限り Panic で異常終了しないようにしたいです

そこで、以下の資料を参考に次の通り修正しました

- https://golangbot.com/go-webassembly-dom-access/
- https://dev.bitolog.com/go-in-the-browser-using-webassembly/

`wasm-calculator2`から新たに`wasm-calculator3`ブランチを切って修正しました

### main.go

修正後のコードを最初に書くと以下の通りです。

```go
package main

import (
	"errors"
	"fmt"
	"strconv"
	"syscall/js"
)

func main() {
	registerCallbacks()
	<-make(chan struct{})
}

func registerCallbacks() {
	js.Global().Set("calcAdd", calculatorWrapper("add"))
	js.Global().Set("calcSubtract", calculatorWrapper("subtract"))
}

func calculatorWrapper(ope string) js.Func {
	calcFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		value1, err := getJSValue(args[0].String())
		if err != nil {
			return wrapResult("", err)
		}
		value2, err := getJSValue(args[1].String())
		if err != nil {
			return wrapResult("", err)
		}
		fmt.Println("value1:", value1, " value2:", value2)

		int1, err := strconv.Atoi(value1)
		if err != nil {
			return wrapResult("", fmt.Errorf("failed to convert value1 to int: %v", err))
		}
		int2, err := strconv.Atoi(value2)
		if err != nil {
			return wrapResult("", fmt.Errorf("failed to convert value2 to int: %v", err))
		}

		var ans int
		switch ope {
		case "add":
			ans = int1 + int2
		case "subtract":
			ans = int1 - int2
		default:
			return wrapResult("", fmt.Errorf("invalid operation: %s", ope))
		}
		fmt.Println("Answer:", ans)

		if err := setJSValue("answer", ans); err != nil {
			return wrapResult("", err)
		}
		return nil
	})
	return calcFunc
}

func getJSValue(elemID string) (string, error) {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return "", errors.New("failed to get document object")
	}

	jsElement := jsDoc.Call("getElementById", elemID)
	if !jsElement.Truthy() {
		return "", fmt.Errorf("failed to getElementById: %s", elemID)
	}

	jsValue := jsElement.Get("value")
	if !jsValue.Truthy() {
		return "", fmt.Errorf("failed to Get value: %s", elemID)
	}
	return jsValue.String(), nil
}

func setJSValue(elemID string, value interface{}) error {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return errors.New("failed to get document object")
	}

	jsElement := jsDoc.Call("getElementById", elemID)
	if !jsElement.Truthy() {
		return fmt.Errorf("failed to getElementById: %s", elemID)
	}
	jsElement.Set("innerHTML", value)
	return nil
}

func wrapResult(result string, err error) map[string]interface{} {
	return map[string]interface{}{
		"error":    err.Error(),
		"response": result,
	}
}

```

#### 説明

分かりやすいところから書きます

##### 1. `textToStr`関数を修正して`getJSValue`に改名

```go
func getJSValue(elemID string) (string, error) {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return "", errors.New("failed to get document object")
	}
略
}
```

- [Truthy](https://pkg.go.dev/syscall/js#Value.Truthy) メソッドはオブジェクトが`false, 0, "", null, undefined, NaN`のどれかの時に`false`を返します
- これを使うことで Panic を起こす前にエラーを返して呼び出しもとでエラーハンドリングできるようになります
- 関数名はより汎用的に`getJSValue`にしました

##### 2. `printAnswer`関数を修正して`setJSValue`に改名

```go
func setJSValue(elemID string, value interface{}) error {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return errors.New("failed to get document object")
	}

	jsElement := jsDoc.Call("getElementById", elemID)
	if !jsElement.Truthy() {
		return fmt.Errorf("failed to getElementById: %s", elemID)
	}
	jsElement.Set("innerHTML", value)
	return nil
}
```

- こちらも`getJSValue`と同様に[Truthy](https://pkg.go.dev/syscall/js#Value.Truthy) で逐一判定するようにしました
- また、値を設定したい要素の ID を`elemID`として、設定する値を`value`として引数にすることで任意の ID に対して設定できるようにしました
- 合わせて関数名も print よりも set の方がふさわしいことと、より汎用的にするため`setJSValue`に変えました

##### 3. `add`と`subtract`関数を統合して`calculatorWrapper`でラップ

```go
func calculatorWrapper(ope string) js.Func {
	calcFunc := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		value1, err := getJSValue(args[0].String())
		if err != nil {
			return wrapResult("", err)
		}
		value2, err := getJSValue(args[1].String())
		if err != nil {
			return wrapResult("", err)
		}
		fmt.Println("value1:", value1, " value2:", value2)

		int1, err := strconv.Atoi(value1)
		if err != nil {
			return wrapResult("", fmt.Errorf("failed to convert value1 to int: %v", err))
		}
		int2, err := strconv.Atoi(value2)
		if err != nil {
			return wrapResult("", fmt.Errorf("failed to convert value2 to int: %v", err))
		}

		var ans int
		switch ope {
		case "add":
			ans = int1 + int2
		case "subtract":
			ans = int1 - int2
		default:
			return wrapResult("", fmt.Errorf("invalid operation: %s", ope))
		}
		fmt.Println("Answer:", ans)

		if err := setJSValue("answer", ans); err != nil {
			return wrapResult("", err)
		}
		return nil
	})
	return calcFunc
}

func wrapResult(result string, err error) map[string]interface{} {
	return map[string]interface{}{
		"error":    err.Error(),
		"response": result,
	}
}

```

- 今までは`js.FuncOf`の中身の関数を`add`や`subtract`としていましたが、それらをラップして`calculatorWrapper`にしました
- これにより、`js.FuncOf`のインターフェースに縛られず、今回の`ope`のように自由に引数を与えることができます
- 今回の場合は、`add`と`subtract`には共通部分が多かったのでこれらを統合して、演算部分だけ`ope`に応じて`switch`で条件分岐させるようにしました

`wrapResult`:

- `getJSValue`や`setJSValue`で返したエラーと返り値をこれでラップしています
- `map[string]interface{}`として返すことで、後述の javascript でエラーハンドリングできるようになります
- 今回`wrapResult`の中の`response`は全部空にしているので使いません。コールバック関数から値を返したい場合はここに値を設定します

##### 4. `registerCallbacks`の中で引数を指定して`calculatorWrapper`を呼ぶ

```go
func main() {
	registerCallbacks()
	<-make(chan struct{})
}

func registerCallbacks() {
	js.Global().Set("calcAdd", calculatorWrapper("add"))
	js.Global().Set("calcSubtract", calculatorWrapper("subtract"))
}
```

- `calculatorWrapper`で統合したので、`add`と`subtract`は与える引数の違いだけになりました
  - `calcAdd`と`calcSubtract` は後述の`index.html`の javascript で使います
- `<-make(chan struct{})`ここは、channel の定義とまとめたほうが簡潔なのでこのようにしました

### index.html

index.html は以下のように修正しました。

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
		<input type="text" id="value1" />
		<input type="text" id="value2" />

		<button onClick="addOrErr('value1', 'value2');" id="addButton">Add</button>
		<button onClick="subtractOrErr('value1', 'value2');" id="subtractButton">
			Subtract
		</button>

		<div align="left">answer:</div>
		<div id="answer"></div>

		<script>
			function checkError(result) {
				if (result != null && "error" in result) {
					console.log("Go return value", result);
					answer.innerHTML = "";
					alert(result.error);
				}
			}

			var addOrErr = function (value1, value2) {
				var result = calcAdd(value1, value2);
				checkError(result);
			};
			var subtractOrErr = function (value1, value2) {
				var result = calcSubtract(value1, value2);
				checkError(result);
			};
		</script>
	</body>
</html>
```

#### 説明

- いままで`onClick`で直接 Go で書いた`add`コールバック関数を呼び出していましたが、ここでは`addOrErr`という新しく定義した関数を呼び出しています
- `addOrErr`の中身を分かりやすいように`checkError`部分を展開して書くと以下の通りです

```js
var addOrErr = function (value1, value2) {
	var result = calcAdd(value1, value2);
	if (result != null && "error" in result) {
		console.log("Go return value", result);
		answer.innerHTML = "";
		alert(result.error);
	}
};
```

- この関数は、テキスト欄から入力された`value1`, `value2`を引数として取ります
- 内部で、Go 側で用意した`calcAdd`コールバック関数を呼び出して`result`を返します
- この`result`には`wrapResult`で入れたマップデータが入っています
- そこで、`result.error`を見ることで Go 側の処理でエラーを返したかどうかが判定できます
- ここでは、エラーがある場合は answer の値を空にして、alert でポップアップを出すようにしています
- 【注意点】今回は、`answer`が div の HTML タグなので`innerHTML`を使っていますが、もし`answer`が`input`や`textarea`などの入力フォームの場合は`answer.value = "";`とするのが正しいです

### 実行結果

テキスト欄に、`5`と`3`を入れて`Add`ボタンを押すと以下のように`8`が表示されます（計算機２と同じ）
![image](https://user-images.githubusercontent.com/18366858/150023478-828facb9-973e-4a6e-9b40-e6e1af12e053.png)

テキスト欄に、`5`と`3`を入れて`Subtract`ボタンを押すと以下のように`2`が表示されます（計算機２と同じ）

![image](https://user-images.githubusercontent.com/18366858/150023507-74f9e023-7e06-4261-82d0-8973b14e3ce6.png)

`5`の代わりに`a`などの数値変換できない文字を入れると、
Go で設定した`failed to convert value1 to int: strconv.Atoi: parsing "a": invalid syntax` のエラーがポップアップとして表示されます。

また、Console に`Go return value`が表示されていることが分かります

![image](https://user-images.githubusercontent.com/18366858/150023588-bd920961-15b2-45b3-8548-3b5c368ce053.png)

ポップアップを閉じると`answer`の中身が消えています

- `answer`が空になってからポップアップが表示されると思っていましたがよしとします
- ここの実行順序は分かっていません

![image](https://user-images.githubusercontent.com/18366858/150216750-2b5e8bc3-e17a-4346-a3a6-3d02cae73e42.png)

以上で、エラーハンドリングまで対応できるようになりました

# Unixtime 変換ツール

上の加算減算しかできない計算機より少しは使い道のありそうな、Unixtime を JST の日付に変換するツールを作ってみました

いきなりコードを載せると以下の通りです

unixtime.go

```go
package main

import (
	"errors"
	"fmt"
	"strconv"
	"syscall/js"
	"time"
)

func main() {
	unixtime()

	<-make(chan struct{})
}

func unixtime() {
	// time zoneを最初に表示させる
	js.Global().Call("queueMicrotask", js.FuncOf(setTimeZone))
	// 二度と使わない関数はメモリを解放する
	js.FuncOf(setTimeZone).Release()

	// 一定時間おきにclockを呼び出す
	js.Global().Call("setInterval", js.FuncOf(clock), "200")

	getElementByID("in").Call("addEventListener", "input", js.FuncOf(convTime))
}

func setTimeZone(this js.Value, args []js.Value) interface{} {
	t := time.Now()
	zone, _ := t.Zone()
	return setJSValue("time_zone", fmt.Sprintf("(%s)", zone))
}

func setJSValue(elemID string, value interface{}) error {
	jsDoc := js.Global().Get("document")
	if !jsDoc.Truthy() {
		return errors.New("failed to get document object")
	}

	jsElement := jsDoc.Call("getElementById", elemID)
	if !jsElement.Truthy() {
		return fmt.Errorf("failed to getElementById: %s", elemID)
	}
	jsElement.Set("innerHTML", value)
	return nil
}

func getElementByID(targetID string) js.Value {
	return js.Global().Get("document").Call("getElementById", targetID)
}

func clock(this js.Value, args []js.Value) interface{} {
	nowStr, nowUnix := getNow(time.Now())

	getElementByID("clock").Set("textContent", nowStr)
	getElementByID("clock_unixtime").Set("textContent", nowUnix)
	return nil
}

func convTime(this js.Value, args []js.Value) interface{} {
	in := getElementByID("in").Get("value").String()
	date, err := unixtimeToDate(in)
	if err != nil {
		getElementByID("out").Set("value", js.ValueOf("不正な時刻です"))
		return nil
	}
	getElementByID("out").Set("value", js.ValueOf(date))
	return nil
}

func getNow(now time.Time) (string, string) {
	s := now.Format("2006-01-02 15:04:05")
	unix := now.Unix()
	return s, fmt.Sprintf("%d", unix)
}

func unixtimeToDate(s string) (string, error) {
	unixtime, err := strconv.Atoi(s)
	if err != nil {
		return "", err
	}
	date := time.Unix(int64(unixtime), 0)
	layout := "2006-01-02 15:04:05" // Goの時刻フォーマットではこれで時分秒まで取れる
	return date.Format(layout), nil
}

```

index.html

```html
<html>
	<head>
		<meta charset="utf-8" />
		<title>unixtime</title>
		<link rel="shortcut icon" href="#" />
		<script src="wasm_exec.js"></script>
		<script>
			const go = new Go();
			WebAssembly.instantiateStreaming(
				fetch("unixtime.wasm"),
				go.importObject
			).then((result) => {
				go.run(result.instance);
			});
		</script>
	</head>
	<body>
		<h1>UnixTimeを日付に変換するツール</h1>
		<table border="1" align="center" width="600" height="100">
			<tr align="center">
				<td>
					現在時刻<br />
					<div id="time_zone">time_zone</div>
				</td>
				<td><div id="clock"></div></td>
				<td><div id="clock_unixtime"></div></td>
			</tr>
		</table>

		<hr />
		<table border="1" align="center" width="300" height="200">
			<tr align="center">
				<td valign="middle">変換対象の時刻</td>
				<td><input type="text" id="in" /></td>
			</tr>
			<tr align="center">
				<td valign="middle">変換後の時刻</td>
				<td>
					<input type="text" id="out" />
					<button onclick="document.getElementById('out').value = ''">
						Clear
					</button>
				</td>
			</tr>
		</table>
	</body>
</html>
```

## ツールの概要

このツールでは大きく分けて 3 つの機能を作りました

1. リアルタイムで現在時刻を表示し続ける機能
2. テキスト欄に入力された Unixtime を日付時分秒に変換する機能
3. タイムゾーンの表示

順番に見ていきます

### 説明 1. リアルタイムで現在時刻を表示し続ける機能

Go では`clock`関数がこの機能を担当します

- まず`getNow(time.Now())`で現在時刻を取得して、それをもとに日付と時分秒の`nowStr`と Unixtime の`nowUnix`を作成します
- これらをそれぞれ、`getElementByID`で取得した HTML のタグ、`clock`と`clock_unixtime`に設定しています
- ポイントは、この`clock`関数を、[setInterval](https://developer.mozilla.org/ja/docs/Web/API/setInterval)を使って 200 ミリ秒ごとに実行されるようにしていることです
  - `js.Global().Call("setInterval", js.FuncOf(clock), "200")`
- これにより、Web ツールの表示中、200 ミリ秒ごとに現在時刻が更新されるようになります

HTML 側は以下が対応します

```html
<table border="1" align="center" width="600" height="100">
	<tr align="center">
		<td>現在時刻</td>
		<td><div id="clock"></div></td>
		<td><div id="clock_unixtime"></div></td>
	</tr>
</table>
```

### 説明 2. テキスト欄に入力された Unixtime を日付時分秒に変換する機能

Go では、`convTime`関数がこの機能を担当します

- この関数では、HTML の`in`テキスト欄に入力された文字を`unixtimeToDate`関数で変換し、変換後の文字列を HTML の`out`テキスト欄に設定します
- この時、Unixtime として間違ったものを`in`に入力すると、`out`に`不正な時刻です`と出すようにしました
- ポイントは、[addEventListener](https://developer.mozilla.org/ja/docs/Web/API/EventTarget/addEventListener)を使って、`in`テキスト欄に入力があったら（`input`があったら）`convTime`が実行されるようにしたことです
- これにより、`in`に入力したのと同時に`out`に変換後の値が表示されるようになります

HTML 側は以下のコードが対応します

```html
<table border="1" align="center" width="300" height="200">
	<tr align="center">
		<td valign="middle">変換対象の時刻</td>
		<td><input type="text" id="in" /></td>
	</tr>
	<tr align="center">
		<td valign="middle">変換後の時刻</td>
		<td>
			<input type="text" id="out" />
			<button onclick="document.getElementById('out').value = ''">Clear</button>
		</td>
	</tr>
</table>
```

- テキスト欄`out`の文字をクリアするボタンをつけたくなったのですが、これは次のように直下に書いたほうが(Go 側で実装して呼び出すよりも)簡単なのでこうしました
- `<button onclick="document.getElementById('out').value = ''">Clear</button>`

### 説明 3. タイムゾーンの表示

- `現在時刻`の下に、タイムゾーンを表示させました。
- Go の time パッケージの`Zone`メソッドを使って取得したものを HTML の`time_zone`タグに出しています
- ここでは、「計算機３」の時に作った`setJSValue`関数を転用しました
- unixtime 全体に言えますが、ここではコードのわかりやすさを優先して、「計算機３」で用いたようなエラーハンドリングはここではしていません
- ここで、ページの読み込み時に`queueMicrotask`を使用しました
  - この`queueMicrotask`を使った経緯は長くなるので詳しくは後に説明を書きましたが、ここで簡単に説明すると、実行したい処理を`キュー`につめて後で実行されるようにしています

```
js.Global().Call("queueMicrotask", js.FuncOf(setTimeZone))
js.FuncOf(setTimeZone).Release()
```

- やっていることは上の`setInterval`とそっくりで、`setInterval`が定期的に実行されるのに対して、こちらは単発での実行となります

- `setTimeZone`のように、一度呼びだされたら二度と使わない関数は、`Release`メソッドを使ってメモリを解放しておくとメモリの節約になってよいので`js.FuncOf(setTimeZone).Release()`を後ろに書いておきます

参考

- https://pkg.go.dev/syscall/js#Func.Release
- https://zenn.dev/nobonobo/books/85e605893d44ebe7dd3f/viewer/b5ac64d9135e123e367a

## 動作確認

このプログラムのビルドと実行方法は以下の通りです

```
[~/go/src/github.com/ludwig125/githubpages/docs/unixtime] $GOOS=js GOARCH=wasm go build -o unixtime.wasm
```

unixtime.wasm を出力したバイナリファイル名としました

サーバを実行します

```
goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'
```

Web ページを最初に見たときはこんな感じです

現在時刻の部分は 200 ミリ秒ごとにリアルタイムで現在時刻の日付時分秒と Unixtime を表示し続けます

- 「現在時刻」の下に`UTC+9`とタイムゾーンが表示されることも確認できます。

![image](https://user-images.githubusercontent.com/18366858/153731707-6892e7b6-fb53-4e58-9d78-62ff320a3d2d.png)

「変換対象の時刻」のテキスト欄に Unixtime を入力すると、「変換後の時刻」に変換後の日付時分秒が出力されます

![image](https://user-images.githubusercontent.com/18366858/153731624-f619a432-8d15-40ee-bfaf-705b105a640d.png)

日付に変換できない文字を入れると、エラー文が表示されることも確認できます

![image](https://user-images.githubusercontent.com/18366858/153731636-8cab15f1-7775-40a8-973a-abed8e54fd28.png)

- `Clear`ボタンを押すとこの文字は消えます

![image](https://user-images.githubusercontent.com/18366858/153731642-a3f729a7-2e63-4978-a23b-3165c78af492.png)

# WebAssembly.instantiateStreaming()が Promise であるということと、queueMicrotask を使った理由について

`queueMicrotask`について上では簡単に説明しただけだったのですが、
ここは個人的にものすごくはまった個所なので少し詳しく説明します。

これは、私が Javascript 未経験だったことも大きいので、詳しい方は読み飛ばしていい箇所です。

#### WebAssembly.instantiateStreaming()は Promise

まず、この記事で何回も書いてきた WASM ファイルのロード部分をあらためて書きます。

```html
        <script src="wasm_exec.js"></script>
		<script>
			const go = new Go();
			WebAssembly.instantiateStreaming(
				fetch("XXX.wasm"),
				go.importObject
			).then((result) => {
				go.run(result.instance);
			});
```

ここで使っている`WebAssembly.instantiateStreaming`ですが、

https://developer.mozilla.org/ja/docs/Web/JavaScript/Reference/Global_Objects/WebAssembly/instantiateStreaming

> 返値
> Promise で、次の 2 つのフィールドを持つ ResultObject で解決します。

公式ドキュメントのこちらの記載のとおり、
`WebAssembly.instantiateStreaming`は Promise、つまり非同期で実行されます。

Promise 処理が成功したら`then`のあとの部分が実行されます。

Promise については以下の記事などが詳しいですが、

- https://developer.mozilla.org/ja/docs/Web/JavaScript/Guide/Using_promises
- https://qiita.com/cheez921/items/41b744e4e002b966391a

また、`go.importObject`や`go.run`ですが、これは`wasm_exec.js`に定義されたもので、Go ファイルに書いた関数を読み込む部分と実行する部分となります。

- https://github.com/golang/go/blob/master/misc/wasm/wasm_exec.js

つまり、`WebAssembly.instantiateStreaming`部分でやっていることをあらためて説明すると、

1. `WebAssembly.instantiateStreaming`で WASM ファイルをフェッチして Go 関数を Import する処理を Promise で実行
2. Promise が成功したら then 内の Go の関数が実行

となります。

ここまで当然のことを書いているようですが、
ここで重要なのは、Go に書いた任意の関数を実行しようとしても、
`then`内に定義しないと「**まだその関数が認識されない可能性がある**」ということです。

以下問題となる例を書きます。

## ページ読み込み時に Go の 関数が実行できない問題

上の Unixtime ツールで作成した`setTimeZone`関数は、Web ページ読み込み時にページが実行される地域のタイムゾーン(以下の `UTC+9` 部分)を Web ページに設定するためにつくりました。

![image](https://user-images.githubusercontent.com/18366858/153731751-e26e2dda-fb70-4aa4-99f4-3cdb79b2095e.png)

一般的には、Web ページ読み込み時に Javascript の関数を即座に実行する方法として、`onload`や、`DOMContentLoaded`を使った方法が多く見つかります。

最初、`setTimeZone`関数をこの方法で実行させようとしてうまくいかずはまりました。

### うまくいかない例

`setTimeZone`という Go の関数を Javascript 側で実行させるために、前述までのイベント処理と同じく、
Go 側で以下のように`setTimeZone`を Javascript の関数`setTimeZoneFunc`として登録します。

```go
js.Global().Set("setTimeZoneFunc", js.FuncOf(setTimeZone))
```

- ここで`setTimeZone`に`setTimeZoneFunc`という名前をつけているのは、単にどちらを指しているのか分かりやすくするためです

この関数を Web ページの読み込み時に実行させるために、
HTML のヘッダー部分に以下のように`window.onload`や、`document.addEventListener("DOMContentLoaded", 関数)`を書いて実行させると、ブラウザのコンソールは次のようになります。

```html
<script>
	const go = new Go();
	WebAssembly.instantiateStreaming(
		fetch("unixtime.wasm"),
		go.importObject
	).then((result) => {
		go.run(result.instance);
	});

	window.onload = function () {
		console.log("test1");
	};
	document.addEventListener("DOMContentLoaded", function () {
		console.log("test2");
	});
	setTimeZoneFunc();
</script>
```

![image](https://user-images.githubusercontent.com/18366858/153741319-44f3e039-8d4b-433a-97d1-83eeabdc87cc.png)

`console.log`に書いた文字は表示されるのに、Go 側で定義した`setTimeZoneFunc`は

```
Uncaught ReferenceError: setTimeZoneFunc is not defined
```

と、関数が存在しないというエラーが出てしまいました。
(`time_zone`の div タグは置き換わらずにそのままです)

この理由は、上で書いた通り`WebAssembly.instantiateStreaming`が Promise で非同期の呼び出しとなっていて、
`setTimeZoneFunc`を実行されたタイミングではまだロードが終わっておらずこの関数が認識されないためです。

#### 【脱線】`test1`と`test2`の実行順序について

ちなみに上の例で、`test1`よりも`test2`の方を後に書いているのに、ブラウザで順序が入れ変わっている理由ですが、
onload がページや画像などのリソースを読み込んでから処理を実行されるのに対し、DOMContentLoaded は HTML の読み込みと解析が完了したとき、スタイルシート、画像などの読み込みが完了するのを待たずに実行するためです。

以下のページが詳しいです。

- https://developer.mozilla.org/ja/docs/Web/API/Window/DOMContentLoaded_event

#### 【補足】ボタンのクリックなどのイベント処理で関数がうまく実行できた理由

上のように、Javascript でページ読み込み時に Go の関数の呼び出しに失敗して`ReferenceError`が出ましたが、
それまでに紹介したボタンのクリックやテキスト欄への入力では、Go の関数が呼び出せました。

この理由は単純で、ボタンのクリックなどを実行する頃には、Go の関数のロードが終わっていて呼び出せる状態になったからです。

実際、上で `Uncaught ReferenceError: setTimeZoneFunc is not defined`と出た直後に、
コンソールに`setTimeZoneFunc()`と入力すると、この時点ではもうロードが終わっていて、正しく実行されます。

`time_zone`部分が`UTC+9`に変わりました。

![image](https://user-images.githubusercontent.com/18366858/153741311-ec1fe73e-fd38-462b-8463-e1b682be97b2.png)

同様の理由で、Javascript で意図的に Sleep をさせたあとに Go の関数を呼び出しても成功します。

以下では、Promise で`setTimeout`をすることで、3 秒待ってから`setTimeZoneFunc`を呼び出すコードを書きました。
３秒も待てばロードが終わるので、呼び出しに失敗することがありません。
ただし、これは厳密に `WebAssembly.instantiateStreaming`の完了を待っているわけではないので良いコードとは言えません。

```js
		<script>
			const go = new Go();
			WebAssembly.instantiateStreaming(
				fetch("unixtime.wasm"),
				go.importObject
			).then((result) => {
				go.run(result.instance);
			});

			async function waitGoLoad() {
				console.log("wait 3 seconds...");
				await new Promise((s) => setTimeout(s, 3000));
				setTimeZoneFunc();
			}
			waitGoLoad();
```

## ページ読み込み時に Go の関数を実行させる方法 その１

もっとも単純な解決方法は、
`WebAssembly.instantiateStreaming`の Promise が成功した後、つまり`then`のなかの`go.run(result.instance);`のあとに`setTimeZoneFunc`を設定することです。

こうすれば確実に Go 関数のロードが完了しているので、問題なく呼び出すことができます。

```html
<head>
	略
	<script src="wasm_exec.js"></script>
	<script>
		const go = new Go();
		WebAssembly.instantiateStreaming(
			fetch("unixtime.wasm"),
			go.importObject
		).then((result) => {
			go.run(result.instance);

			setTimeZoneFunc();
		});
	</script>
</head>
```

- https://stackoverflow.com/questions/56398142/is-it-possible-to-explicitly-call-an-exported-go-webassembly-function-from-js

上の記事のように、`go.run(result.instance);`後に Web ページ読み込み時に必要な処理を書いていく方法は他にもいくつか見つけたのですが、今回は次の`queueMicrotask`を使う方法を採用しました。

## ページ読み込み時に Go の関数を実行させる方法 その２

今回の用途では上の方法でも良かったのですが、

もしこの方法で他の処理も書いていくと`<head>`の`<script>`部分がどんどん肥大化していくことになります。
個人的にはこの部分はシンプルにしたい思いがありました。

また、Unixtime ツールの機能のうち、
「1. リアルタイムで現在時刻を表示し続ける機能」が Go の`js.Global().Call("setInterval", js.FuncOf(clock), "200")`で完結しているのに、「3. タイムゾーンの表示」を HTML 側でも呼び出さないといけないのがどうにも気に入りませんでした。

そこで、`queueMicrotask`を使う方法にしました。

`queueMicrotask`の仕様は以下が詳しいです

- https://developer.mozilla.org/ja/docs/Web/API/HTML_DOM_API/Microtask_guide
- https://developer.mozilla.org/en-US/docs/Web/API/queueMicrotask

また、そもそも Macrotasks と Microtasks について知らなかったので以下の記事が大変参考になりました。

- https://hidekazu-blog.com/javascript-macrotasks-microtasks/
- https://ja.javascript.info/event-loop#ref-473
- https://christina04.hatenablog.com/entry/2017/03/13/190000
- https://tech.wwwave.jp/entry/javascript-async-execution

詳しい説明は上の記事に譲るとして、ここでは結論として、`queueMicrotask`関数に`setTimeZone`を登録しておくことで、Go の実行時に即時に`setTimeZone`を実行することができるようになります。

また、蛇足ですが、上で紹介した
`js.Global().Call("setInterval", js.FuncOf(clock), "200")`は、200 ミリ秒ごとに`clock`を呼び出しているので、Web ページ表示後最初の 200 ミリ秒間、一瞬だけ時刻の部分が空になる瞬間があります。

これを防ぐ方法として、`clock`に対しても以下のように`queueMicrotask`を使うことで、
Web ページ読み込み時に最初にすぐに`clock`を実行し、そのあと 200 ミリ秒毎に実行されることで、一瞬空になる瞬間をなくすことができます。

```
js.Global().Call("queueMicrotask", js.FuncOf(clock))
js.Global().Call("setInterval", js.FuncOf(clock), "200")
```

繰り返しですが、私は Javascript 初心者なので、この`queueMicrotask`を使った方法が最適なのかどうかまでは確認していません。

# TinyGo への置き換え

## TinyGo の実行方法

上で作った Unixtime ツールを TinyGo に置き換えてみます。

以下の方法で TinyGo として Buid できます。（通常の Go に比べて若干 Build に時間がかかるような気がします）

```
$tinygo build -o unixtime.wasm -target wasm unixtime.go
```

```
cp $(tinygo env TINYGOROOT)/targets/wasm_exec.js .
```

これだけで TinyGo として WASM で実行できます。

```
[~/go/src/github.com/ludwig125/githubpages/unixtime_tinygo] $goexec 'http.ListenAndServe(`:8080`, http.FileServer(http.Dir(`.`)))'

```

ただ、http://localhost:8080/ を見ると、一見問題ないように見えますが、
「変換対象の時刻」に Unixtime を入れると Console にエラーがでます。
（処理自体は問題なく行われます）

![image](https://user-images.githubusercontent.com/18366858/154571515-af59001b-8c84-4ef6-9083-0ddf4984e017.png)

エラー：`syscall/js.finalizeRef not implemented`

このエラー原因について詳しくは以下を見ると良いのですが、

- https://github.com/tinygo-org/tinygo/issues/1140#issuecomment-674425070

TinyGo のバグなので、TinyGo の`wasm_exec.js`が修正されるまでは、以下のように`index.html`側に書いておくとこのエラーがなくなります。

```js
const go = new Go();

// TinyGoのバグを無視するため
// https://github.com/tinygo-org/tinygo/issues1140#issuecomment-671261465
go.importObject.env["syscall/js.finalizeRef"] = ()=> {};

WebAssembly.instantiateStreaming(
	fetch("unixtime.wasm"),
	go.importObject
).then((result) => {
```

参考：

- https://blog.suborbital.dev/foundations-wasm-in-golang-is-fantastic

これでエラー文が出なくなります。

## TinyGo のバイナリサイズ

２つのバイナリサイズを比べてみます

```
[~/go/src/github.com/ludwig125/githubpages/unixtime] $GOOS=js GOARCH=wasm go build -o unixtime.wasm
[~/go/src/github.com/ludwig125/githubpages/unixtime] $ls -l
合計 2096
-rw-r--r-- 1 ludwig125 ludwig125    1247  2月 14 06:58 index.html
-rw-r--r-- 1 ludwig125 ludwig125    2103  2月 18 06:08 unixtime.go
-rwxr-xr-x 1 ludwig125 ludwig125 2113909  2月 18 06:08 unixtime.wasm*
-rw-r--r-- 1 ludwig125 ludwig125   18346  2月 14 06:10 wasm_exec.js
```

```
[~/go/src/github.com/ludwig125/githubpages/unixtime_tinygo] $tinygo build -o unixtime.wasm -target wasm unixtime.go
[~/go/src/github.com/ludwig125/githubpages/unixtime_tinygo] $ls -l
合計 464
-rw-r--r-- 1 ludwig125 ludwig125   1437  2月 17 06:39 index.html
-rw-r--r-- 1 ludwig125 ludwig125   2103  2月 18 06:08 unixtime.go
-rwxr-xr-x 1 ludwig125 ludwig125 447857  2月 18 06:09 unixtime.wasm*
-rw-r--r-- 1 ludwig125 ludwig125  15929  2月 14 06:30 wasm_exec.js
```

私の環境では、ほぼ同じコードでも、TinyGo は Go と比べて`unixtime.wasm*`のバイナリサイズが 1/4 以下になっていました。

## TinyGo の速度

バイナリサイズが小さいということは、当然 WASM として Fetch したり Load するのも速くなるはずです。

通常の Go と TinyGo の Load までの時間を計測するために、それぞれの`index.html`に以下のコードを追加してみます。

```js
<script>
	var start = performance.now(); // 追加部分

	const go = new Go();

	WebAssembly.instantiateStreaming(
		fetch("unixtime.wasm"),
		go.importObject
	).then((result) => {
		go.run(result.instance);

		var end = performance.now(); // 追加部分
		console.log("latency of load and run wasm %f ms", end - start); // 追加部分
	});
</script>
```

https://developer.mozilla.org/ja/docs/Web/API/Performance/now

こちらのパフォーマンス計測用の関数を使います

- `WebAssembly.instantiateStreaming`の前を`start`
- `go.run(result.instance);`の後を`end`

としてこの差分を測ってみます。

ついでに、Go の方の関数にも Latency を計測するために以下の部分を追記します。

```go
func convTime(this js.Value, args []js.Value) interface{} {
	start := time.Now()
	defer func() {
		fmt.Println("convTime latency:", time.Since(start))
	}()

	略
```

これで、Go と TinyGoUnixtime の Web ページをそれぞれ順番に見てみます。

通常の Go
![image](https://user-images.githubusercontent.com/18366858/154574625-c6809e7d-f780-47c1-8730-06cec93b11ba.png)

TinyGo
![image](https://user-images.githubusercontent.com/18366858/154574838-4e8551ec-b3e8-4ece-86b6-42c06f8832fd.png)

注意点

- Go のあとに TinyGo のページを読み込みなおすときは、Chrome のキャッシュに残っていておかしなエラーが出る場合があります。この場合はキャッシュをクリアしてページを再読み込みするために、`Ctrl+Shift+R`でページを更新するといいです

WASM の Fetch から実行までの時間は

- Go: 52.10000002384186 ms
- TinyGo: 16 ms

となりました。

やはり、起動までの時間は TinyGo の方が短くなっています。
今回は小さなプログラムなので、この程度の差ですが、大きなプログラムになると実行までの時間はさらに変わってくるかも知れません。

一方で、`convTime`の実行速度はあまり変わりませんでした。
これは意外でした。

ひとたびバイナリとして読み込んでメモリに乗ってしまえばあとはそんなに変わらないものなのか、それとも実行している関数がそんなに違いが見られる類のものではなかったのかも知れませんが分かりません。

### export を利用した TinyGo コードの書き換え

TinyGo は Go と同じコードをそのまま使えますが、TinyGo ならではの`export`の機能を使うとコードをより直接に呼びだすことができます

https://tinygo.org/docs/guides/webassembly/

> If you have used explicit exports, you can call them by invoking them under the wasm.exports namespace. See the export directory in the examples for an example of this.

とあるとおり、以下のように Go の関数に`//export 関数名`をつけるだけで、なんと Javascript 側から呼びだすことができます。

```go
//export multiply
func multiply(x, y int) int {
    return x * y;
}
```

- ここで、`//export`の`//`と`export`の間に半角スペースを入れると認識されないので、くっつけて書くことを注意してください\*\*
- ちなみに`//export`は以前は`//go:export`でしたが、2020 年に変わったので少し古い資料を見ると`//go:export`となっていることがあります
  - https://github.com/tinygo-org/tinygo/pull/1025

javascript からの呼び出し方法

```js
// Calling the multiply function:
console.log("multiplied two numbers:", wasm.exports.multiply(5, 3));
```

この`multiply`関数はこれまでの WASM の Go の書き方の
`multiply(this js.Value, args []js.Value) interface{}` のような形にしなくて済むというのが最大の利点です。

**`//export`を使った場合の大きな問題点もあるのですがそれは後述します**

この機能を使うと、Unixtime ツールの例えば`setTimeZone`関数は以下のようにシンプルになり、

```go
//export setTimeZone
func setTimeZone() {
	t := time.Now()
	zone, _ := t.Zone()
	setJSValue("time_zone", fmt.Sprintf("(%s)", zone))
}
```

index.html 側では以下のように呼びだすことができます。

```js
const go = new Go();
WebAssembly.instantiateStreaming(fetch("unixtime.wasm"), go.importObject).then(
	(result) => {
		go.run(result.instance);

		result.instance.exports.setTimeZone();
	}
);
```

この方式で、`go.run(result.instance);`のあとに必要な処理をつらつら書いても良いのですが、これだと`index.html`の`<head>`の`<script>`部分が肥大するので、以下の資料を参考に`index.js`ファイルに切り出してみます。

- https://wasmbyexample.dev/examples/hello-world/hello-world.go.en-us.html

```go
package main

import (
	"errors"
	"fmt"
	"strconv"
	"syscall/js"
	"time"
)

func main() {}

//export setTimeZone
func setTimeZone() {
	t := time.Now()
	zone, _ := t.Zone()
	setJSValue("time_zone", fmt.Sprintf("(%s)", zone))
}

func setJSValue(elemID string, value interface{}) error {
	// 元と同じ
}

func getElementByID(targetID string) js.Value {
	// 元と同じ
}

//export clock
func clock() {
	nowStr, nowUnix := getNow(time.Now())

	getElementByID("clock").Set("textContent", nowStr)
	getElementByID("clock_unixtime").Set("textContent", nowUnix)
}

//export convTime
func convTime() {
	in := getElementByID("in").Get("value").String()
	date, err := unixtimeToDate(in)
	if err != nil {
		getElementByID("out").Set("value", js.ValueOf("不正な時刻です"))
		return
	}
	getElementByID("out").Set("value", js.ValueOf(date))
}

// 以降、元と同じ
```

「`//export`」を使うことでかなりシンプルになりました。
TinyGo の export を使えば Javascript 側から Go の関数を直接呼びだすことができます。
コールバック関数が呼び出されたときのために Go のプログラムを永久に終わらせないようにするために、`main`関数内でチャネルを使っていましたがその必要もなくなりました。

Go の関数の呼び出し側である、HTML と Javascript も修正します。

前述の通り head 部分を見やすくするために、以下を参考に修正しました。

- https://wasmbyexample.dev/examples/hello-world/hello-world.go.en-us.html

まず、WASM ファイルのインスタンス生成部分を別のファイルにします。

instantiateWasm.js

```js
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
```

- [公式ドキュメント](https://github.com/golang/go/wiki/WebAssembly#getting-started)の[polyfill](https://github.com/golang/go/blob/b2fcfc1a50fbd46556f7075f7f1fbf600b5c9e5d/misc/wasm/wasm_exec.html#L17-L22)を使う場合は以下のようになります
- ここでは省略しました。

```js
// polyfillを定義した場合
if (WebAssembly.instantiateStreaming) {
	response = await WebAssembly.instantiateStreaming(
		fetch(wasmModuleUrl),
		importObject
	);
} else {
	const fetchAndInstantiateTask = async () => {
		const wasmArrayBuffer = await fetch(wasmModuleUrl).then((response) =>
			response.arrayBuffer()
		);
		return WebAssembly.instantiate(wasmArrayBuffer, importObject);
	};
	response = await fetchAndInstantiateTask();
}
```

一方、呼び出し側の`index.html`から、WASM の呼び出し部分を切り出して別のファイルにすると以下のようになります。

index.js

```js
import { wasmBrowserInstantiate } from "./instantiateWasm.js";

const go = new Go(); // Defined in wasm_exec.js. Don't forget to add this in your index.html.

// TinyGoのバグを無視するため
// https://github.com/tinygo-org/tinygo/issues/1140#issuecomment-671261465
go.importObject.env["syscall/js.finalizeRef"] = () => {};

const runWasm = async () => {
	// Get the importObject from the go instance.
	const importObject = go.importObject;

	// wasm moduleのインスタンスを作成
	const wasmModule = await wasmBrowserInstantiate(
		"./unixtime.wasm",
		importObject
	);

	go.run(wasmModule.instance);

	wasmModule.instance.exports.setTimeZone();
	setInterval(wasmModule.instance.exports.clock, 200);
	document
		.getElementById("in")
		.addEventListener("input", wasmModule.instance.exports.convTime);
};
runWasm();
```

- インスタンス生成部分とメインの処理部分を分離して分かりやすくなりました
- （`wasmModule.instance.exports.`部分がやや鬱陶しいですが、）Go の関数を Javascript ネイティブの関数のように扱うことができるようになったため、実行方法も Javascript の書き方になっています

最後に、この`index.js`を`index.html`から呼びだせば終わりです。

```html
<head>
	<meta charset="utf-8" />
	<title>unixtime</title>
	<link rel="shortcut icon" href="#" />
	<script src="./wasm_exec.js"></script>
	<script type="module" src="./index.js"></script>
</head>
```

かなり見やすくなったかと思います。

### export を使った関数の限界

とても素敵な機能に思える TinyGo の`//export`ですが、これを書いている 2022 年 3 月の現時点ではとても大きな問題があります。

それは、**WASM では直接文字列をやりとりできないことです**

以下が WASM の扱える型の種類です。
https://github.com/WebAssembly/design/blob/main/Semantics.md#types

```
WebAssembly has the following value types:

i32: 32-bit integer
i64: 64-bit integer
f32: 32-bit floating point
f64: 64-bit floating point
```

そのため、例えば以下のような方法で直接文字列を関数に渡したり返してもらうことはできません

```go
// 以下のようにTinyGoで関数を使うことはできない

//export printMessage
func printMessage(s string) { // stringを受け取ることができない
	fmt.Println("hello:", s)
}

//export returnString
func returnString() string {
	return "hello" // stringを返すこともできない
}
```

int 型は扱えるので、変数のアドレスと長さを計算してそれを関数に渡す方法があるにはありますが、とても分かりやすいとは言えません。

参考

- https://github.com/tinygo-org/tinygo/issues/645
- https://github.com/tinygo-org/tinygo/issues/411#issuecomment-503066868
- https://www.alcarney.me/blog/2020/passing-strings-between-tinygo-wasm/
- https://stackoverflow.com/questions/41353389/how-can-i-return-a-javascript-string-from-a-webassembly-function
- https://github.com/tinygo-org/tinygo/issues/1824
- https://wasmbyexample.dev/examples/webassembly-linear-memory/webassembly-linear-memory.go.en-us.html
- https://nulab.com/ja/blog/nulab/basic-webassembly-begginer/
- https://zenn.dev/summerwind/articles/96f2aae05b6614

また仮に文字列を１つ渡せても２つ以上はできないので、その場合は json などでデコードして渡す必要があります

TinyGo での json 参考

- https://github.com/tinygo-org/tinygo/issues/447
- https://github.com/mailru/easyjson
- https://www.sambaiz.net/article/193/
- https://stackoverflow.com/questions/40587860/using-easyjson-with-golang/44757748
- https://github.com/tinygo-org/tinygo/pull/2314

以上の理由から、TinyGo でも`//export`を使いまくるわけにはいかず、文字列のやり取りをする際は素直に js パッケージを使って Javascript とやり取りしたほうが便利な場面が多そうです。

# -------------------------------

今回、WASM でやりたいことを実現するのにこのような方法を用いましたが、実をいうとこれが最善手なのか分かっていません。
私が Javascript や WASM の賢い書き方に詳しくないだけの可能性もありますが、
とりあえず納得いくものが得られたのでこれで完成とします。

# WASM の

前述の「計算機３」のコードを`docs/calc3`に配置して、上の Unixtime 変換ツールを

```html
<html>
	<head>
		<meta charset="utf-8" />
		<link rel="shortcut icon" href="#" />
	</head>
	<body>
		<ul>
			<li><a href="./calc3/calc3.html">Calculator</a></li>
			<li><a href="./unixtime/unixtime.html">Unixtime変換ツール</a></li>
		</ul>
	</body>
</html>
```

# 参考

https://www.aaron-powell.com/posts/2019-02-05-golang-wasm-2-writing-go/ js
https://hmaster.net/table4.html http://mh.rgr.jp/memo/mh0025.htm wasm clock
https://github.com/Yaoir/ClockExample-Go-WebAssembly リアルタイム時刻
https://ja.javascript.info/events-change-input
https://www.w3schools.com/howto/howto_html_clear_input.asp ← わかりやすかったです
https://dev.bitolog.com/go-in-the-browser-using-webassembly/
https://golangbot.com/go-webassembly-dom-access/
https://github.com/golangbot/webassembly/blob/tutorial2/cmd/wasm/main.go
https://tinygo.org/docs/guides/webassembly/
