updated at 2020-03-03

# Goでヘキサゴナルアーキテクチャ

[Go](https://qiita.com/tags/go)[HexagonalArchitecture](https://qiita.com/tags/hexagonalarchitecture)

More than 1 year has passed since last update.

## はじめに

『Standard Go Project Layout』と『ヘキサゴナルアーキテクチャ』を参考にサンプルプロジェクトを作ってみました。

トランザクション周りも取り扱います。

## 『Standard Go Project Layout』とは

↓これです。
[Standard Go Project Layout](https://github.com/golang-standards/project-layout)

上記の内容を日本語で簡潔にまとめてくださってる記事もありました。
[Goにはディレクトリ構成のスタンダードがあるらしい。](https://qiita.com/sueken/items/87093e5941bfbc09bea8)

別の記事になりますが、こちらもとても参考になりました。
[Practical Go: Real world advice for writing maintainable Go programs](https://dave.cheney.net/practical-go/presentations/qcon-china.html)

## ヘキサゴナルアーキテクチャとは

↓これです。
[ヘキサゴナルアーキテクチャ(Hexagonal architecture翻訳)](https://blog.tai2.net/hexagonal_architexture.html)

本家サイトへのリンクも張りたかったのですが、現在工事中とのことでした。。。

ヘキサゴナルアーキテクチャはレイヤードアーキテクチャ・オニオンアーキテクチャ・クリーンアーキテクチャの並びで語られることが多いですが、これらの中だと個人的にはヘキサゴナルアーキテクチャが簡潔かつ重要なポイントをわかりやすく解説しているなぁという印象です。

クリーンアーキテクチャはこれらの共通要素を抽出しているものなので、少し抽象度が上がり過ぎている感じがします。いきなりクリーンアーキテクチャに行くよりも、まずはヘキサゴナルアーキテクチャを勉強してみるとイメージが湧きやすいんじゃないかなと今回感じました。

## ヘキサゴナルアーキテクチャにおける重要ポイント

上記の翻訳記事のなかで個人的に重要だと思う箇所を抜き出していきたいと思います。

> アプリケーションを、ユーザー、プログラム、自動テストあるいはバッチスクリプトから、同じように駆動できるようにする。

プログラムを起動するのがコマンドライン・HTTPリクエスト・別のプログラム（ライブラリとして利用）・バッチスクリプトであるかに関わらず動くようにするということですね。

------

> イベントが外側の世界からポートに届くと、特定テクノロジーのアダプターが、利用可能な手続き呼び出しか、メッセージにそれを変換して、アプリケーションに渡す。よろこばしいことに、アプリケーションは、入力デバイスの正体を知らない。

「外界 → アプリケーション」の向きの場合は、アダプターのことを「Controller」という名称でよく呼んでいる気がします。

[![theonionarchitecturepart3_67c4image05.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F1d82d82e-1ce7-16ff-7fba-cd6ddc2df796.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=1418e24b8c008ed53d2387278a88c22c)](https://camo.qiitausercontent.com/a94143e5b967f4917b879d1671cb1eac815972f5/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f31643832643832652d316365372d313666662d376662612d6364366464633264663739362e706e67)

------

> アプリケーションがなにかを送る必要があるとき、それはポートを通じてアダプターに送られて、受信側のテクノロジーが必要とする信号を生む(人力であれ自動であれ)。

[![CleanArchitecture.jpg](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2Fbca32a06-aeae-aa6f-877b-9b1638cfb90a.jpeg?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=7b2ca79cd8651f9322dae47665abd700)](https://camo.qiitausercontent.com/82d14bf33328e790b53b1200bad1bb8b2aab4620/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f62636133326130362d616561652d616136662d383737622d3962313633386366623930612e6a706567)

------

> アプリケーションは、データを取得するために外部のエンティティーと通信する。そのプロトコルの典型は、データベースプロトコルだ。アプリケーションの観点からは、もしデータベースがSQLデータベースから、フラットなファイルや、その他のデータベースに移行しても、APIとの会話は変わるべきではない。ゆえに、同じポートへの追加のアダプターは、SQLアダプター、フラットファイルアダプター、そしてもっとも重要なものとして、「モック」データベースのアダプターを含む。これは、メモリ内に居座るもので、実際のデータベースの存在にまったく依存しない。

------

> 多くのアプリケーションは、ポートを2つだけ持つ: ユーザー側の対話と、データベース側の対話だ。

------

> 「ポートとアダプター」という用語は、素描のパーツの「目的」を強調している。ポートは、目的の会話を識別する。典型的には、どのひとつのポートにも複数のアダプターがあるだろう。それらは、ポートに差し込まれるさまざまな技術のためのものだ。

## ヘキサゴナルアーキテクチャで実装してみる

それでは実際にヘキサゴナルアーキテクチャでプログラムを書いてみるとどのようになるかを試していきます。今回特に注目するのは次の3点です

- 「外界 → アプリケーション」の管理
- 「アプリケーション → 外界」の管理
- トランザクションの管理

また、今回の完成形のディレクトリ構成は次のようになりましたので先に掲載しておきます。

[![スクリーンショット 2019-11-17 13.29.49.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F888bb7e2-9ce7-6aa2-488e-6033d4573bb2.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=54d48520a57e782a6a39e358c4b76023)](https://camo.qiitausercontent.com/41947a8d80506344cc06b1c883974e91792007ed/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f38383862623765322d396365372d366161322d343838652d3630333364343537336262322e706e67)

ソースコードの全体はGitHubリポジトリにアップしていますので、確認したい方はぜひ。
https://github.com/rema424/hexample

今回はプロジェクト名を `hexample` として実装しています。

### 1. 「外界 → アプリケーション」

ヘキサゴナルアーキテクチャにおいては、アプリケーションのコアのロジックをコマンドライン・HTTPリクエスト・バッチプログラム・別のプログラム（ライブラリとして利用）などから同じように呼び出せるようにすることを目指します。

今回は「コマンドライン」「HTTPリクエスト」「ライブラリ」の3つの呼び出し方に対応するプログラムを作ってみます。

今回はコアロジックを `internal/` ディレクトリに隠蔽するようにしてみます。

#### 1.1. コアロジックを作る

`internal/` の中に新規のパッケージを作り、コードを書いていきます。なお、本記事では `internal/` ディレクトリに配置するパッケージの名前を `service1`、`service2`... のように名付けていきます。[Practical Go: Real world advice for writing maintainable Go programs](https://dave.cheney.net/practical-go/presentations/qcon-china.html)の記事によると、Goのパッケージ名はそのパッケージが何を提供するかで決めるのが推奨とのことです。

> Name your package for what it provides, not what it contains.

今回はシーケンス番号を付与してパッケージ名としていますが、本来であれば提供するサービスの内容をディレクトリ名にするのが良さそうです。

ドメイン駆動設計を実践する場合には `internal/` ディレクトリの中に作成するパッケージは「境界づけられたコンテキスト」に基づくようにすると上手くいくような気がします。

[ドメイン分析を使用したマイクロサービスのモデル化 | Microsoft Docs](https://docs.microsoft.com/ja-jp/azure/architecture/microservices/model/domain-analysis#define-bounded-contexts)

実際のソースコードです。アプリケーションコアロジックとは言っても、まずは標準出力に文字を表示するだけのプログラムです。

[![スクリーンショット 2019-11-17 14.04.06.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2Fa72bc63e-faf0-39c8-f656-eba240b34b6e.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=ef0d0ad7d703d3214c2db0be207fbab4)](https://camo.qiitausercontent.com/0723573691cbb1a5f7d8908f4eef0fa09b009a5f/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f61373262633633652d666166302d333963382d663635362d6562613234306233346236652e706e67)

internal/service1/service1.go

```
package service1

import (
    "context"
    "fmt"
)

// AppCoreLogicIn .
type AppCoreLogicIn struct {
    From    string
    Message string
}

// AppCoreLogic .
func AppCoreLogic(ctx context.Context, in AppCoreLogicIn) {
    fmt.Println("--------------------------------------------------")
    fmt.Println("service1:")
    fmt.Println("this is application core logic.")
    fmt.Printf("from: %s, message: %s\n", in.From, in.Message)
    fmt.Println("--------------------------------------------------")
}
```

#### 1.2. コマンドラインツールからコアロジックを呼び出す

まずはコマンドラインツールからコアロジックを呼び出してみます。今回は [cobra](https://github.com/spf13/cobra) を利用してコマンドラインツールを作成します。

プログラムを起動する際のエントリーポイント（main.go）は `cmd/` ディレクトリ配下に設置することが一般的なようです。なお、`cmd/` の中には今後HTTPリクエストを受け付けるためのプログラムのエントリーポイントも設置するので、`cmd/cli/` というサブディレクトリを作成してこちらに main.go を配置します。

[![スクリーンショット 2019-11-17 14.04.45.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F8eeb0114-73d4-a0be-937b-41d515c9f4a6.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=fe9d81b17008f79b825d89c367018ef2)](https://camo.qiitausercontent.com/db8434768d2ac65ef39bcbf9dd1db735b2782626/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f38656562303131342d373364342d613062652d393337622d3431643531356339663461362e706e67)

cmd/cli/main.go

```
package main

import "github.com/rema424/hexample/cmd/cli/cmd"

func main() {
    cmd.Execute()
}
```

cmd/cli/cmd/root.go

```
package cmd

import (
    "context"
    "fmt"
    "os"

    "github.com/rema424/hexample/internal/service1"
    "github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Run: func(cmd *cobra.Command, args []string) {
        var msg string
        if len(args) != 0 {
            msg = args[0]
        } else {
            msg = "Hello, from cli!"
        }

        arg := service1.AppCoreLogicIn{
            From:    "cli",
            Message: msg,
        }

        service1.AppCoreLogic(context.Background(), arg)
    },
}

// Execute ...
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}
```

`cmd/cli/cmd/root.go` が今回アダプターの役割を担っています。アプリケーションコアは「パラメータがコマンドライン引数で渡ってくる」ということは知らないので、代わりにこのアダプターがコマンドライン引数を解釈して、実際のアプリケーションコアにパラメータとして渡しています。

なお、ディレクトリ階層に `cmd/.../cmd/...` と 2 回 cmd が現れるのが気になりますが、これが cobra のデフォルトなのでしょうがないですね・・・

また、`cmd/` ディレクトリ配下に main.go 以外のプログラム（つまりアダプター）を配置して良いものかどうかも意見が別れる部分かなと思います。今回は `cmd/` 配下にアダプターも配置するようにしました。例えば Go のテストコードジェネレータである [cweill/gotests](https://github.com/cweill/gotests) もこのディレクトリ構成です。

それではプログラムを起動してみます。

```
$ go run cmd/cli/main.go "おはようございます"
--------------------------------------------------
service1:
this is application core logic.
from: cli, message: おはようございます
--------------------------------------------------
```

[![スクリーンショット 2019-11-17 14.21.39.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F7edf3966-d575-2fe6-0627-b85b523cbee2.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=0ed0c2154a7f0e881814dd914aab4c33)](https://camo.qiitausercontent.com/5b6f87c14a04de32819021cfc6ec44ac52c6688f/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f37656466333936362d643537352d326665362d303632372d6238356235323363626565322e706e67)

アプリケーションロジックを呼び出すことができました。

#### 1.3. Webサーバーからコアロジックを呼び出す

次に、HTTP リクエストでアプリケーションのコアロジックを起動してみます。今回は [echo](https://github.com/labstack/echo) を利用してWebサーバーを立ち上げます。

`cmd/http/` ディレクトリにプログラムを追加していきます。

[![スクリーンショット 2019-11-17 14.28.10.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F38b4a3de-265d-e382-35d2-05cb7edb21ba.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=0c41fdd2f32d6e852e0b6befc0073dcf)](https://camo.qiitausercontent.com/090c9f37948a3cce1be486b8f814b13be25f58a4/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f33386234613364652d323635642d653338322d333564322d3035636237656462323162612e706e67)

cmd/http/main.go

```
package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/rema424/hexample/cmd/http/controller"

    "github.com/labstack/echo/v4"
    "github.com/labstack/echo/v4/middleware"
)

var e = createMux()

func main() {
    http.Handle("/", e)

    port := os.Getenv("PORT")
    if port == "" {
        port = "8080"
        log.Printf("Defaulting to port %s", port)
    }

    log.Printf("Listening on port %s", port)
    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

func init() {
    ctrl := &controller.Controller{}

    e.GET("/:message", ctrl.HandleMessage)
}

func createMux() *echo.Echo {
    e := echo.New()

    e.Use(middleware.Recover())
    e.Use(middleware.Logger())
    e.Use(middleware.Gzip())

    return e
}
```

cmd/http/controller/controller.go

```
package controller

import (
    "github.com/rema424/hexample/internal/service1"

    "github.com/labstack/echo/v4"
)

// Controller ...
type Controller struct{}

// HandleMessage ...
func (ctrl *Controller) HandleMessage(c echo.Context) error {
    msg := c.Param("message")
    if msg == "" {
        msg = "Hello, from http!"
    }

    arg := service1.AppCoreLogicIn{
        From:    "http",
        Message: msg,
    }

    service1.AppCoreLogic(c.Request().Context(), arg)
    return nil
}
```

それではWebサーバーを起動して curl でリクエストを実行してみます。パラメータはパスパラメータで渡します。

```
# Webサーバー起動
go run cmd/http/main.go

# 別の shell を立ち上げてリクエストを実行
curl localhost:8080/こんにちは

# Webサーバーを起動した方の shell に以下の出力
--------------------------------------------------
service1:
this is application core logic.
from: http, message: こんにちは
--------------------------------------------------
```

[![スクリーンショット 2019-11-17 14.36.57.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F8638505c-8bcd-27cf-b8a9-d6a18dd625ae.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=dc7657debc488598e5fb26f9b7305ca2)](https://camo.qiitausercontent.com/d3d19fd4f1452fe8607e9c439705c4dc5801b391/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f38363338353035632d386263642d323763662d623861392d6436613138646436323561652e706e67)

コアロジックを呼び出すことができました。

コアロジックではパラメータがパスパラメータで渡ってくるということを知りませんが、アダプターがこの解釈の部分を担っています。このように、アプリケーションとユーザーの間にアダプターを挟むことで、どのような形で処理がリクエストされる場合でも、コアロジック側の実装を変えることなく新しい呼び出し元に対応することができます。

#### 1.4. 別プログラムからコアロジックを呼び出す

要するにライブラリとしての利用です。

ライブラリとして `internal/` 配下のロジックを使うには、プロジェクトのトップ階層にプロジェクト名と同名のファイルを作成し、これをアダプターとして利用するパターンが多いようです。

[![スクリーンショット 2019-11-17 14.42.53.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2Fb96cf378-2cb7-9b45-e9ae-5b8fd33968e8.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=0f33640a2aab068a44bbbec914a987df)](https://camo.qiitausercontent.com/c059369e736bffb2098b035fa1ae4b89366ee54b/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f62393663663337382d326362372d396234352d653961652d3562386664333339363865382e706e67)

hexample.go

```
package hexample

import (
    "context"

    "github.com/rema424/hexample/internal/service1"
)

// Run ...
func Run(ctx context.Context, msg string) {
    if msg == "" {
        msg = "Hello, from external pkg!"
    }
    arg := service1.AppCoreLogicIn{
        From:    "external pkg",
        Message: msg,
    }
    service1.AppCoreLogic(ctx, arg)
}
```

go.mod の module 名がこのパッケージ名に合致しているか確認します。

[![スクリーンショット 2019-11-17 14.46.13.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F1186fe26-a147-65ef-1d6d-1071099a2b0c.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=51db174f992ba859624ec756d1cfc214)](https://camo.qiitausercontent.com/3c6d432454efcf8374ede8b40c58c2328b2c6fce/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f31313836666532362d613134372d363565662d316436642d3130373130393961326230632e706e67)

あとはこのソースコードを GitHub か何かにアップして公開し、別のプロジェクトでインポートして利用します。

```
package mainimport (    "context"    "github.com/rema424/hexample")func main() {    ctx := context.Background()    msg := "こんばんは"    hexample.Run(ctx, msg)}
```

このプログラムを実行すると次のようになります。

```
go run main.go--------------------------------------------------service1:this is application core logic.from: external, message: こんばんは--------------------------------------------------
```

[![スクリーンショット 2019-11-17 14.52.24.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2Fe33ca3e0-0e3c-5bc9-8b55-5fe1d8dcd445.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=8f73109598dc0b765c4facb81d44d6de)](https://camo.qiitausercontent.com/d9da3863a644dea873247f8180bfb1e3bf9ab5b7/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f65333363613365302d306533632d356263392d386235352d3566653164386463643434352e706e67)

コアロジックを呼び出すことができました。

今回はコマンドラインツール・HTTPリクエスト・ライブラリからのアプリケーションの利用を見てみましたが、同様にアダプターを追加していけば他のプロトコルにも対応させることができます。このような場合でもアプリケーションのコアのソースコードを書き換える必要はありません。アダプターがこれらに対応します。

### 2. 「アプリケーション → 外界」

次に「アプリケーション → 外界」の向き、つまり「アプリケーション → データベース」について見ていきます。ユーザーからの入力をアプリケーションが知らなくてよかったように、データベースについてもアプリケーションは知らずに済むように作っていきます。

これを実現するためには「依存関係逆転の原則（DIP, dependency inversion principle）」または「依存性注入（dependency injection）」と呼ばれる技法を使います。ここら辺はソースコードを見た方が早いですね。

今回の例では次のようなプログラムを作ります。

- `Person` の登録
- `Person` の1件取得
- ユーザー側は HTTP リクエストに対応
- 本番 DB として MySQL を利用
- モック DB としてメモリ（Goのmap）を利用

リレーショナルデータベースのスキーマは次の通りです。

```
create table if not exists person (  id bigint auto_increment,  name varchar(255),  email varchar(255),  primary key (id));
```

#### 2.1. アプリケーションコアの実装

まずはアプリケーションのコアを作成していきます。アダプター（Gateway）にはまだ手を付けません。新しいパッケージを作成して開発します。

[![スクリーンショット 2019-11-17 15.10.33.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2Fa803b31f-2ea5-138f-fb4e-807fc6b73a78.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=e00fb27d1e70656676fb4d9d1a9d8a13)](https://camo.qiitausercontent.com/9d40a80aeb91966ed60b8cbdceeac6f6068151a2/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f61383033623331662d326561352d313338662d666234652d3830376663366237336137382e706e67)

internal/service2/model.go

```
package service2// Person ...type Person struct {    ID    int64  `db:"kokoha"`    Name  string `db:"tekitode"` // sql.NullString はインフラに結合するので使わない    Email string `db:"yoiyo"`}
```

モデルにはオブジェクトを定義していきます。オブジェクト特有の振る舞い（メソッド）もここに記述しますが、今回のプログラムは簡素なためメソッドはありません。重要なのはインフラの知識を持ち込まないことです。dbタグがありますが、ここは適当でいいです。モデルがDBの事情に合わせるのではなく、アダプター（Gateway）がこのdbタグと実際のカラム名の調整を行います。

internal/service2/repository.go

```
package service2import "context"// Repository ...type Repository interface {    RegisterPerson(context.Context, Person) (Person, error)    GetPersonByID(context.Context, int64) (Person, error)}
```

リポジトリは今回インタフェースとして利用するのでこれくらいです。アプリケーションコアなのでここにもインフラに関する知識は出てきません。

internal/service2/provider.go

```
package service2import "context"// Provider ...type Provider struct {    r Repository}// NewProvider ...func NewProvider(r Repository) *Provider {    return &Provider{r}}// RegisterPerson ...func (p *Provider) RegisterPerson(ctx context.Context, name, email string) (Person, error) {    psn := Person{        Name:  name,        Email: email,    }    psn, err := p.r.RegisterPerson(ctx, psn)    if err != nil {        return Person{}, err    }    return psn, nil}// GetPersonByID ...func (p *Provider) GetPersonByID(ctx context.Context, id int64) (Person, error) {    psn, err := p.r.GetPersonByID(ctx, id)    if err != nil {        return Person{}, err    }    return psn, nil}
```

プロバイダーは提供するサービス内容を記述していきます。クリーンアーキテクチャでは `Use Case` ですね。 今回は 『[Practical Go: Real world advice for writing maintainable Go programs](https://dave.cheney.net/practical-go/presentations/qcon-china.html)』の中から言葉をもらって `Provider` にしてみました。

なお、コントローラーに対してモデルの型を公開するかどうか（引数や戻り値にモデルの型を使う）については議論があるかと思います。今回はアダプターに対してモデルの型を公開してもよいという方針でプログラムを作っています。（DB側のアダプターであるGatewayにはモデルの型を公開することになりますし。）

#### 2.2. MySQL用アダプターの実装

Repository のインタフェースを満たすように、MySQL と疎通するための Gateway を実装します。なお、今回は [sqlx](https://github.com/jmoiron/sqlx) を O/R マッパーとして利用します。

[![スクリーンショット 2019-11-17 15.33.37.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F6c012479-29d6-27fe-0eef-83dc8dd7e2b9.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=ed8e0b24ab160eeaf38cb9f7a1f0ab8b)](https://camo.qiitausercontent.com/d9423ab83a0a11ef323e1b5b2cc5d1dbe2aad046/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f36633031323437392d323964362d323766652d306565662d3833646338646437653262392e706e67)

internal/service2/gateway.go

```
package service2import (    "context"    "github.com/jmoiron/sqlx")// Gateway ...type Gateway struct {    db *sqlx.DB}// NewGateway ...func NewGateway(db *sqlx.DB) Repository {    return &Gateway{db}}// RegisterPerson ...func (r *Gateway) RegisterPerson(ctx context.Context, p Person) (Person, error) {    q := `INSERT INTO person (name, email) VALUES (:tekitode, :yoiyo);`    res, err := r.db.NamedExecContext(ctx, q, p)    if err != nil {        return Person{}, err    }    id, err := res.LastInsertId()    if err != nil {        return Person{}, err    }    p.ID = id    return p, nil}// GetPersonByID ...func (r *Gateway) GetPersonByID(ctx context.Context, id int64) (Person, error) {    // DB上のnull対策はここで実装する    q := `SELECT  COALESCE(id, 0) AS 'kokoha',  COALESCE(name, '') AS 'tekitode',  COALESCE(email, '') AS 'yoiyo'FROM personWHERE id = ?;`    var p Person    err := r.db.GetContext(ctx, &p, q, id)    return p, err}
```

Gateway はインフラにどっぷり浸かります。DB の null への対策や、モデルのフィールド名とレコードのカラム名の繋ぎ混みなどを担います。

なお `NewGateway()` 関数は、戻り値の型は Repository インタフェースとなっていますが、実際に返却しているのは Gateway 構造体（のポインタ）になっています。

#### 2.3. ユーザー側アダプター（controller）の実装とルーティングの追加

現在の話の本筋ではありませんが、プログラムの実行のために実装します。

[![スクリーンショット 2019-11-17 15.44.46.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F73357aff-3293-3779-8b68-32c8bb07a1af.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=5454f58b5436acab6a6f3f7d32b2f295)](https://camo.qiitausercontent.com/e6fe30062ee6cee549c589563de58a6853dc2836/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f37333335376166662d333239332d333737392d386236382d3332633862623037613161662e706e67)

cmd/http/controller/controller2.go

```
package controllerimport (    "net/http"    "strconv"    "github.com/rema424/hexample/internal/service2"    "github.com/labstack/echo/v4")// Controller2 ...type Controller2 struct {    p *service2.Provider}// NewController2 ...func NewController2(p *service2.Provider) *Controller2 {    return &Controller2{p}}// HandlePersonRegister ...// curl -X POST -H 'Content-type: application/json' -d '{"name": "Alice", "email": "alice@example.com"}' localhost:8080/peoplefunc (ctrl *Controller2) HandlePersonRegister(c echo.Context) error {    in := struct {        Name  string `json:"name"`        Email string `json:"email"`    }{}    if err := c.Bind(&in); err != nil {        return c.JSON(http.StatusBadRequest, err.Error())    }    // TODO: implement    // if err := c.Validate(&in); err != nil {    //  return c.JSON(http.StatusUnprocessableEntity, err.Error())    // }    ctx := c.Request().Context()    psn, err := ctrl.p.RegisterPerson(ctx, in.Name, in.Email)    if err != nil {        return c.JSON(http.StatusInternalServerError, err.Error())    }    return c.JSON(http.StatusOK, psn)}// HandlePersonGet ...// curl localhost:8080/people/999func (ctrl *Controller2) HandlePersonGet(c echo.Context) error {    id, err := strconv.Atoi(c.Param("personID"))    if err != nil {        return c.JSON(http.StatusUnprocessableEntity, err.Error())    }    ctx := c.Request().Context()    psn, err := ctrl.p.GetPersonByID(ctx, int64(id))    if err != nil {        return c.JSON(http.StatusInternalServerError, err.Error())    }    return c.JSON(http.StatusOK, psn)}
```

サーバー側で発生したエラーをそのままクライアントに返却してしまっていますが、サンプルということでお見逃しください![:pray_tone2:](https://cdn.qiita.com/emoji/twemoji/unicode/1f64f-1f3fc.png)

cmd/http/main.go

```
package mainimport (    "fmt"    "log"    "net/http"    "os"    "github.com/rema424/hexample/cmd/http/controller"    "github.com/rema424/hexample/internal/service2"    "github.com/rema424/hexample/pkg/mysql"    "github.com/labstack/echo/v4"    "github.com/labstack/echo/v4/middleware")var e = createMux()func main() {    http.Handle("/", e)    port := os.Getenv("PORT")    if port == "" {        port = "8080"        log.Printf("Defaulting to port %s", port)    }    log.Printf("Listening on port %s", port)    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))}func init() {    // Mysql    c := mysql.Config{        Host:                 os.Getenv("DB_HOST"),        Port:                 os.Getenv("DB_PORT"),        User:                 os.Getenv("DB_USER"),        DBName:               os.Getenv("DB_NAME"),        Passwd:               os.Getenv("DB_PASSWORD"),        AllowNativePasswords: true,    }    db, err := mysql.Connect(c)    if err != nil {        log.Fatalln(err)    }    ctrl := &controller.Controller{}    // DI    gateway2 := service2.NewGateway(db)    provider2 := service2.NewProvider(gateway2)    ctrl2 := controller.NewController2(provider2)    e.GET("/:message", ctrl.HandleMessage)    e.GET("/people/:personID", ctrl2.HandlePersonGet)    e.POST("/people", ctrl2.HandlePersonRegister)}func createMux() *echo.Echo {    e := echo.New()    e.Use(middleware.Recover())    e.Use(middleware.Logger())    e.Use(middleware.Gzip())    return e}
```

`init()` 関数で DI とルーティングの追加を行なっています。

また、データベースとの接続には `pkg/mysql/` というパッケージを利用しています。今回しれっと追加したソースコードです。

[![スクリーンショット 2019-11-17 15.57.26.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2Ffe3d91b9-e6ea-c30b-90a6-bccc97d2ed2a.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=21483acba50effd48e55efc291c9280d)](https://camo.qiitausercontent.com/ece2edad5c91faf7c1e1c811ef39548f461754d3/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f66653364393162392d653665612d633330622d393061362d6263636339376432656432612e706e67)

pkg/mysql/mysql.go

```
package mysqlimport (    "time"    my "github.com/go-sql-driver/mysql"    "github.com/jmoiron/sqlx")// Config ...type Config struct {    User                 string    Passwd               string    Host                 string    Port                 string    Net                  string    Addr                 string    DBName               string    Collation            string    InterpolateParams    bool    AllowNativePasswords bool    ParseTime            bool    MaxOpenConns         int    MaxIdleConns         int    ConnMaxLifetime      time.Duration}func (c Config) build() Config {    if c.User == "" {        c.User = "root"    }    if c.Net == "" {        c.Net = "tcp"    }    if c.Host == "" {        c.Host = "127.0.0.1"    }    if c.Port == "" {        c.Port = "3306"    }    if c.Addr == "" {        c.Addr = c.Host + ":" + c.Port    }    if c.Collation == "" {        c.Collation = "utf8mb4_bin"    }    if c.MaxOpenConns < 0 {        c.MaxOpenConns = 30    }    if c.MaxIdleConns < 0 {        c.MaxIdleConns = 30    }    if c.ConnMaxLifetime < 0 {        c.ConnMaxLifetime = 60 * time.Second    }    return c}// Connect .func Connect(c Config) (*sqlx.DB, error) {    c = c.build()    mycfg := my.Config{        User:                 c.User,        Passwd:               c.Passwd,        Net:                  c.Net,        Addr:                 c.Addr,        DBName:               c.DBName,        Collation:            c.Collation,        InterpolateParams:    c.InterpolateParams,        AllowNativePasswords: c.AllowNativePasswords,        ParseTime:            c.ParseTime,    }    dbx, err := sqlx.Connect("mysql", mycfg.FormatDSN())    if err != nil {        return nil, err    }    dbx.SetMaxOpenConns(c.MaxOpenConns)    dbx.SetMaxIdleConns(c.MaxIdleConns)    dbx.SetConnMaxLifetime(c.ConnMaxLifetime)    return dbx, nil}
```

Go では アプリケーションのコアではないソースコードは `pkg` というディレクトリの中に作成するのが慣習となっているようです。他の言語やフレームワークだと `lib` や `util` などと名付けられているかと思います。ここに配置されるソースコードは別のプロジェクトからでも利用が可能なソースコードです。現在のアプリケーションでのみ有効的に利用できる処理であれば `internal/` 配下に設置した方がいいかもしれません。この例で有名なのが Docker の namesgenerator （現: Moby Project）でしょうか。

- [moby/pkg/namesgenerator](https://github.com/moby/moby/tree/master/pkg/namesgenerator)
- [Dockerコンテナのおもしろい名前](https://deeeet.com/writing/2014/07/15/docker-container-name/)

それではデータベースや環境変数をよしなに準備した上で、Webサーバーを起動し curl でアプリケーションの処理を呼び出してみます。

```
$ go run cmd/http/main.go2019/11/17 16:11:02 Defaulting to port 80802019/11/17 16:11:02 Listening on port 8080$ curl -X POST -H 'Content-type: application/json' -d '{"name": "Alice", "email": "alice@example.com"}' localhost:8080/people{"ID":1,"Name":"Alice","Email":"alice@example.com"}$ curl -X POST -H 'Content-type: application/json' -d '{"name": "Bob", "email": "bob@example.com"}' localhost:8080/people{"ID":2,"Name":"Bob","Email":"bob@example.com"}$ curl localhost:8080/people/1{"ID":1,"Name":"Alice","Email":"alice@example.com"}$ curl localhost:8080/people/2{"ID":2,"Name":"Bob","Email":"bob@example.com"}
```

[![スクリーンショット 2019-11-17 16.17.16.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F07e58280-88b7-d1cf-ca09-de10d27b35e7.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=eb8f9e3b70f5b10a12e2accc71889562)](https://camo.qiitausercontent.com/c5a7a4ff228adcdebe2c95732e026882d69876ad/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f30376535383238302d383862372d643163662d636130392d6465313064323762333565372e706e67)

[![スクリーンショット 2019-11-17 16.15.47.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2Fc1da8092-6a19-b83b-cfda-61d0d378dc91.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=d1887277151c5bb4ebd41f9cb5865cad)](https://camo.qiitausercontent.com/55331701e5c45a6d521d1977dcb0384821d43560/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f63316461383039322d366131392d623833622d636664612d3631643064333738646339312e706e67)

Gatewayを通じてアプリケーションがデータベースと疎通できました。

#### 2.4. MockDB用アダプターの実装

それでは次にメモリを利用した MockDB のアダプターを実装していきます。並行アクセスに対処するため、相互排他（Mutual eXclusion）を用いて制御します。また、RDBにおける自動採番の代わりに擬似乱数を用いてIDを発行します。

[![スクリーンショット 2019-11-17 18.57.23.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2Fcc465e1f-5d0f-cd0b-5a99-bfc6a7719a01.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=2a3b6e50c60f65b9a420e660caa87d9a)](https://camo.qiitausercontent.com/4f2192bafdd792fd053f6c8f6b37e9668f848a35/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f63633436356531662d356430662d636430622d356139392d6266633661373731396130312e706e67)

**追記2（2019/11/20）**

今回作ったの MockDB + MockGateway じゃなくて MockDB + FakeGateway だったかもしれません。

[Test Doubles — Fakes, Mocks and Stubs.](https://blog.pragmatists.com/test-doubles-fakes-mocks-and-stubs-1a7491dfa3da)

internal/service2/mock_gateway.go

```
package service2

import (
    "context"
    "fmt"
    "math/rand"
    "sync"
    "time"
)

var src = rand.NewSource(time.Now().UnixNano())

// MockGateway ...
type MockGateway struct {
    db *MockDB
}

// MockDB ...
type MockDB struct {
    mu   sync.RWMutex
    data map[int64]Person
}

// NewMockDB ...
func NewMockDB() *MockDB {
    return &MockDB{data: make(map[int64]Person)}
}

// NewMockGateway ...
func NewMockGateway(db *MockDB) Repository {
    return &MockGateway{db}
}

// RegisterPerson ...
func (r *MockGateway) RegisterPerson(ctx context.Context, p Person) (Person, error) {
    r.db.mu.Lock()
    defer r.db.mu.Unlock()

    // 割り当て可能なIDを探す
    var id int64
    for {
        id = src.Int63()
        _, ok := r.db.data[id]
        if !ok {
            break
        }
    }

    p.ID = id
    r.db.data[p.ID] = p

    return p, nil
}

// GetPersonByID ...
func (r *MockGateway) GetPersonByID(ctx context.Context, id int64) (Person, error) {
    r.db.mu.RLock()
    defer r.db.mu.RUnlock()

    if p, ok := r.db.data[id]; ok {
        return p, nil
    }
    return Person{}, fmt.Errorf("person not found - id: %d", id)
}
```

続いて main.go において DI の部分を修正し、Provider が Reository として Gateway ではなく MockGateway を利用するようにします。

cmd/http/main.go

```
func init() {
    // Mysql
    // c := mysql.Config{
    //  Host:                 os.Getenv("DB_HOST"),
    //  Port:                 os.Getenv("DB_PORT"),
    //  User:                 os.Getenv("DB_USER"),
    //  DBName:               os.Getenv("DB_NAME"),
    //  Passwd:               os.Getenv("DB_PASSWORD"),
    //  AllowNativePasswords: true,
    // }

    // db, err := mysql.Connect(c)
    // if err != nil {
    //  log.Fatalln(err)
    // }

    // DI
    // gateway2 := service2.NewGateway(db)
    // provider2 := service2.NewProvider(gateway2)
    mockGateway2 := service2.NewMockGateway(service2.NewMockDB())
    provider2 := service2.NewProvider(mockGateway2)

    ctrl := &controller.Controller{}
    ctrl2 := controller.NewController2(provider2)

    e.GET("/:message", ctrl.HandleMessage)
    e.GET("/people/:personID", ctrl2.HandlePersonGet)
    e.POST("/people", ctrl2.HandlePersonRegister)
}
```

Webサーバーを起動して curl でプログラムの処理を呼んでみます。なお、ご自身で試される際にはIDが擬似乱数になっている点にご注意ください。

```
$ go run cmd/http/main.go$ curl -X POST -H 'Content-type: application/json' -d '{"name": "Alice", "email": "alice@example.com"}' localhost:8080/people{"ID":4604021376263565598,"Name":"Alice","Email":"alice@example.com"}$ curl -X POST -H 'Content-type: application/json' -d '{"name": "Bob", "email": "bob@example.com"}' localhost:8080/people{"ID":6891153250004441175,"Name":"Bob","Email":"bob@example.com"}$ curl localhost:8080/people/4604021376263565598{"ID":4604021376263565598,"Name":"Alice","Email":"alice@example.com"}$ curl localhost:8080/people/6891153250004441175{"ID":6891153250004441175,"Name":"Bob","Email":"bob@example.com"}
```

[![スクリーンショット 2019-11-17 19.11.56.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F528ef51b-c188-e2c3-fb79-0e1968ced049.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=74da075c1eea6557861abe58682296db)](https://camo.qiitausercontent.com/677620f35f1aa0ce6f83430e661f55cc311ca700/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f35323865663531622d633138382d653263332d666237392d3065313936386365643034392e706e67)

MockDB を利用してもアプリケーションのコアロジックを実行することができました。

前のセクションと合わせて「外界 → アプリケーション」「アプリケーション → 外界」の両方向において、アプリケーションのコアロジックをいじることなく外界とコミュニケーションがとれるようになりました。

次のセクションでは「トランザクション管理」について見ていきます。

### 3. トランザクション管理

それではトランザクション管理方法の手法について考えていきます。

今回はサンプルプログラムとして「銀行口座における送金プログラム」を作ってみます。リレーショナルデータベースにおけるスキーマは次の通りです。

```
create table if not exists account (  id bigint auto_increment,  balance int,  primary key (id));
```

#### 3.1. トランザクションを管理するのはどのレイヤーか？

プログラムの作成に先立って「トランザクションを管理するのはどのレイヤーか」ということについて考えてみます。本記事に沿って考えると、トランザクションの開始を宣言するのは `Provider` なのか `Gateway` なのかということです。

『エリック・エヴァンスのドメイン駆動設計』ではアプリケーション層（本記事における Provider ）がトランザクションの開始を宣言するように記載されています。「第4章 ドメインを隔離する」の図4.1のシーケンス図で示されているので、書籍をお持ちの方は確認してみてください。本記事でもこれに習って Provider がトランザクションの開始を宣言できるように実装します。（なお、後述しますが、「トランザクションの開始は必ずアプリケーション層が通知しなければならない」とするのは問題があるように思います。場合によっては Gateway がトランザクションを開始したほうが良いケースもあるのでないかと感じています。）

#### 3.2. アプリケーションコアの実装

アプリケーションコアはインフラの知識が入り込まないように注意します。 Repository において `RunInTransaction()` というメソッドを定義していますが、引数・戻り値にもインフラに関する知識は入り込んでいません。

[![スクリーンショット 2019-11-17 19.52.34.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F734e75a8-6d58-522d-ccbb-4e7f9a67ada5.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=91741e47022783d73f9db057f6979373)](https://camo.qiitausercontent.com/21c0a23563dc5018f250f838813ac041cb2fc01e/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f37333465373561382d366435382d353232642d636362622d3465376639613637616461352e706e67)

internal/service3/model.go

```
package service3// Account ...type Account struct {    ID      int64 `db:"aikawarazu"`    Balance int   `db:"tekitode"`}// IsSufficient ...func (a *Account) IsSufficient(ammount int) bool {    return a.Balance >= ammount}// Transfer ...func (a *Account) Transfer(ammount int, to *Account) {    a.Balance -= ammount    to.Balance += ammount}
```

internal/service3/provider.go

```
package service3import "context"// Repository ...type Repository interface {    RunInTransaction(context.Context, func(context.Context) (interface{}, error)) (interface{}, error)    OpenAccount(ctx context.Context, initialAmmount int) (Account, error)    GetAccountsForTransfer(ctx context.Context, fromID, toID int64) (from, to Account, err error)    UpdateBalance(ctx context.Context, a Account) (Account, error)}
```

internal/service3/provider.go

```
package service3import (    "context"    "fmt")// Provider ...type Provider struct {    r Repository}// NewProvider ...func NewProvider(r Repository) *Provider {    return &Provider{r}}// OpenAccount ...func (p *Provider) OpenAccount(ctx context.Context, initialAmmount int) (Account, error) {    if initialAmmount <= 0 {        return Account{}, fmt.Errorf("provider: initial ammount must be greater than 0")    }    account, err := p.r.OpenAccount(ctx, initialAmmount)    if err != nil {        return Account{}, err    }    return account, nil}// Transfer ...func (p *Provider) Transfer(ctx context.Context, ammount int, fromID, toID int64) (from, to Account, err error) {    if fromID == toID {        return Account{}, Account{}, fmt.Errorf("provider: cannot transfer money to oneself")    }    type Accounts struct {        from Account        to   Account    }    // トランザクションで実行したい処理をまとめる    txFn := func(ctx context.Context) (interface{}, error) {        // 送金元、送金先の口座を取得する        from, to, err := p.r.GetAccountsForTransfer(ctx, fromID, toID)        if err != nil {            return Accounts{}, err        }        // 送金元の残高を確認        if !from.IsSufficient(ammount) {            return Accounts{}, fmt.Errorf("provider: balance is not sufficient - accountID: %d", from.ID)        }        // 送金する        from.Transfer(ammount, &to)        // 送金元の残高を更新する        from, err = p.r.UpdateBalance(ctx, from)        if err != nil {            return Accounts{}, err        }        // 送金先の残高を更新する        to, err = p.r.UpdateBalance(ctx, to)        if err != nil {            return Accounts{}, err        }        return Accounts{from: from, to: to}, nil    }    // トランザクションでまとめて処理を実行    v, err := p.r.RunInTransaction(ctx, txFn)    if err != nil {        return Account{}, Account{}, err    }    val, ok := v.(Accounts)    if !ok {        return Account{}, Account{}, fmt.Errorf("provider: an error occurs - transfer")    }    return val.from, val.to, nil}
```

#### 3.3. MockDB用アダプターの実装

今回は先に MockDB の方の Gateway を実装します。

インメモリの MockDB におけるトランザクションは context と mutual exclusion を利用して実装していきます。なお、今回は RDB のようなロールバックは実装せず、Redis のように間に他の処理が入らないようにするトランザクションを実装します。

internal/service3/mock_gateway.go

```
package service3import (    "context"    "fmt"    "math/rand"    "sync"    "time")type ctxKey stringconst txCtxKey ctxKey = "transaction"var src = rand.NewSource(time.Now().UnixNano())// MockGateway ...type MockGateway struct {    db *MockDB}// NewMockGateway ...func NewMockGateway(db *MockDB) Repository {    return &MockGateway{db}}// MockDB ...type MockDB struct {    mu   sync.RWMutex    data map[int64]Account}// NewMockDB ...func NewMockDB() *MockDB {    return &MockDB{data: make(map[int64]Account)}}// OpenAccount ...func (g *MockGateway) OpenAccount(ctx context.Context, initialAmmount int) (Account, error) {    if !isInTx(ctx) {        g.db.mu.Lock()        defer g.db.mu.Unlock()    }    // 割り当て可能なIDを探す    var id int64    for {        id = src.Int63()        _, ok := g.db.data[id]        if !ok {            break        }    }    a := Account{ID: id, Balance: initialAmmount}    g.db.data[id] = a    return a, nil}// GetAccountsForTransfer ...func (g *MockGateway) GetAccountsForTransfer(ctx context.Context, fromID, toID int64) (from, to Account, err error) {    if !isInTx(ctx) {        g.db.mu.Lock()        defer g.db.mu.Unlock()    }    var ok bool    from, ok = g.db.data[fromID]    if !ok {        return Account{}, Account{}, fmt.Errorf("gateway: account not found - accoutID: %d", fromID)    }    to, ok = g.db.data[toID]    if !ok {        return Account{}, Account{}, fmt.Errorf("gateway: account not found - accoutID: %d", toID)    }    return from, to, nil}// UpdateBalance ...func (g *MockGateway) UpdateBalance(ctx context.Context, a Account) (Account, error) {    if !isInTx(ctx) {        g.db.mu.Lock()        defer g.db.mu.Unlock()    }    g.db.data[a.ID] = a    return a, nil}// RunInTransaction ...func (g *MockGateway) RunInTransaction(ctx context.Context, txFn func(context.Context) (interface{}, error)) (interface{}, error) {    // 多重トランザクションはエラーとする    if isInTx(ctx) {        return nil, fmt.Errorf("gateway: detect nested transaction")    }    // context をデコレートして transaction context を生成する    txCtx := genTxCtx(ctx)    // ロックを取得する（ロックの取得から解放までの間がトランザクションとなる）    g.db.mu.Lock()    defer g.db.mu.Unlock()    // transaction 処理を実行する    return txFn(txCtx)}func isInTx(ctx context.Context) bool {    if val, ok := ctx.Value(txCtxKey).(bool); ok {        return val    }    return false}func genTxCtx(ctx context.Context) context.Context {    if isInTx(ctx) {        return ctx    }    return context.WithValue(ctx, txCtxKey, true)}
```

#### 3.4. ユーザー側アダプター（controller）の実装とルーティングの追加

[![スクリーンショット 2019-11-17 20.07.56.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F2cb6277a-d63b-b39a-c474-81101f1df0d8.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=27764c2f38936a53baf051901b40d1db)](https://camo.qiitausercontent.com/62724ca7cecc7770bbcd7cc1199c6525e24ad6fd/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f32636236323737612d643633622d623339612d633437342d3831313031663164663064382e706e67)

cmd/http/controller/controller3.go

```
package controllerimport (    "net/http"    "github.com/labstack/echo/v4"    "github.com/rema424/hexample/internal/service3")// Controller3 ...type Controller3 struct {    p *service3.Provider}// NewController3 ...func NewController3(p *service3.Provider) *Controller3 {    return &Controller3{p}}// HandleAccountOpen ...// curl -X POST -H 'Content-type: application/json' -d '{"ammount": 1000}' localhost:8080/accountsfunc (ctrl *Controller3) HandleAccountOpen(c echo.Context) error {    in := struct {        Ammount int `json:"ammount"`    }{}    if err := c.Bind(&in); err != nil {        return c.JSON(http.StatusBadRequest, err.Error())    }    // TODO: implement    // if err := c.Validate(&in); err != nil {    //  return c.JSON(http.StatusUnprocessableEntity, err.Error())    // }    ctx := c.Request().Context()    psn, err := ctrl.p.OpenAccount(ctx, in.Ammount)    if err != nil {        return c.JSON(http.StatusInternalServerError, err.Error())    }    return c.JSON(http.StatusOK, psn)}// HandleMoneyTransfer ...// curl -X POST -H 'Content-type: application/json' -d '{"fromId": , "toId": , "ammount": 1000}' localhost:8080/accounts/transferfunc (ctrl *Controller3) HandleMoneyTransfer(c echo.Context) error {    in := struct {        FromAccountID int64 `json:"fromId"`        ToAccountID   int64 `json:"toId"`        Ammount       int   `json:"ammount"`    }{}    if err := c.Bind(&in); err != nil {        return c.JSON(http.StatusBadRequest, err.Error())    }    // TODO: implement    // if err := c.Validate(&in); err != nil {    //  return c.JSON(http.StatusUnprocessableEntity, err.Error())    // }    ctx := c.Request().Context()    from, to, err := ctrl.p.Transfer(ctx, in.Ammount, in.FromAccountID, in.ToAccountID)    if err != nil {        return c.JSON(http.StatusInternalServerError, err.Error())    }    return c.JSON(http.StatusOK, map[string]interface{}{"from": from, "to": to})}
```

cmd/http/main.go

```
func init() {    // Mysql    // c := mysql.Config{    //  Host:                 os.Getenv("DB_HOST"),    //  Port:                 os.Getenv("DB_PORT"),    //  User:                 os.Getenv("DB_USER"),    //  DBName:               os.Getenv("DB_NAME"),    //  Passwd:               os.Getenv("DB_PASSWORD"),    //  AllowNativePasswords: true,    // }    // db, err := mysql.Connect(c)    // if err != nil {    //  log.Fatalln(err)    // }    // acsr, err := sqlxx.Open(db)    // if err != nil {    //  log.Fatalln(err)    // }    // Service2    // gateway2 := service2.NewGateway(db)    // provider2 := service2.NewProvider(gateway2)    mockGateway2 := service2.NewMockGateway(service2.NewMockDB())    provider2 := service2.NewProvider(mockGateway2)    // Service3    // gateway3 := service3.NewGateway(acsr)    // provider3 := service3.NewProvider(gateway3)    mockGateway3 := service3.NewMockGateway(service3.NewMockDB())    provider3 := service3.NewProvider(mockGateway3)    ctrl := &controller.Controller{}    ctrl2 := controller.NewController2(provider2)    ctrl3 := controller.NewController3(provider3)    e.GET("/:message", ctrl.HandleMessage)    e.GET("/people/:personID", ctrl2.HandlePersonGet)    e.POST("/people", ctrl2.HandlePersonRegister)    e.POST("/accounts", ctrl3.HandleAccountOpen)    e.POST("/accounts/transfer", ctrl3.HandleMoneyTransfer)}
```

Webサーバーを起動して curl でリクエストを実行してみます。

```
$ go run cmd/http/main.go$ curl -X POST -H 'Content-type: application/json' -d '{"ammount": 1000}' localhost:8080/accounts{"ID":6604275530202776837,"Balance":1000}$ curl -X POST -H 'Content-type: application/json' -d '{"ammount": 1000}' localhost:8080/accounts{"ID":8605590474089424096,"Balance":1000}$ curl -X POST -H 'Content-type: application/json' -d '{"fromId": 6604275530202776837, "toId": 8605590474089424096, "ammount": 300}' localhost:8080/accounts/transfer{"from":{"ID":6604275530202776837,"Balance":700},"to":{"ID":8605590474089424096,"Balance":1300}}$ curl -X POST -H 'Content-type: application/json' -d '{"fromId": 6604275530202776837, "toId": 8605590474089424096, "ammount": 300}' localhost:8080/accounts/transfer{"from":{"ID":6604275530202776837,"Balance":400},"to":{"ID":8605590474089424096,"Balance":1600}}$ curl -X POST -H 'Content-type: application/json' -d '{"fromId": 6604275530202776837, "toId": 8605590474089424096, "ammount": 300}' localhost:8080/accounts/transfer{"from":{"ID":6604275530202776837,"Balance":100},"to":{"ID":8605590474089424096,"Balance":1900}}$ curl -X POST -H 'Content-type: application/json' -d '{"fromId": 6604275530202776837, "toId": 8605590474089424096, "ammount": 300}' localhost:8080/accounts/transfer"provider: balance is not sufficient - accountID: 6604275530202776837"$ curl -X POST -H 'Content-type: application/json' -d '{"fromId": 6604275530202776837, "toId": 8605590474089424096, "ammount": 100}' localhost:8080/accounts/transfer{"from":{"ID":6604275530202776837,"Balance":0},"to":{"ID":8605590474089424096,"Balance":2000}}
```

[![スクリーンショット 2019-11-17 20.23.08.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F890b5e22-e8fd-195b-8837-a96a27f18fb6.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=7171c155730a380f7e4d1019cdf11b1c)](https://camo.qiitausercontent.com/f740ac8137ce584c34ddd6cbd6d7ad0a2f7d43b4/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f38393062356532322d653866642d313935622d383833372d6139366132376631386662362e706e67)

#### 3.5. MySQL用アダプターの実装

次に MySQL におけるトランザクション管理の実装をしていきます。なんですが、、アプリケーションコアにDBの知識を漏らさずにトランザクションを開始するための実装が少し面倒だったので、トランザクション管理部分だけライブラリ化してしまいました。sqlx のラッパーになっています。ソースコードを確認したい方はリポジトリの方を覗いてみてください。テストなんてしてな（ry



[rema424/sqlxxhttps://github.com![img](https://opengraph.githubassets.com/7ad3e534f4161485749221002e7e5d3c3f94e8f21ea150a4ec1af35d18909394/rema424/sqlxx)](https://github.com/rema424/sqlxx)



internal/service3/gateway.go

```
package service3

import (
    "context"
    "fmt"
    "log"

    "github.com/rema424/sqlxx"
)

// Gateway ...
type Gateway struct {
    db *sqlxx.Accessor
}

// NewGateway ...
func NewGateway(db *sqlxx.Accessor) Repository {
    return &Gateway{db}
}

// OpenAccount ...
func (g *Gateway) OpenAccount(ctx context.Context, initialAmmount int) (Account, error) {
    q := `INSERT INTO account (balance) VALUES (?);`

    res, err := g.db.Exec(ctx, q, initialAmmount)
    if err != nil {
        return Account{}, err
    }

    id, err := res.LastInsertId()
    if err != nil {
        return Account{}, nil
    }

    return Account{ID: id, Balance: initialAmmount}, nil
}

// GetAccountsForTransfer ...
func (g *Gateway) GetAccountsForTransfer(ctx context.Context, fromID, toID int64) (from, to Account, err error) {
    // 送金に関わるアカウントはロックをかけて（FOR UPDATE）取得する
    q := `
SELECT
  COALESCE(id, 0) AS 'aikawarazu',
  COALESCE(balance, 0) AS 'tekitode'
FROM account
WHERE id = ? OR id = ?
FOR UPDATE;
`
    var dest []Account
    if err := g.db.Select(ctx, &dest, q, fromID, toID); err != nil {
        return from, to, err
    }

    if len(dest) != 2 {
        return from, to, fmt.Errorf("gateway: account not found for transfer")
    }

    for _, a := range dest {
        if a.ID == fromID {
            from = a
        } else if a.ID == toID {
            to = a
        }
    }

    return from, to, nil
}

// UpdateBalance ...
func (g *Gateway) UpdateBalance(ctx context.Context, a Account) (Account, error) {
    q := `UPDATE account SET balance = :tekitode WHERE id = :aikawarazu;`
    _, err := g.db.NamedExec(ctx, q, a)
    if err != nil {
        return Account{}, err
    }
    return a, nil
}

// RunInTransaction ...
func (g *Gateway) RunInTransaction(ctx context.Context, txFn func(context.Context) (interface{}, error)) (interface{}, error) {
    v, err, rlbkErr := g.db.RunInTx(ctx, txFn)
    if rlbkErr != nil {
        log.Printf("gateway: failed to rollback - err: %s\n", rlbkErr.Error())
    }
    return v, err
}
```

main.go の DI 部分、トランザクション管理用のライブラリ利用部分を書き換えます。

cmd/http/main.go

```
func init() {    // Mysql    c := mysql.Config{        Host:                 os.Getenv("DB_HOST"),        Port:                 os.Getenv("DB_PORT"),        User:                 os.Getenv("DB_USER"),        DBName:               os.Getenv("DB_NAME"),        Passwd:               os.Getenv("DB_PASSWORD"),        AllowNativePasswords: true,    }    db, err := mysql.Connect(c)    if err != nil {        log.Fatalln(err)    }    acsr, err := sqlxx.Open(db)    if err != nil {        log.Fatalln(err)    }    // Service2    // gateway2 := service2.NewGateway(db)    // provider2 := service2.NewProvider(gateway2)    mockGateway2 := service2.NewMockGateway(service2.NewMockDB())    provider2 := service2.NewProvider(mockGateway2)    // Service3    gateway3 := service3.NewGateway(acsr)    provider3 := service3.NewProvider(gateway3)    // mockGateway3 := service3.NewMockGateway(service3.NewMockDB())    // provider3 := service3.NewProvider(mockGateway3)    ctrl := &controller.Controller{}    ctrl2 := controller.NewController2(provider2)    ctrl3 := controller.NewController3(provider3)    e.GET("/:message", ctrl.HandleMessage)    e.GET("/people/:personID", ctrl2.HandlePersonGet)    e.POST("/people", ctrl2.HandlePersonRegister)    e.POST("/accounts", ctrl3.HandleAccountOpen)    e.POST("/accounts/transfer", ctrl3.HandleMoneyTransfer)}
```

データベースにDDLを適用してからWebサーバーを起動し、curlでリクエストを飛ばします。

```
$ go run cmd/http/main.go$ curl -X POST -H 'Content-type: application/json' -d '{"ammount": 1000}' localhost:8080/accounts{"ID":1,"Balance":1000}$ curl -X POST -H 'Content-type: application/json' -d '{"ammount": 1000}' localhost:8080/accounts{"ID":2,"Balance":1000}$ curl -X POST -H 'Content-type: application/json' -d '{"fromId": 1, "toId": 2, "ammount": 300}' localhost:8080/accounts/transfer{"from":{"ID":1,"Balance":700},"to":{"ID":2,"Balance":1300}}$ curl -X POST -H 'Content-type: application/json' -d '{"fromId": 1, "toId": 2, "ammount": 300}' localhost:8080/accounts/transfer{"from":{"ID":1,"Balance":400},"to":{"ID":2,"Balance":1600}}$ curl -X POST -H 'Content-type: application/json' -d '{"fromId": 1, "toId": 2, "ammount": 300}' localhost:8080/accounts/transfer{"from":{"ID":1,"Balance":100},"to":{"ID":2,"Balance":1900}}$ curl -X POST -H 'Content-type: application/json' -d '{"fromId": 1, "toId": 2, "ammount": 300}' localhost:8080/accounts/transfer"provider: balance is not sufficient - accountID: 1"
```

[![スクリーンショット 2019-11-17 20.35.31.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2Fbba88e99-69d1-ad04-0b90-c42942eab500.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=56409c58a1871b5dcfcaa26916e85b92)](https://camo.qiitausercontent.com/0c7d9cda52759e19ef36c5444ab6c51349c53534/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f62626138386539392d363964312d616430342d306239302d6334323934326561623530302e706e67)

[![スクリーンショット 2019-11-17 20.35.57.png](https://qiita-user-contents.imgix.net/https%3A%2F%2Fqiita-image-store.s3.ap-northeast-1.amazonaws.com%2F0%2F237863%2F70ec0740-ca2b-ac54-8555-197e7f263692.png?ixlib=rb-4.0.0&auto=format&gif-q=60&q=75&s=fb385482a58a71a586e91afe38a3a199)](https://camo.qiitausercontent.com/aa8a332db78442deb95e08433d9e6b1e3becc3bf/68747470733a2f2f71696974612d696d6167652d73746f72652e73332e61702d6e6f727468656173742d312e616d617a6f6e6177732e636f6d2f302f3233373836332f37306563303734302d636132622d616335342d383535352d3139376537663236333639322e706e67)

アプリケーションコアがインフラの知識を持つことなくトランザクションの実装ができました。

以上でサンプルアプリケーションの実装は終わりです。

## google/wire による DI について

[google/wire](https://github.com/google/wire) というツールを使うと DI を簡単に行うことができます。`wire` を利用した DI の実装のサンプルを GitHub リポジトリに含めているので興味がある方は覗いてみてください。



[rema424/hexamplehttps://github.com![img](https://opengraph.githubassets.com/24e233e3fa819ddb13be3e2ff6e823340944e32533e806555f60bf863d67e0a3/rema424/hexample)](https://github.com/rema424/hexample/tree/master/cmd/httpwire)



## 悩み1: トランザクションの管理はアプリケーション層？

記事の途中でも言及しましたが、トランザクションの管理をどこで行うかということは悩ましいです。エヴァンスはアプリケーション層でのトランザクション管理を提示していますが、場合によっては Gateway でトランザクションを管理したいこともあります。

例えば一対多の関係を持つオブジェクト群の保存処理です。例えば「注文」と「注文詳細」の登録です。

```
type Order struct {
    ID           int64
    TotalPrice   int
    OrderDetails []OrderDetail
}

type OrderDetail struct {
    ID            int64
    Product       Product
    Quantity      int
    SubTotalPrice int
}

type Product struct {
    ID    int64
    Name  string
    Price int
}
```

上記の `Order` オブジェクト（集約）をリレーショナルデータベースに登録しようと思ったら、`order` テーブルと `order_detail` テーブルに分けて保存することになるかと思います。この処理はもちろん同一トランザクション内で実行されます。

一方で、リレーショナルデータベースではなく本記事のサンプルプログラムのようにGoアプリケーションのメモリ上に保存する場合を考えてみます。この場合は `Order` と `OrderDetail` を分けて保存する必要はなく、`Order` オブジェクトをそのまま保存できます。この場合にはトランザクションを開始する必要はありません。

```
type MockDB struct {    mu   sync.RWMutex    data map[int64]Order}
```

このケースでトランザクションを獲得して処理を行うかどうかは、アプリケーション層の関心の範囲内ではなく、インフラの関心の範囲です。よって、トランザクションの管理を Provider で行うのは好ましくありません。

『エリック・エヴァンスのドメイン駆動設計』「第6章 ドメインオブジェクトのライフサイクル」には次の記載があります。

> トランザクションの制御をクライアントに委ねること。リポジトリはデータベースに対する挿入と削除を行うが、通常は何もコミットしない。例えば、保存した後にはコミットしたくなるが、おそらくクライアントには、作業ユニット（unit of work）を正しく開始し、コミットするためのコンテキストがある。トランザクション管理は、リポジトリが手を出さないでいる方が単純になる。

しかしながら、場合によっては Gateway でトランザクションを管理した方がいいのではないでしょうか。

**追記1（2019/11/19）**

リレーショナルデータベースの **AUTOCOMMIT** を有効にしているか無効にしているかで事情が少し変わる気がしてきました。AUTOCOMMIT を OFF に設定し、書き込み処理の際には必ず明示的にコミットしなければならない（RunInTransactionを呼ばなければならない）というポリシーの元で開発するなら、トランザクション管理は Provider に一元化できるかもしれません。

## 悩み2: 複数のデータストアに保存する場合

文字列や数値といったプリミティブなデータはリレーショナルデータベースに、画像などのバイナリデータはオブジェクトストレージ（GCS・S3）に保存するというケースは多いかと思います。もしくはDBにデータを保存しつつ、同一の処理の流れでメッセージングキューにもデータを送りたい場合もあります。

このように複数のデータストアが登場するとき、データストアごとにアダプター（Gateway）を作成すべきでしょうか？

『エリック・エヴァンスのドメイン駆動設計』「第6章 ドメインオブジェクトのライフサイクル」には次の記載があります。

> 完全にインスタンス化されたオブジェクトかオブジェクトのコレクションを戻すメソッドを提供すること。それによって、実際のストレージや問い合わせの技術をカプセル化すること。実際に直接的なアクセスを必要とする集約ルートに対してのみ、リポジトリを提供すること。

最後の一文から読み取れるのは、リポジトリは集約に関心を向けて作成されるものであって、ストレージに関心を向けて作成されるものではないということです。

「MysqlRepositoryImpl」「DynamoDBAccessor」のようにストレージごとに Gateway を作成するのではなく、単一の Gateway の中に複数のストレージへのアクセス手法をカプセル化するのがいいのかなと現時点では考えています。

## おわりに

Go にはデファクトスタンダードと呼ばれるフレームワークがなく、ディレクトリ構成に悩んでいる方が多い印象です。僕もその1人でした。

アプリケーションの規模にも依るのでこれと言った正解はないかとは思いますが、常に学習を続けてメンテナンス性の高い
ソフトウェアを作っていきたいです。

以上で本記事は終わりです。

## ソースコード全体（再掲）



[rema424/hexamplehttps://github.com![img](https://opengraph.githubassets.com/24e233e3fa819ddb13be3e2ff6e823340944e32533e806555f60bf863d67e0a3/rema424/hexample)](https://github.com/rema424/hexample)



## 参考

- [Standard Go Project Layout](https://github.com/golang-standards/project-layout)
- [Goにはディレクトリ構成のスタンダードがあるらしい。](https://qiita.com/sueken/items/87093e5941bfbc09bea8)
- [Practical Go: Real world advice for writing maintainable Go programs](https://dave.cheney.net/practical-go/presentations/qcon-china.html)
- [ヘキサゴナルアーキテクチャ(Hexagonal architecture翻訳)](https://blog.tai2.net/hexagonal_architexture.html)
- [クリーンアーキテクチャ(The Clean Architecture翻訳)](https://blog.tai2.net/the_clean_architecture.html)
- [The Onion Architecture](https://jeffreypalermo.com/2008/07/the-onion-architecture-part-1/)
- [ドメイン分析を使用したマイクロサービスのモデル化 | Microsoft Docs](https://docs.microsoft.com/ja-jp/azure/architecture/microservices/model/domain-analysis#define-bounded-contexts)
- [一般的な Web アプリケーション アーキテクチャ | Microsoft Docs](https://docs.microsoft.com/ja-jp/dotnet/architecture/modern-web-apps-azure/common-web-application-architectures)
- [Dockerコンテナのおもしろい名前](https://deeeet.com/writing/2014/07/15/docker-container-name/)
- [ファイル構成 - React](https://ja.reactjs.org/docs/faq-structure.html)
- [『エリック・エヴァンスのドメイン駆動設計』](https://www.shoeisha.co.jp/book/detail/9784798126708)
- [『実践ドメイン駆動設計』](https://www.shoeisha.co.jp/book/detail/9784798131610)