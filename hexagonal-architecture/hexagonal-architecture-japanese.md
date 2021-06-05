[blog.tai2.net](https://blog.tai2.net/)

# ヘキサゴナルアーキテクチャ(Hexagonal architecture翻訳)

[tai2](https://tai2.net/) 2015-10-09

[Alistair Cockburn](http://alistair.cockburn.us/) による [Hexagonal architecture](http://alistair.cockburn.us/Hexagonal+architecture) の翻訳です。PoEAAで言及されていることから、2002年ごろにはすでにC2 Wikiにページがあった模様。似たようなアーキテクチャである [クリーンアーキテクチャ](https://blog.tai2.net/the_clean_architecture.html) も翻訳したので参考にしてください。

この記事は著者から許可を得て公開しています。Thanks to Alistair Cockburn!

目次

- パターン: Ports and Adapters (構造に関するパターン)
  - [意図](https://blog.tai2.net/hexagonal_architexture.html#id4)
  - [動機](https://blog.tai2.net/hexagonal_architexture.html#id5)
  - [解決法の本質](https://blog.tai2.net/hexagonal_architexture.html#id6)
  - [構造](https://blog.tai2.net/hexagonal_architexture.html#id7)
  - サンプルコード
    - [ステージ1: FIT アプリ 定数をモックデータベースとして](https://blog.tai2.net/hexagonal_architexture.html#fit)
    - [ステージ2: UI アプリ 定数をモックデータベースとして](https://blog.tai2.net/hexagonal_architexture.html#ui)
    - [ステージ3: (FITまたはUI) アプリ モックデータベース](https://blog.tai2.net/hexagonal_architexture.html#fitui)
  - 応用ノート
    - [左右の非対称性](https://blog.tai2.net/hexagonal_architexture.html#id10)
    - [ユースケースとアプリケーションの境界](https://blog.tai2.net/hexagonal_architexture.html#id11)
    - [ポートはいくつ?](https://blog.tai2.net/hexagonal_architexture.html#id12)
    - [既知の用例](https://blog.tai2.net/hexagonal_architexture.html#id13)
    - [Mac, Windows, Google, Flickr, Web 2.0](https://blog.tai2.net/hexagonal_architexture.html#mac-windows-google-flickr-web-2-0)
    - [ストアード出力](https://blog.tai2.net/hexagonal_architexture.html#id14)
    - [C2-wikiからの匿名の例](https://blog.tai2.net/hexagonal_architexture.html#c2-wiki)
    - [分散された、大きなチームでの開発](https://blog.tai2.net/hexagonal_architexture.html#id15)
    - [UIとアプリケーションロジックの開発を分割する](https://blog.tai2.net/hexagonal_architexture.html#id16)
  - 関連するパターン
    - [Adapter](https://blog.tai2.net/hexagonal_architexture.html#adapter)
    - [Model-View-Controller](https://blog.tai2.net/hexagonal_architexture.html#model-view-controller)
    - [Mock Objects と Loopback](https://blog.tai2.net/hexagonal_architexture.html#mock-objects-loopback)
    - [Pedestals](https://blog.tai2.net/hexagonal_architexture.html#pedestals)
    - [Checks](https://blog.tai2.net/hexagonal_architexture.html#checks)
    - [Dependency Inversion (Dependency Injection) と SPRING](https://blog.tai2.net/hexagonal_architexture.html#dependency-inversion-dependency-injection-spring)
  - [謝辞](https://blog.tai2.net/hexagonal_architexture.html#id18)
  - [リファレンスと、関連した読み物](https://blog.tai2.net/hexagonal_architexture.html#id19)

------

UIやデータベースがなくても動くようにアプリケーションを作ること。そうすれば、アプリケーションに対して自動化された回帰テストを動かすことができるし、データベースが利用できなくなっても動作するし、ユーザーの介入がなくともアプリケーション群を結合できる。

![Microscope](https://blog.tai2.net/images/hexagonal_architecture/3005.jpeg)

## [パターン: Ports and Adapters (構造に関するパターン)](https://blog.tai2.net/hexagonal_architexture.html#id20)

**別名: 「Ports & Adapter」**

**別名: 「Hexagonal Architecture」**

### [意図](https://blog.tai2.net/hexagonal_architexture.html#id21)

アプリケーションを、ユーザー、プログラム、自動テストあるいはバッチスクリプトから、同じように駆動できるようにする。そして、実際のランタイムデバイスとデータベースから隔離して、開発とテストをできるようにする。

イベントが外側の世界からポートに届くと、特定テクノロジーのアダプターが、利用可能な手続き呼び出しか、メッセージにそれを変換して、アプリケーションに渡す。よろこばしいことに、アプリケーションは、入力デバイスの正体を知らない。アプリケーションがなにかを送る必要があるとき、それはポートを通じてアダプターに送られて、受信側のテクノロジーが必要とする信号を生む(人力であれ自動であれ)。アプリケーションは、実際に他方のアダプターの正体を知ることはなしに、全側面のアダプターと意味的に完全なやりとりをする。

![Figure 1](https://blog.tai2.net/images/hexagonal_architecture/2301.gif)

図1

### [動機](https://blog.tai2.net/hexagonal_architexture.html#id22)

数年来、ソフトウェアアプリケーションで一番怖いことのひとつは、ビジネスロジックがユーザーインターフェイスコードに侵入することだった。これが引き起す問題は、3つある:

- はじめに、システムを自動テストスイートで綺麗にテストすることができない。なぜなら、テストを必要とするロジックの部分が、フィールドサイズやボタン配置など、頻繁に変わるビジュアルの詳細に依存するからだ。
- 同じ理由により、人間駆動のシステム使用から、バッチ処理システムに移行することが不可能になる。
- これも同じ理由から、他のプログムからプログラムを駆動したくなったときに、そうすることが難しいか、または不可能になる。

(多くの組織によって繰り返し)試みられた解法は、アーキテクチャに新しい層を足すことだ。そのときには、今度は、本当に絶対に、ビジネスロジックが新しいレイヤーに置かれることはないという取り決めをする。しかしながら、取り決めへの違反が起きたときに検出する仕組みはなく、組織は、数年後、新しいレイヤーがビジネスロジックでとっちらかっており、同じ問題が起きたことに気付く。

アプリケーションの提供する機能の部品すべてを、API(application programmed interface)ないし関数呼び出しと通して利用できるとしたらどうなるか、想像してみよう。この状況では、テストないし品質保証部は、新しいコードが以前動いていた機能を壊したときにそれを検出するために、アプリケーションに対して自動化されたテストスクリプトを走らせることができる。ビジネスエキスパートは、GUIの詳細が確定する前に、自動化されたテストケースを作成できる。それは、プログラマーに、自分の作業が正しく完了したことを教える(そして、それらのテストは、テスト部署によって実行されることになる)。APIだけが利用可能でも、アプリケーションを「ヘッドレス」モード[1](https://blog.tai2.net/hexagonal_architexture.html#sf-hexagonal_architexture-1)でデプロイすることができる。そして、他のプログラムは、その機能を利用することができる。これは、複雑なアプリケーションスイートの設計全体をシンプルにできるし、また、B2Bサービスアプリケーション群が、互いに、ウェブ経由で人力を介さずに、利用しあうことを許す。最後に、自動化された回帰テストは、ビジネスロジックをプレゼンテーションレイヤーから隔離しておくという取り決めへの違反を検出できる。組織は、ロジックの漏洩を検出して訂正できる。

興味深い同様の問題が、アプリケーションの「反対側」と通常考えられている部分にも存在する。アプリケーションのロジックが、外部のデータベースや他のサービスと結びつけられている部分だ。データベースサーバーが落ちたとき、あるいは、大規模な改善や置き換え中のときなど、プログラマーは作業をすることができない。なぜなら、作業がデータベースの存在と結びついているからだ。これは、遅延のコストと、しばしば人々の間の嫌な雰囲気を生じさせる。

二つの問題が関連していることは明らかではない、しかし、解決法の本質で見るように、これらの間には対称性がある。

### [解決法の本質](https://blog.tai2.net/hexagonal_architexture.html#id23)

ユーザーサイドとサーバーサイドの問題は、どちらも、実際には、設計とプログラミングにおける同様の誤りから生じている。ビジネスロジックと、外部エンティティーとのやりとりが、絡み合っているのだ。利用すべき非対称性は、アプリケーションの「左側」と「右側」ではなく、アプリケーションの「内側」と「外側」だ。従うべきルールは、「内側」の部分にあるコードが「外側」の部分に漏れ出さないようにすべき、ということだ。

左右または上下の非対称性からはしばらく離れて、アプリケーションが、「ポート」を越えて外部のエージェントと通信することを見よう。「ポート」という語には、オペレーティングシステムの「ポート」を想起させることが期待される。それは、ポートのプロトコルに従うデバイスが、差し込まれる場所だ。そして、電子ガジェットの「ポート」、ここでもまた、機械的かつ電気的なプロトコルに適合するデバイスが、差し込まれる。

- ポート用のプロトコルは、2つのデバイスの会話を目的として、与えられる。

このプロトコルは、アプリケーションプログラムインターフェイス(API)の形を取る。

各外部デバイスには「アダプタ」があり、それは、API定義をデバイスが必要とする信号に変える、逆もまた然り。グラフィカルユーザーインターフェイスすなわちGUIは、人間の動作をポートのAPIと対応付けるアダプタの例だ。同じポートに適合するその他のアダプタは、FIT[2](https://blog.tai2.net/hexagonal_architexture.html#sf-hexagonal_architexture-2)やFitnessのようなテストハーネス、バッチドライバー、あるいは、企業やネットをまたがるアプリケーション間の通信で必要とされるあらゆるコードだ。

アプリケーションの他方では、アプリケーションは、データを取得するために外部のエンティティーと通信する。そのプロトコルの典型は、データベースプロトコルだ。アプリケーションの観点からは、もしデータベースがSQLデータベースから、フラットなファイルや、その他のデータベースに移行しても、APIとの会話は変わるべきではない。ゆえに、同じポートへの追加のアダプターは、SQLアダプター、フラットファイルアダプター、そしてもっとも重要なものとして、「モック」データベースのアダプターを含む。これは、メモリ内に居座るもので、実際のデータベースの存在にまったく依存しない。

多くのアプリケーションは、ポートを2つだけ持つ: ユーザー側の対話と、データベース側の対話だ。これは、非対称的な様相をもたらすので、アプリケーションを1次元、3,4,あるいは5層のスタックアーキテクチャで構築するのが自然だと思わせる。

これらの素描には2つの問題がある。はじめに、そしてもっとも悪いのは、人々がレイヤー素描の「線」を深刻に受け取らない傾向があるということだ。かれらは、アプリケーションロジックをレイヤー境界を越えて侵食させ、上述した問題を生む。2番目に、アプリケーションには、2つ以上のポートがあるかもしれないということだ、そうなると、そのアーキテクチャは、1次元レイヤーの素描に適合しない。

ヘキサゴナル(またはPorts and adapters)アーキテクチャーでは、こうした状況において対称なものがなにもないことによって、問題を解決する: 内部にはアプリケーションがあり、いくつかのポートごしに外部のものと通信する。アプリケーションの外側のものは、対称的に扱うことができる。

六角形は、視覚的に、

1. 内側と外側の非対称性と、ポートの似たような特性(1次元のレイヤーの絵と、それが想起させるものから完全に離れるために)と、
2. 定義された数の異なるポートの存在 ー 2,3,あるいは4つの(4が、わたしがこれまで遭遇した中では一番多かった)

に焦点を当てるよう意図されている。

この六角形は、6という数字が重要だから六角形なのではなく、人々が、必要に応じて、ポートとアダプターを挿入するための余分を素描に持たせ、1次元レイヤーの素描に制限されないようにするから、六角形なのだ。ヘキサゴナルアーキテクチャという用語は、この視覚効果から来ている。

「ポートとアダプター」という用語は、素描のパーツの「目的」を強調している。ポートは、目的の会話を識別する。典型的には、どのひとつのポートにも複数のアダプターがあるだろう。それらは、ポートに差し込まれるさまざまな技術のためのものだ。典型的には、これには、留守番電話、人間の声、プッシュホン、グラフィカルユーザーインターフェイス、テストハーネス、バッチドライバー、HTTPインターフェイス、プログラムからプログラムへの直接インターフェイス、(インメモリ)モックデータベース、実際のデータベース(おそらく、開発用、テスト用、実利用用で異なるもの)が含まれる。

応用ノートでは、左右の非対称性について再度述べる。しかしながら、このパターンの主たる目的は、内側と外側の非対称性にフォーカスすることであり、つかの間、外部の要素がアプリケーションの観点からは等しいふりをしているのだ。

### [構造](https://blog.tai2.net/hexagonal_architexture.html#id24)

![Figure 2](https://blog.tai2.net/images/hexagonal_architecture/2302.gif)

図2

図2は、2つのアクティブなポートと、各ポートに複数のアダプターを持つアプリケーションを示している。2つのポートは、アプリケーション制御側と、データ取り出し側だ。この素描は、アプリケーションが、自動化されたシステムレベルの回帰テスト、人間のユーザー、リモートHTTPアプリケーション、あるいは、他のローカルアプリケーションから、同じように駆動されることを示している。データ側では、アプリケーションは、外部のデータベースから分離して実行されるよう構成することができる。これには、インメモリのオラクル(すなわち「モック」)データベースの置き換えを利用する。あるいは、テストまたはランタイムのデータベースに対して、動かすことができる。アプリケーションの機能的な仕様は、(もしかするとユースケース内で)内側の六角形のインターフェイスに対して作られるのであって、使われるかもしれない外部のテクノロジーに対してではない。

![Figure 3](https://blog.tai2.net/images/hexagonal_architecture/2303.gif)

図3

図3は、同じアプリケーションを3レイヤーのアーキテクチャーに対応付けた素描を示している。簡単のために、素描では、各ポートにつき2つのアダプターしか見せていない。この素描は、複数のアダプターが、いかにして上下のレイヤーに適合するか、そして、システム開発の間にいろいろなアダプターが使われるシーケンスを示すことを意図している。数字の付けられた矢印は、チームがアプリケーションの開発と使用をするかもしれない順番を示している。

1. FITテストハーネスを使ってアプリケーションを駆動する、そして、モック(インメモリ)データベースを実際のデータベースの代替として使う。
2. GUIをアプリケーションに追加しつつ、依然モックデータベースを使う。
3. 統合テスト中、自動化されたテストスクリプト(例えばCruise Control[3](https://blog.tai2.net/hexagonal_architexture.html#sf-hexagonal_architexture-3)から)で、アプリケーションをテストデータを保持した実際のデータベースに対して駆動させる。
4. 実際の利用で、アプリケーションを使う人が、生きたデータベースにアクセスする。

### [サンプルコード](https://blog.tai2.net/hexagonal_architexture.html#id25)

Ports & Adaptersのデモをするのにもっとも簡単なアプリケーションが、幸運にもFITのドキュメントに付いてくる。シンプルな割引計算のアプリケーションだ:

```
discount(amount) = amount * rate(amount);
```

我々のバージョンでは、合計額はユーザーから、レートはデータベースから来るので、ポートは2つになるだろう。段階に分けて実装する:

- テストを使って、しかし、モックデータベースの代わりに定数レートで
- それから、GUIを使って
- それから、実際のデータベースと交換できるモックデータベースとを使って

IHCのGyan Sharma、この例のコードを提供してくれてありがとう。

#### [ステージ1: FIT アプリ 定数をモックデータベースとして](https://blog.tai2.net/hexagonal_architexture.html#id26)

まずはじめに、テストケースをHTMLのtableとして作る(これについてはFITのドキュメントを見よ):

| TestDiscounter |            |
| -------------- | ---------- |
| amount         | discount() |
| 100            | 5          |
| 200            | 10         |

カラム名が、我々のプログラムでは、クラスと関数名になることに注意。FITには、プログラマ的なスタイルを排除する方法があるが、この記事では、そのまま残しておくほうが簡単だ。

テストデータどんなものになるかわかったら、ユーザー側のアダプターを作る。FITといっしょに配布されているColumnFixtureだ:

```
import fit.ColumnFixture;
public class TestDiscounter extends ColumnFixture
{
   private Discounter app = new Discounter();
   public double amount;
   public double discount()
   { return app.discount(amount); }
}
```

実際のところ、ここにあるのは、すべてアダプタのためのものだ。これまでのところ、テストはコマンドラインから実行する(必要なパスについてはFITの本を見よ)。我々はこのようにした:

```
set FIT_HOME=/FIT/FitLibraryForFit15Feb2005
java -cp %FIT_HOME%/lib/javaFit1.1b.jar;%FIT_HOME%/dist/fitLibraryForFit.jar;src;bin
fit.FileRunner test/Discounter.html TestDiscount_Output.html
```

FITは、出力ファイルを色付きで作成して、なにがパスしたのか見せてくれる(あるいは、どこかでtypoした場合には、なにが失敗したのか)。

この時点で、コードはチェックインし、Cruise Controlやあなたの自動ビルドマシンに仕込んで、ビルドおよびテストスイートに入れる準備ができている。

#### [ステージ2: UI アプリ 定数をモックデータベースとして](https://blog.tai2.net/hexagonal_architexture.html#id27)

わたしは、あなたに自身のUIを作って、それに割引アプリケーションを駆動させてもらうつもりだ。ここに入れるには少々長いコードになるからだ。コードのキーになる行は、このようなものだ:

```
...
 Discounter app = new Discounter();
public void actionPerformed(ActionEvent event)
{
    ...
   String amountStr = text1.getText();
   double amount = Double.parseDouble(amountStr);
   discount = app.discount(amount));
   text3.setText( "" + discount );
   ...
```

この時点で、アプリケーションは、デモと回帰テストができる。ユーザー側のアダプターは両方動いている。

#### [ステージ3: (FITまたはUI) アプリ モックデータベース](https://blog.tai2.net/hexagonal_architexture.html#id28)

データベース側の置き換え可能なアダプターを作るために、リポジトリへの「インターフェイス」を作る。モックデータベースや実際のサービスオブジェクトを生成する「RepositoryFactory」と、データベースのインメモリモックだ。

```
public interface RateRepository
{
   double getRate(double amount);
 }
public class RepositoryFactory
{
   public RepositoryFactory() {  super(); }
   public static RateRepository getMockRateRepository()
   {
      return new MockRateRepository();
   }
}
public class MockRateRepository implements RateRepository
{
   public double getRate(double amount)
   {
      if(amount <= 100) return 0.01;
      if(amount <= 1000) return 0.02;
      return 0.05;
    }
}
```

このアダプターを割引アプリケーションに仕込むために、使用するリポジトリアダプターを受け入れるように、アプリケーション自体を更新する必要がある。そして、(FITまたはUI)ユーザー側アダプターに、使用するリポジトリ(実またはモック)をアプリケーション自体のコンストラクタへと渡させる。これが、更新されたアプリケーションと、モックリポジトリを渡すFITアダプターだ(モックか実リポジトリのアダプターどちらを渡すのか選べるFITアダプターのコードは、長いわりに、新しい情報が増えるわけでもないので、ここではそのバージョンは省略する)。

```
import repository.RepositoryFactory;
import repository.RateRepository;
public class Discounter
{
   private RateRepository rateRepository;
   public Discounter(RateRepository r)
   {
      super();
      rateRepository = r;
    }
   public double discount(double amount)
   {
      double rate = rateRepository.getRate( amount );
      return amount * rate;
    }
}
import app.Discounter;
import fit.ColumnFixture;
public class TestDiscounter extends ColumnFixture
{
   private Discounter app =
       new Discounter(RepositoryFactory.getMockRateRepository());
   public double amount;
   public double discount()
   {
      return app.discount( amount );
   }
}
```

これで、もっとも簡単なバージョンのヘキサゴナルアーキテクチャの実装を終える。

RubyとRackをブラウザの用例に使った異なる実装としては、https://github.com/totheralistair/SmallerWebHexagon を見よ。

### [応用ノート](https://blog.tai2.net/hexagonal_architexture.html#id29)

#### [左右の非対称性](https://blog.tai2.net/hexagonal_architexture.html#id30)

ports and adaptersパターンは、意図的に、すべてのポートが基本的に類似しているふりをしながら書かれている。このようなふりをすることは、アーキテクチャレベルで有益だ。実装においては、ポートとアダプターには2種類のものがあることがわかる。すぐに明らかになる理由から、わたしが、「プライマリ」と「セカンダリ」と呼ぶものだ。これらは、「駆動する」アダプターと「駆動される」アダプターと呼ばれることもある。

懸命な読者は気付くだろうが、与えられた例ではすべて、FITフィクスチャは左側にあり、モックが右側にある。3層アーキテクチャでは、FITは、層の最上位にあり、モックは最下層にある。

これは、「プライマリアクター」と「セカンダリアクター」のユースケースから来たアイデアと関連する。「プライマリアクター」は、アプリケーションを駆動するアクターだ(アプリケーションの公開している機能のひとつを実行させるために、アクティブでない状態から起こす)。「セカンダリアクター」は、アプリケーションが駆動するもので、そこから解答を得るか、単に通知する。「プライマリ」と「セカンダリ」の違いは、だれが起動するのか、あるいは、だれが会話の責任を持つのか、ということだ。

「プライマリ」アクターを置き換えるのに自然なテスト用アダプターは、FITだ。このフレームワークは、スクリプトを読んで、アプリケーションを駆動するよう設計されたものだからだ。データベースのような「セカンダリ」アクターを置き換えるのに自然なテスト用アダプターは、モックだ。モックは、問合せに答えたり、アプリケーションからのイベントを記録するために設計されたものだからだ。

これらの観測から導かれるのは、システムのユースケース文脈図に従い、「プライマリポート」と「プライマリアダプター」を六角形の左側(ないし上側)に、「セカンダリポート」と「セカンダリアダプター」を六角形の右側(ないし下側)に描くということだ。

プライマリとセカンダリのポート・アダプター間の関係と、FITとモックでの対応する実装は、覚えておいて損はないが、それは、ports and adaptersアーキテクチャを使うことの帰結として使われるべきなのであって、一足飛びにそこにいくべきではない。ports and adapters実装の究極の便益は、アプリケーションを完全に隔離されたモードで動かすことができるということだ。

#### [ユースケースとアプリケーションの境界](https://blog.tai2.net/hexagonal_architexture.html#id31)

ヘキサゴナルアーキテクチャパターンを使って、ユースケースを書く好ましいやりかたを強めるのは、有用だ。 よくある間違いは、ユースケースを書いた結果、各ポートの外側にある技術の親密な知識が入ってしまうことだ。こういったユースケースは、正当にも、長いあいだ業界で悪名を得てきた。読み辛い、退屈、壊れやすい、そして、保守が高くつく。

port and adaptersアーキテクチャを理解すると、ユースケースは、一般にアプリケーション境界(六角形の内側)で書かれるべきということがわかる。外部のテクノロジーと無関係に、アプリケーションによってサポートされた機能やイベントを指定するためだ。これらのユースケースは、短く、読み易く、保守が安く済み、時間が経っても、より安定していられる。

#### [ポートはいくつ?](https://blog.tai2.net/hexagonal_architexture.html#id32)

なにがポートで、なにがそうでないかは、ほとんど好みの問題だ。もっとも極端なものは、すべてのユースケースが、それ自身のポートを与えられて、たくさんのアプリケーションのために数百のポートを作るというものだ。別のものとして、すべてのプライマリポートと、すべてのセカンダリポートを合わせて、左側と右側の2つのポートだけにするということも想像できる。

どちらの極端な例も最適とは思われない。

既知の用例で説明する天気システムには、4つの自然なポートがある: 天気フィード、管理者、通知を受ける購読者、購読者のデータベースだ。コーヒーメーカーのコントローラーは、4つの自然なポートを持つ: ユーザー、レシピと価格を保持するデータベース、抽出口、そして硬貨箱だ。病院の医薬システムなら3つかもしれない: 看護婦のためのもの、処方箋データベースのためのもの、そして、コンピューター制御の薬受取機のためのもの。

「間違った」ポートの数を選んだとしても、とくだんダメージがあるようには思われない、なのでこれは直感の問題として残される。わたしの選択は、2,3,4ポートの小さい数字を好む傾向がある。これは上記や、既知の用例で説明される通りだ。

#### [既知の用例](https://blog.tai2.net/hexagonal_architexture.html#id33)

![Figure 4](https://blog.tai2.net/images/hexagonal_architecture/2304.gif)

図4

図4は、4つのポートと、各ポートに複数のアダプターを持つアプリケーションを示している。これは、国立気象局からの、地震、竜巻、家事と洪水についての警報を聴取し、電話や留守番電話で人々に通知するアプリケーションに由来した。このシステムについて議論したとき、システムのインターフェイスは、「目的と結びついた技術」によって特定され、議論された。そこには、有線で届くトリガーデータのためのインターフェイスがあった。それは、留守番電話に送られる通知データのためのインターフェイス、GUIで実装された管理インターフェイス、そして、購読者データを取得するためのデータベースインターフェイスだった。

人々は奮闘していた、なぜなら、気象局からのHTTPインターフェイス、購読者へのEメールインターフェイスを追加する必要があったからだ、そして、成長するアプリケーションスイートを異なる顧客購買嗜好のために組み合わたり、分割する方法を見付けなければならなかった。かれらが目の前にある保守とテストの悪夢に恐怖したのは、別のバージョンをすべての組合わせと順列のために実装、テストそして保守しなければならなかったからだ。

かれらの設計上の変化は、システムのインターフェイスを、技術というよりは「目的」から組織し、そして、技術をアダプターによって(すべての側面において)置き換え可能にするということだった。即座に、HTTPフィードとEメール通知の能力を入れられることに気付いた(新しいアダプターは、図の中で点線とともに描かれている)。各アプリケーションをAPIを通じてヘッドレスモードで実行できるようにすることで、アプリ追加アダプターを追加して、サブアプリケーションを必要に応じて接続し、アプリケーションスイートをばらすことができた。最後に、テストとモックアダプターを適切に配置し、各アプリケーションを完全に隔離環境で実行できるようにすることで、スタンドアローンの自動化されたスクリプトで、アプリケーションを回帰テストできる能力を得た。

#### [Mac, Windows, Google, Flickr, Web 2.0](https://blog.tai2.net/hexagonal_architexture.html#id34)

1990年代初頭、ワープロアプリケーションのようなMachintoshアプリケーションは、API駆動のインターフェイスを備える必要があった。アプリケーションとユーザーの書いたスクリプトが、アプリケーションの全機能にアクセスできるようにするためだ。Windowsデスクトップアプリケーションも同じ能力を進化させてきた(どちらが先だったか言えるような歴史的知識は持ち合わせていないが、どちらだろうが、話の要点とは関係ない)。

現在(2005年)のウェブアプリケーションにおけるトレンドは、APIを公開して、他のウェブアプリケーションが直接それらのAPIにアクセスできるようにすることだ。ゆえに、地域の犯罪データをGoogleマップを通じて公開することや、Flickrの写真をアーカイブしたり注釈をつけたりする能力を持ったウェブアプリケーションを作成することが可能だ。

これらは、どれも「プライマリ」ポートのAPIを可視化することについての例だ。セカンダリポートについての情報は、ここには見られない。

#### [ストアード出力](https://blog.tai2.net/hexagonal_architexture.html#id35)

この例は、C2 wikiで、 Willem Bogaertsによって書かれた:

「わたしも似たようなことに遭遇したが、それは主に、アプリケーションレイヤーが、管理すべきでないものまで管理する一種の電話交換機になってしまう、強い傾向を持っていたからだった。アプリケーションは出力を生成し、ユーザーに表示して、その後、出力を保存する可能性もあった。主な問題は、常に保存する必要はない、ということだった。だから、アプリケーションは出力を生成し、バッファしてからユーザーに表示しなければならなかった。そして、ユーザーが出力を保存することを決めたら、アプリケーションはバッファを取り出し、それを実際に保存する。

わたしは、これがまったく好きではなかった。そして、解決法が受かんだ: ストレージ機能付きの表示制御部を持つということだ。もはや、アプリケーションは、出力を異なる方向に向けないのみならず、単に表示制御部に出力する。答えをバッファして、ユーザーに保存の機会を与えるのは、表示制御部だ。

伝統的なレイヤー構造のアーキテクチャは、『UI』と『ストレージ』を異なるものとして強調する。Port and Adapterアーキテクチャは、出力が、単に再度『出力』されるよう強制できる。」

#### [C2-wikiからの匿名の例](https://blog.tai2.net/hexagonal_architexture.html#id36)

「わたしが働いていたあるプロジェクトでは、コンポーネントステレオシステムのシステムメタファーを使っていた。各コンポーネントには、定義されたインターフェイスがあり、それぞれが特定の目的を持っていた。すると、簡単なケーブルとアダプターを使って、ほとんど制限なくコンポーネントを接続することができる」

#### [分散された、大きなチームでの開発](https://blog.tai2.net/hexagonal_architexture.html#id37)

これは、まだ試験的な用法なので、このパターンの用例として入れるのは、おそらく適切ではない。しかしながら、考えてみるのはおもしろい。

別の地域にあるチームが、全員ヘキサゴナルアーキテクチャを構築する。チームは、アプリケーションあるいはコンポーネントが、スタンドアロンモードでテストできるように、FITとモックを使う。Cruise Controlのビルドは30分ごとに走り、すべてのアプリケーションを FITとモックの組合せで走らせる。アプリケーションサブシステムとデータベースが完璧になったら、モックがテストデータベースと置き換えられる。

#### [UIとアプリケーションロジックの開発を分割する](https://blog.tai2.net/hexagonal_architexture.html#id38)

これは、まだ早期のトライアルなので、このパターンの用例として数には入れられない。しかしながら、考えてみるのはおもしろい。

UIデザインが不安定なのは、駆動する技術やメタファーをまだ決めていないからだ。バックエンドサービスアーキテクチャは、未決定で、実際、次の数ヶ月で何度か変わるかもしれない。にもかかわらず、プロジェクトは公式に開始され、時間は過ぎていく。

アプリケーションチームは、アプリケーションを隔離し、そして、テスト可能で、デモ可能な機能をユーザーに見せるために、FITテストとモックを作成する。UIとバックエンドサービスが最終的に決まるころには、それらの要素をアプリケーションに追加するのは、「容易であるべき」だ。これがどう機能するのか学びたければ、乞うご期待(もしくは、自分で試して、わたしに教えるために書くとか)。

### [関連するパターン](https://blog.tai2.net/hexagonal_architexture.html#id39)

#### [Adapter](https://blog.tai2.net/hexagonal_architexture.html#id40)

「デザインパターン」本は、一般的な「Adapter」パターンの説明を収録している: 「クラスのインターフェイスを、クライアントが期待する異なったインターフェイスに変換する」 ports and adaptersパターンは、「Adapter」パターンのひとつの用例だ。

#### [Model-View-Controller](https://blog.tai2.net/hexagonal_architexture.html#id41)

MVCパターンは、1974の早い時期にSmalltalkプロジェクトで実装された。何年にも渡り、Model-InteractorやModel-View-Presenterのような、さまざまなバリエーションが供されてきた。いずれも、ports and adaptersの、セカンダリポートではなく、プライマリポートを実装している。

#### [Mock Objects と Loopback](https://blog.tai2.net/hexagonal_architexture.html#id42)

モックオブジェクトは、他のオブジェクトの挙動をテストするための"2重のエージェント"だ。はじめに、モックオブジェクトは、インターフェイスやクラスの擬似的な実装として振舞い、ほんとうの実装の外向けの振舞いを模倣する。二番目に、モックオブジェクトは、他のオブジェクトが、そのメソッドとどのようにやりとするかを監視し、規定の、期待される実際の振舞いと比較する。齟齬が起きると、モックオブジェクトは、テストに割り込んで、状況を報告することができる。テスト中齟齬が発見されなければ、テスターから呼ばれた検証メソッドは、すべて期待と合致したことを保証する。さもなくば、失敗が報告される。 [http://MockObjects.com](http://mockobjects.com/) より。

モックオブジェクトのアジェンダに沿って完全に実装されるなら、モックオブジェクトは、外部インターフェイスのみにとどまらず、アプリケーション全体を通して利用される。モックオブジェクトムーブメントの主要な論点は、個別のクラスとオブジェクトレベルで、指定されたプロトコロルを満たせるということだ。わたしは、彼等の「モック」という語を、外部のセカンダリの役割を演じるものへの、インメモリーな代替の、最も簡単な説明として借用している。

Loopbackパターンは、外部デバイスのための内部の代替を作成する、明示的なパターンだ。

#### [Pedestals](https://blog.tai2.net/hexagonal_architexture.html#id43)

「Patterns for generating a layers architecture」の中で、Barry Rubelは、制御ソフトウェアにおいて対象な軸を作ることについてのパターンを記述した。これは、ports and adaptersに非常に似ている。「Pedestal」[4](https://blog.tai2.net/hexagonal_architexture.html#sf-hexagonal_architexture-4)パターンは、システムの各ハードウェアデバイスを表すオブジェクトの実装を必要とし、それらのオブジェクトを制御レイヤーで繋ぐ。「Pedestal"パターンは、ヘキサゴナルアーキテクチャのどちらかの側を記述するのに使えるが、アダプター間の類似性をまだ強調してはいない。また、機械制御環境のために書かれており、ITアプリケーションにこのパターンを適用するのは、それほど容易ではない。

#### [Checks](https://blog.tai2.net/hexagonal_architexture.html#id44)

Ward Cunninghamのユーザー入力エラーを検出し扱うためのパターン言語で、内側の六角形境界をまたがってエラーハンドリングするのに良い。

#### [Dependency Inversion (Dependency Injection) と SPRING](https://blog.tai2.net/hexagonal_architexture.html#id45)

Bob Martin の依存関係逆転の原則(Martin Fowlerからは、依存性注入(Dependency Injection)とも呼ばれている)は、「高レベルのモジュールは、低レベルのモジュールに依存すべきでない。ともに、抽象に依存すべきだ。抽象は、詳細に依存すべきではない。詳細が抽象に依存すべきだ」と述べている。Martin Fowlerによる「Dependency Injection」パターンは、いくらか実装を与えている。これらは、入れ替え可能な、セカンダリーアクターアダプターをいかにして作成するかを示す。コードは、この記事のサンプルコードのように、直接型付けすることができる。あるいは、設定ファイルを使って、SPRINGフレームワークに同等のコードを生成させるやりかたがある。

### [謝辞](https://blog.tai2.net/hexagonal_architexture.html#id46)

Intermountain Health CareのGyan Sharma、ここで使ったサンプルコードを提供してくれてありがとう。 書籍「Object Design」のRebecca Wirfs-Brockありがとう。この本を「デザインパターン」本の「Adapter」パターンといっしょに読むことで、六角形がなんであるのかを理解する助けになった。Ward’s wikの人々もありがとう。彼等は、何年にもわたって、パターンについてコメントを提供してくれた(とくに、 Kevin Rutherfordの http://silkandspinach.net/blog/2004/07/hexagonal_soup.html)

### [リファレンスと、関連した読み物](https://blog.tai2.net/hexagonal_architexture.html#id47)

- FIT, A Framework for Integrating Testing: Cunningham, W., online at [http://fit.c2.com](http://fit.c2.com/), and Mugridge, R. and Cunningham, W., ‘’Fit for Developing Software’’, Prentice-Hall PTR, 2005.
- The ‘’Adapter’’ pattern: in Gamma, E., Helm, R., Johnson, R., Vlissides, J., ‘’Design Patterns’’, Addison-Wesley, 1995, pp. 139-150.
- The ‘’Pedestal’’ pattern: in Rubel, B., “Patterns for Generating a Layered Architecture”, in Coplien, J., Schmidt, D., ‘’PatternLanguages of Program Design’’, Addison-Wesley, 1995, pp. 119-150.
- The ‘’Checks’’ pattern: by Cunningham, W., online at http://c2.com/ppr/checks.html
- The ‘’Dependency Inversion Principle’‘: Martin, R., in ‘’Agile Software Development Principles Patterns and Practices’’, Prentice Hall, 2003, Chapter 11: “The Dependency-Inversion Principle”, and online at http://www.objectmentor.com/resources/articles/dip.pdf
- The ‘’Dependency Injection’’ pattern: Fowler, M., online at http://www.martinfowler.com/articles/injection.html
- The ‘’Mock Object’’ pattern: Freeman, S. online at [http://MockObjects.com](http://mockobjects.com/)
- The ‘’Loopback’’ pattern: Cockburn, A., online at http://c2.com/cgi/wiki?LoopBack
- ‘’Use cases:’’ Cockburn, A., ‘’Writing Effective Use Cases’’, Addison-Wesley, 2001, and Cockburn, A., “Structuring Use Cases with Goals”, online at http://alistair.cockburn.us/crystal/articles/sucwg/structuringucswithgoals.htm

1. (訳注) GUIなしバージョン [↩](https://blog.tai2.net/hexagonal_architexture.html#sf-hexagonal_architexture-1-back)
2. (訳注) MS Wordなどで作成したHTMLのテーブルとして記述されたフィクスチャを元にテストケースを自動生成して走らせるツール。顧客のドメイン知識を活用して、早期から開発に参加してもらうことができる。http://fit.c2.com/wiki.cgi?IntroductionToFit [↩](https://blog.tai2.net/hexagonal_architexture.html#sf-hexagonal_architexture-2-back)
3. (訳注) CIツール http://cruisecontrol.sourceforge.net/ [↩](https://blog.tai2.net/hexagonal_architexture.html#sf-hexagonal_architexture-3-back)
4. (訳注) 台座、という意味 [↩](https://blog.tai2.net/hexagonal_architexture.html#sf-hexagonal_architexture-4-back)

[LICENSE](https://blog.tai2.net/pages/license.html) 

<iframe id="twitter-widget-0" scrolling="no" frameborder="0" allowtransparency="true" allowfullscreen="true" class="twitter-share-button twitter-share-button-rendered twitter-tweet-button" title="Twitter Tweet Button" src="https://platform.twitter.com/widgets/tweet_button.06c6ee58c3810956b7509218508c7b56.ja.html#dnt=false&amp;id=twitter-widget-0&amp;lang=ja&amp;original_referer=https%3A%2F%2Fblog.tai2.net%2Fhexagonal_architexture.html&amp;size=m&amp;text=%E3%83%98%E3%82%AD%E3%82%B5%E3%82%B4%E3%83%8A%E3%83%AB%E3%82%A2%E3%83%BC%E3%82%AD%E3%83%86%E3%82%AF%E3%83%81%E3%83%A3(Hexagonal%20architecture%E7%BF%BB%E8%A8%B3)&amp;time=1622898158280&amp;type=share&amp;url=https%3A%2F%2Fblog.tai2.net%2Fhexagonal_architexture.html&amp;via=__tai2__" style="position: static; visibility: visible; width: 74px; height: 20px;"></iframe>

 