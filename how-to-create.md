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
