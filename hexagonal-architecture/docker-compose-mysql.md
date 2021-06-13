## [Dockerとdocker-composeでMySQLを動かした](https://uzimihsr.github.io/post/2020-11-27-mysql-on-docker/)

Nov 27, 2020

## 環境構築が面倒

DBを使ったアプリケーションを開発するときに毎回MySQLやテーブルを準備するのが嫌だったので,
コンテナを使ってMySQLが動作する環境を簡単に作成できるようにした.

------



## やったことのまとめ

- `MySQL`のコンテナを`Docker`で起動した
- `docker-compose`を使ってコンテナの起動時に初期設定用`SQL`が自動で実行されるようにした

```
docker-compose.yml
version: "3.8"
services:
  mysql:
    image: mysql:5.7.32
    container_name: mysql-container
    environment:
      MYSQL_ROOT_PASSWORD_FILE: /password-root              # パスワードが記述されたファイルのパスを指定
    volumes:
      - ./password-root:/password-root                      # パスワードが記述されたファイル
      - ./member.sql:/docker-entrypoint-initdb.d/member.sql # MySQLの起動時に実行したい初期設定用SQLファイル
    ports:
      - 3306:3306
# コンテナを直接起動する例
$ docker container run -d -p 3306:3306 --name mysql-container -e MYSQL_ROOT_PASSWORD=<rootユーザのパスワードに設定したい文字列> mysql:5.7.32

# docker-composeで起動する例
# 上記docker-compose.ymlの他にpassword-root(rootパスワードを記述したファイル), member.sql初期設定用SQLファイルを用意して実行する
$ docker-compose up -d

# ホストOSからDBを操作するにはコンテナに入るか, 別のコンテナを建てるか, ホストOSのクライアントを使う
$ docker container exec -it mysql-container mysql -u root -p
$ docker container run --rm -it mysql:5.7.32 mysql -h host.docker.internal -P 3306 -u root -p
$ mysql -h 127.0.0.1 -P 3306 -u root -p
```

## つかうもの

- macOS Mojave 10.14

- Docker Desktop for Mac

  - Version 2.5.0.0
  - Kubernetes: v1.19.3
  - Docker version 19.03.13
  - docker-compose version 1.27.4

- MySQL

   

  (Docker image)

  - Server version: 5.7.32

## やったこと

- [Dockerで起動する](https://uzimihsr.github.io/post/2020-11-27-mysql-on-docker/#dockerで起動する)
- [docker-composeで起動する](https://uzimihsr.github.io/post/2020-11-27-mysql-on-docker/#docker-composeで起動する)

### Dockerで起動する

`MySQL`の公式`Docker image`[1](https://uzimihsr.github.io/post/2020-11-27-mysql-on-docker/#fn:1)が用意されているので, これを使う.

```bash
# localhostからアクセスできるよう3306ポートを割り当てて起動
# 環境変数(MYSQL_ROOT_PASSWORD)でrootユーザのパスワード(hogehoge)を設定する
$ docker container run -d -p 3306:3306 --name mysql-container -e MYSQL_ROOT_PASSWORD=hogehoge mysql:5.7.32

# コンテナに入ってMySQLクライアントを起動
$ docker container exec -it mysql-container mysql -u root -p
Enter password: # hogehoge
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 3
Server version: 5.7.32 MySQL Community Server (GPL)

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| information_schema |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
4 rows in set (0.01 sec)

mysql> exit
Bye

# ホストOSのMySQLクライアントからも接続できる
$ mysql -h 127.0.0.1 -P 3306 -u root -p
Enter password: # hogehoge
...
mysql> exit
Bye

# 使い捨てのコンテナからホストOSのlocalhost:3306で起動しているMySQLのコンテナに接続
$ docker container run --rm -it mysql:5.7.32 mysql -h host.docker.internal -P 3306 -u root -p
Enter password: # hogehoge
...
mysql> exit
Bye

# クライアント用のコンテナを使用するパターン
# 事前にdocker network(bridge)内のmysql-containerのIPを確認する
$ docker container inspect mysql-container -f "{{.NetworkSettings.IPAddress}}"
172.17.0.2
$ docker container run --rm -it mysql:5.7.32 mysql -h 172.17.0.2 -P 3306 -u root -p
Enter password: # hogehoge
...
mysql> exit
Bye

# コンテナを削除
$ docker container rm -f mysql-container
```

流石に`Docker`なら簡単に動かせた.

### docker-composeで起動する

もうちょっと複雑なことがしたいので, `docker-compose`で起動してみる.

```yaml
version: "3.8"
services:
  mysql:
    image: mysql:5.7.32
    container_name: mysql-container
    environment:
      MYSQL_ROOT_PASSWORD_FILE: /password-root              # パスワードが記述されたファイルのパスを指定
    volumes:
      - ./password-root:/password-root                      # パスワードが記述されたファイル
      - ./member.sql:/docker-entrypoint-initdb.d/member.sql # MySQLの起動時に実行したい初期設定用SQLファイル
    ports:
      - 3306:3306
DROP DATABASE IF EXISTS mydb;
CREATE DATABASE IF NOT EXISTS mydb;
USE mydb;


DROP TABLE IF EXISTS members;
CREATE TABLE IF NOT EXISTS members (
  id INT PRIMARY KEY AUTO_INCREMENT,
  name VARCHAR(10),
  team VARCHAR(10),
  role VARCHAR(10),
  age INT
);


INSERT IGNORE INTO members (name, team, role, age) VALUES ('Alice', 'A', 'manager', 30);
INSERT IGNORE INTO members (name, team, role, age) VALUES ('Ben', 'B', 'manager', 50);
INSERT IGNORE INTO members (name, team, role, age) VALUES ('Charlie', 'A', 'member', 40);
INSERT IGNORE INTO members (name, team, role, age) VALUES ('Daniel', 'A', 'member', 30);
INSERT IGNORE INTO members (name, team, role, age) VALUES ('Emily', 'A', 'member', 20);
INSERT IGNORE INTO members (name, team, role, age) VALUES ('Florence', 'A', 'member', 30);
INSERT IGNORE INTO members (name, team, role, age) VALUES ('George', 'A', 'trainee', 20);
INSERT IGNORE INTO members (name, team, role, age) VALUES ('Harry', 'B', 'member', 40);
INSERT IGNORE INTO members (name, team, role, age) VALUES ('Isabel', 'B', 'member', 40);
INSERT IGNORE INTO members (name, team, role, age) VALUES ('Jack', 'B', 'trainee', 20);
INSERT IGNORE INTO members (name, team, role, age) VALUES ('Katie', 'B', 'trainee', 20);
# パスワードファイルの作成
$ vim password-root # パスワード文字列(hogehoge)を入力する
$ cat password-root
hogehoge

# パスワードファイル(password-root), 初期設定用SQLファイル(member.sql), docker-compose.ymlが存在する状態
$ ls
docker-compose.yml member.sql         password-root

# コンテナの起動
$ docker-compose up -d
Starting mysql-container ... done

# データベースの確認
# mysql -h 127.0.0.1 -P 3306 -u root -p でも同じことができる(ホストOSのMySQLクライアントを使う場合)
$ docker-compose exec mysql mysql -u root -p
Enter password: # hogehoge
...
mysql> use mydb
Database changed

mysql> SELECT * FROM members;
+----+----------+------+---------+------+
| id | name     | team | role    | age  |
+----+----------+------+---------+------+
|  1 | Alice    | A    | manager |   30 |
|  2 | Ben      | B    | manager |   50 |
|  3 | Charlie  | A    | member  |   40 |
|  4 | Daniel   | A    | member  |   30 |
|  5 | Emily    | A    | member  |   20 |
|  6 | Florence | A    | member  |   30 |
|  7 | George   | A    | trainee |   20 |
|  8 | Harry    | B    | member  |   40 |
|  9 | Isabel   | B    | member  |   40 |
| 10 | Jack     | B    | trainee |   20 |
| 11 | Katie    | B    | trainee |   20 |
+----+----------+------+---------+------+
11 rows in set (0.00 sec)
# 初期設定用SQLが実行されているので最初からデータベースとテーブルが作成されている
```

ポイントは2つ.

1つめはrootユーザのパスワードをファイルから読み込む形にしていること.

先程使用した環境変数の`MYSQL_ROOT_PASSWORD`ではなく`MYSQL_ROOT_PASSWORD_FILE`でファイルのパスを指定すればその中身をパスワードとして設定できるので,
パスワード文字列(**hogehoge**)が記述されたファイル(`password-root`)をホストOSからコンテナにマウントしてそのパスを`MYSQL_ROOT_PASSWORD_FILE`で指定するようにしている.

こうすると`docker-compose.yml`からパスワードの情報を分離できて,
そのままリポジトリとかで共有できるようになる.

2つめはデータベースの初期設定用の`SQL`がコンテナの起動時に勝手に実行されるようにしていること.

コンテナのディレクトリ`/docker-entrypoint-initdb.d`の配下にある拡張子が`.sh`, `.sql`, `.sql.gz`のスクリプトは勝手に実行されるので,
ここに実行したい`SQL`ファイル(`member.sql`)をマウントしている.

こうするとわざわざ起動後のコンテナに入って初期設定用の`SQL`を実行する必要がなくなり,
`docker-compose`を立ち上げるだけで使いたいデータベースとテーブルが準備できるので便利.

## おわり

`Docker`と`docker-compose`を使って`MySQL`を起動する手順を試してみた.

特に`docker-compose`を使う方法だと設定ファイルさえ用意すればコマンド1つでレコード入りのテーブルが簡単に用意できるのでめっちゃ便利だと思う.