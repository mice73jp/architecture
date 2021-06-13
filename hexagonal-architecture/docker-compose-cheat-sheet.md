# Docker-ComposeでDockerのベース環境を簡単構築！[Sample:MySQL]

[Docker-ComposeでDockerのベース環境を簡単構築！[Sample:MySQL]](https://qiita.com/Shunya078/items/17f42db80b787a9d1c59)

# 目的

docker-composeによって、簡単にDocker環境構築をしてみよう、と言う記事です。

自分が環境を先日作った際に、あまりyml文内の意味やオプション、コマンドって結局どういう意味やねん、、、とプチパニックになったのでまとめ直したものになります。そのため知識、かつ環境の構築がこの記事のみで完結できるような構成を目指しました。

あまりMarkDown記法になれていないこともあるため、少し拙い文章になるかと思いますが、読んでいただけると幸いです。

# サクッと作るぞ！チートシート

この記事をみてサクッと環境作りたいわ、って人のために
まず実行するためだけの手順を完全に説明を省いて紹介していきます。
今回はサンプルのためにMySQL環境を作成しますが、
柔軟に変更できるようにコメントにて、可変できる部分は特記しています。

## 構成

構成は以下の通りに作成しました。

```
├── docker
│   └── db
│       ├── data
│       ├── mysql-config.cnf
│       └── init
│           ├── 001-create-tables.sql
│           └── init-database.sh
└── docker-compose.yml
```

## ソースコード

### 1. docker-compose.yml

アプリケーションを構成するサービスを`docker-compose.yml`に書きます。

今回はMySQLにあたりますね。

これを使用するコードディレクトリの一番親のディレクトリに生成します。

```
# docker-compose.yml
version: '3'
services:
  # (例1)ここに使用するサービスを書きます
  mysql:
    # image > コンテナ実行時の元になるイメージです。使うサービスのリポジトリ名だったりを書きます。
    image: mysql:latest
    # volume > 仮想環境上でのファイルをパスを指定してアクセスできるようにします。パスはこのymlファイルがあるディレクトリ基準です
    volumes:
      - ./mysql/data:/var/lib/mysql
      - ./mysql/mysql-config.cnf:/etc/mysql/conf.d/my.cnf
      - ./mysql/init:/docker-entrypoint-initdb.d
        # environment > 環境変数を指定します。Compose実行時に指定されるものにあたります。
    environment:
      MYSQL_ROOT_PASSWORD: password
      MYSQL_DATABASE: database
      MYSQL_USER: user
      MYSQL_PASSWORD: password
      TZ: 'Asia/Tokyo'

    # (例2)goを使う場合であれば以下のようになります。細かい詳細は解説にて説明、参照します。
    go:
        build:
          context: ./go/
          target: dev
        volumes:
          - ./go/src:/go/src:cached
        working_dir: /go/src
        ports:
          - 8080:8080
        tty: true
        env_file:
          - ./go/.env
```

MySQL以外の環境を構築する場合は3. へ飛んでください。

## 2. MySQL環境用のソースコード

MySQLの場合に初期DBを用意する必要があるので、設定ファイル(`mysql-config.cnf`)と初期テーブル(`001-create-tables.sql`)のコードを書き、それを実行するコード(`init-database.sh`)を書きます。

```
### mysql-config.cnf ###
[mysqld]
character-set-server=utf8mb4
collation-server=utf8mb4_unicode_ci

[client]
default-character-set=utf8mb4
--- 001-create-tables.sql ---
---- drop ----
DROP TABLE IF EXISTS `first_table`;

---- create ----
create table IF not exists `first_table`
(
 `id`               INT(20) AUTO_INCREMENT,
 `name`             VARCHAR(20) NOT NULL,
 `created_at`       Datetime DEFAULT NULL,
 `updated_at`       Datetime DEFAULT NULL,
    PRIMARY KEY (`id`)
) DEFAULT CHARSET=utf8 COLLATE=utf8_bin;
### init-database.sh ###
#!/usr/bin/env bash
#wait for the MySQL Server to come up
#sleep 90s

#run the setup script to create the DB and the schema in the DB
mysql -u docker -pdocker test_database < "/docker-entrypoint-initdb.d/001-create-tables.sql"
```

## 3. Docker起動

imageを指定していない場合(例2)、ビルドを行い、イメージを構築します。この時に参照されるのがよく巷で聞く`Dockerfile`です。今回はMySQL環境を爆速で仕上げるため省略しますが、解説にて説明します。

今回は使用するイメージ(image)を

```
# docker-compose.yml
version: '3'
services:
  # (例1)ここに使用するサービスを書きます
  mysql:
    # image > コンテナ実行時の元になるイメージです。使うサービスのリポジトリ名だったりを書きます。
    image: mysql:latest
```

ここで指定しているので、コンテナを作成し起動するところからで大丈夫です。

```
$ docker-compose up
Creating network "backend_default" with the default driver
Creating backend_mysql_1 ... done
Creating backend_go_1    ... done
...
```

上記コマンドで起動できたら、同じ階層にいる別ターミナルで以下のコマンドによって起動を確認してみてください。

```
$ docker-compose ps   
     Name                   Command             State           Ports         
------------------------------------------------------------------------------
backend_go_1      bash                          Up      0.0.0.0:8080->8080/tcp
backend_mysql_1   docker-entrypoint.sh mysqld   Up      3306/tcp, 33060/tcp
```

ここで起動し、サービス名が表示されていれば構築完了です！早いですね。

最後に以下コマンドによってコンテナとそれに関わるネットワークを停止します。

```
$ docker-compose down
Stopping backend_mysql_1 ... done
Stopping backend_go_1    ... done
Removing backend_mysql_1 ... done
Removing backend_go_1    ... done
Removing network backend_default
```

下記コマンドを叩いたターミナルが

```
$ docker-compose up
Creating network "backend_default" with the default driver
Creating backend_mysql_1 ... done
Creating backend_go_1    ... done
...
mysql_1  | 2020-11-30T14:42:11.612252Z 0 [System] [MY-013172] [Server] Received SHUTDOWN from user <via user signal>. Shutting down mysqld (Version: 8.0.22).
go_1     | root@f1d2304a041d:/go/src# exit
backend_go_1 exited with code 0
mysql_1  | 2020-11-30T14:42:12.324000Z 0 [System] [MY-010910] [Server] /usr/sbin/mysqld: Shutdown complete (mysqld 8.0.22)  MySQL Community Server - GPL.
backend_mysql_1 exited with code 0
```

上記のようにexitできていれば、無事に停止しています。

これでDocker-composeによって仮想環境が構築できました。

## 4. MySQLでsample-tableを生成

仮想環境に入り、init-database.shを叩きます。

コンテナ内で起動中にコマンドを叩く場合は以下を参考にしてください。

```
# コンテナ起動
$ docker-compose up
...

> 別ターミナルでdocker環境に入り、コマンド実行
# 権限付与 > ここでyml中に書いたvolumesで、パスを指定できる、と言うわけです
$ docker-compose exec mysql bash -c "chmod 0775 docker-entrypoint-initdb.d/init-database.sh"
# init-database.shを叩く
$ docker-compose exec mysql bash -c "./docker-entrypoint-initdb.d/init-database.sh"
```

# 解説

最初にアジェンダ書いてた時は、ここにがっつり解説書いていこうと思ったのですが、思ったより上で説明してしまった感が否めないです。書くことなくなってきた。。。まずいぞ。。。

## そもそもdocker-composeって？？

複数のコンテナを、簡易に管理し、実行するDockerアプリケーションのためのツールの1つです。
コマンドを1つ実行するだけで、docker-compose.ymlに定義した設定に基づいて環境を構築してくれるので、本当にいろんな使い方ができます！！、、(らしいですと言うのが本心。お恥ずかしい。)

## ymlで(よく)使用するオプション

### image:

コンテナを実行する時に元となるイメージを指定します。
自分で`Dockerfile`を書いて実行する場合は、次のbuildオプションを使用します。
imageで使用する場合はビルドを行わず、`$ docker-compose up` からで大丈夫です。

### build:

コンテナを実行する時に参照する`Dockerfile`を指定します。外部imageではなく、自分で`Dockerfile`を制作して使う場合はよく以下の構成で使用します。

自分でいちから`Dockerfile`を書く、よりはリポジトリをクローンしてきて修正する、という扱い方が多いと思います。

```
├── docker
├── Dockerfile
└── docker-compose.yml
# docker-compose.yml
version: '3'
services:
  hoge:
    # 同階層のため相対パスで指定
    build: .
```

ここで自分で`Dockerfile`を書いた場合はイメージを構築、すなわちビルドする必要があるので、初期起動前に以下コマンドを叩きます。

```
$ docker-compose build
db uses an image, skipping
Building go ......
```

使用するcontextの指定や、dockerfile生成に関してはリファレンスも参照することをお勧めします。
https://docs.docker.jp/compose/compose-file.html#context

### environment:

主に仮想環境上で用いる環境変数を指定します。
配列、もしくはDictionaryで定義できるので、`$ docker-compose exec`で実行する際にアクセスできるようにAPI鍵や今回のようにMySQLのホストの値を記載します。

### env_file:

上記のenvironmentでは隠したい情報(自分はAPI鍵などはenv_fileを使用して`.gitignore`で隠す、という形をとることが多いです)の場合は、これで環境変数を記載したファイルを指定します。

```
# docker-compose.yml
version: '3'
services:
    go:
        env_file:
          - ./go/.env
```

なお`.env`ファイルは以下のように、`変数 = 値`の形を取ります。

```
# /go/.env
DB_USER=user
DB_PASSWORD=password
DB_HOST=mysql:3306
DB_DATABASE=database
DB_SOCKET=tcp
```

### volumes:

自分のローカルファイルを環境上のファイルに割り当てるように指定することができます。

なお、実行手順にもコメントで書いた通り、ymlファイルがあるディレクトリを基準として相対パスは使用できます。

```
# docker-compose.yml
version: '3'
services:
  mysql:
    volumes:
      # ローカルファイル：仮想環境上のパス となります
      - ./mysql/data:/var/lib/mysql
      - ./mysql/mysql-config.cnf:/etc/mysql/conf.d/my.cnf
      - ./mysql/init:/docker-entrypoint-initdb.d
```

### ports:

ホスト側とコンテナ側、両者のポートを指定することができます。
(なお、コンテナのみの指定も可能)

```
# docker-compose.yml
version: '3'
services:
    go:
        ports:
        # ローカル側：仮想環境側 という感じです
          - 8080:8080
        env_file:
          - ./go/.env
```

## docker-composeで(よく)使うコマンド

説明するコマンドを以下に書いておきます。このほうがササッと使いやすいですよね。

```
$ docker-compose build
$ docker-compose up
$ docker-compose down
$ docker-compose ps
$ docker-compose start
$ docker-compose stop
```

### $ docker-compose build

`image:`を指定せずに`build:`を使用する際に使います。これによりサービスをビルドします。
上記にも書いた通り、`image:`を指定しない場合は、以下の`$ docker-compose up`の前にビルドする必要があります。

```
頻度の高いオプション
--no-cache 構築時にイメージのキャッシュを使わない
```

### $ docker-compose up

仮想環境の立ち上げを行います。
実際はコンテナを構築、作成し、起動、アタッチまでを全てこのコマンドで行ってくれます。

なお、`-d` をつけることによって、バックグラウンドで実行されるため、コンテナは起動し続ける状態を確保できます。CIなどを回す際にはこのオプションを使うと便利です。

```
頻度の高いオプション
-d バックグラウンドでコンテナを実行
```

### $ docker-compose down

起動している仮想環境を停止し、`$ docker-compose up` によって立ち上げたコンテナとネットワークを削除します。

```
頻度の高いオプション
--rmi all 全イメージを削除
```

### $ docker-compose ps

起動しているコンテナの一覧を表示できます。

### $ docker-compose exec

起動している環境に対して、任意のコマンドを実行することができます。
なので、基本的に環境上でプロンプトを動かしたりする際はこれを使用します。

```
# exec後の部分でどのコンテナを叩くか指定します
$ docker-compose exec mysql bash -c "./docker-entrypoint-initdb.d/init-database.sh"

# 以下コマンドのように扱えば、直接シェルを叩くこともできますが、volumesにてデフォルトの階層を指定する必要があります
$ docker-compose exec hogehoge /bin/bash
頻度の高いオプション
-u 指定されたユーザによりコマンドを実行
```

### $ docker-compose start

既存のコンテナをサービスとして起動します。
`$ docker-compose up`すれば起動するので、`$ docker-compose down`ではなく(停止後コンテナを削除するため)、`$ docker-compose stop` などで停止した場合に再び起動する時に使用します。

### $ docker-compose stop

稼働中のコンテナを停止しますが、`$ docker-compose down` とは違い、削除しません。 上記の`$ docker-compose start` コマンドで、再起動できます。

# まとめ

docker-composeによる仮想環境構築でした。
運用だと、`docker-compose.yml`に複数コンテナを組み合わせて一挙に立てて使用することになることがおおいですね。

初日ということでしたが、自分が実装していた時にメモしていたことなどを噛み砕いて書いていくうちにどんどんボリュームが増しちゃいました。(笑)

ただ、語彙的にもとっつかみやすい文章になってるんじゃないでしょうか、、と思っています、、、(そうでなければ僕の語彙力の問題ですね、反省します)

皆様のますますのご活躍の少しでもお手伝いになれば、ということを書き締めたいと思います。

また質問等あればコメントにてご指摘などよろしくお願いします。

# 参考資料

- docker-compose オプション / リファレンス

https://docs.docker.jp/compose/compose-file.html

- docker-compose コマンド / リファレンス

https://docs.docker.jp/compose/reference/index.html