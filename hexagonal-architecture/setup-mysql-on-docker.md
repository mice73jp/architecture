# [Docker] Mac に docker-compose で MySQL 環境を構築する

![img](https://blog.hiros-dot.net/wp-content/uploads/2019/09/Docker.jpg)Docker

[Twitter](https://twitter.com/intent/tweet?text=[Docker]+Mac+に+docker-compose+で+MySQL+環境を構築する&url=https%3A%2F%2Fblog.hiros-dot.net%2F%3Fp%3D10469)[Facebook](https://www.facebook.com/sharer/sharer.php?u=https%3A%2F%2Fblog.hiros-dot.net%2F%3Fp%3D10469&t=[Docker]+Mac+に+docker-compose+で+MySQL+環境を構築する)[はてブ](https://b.hatena.ne.jp/entry/s/blog.hiros-dot.net/?p=10469)[Pocket](https://getpocket.com/edit?url=https://blog.hiros-dot.net/?p=10469)[LINE](https://timeline.line.me/social-plugin/share?url=https%3A%2F%2Fblog.hiros-dot.net%2F%3Fp%3D10469)[コピー](https://blog.hiros-dot.net/?p=10469)

 2021.02.21

現在 Django を勉強しているのですが、データベースに MySQL を利用することとしました。

そこで、docker-compose を使用して MySQL 環境を構築しましたので、その手順を記します。

目次

1. [環境](https://blog.hiros-dot.net/?p=10469#toc1)
2. 事前準備
   1. [ディレクトリの作成](https://blog.hiros-dot.net/?p=10469#toc3)
   2. [MySQL の設定ファイル作成](https://blog.hiros-dot.net/?p=10469#toc4)
   3. [docker-compose ファイルの作成](https://blog.hiros-dot.net/?p=10469#toc5)
3. [作成したディレクトリ構成の確認](https://blog.hiros-dot.net/?p=10469#toc6)
4. [MySQL データベースへの接続](https://blog.hiros-dot.net/?p=10469#toc7)
5. 動作確認
   1. [MySQL の起動確認](https://blog.hiros-dot.net/?p=10469#toc9)
   2. [MySQL への接続](https://blog.hiros-dot.net/?p=10469#toc10)
   3. [MySQL 接続時に command not found: mysql が発生した場合の対応](https://blog.hiros-dot.net/?p=10469#toc11)
   4. [環境変数に MySQL クライアントを追加する](https://blog.hiros-dot.net/?p=10469#toc12)
6. [データベースの作成](https://blog.hiros-dot.net/?p=10469#toc13)
7. [ テーブルの作成](https://blog.hiros-dot.net/?p=10469#toc14)
8. [データの取得](https://blog.hiros-dot.net/?p=10469#toc15)
9. [MySQL から切断する](https://blog.hiros-dot.net/?p=10469#toc16)
10. [まとめ](https://blog.hiros-dot.net/?p=10469#toc17)

## 環境

macOS Big Sur: 11.2.1
Docker: 20.10.2, build 2291f61
Docker-Compose: 1.27.4, build 40524192

## 事前準備

### ディレクトリの作成

docker-compose を使用して MySQL 環境構築する前に、あらかじめディレクトリーを作成します。

以下のディレクトリー構成にします。

```
docker 
  |---mysql
        |---data
        |---sql
```

data ディレクトリは、MySQL のデータ永続化用です。

また、sql ディレクトリは、MySQL 起動時の初期化スクリプト置き場です。

それでは、ターミナルで docker ディレクトリに移動し、以下のようにコマンドを入力してディレクトリを作成します。

```
% mkdir mysql
% mkdir mysql/data
% mkdir mysql/sql
```

### MySQL の設定ファイル作成

続いて、MySQL の設定ファイル my.cnf を作成します。my.cnf ファイルは、先ほど作成した mysql ディレクトリー直下に作成します。

```
$ vi mysql/my.cnf
```

my.cnf は以下の通りです。

```
[mysqld]
character-set-server=utf8mb4
collation-server=utf8mb4_unicode_ci

[client]
default-character-set=utf8mb4                                                 
```

### docker-compose ファイルの作成

最後に docker-compose ファイルを作成します。

docker-compose.yml ファイルは、docker ディレクトリと同じ階層に作成します。

```
% vi ../docker-compose.yml
```

docker-compose.yml は以下のように編集します。

```
version: '3'

services:
  mysqldb:
    image: mysql:5.7
    container_name: mysql_container
    environment:
      MYSQL_ROOT_PASSWORD: root
      MYSQL_DATABASE: my_testdb
      MYSQL_USER: docker
      MYSQL_PASSWORD: docker
      TZ: 'Asia/Tokyo'
    volumes:
      - ./docker/db/data:/var/lib/mysql
      - ./docker/db/my.cnf:/etc/mysql/conf.d/my.cnf
      - ./docker/db/sql:/docker-entrypoint-initdb.d
    ports:
      - "3306:3306"
    command: mysqld --character-set-server=utf8mb4 --collation-server=utf8mb4_unicode_ci
```

## 作成したディレクトリ構成の確認

作成したディレクトリ構成は以下の通りです。

```
docker 
  |---mysql
        |---data
        |---my.cnf
        |---sql
docker-compose.yml
```

## MySQL データベースへの接続

準備ができたら、以下の docker-compose コマンドを実行して MySQL データベースを作成・起動します。

```
% docker-compose up -d
```

## 動作確認

### MySQL の起動確認

以下のように、「docker ps -a」コマンドを実行して、起動していることを確認します。

```
% docker ps -a
CONTAINER ID   IMAGE       COMMAND                  CREATED          STATUS                    PORTS                               NAMES
debde6edc937   mysql:5.7   "docker-entrypoint.s…"   19 seconds ago   Up 18 seconds             0.0.0.0:3306->3306/tcp, 33060/tcp   mysql_container
```

### MySQL への接続

次に、ターミナルから以下のコマンドを実行して MySQL へ接続します。

接続に成功すると、パスワードを聞いてきますので、root ユーザーのパスワードを入力します。

```
% mysql -h 127.0.0.1 -P 3306 -u root -p
Enter password:
```

### MySQL 接続時に command not found: mysql が発生した場合の対応

もし MySQL に接続するときに、以下のようにエラーが発生した場合は、MySQL Client をインストールしていないことが原因と思われます（筆者はこのエラーが発生しました）。

```
zsh: command not found: mysql
```

この場合は、以下のように brew コマンドで インストール可能な MySQL クライアントを探します。

```
% brew search mysql
==> Formulae
automysqlbackup            mysql-client@5.7 ✔         mysql@5.6
mysql                      mysql-connector-c++        mysql@5.7
mysql++                    mysql-sandbox              mysqltuner
mysql-client               mysql-search-replace
```

「mysql@5.7」という文字列があれば、インストール可能です。

以下のコマンドを実行して、MySQL クライアントをインストールします。

```
% brew install mysql-client@5.7
```

### 環境変数に MySQL クライアントを追加する

MySQLクライアントをインストールした場合は、環境変数に追加しておきましょう。

zsh の場合は vi で ~/.zshrc を開きます。

```
vi ~/.zshrc
```

~/.zshrc を開いたら、以下のように EXPORT を追加します。

```
export PATH="$PATH:/usr/local/opt/mysql-client@5.7/bin"
```

bash の場合は vi で ~/.bashrc を開きます。

```
vi ~/.bashrc
```

~/.bashrc を開いたら、以下のように EXPORT を追加します。

```
export PATH="$PATH:/usr/local/opt/mysql-client@5.7/bin"
```

インストールが完了したら、再度接続をしてみましょう。

## データベースの作成

接続ができたら、以下のコマンドを実行して、test_db というデータベースを作成します。

```
mysql> CREATE DATABASE test_db;
Query OK, 1 row affected (0.01 sec)
```

データベースが作成できたら、以下のコマンドを実行してデータベースを切り替えます。

```
mysql> USE test_db;
Database changed
```

##  テーブルの作成

続いて、以下のコマンドを実行して test_tbl を作成します。

```
mysql> CREATE TABLE test_tbl (
    -> id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
    -> name TEXT NOT NULL
    -> );
Query OK, 0 rows affected (0.05 sec)
```

次に、以下のコマンドで、作成したテーブルにデータを追加します。

```
INSERT INTO test_tbl (name) VALUES ("HIRO"),("Steve"),("michel");
```

## データの取得

最後に、以下のコマンドを実行して、テーブルにデータが追加されているかを確認します。

```
mysql> select * from test_tbl;
+----+--------+
| id | name   |
+----+--------+
|  1 | HIRO   |
|  2 | Steve  |
|  3 | michel |
+----+--------+
3 rows in set (0.00 sec)
```

## MySQL から切断する

最後に以下のコマンドを実行して、MySQL から切断しましょう。

```
mysql> \q
Bye
```

## まとめ

思ったより簡単に、Docker を使用して MySQL 環境を構築することができました。

MySQL に接続できない時は「あれ？」となりましたが、単なる MySQL クライアントの入れ忘れだったのでほっとしました。

これで自分専用の MySQL 環境ができたので、プログラミング言語と合わせて勉強を進められそうです。