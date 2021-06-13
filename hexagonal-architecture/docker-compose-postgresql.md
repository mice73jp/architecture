# [Docker PostgreSQLイメージを利用する](https://www.kimullaa.com/entry/2019/12/01/133740)

- [目的](https://www.kimullaa.com/entry/2019/12/01/133740#目的)
- [検証環境](https://www.kimullaa.com/entry/2019/12/01/133740#検証環境)
- [PostgreSQLを起動する](https://www.kimullaa.com/entry/2019/12/01/133740#PostgreSQLを起動する)
- PostgreSQLに接続する
  - [ホストからアクセスする](https://www.kimullaa.com/entry/2019/12/01/133740#ホストからアクセスする)
  - [コンテナを利用してアクセスする パターン1](https://www.kimullaa.com/entry/2019/12/01/133740#コンテナを利用してアクセスする-パターン1)
  - [コンテナを利用してアクセスする パターン2](https://www.kimullaa.com/entry/2019/12/01/133740#コンテナを利用してアクセスする-パターン2)
- [PostgreSQLコンテナの停止](https://www.kimullaa.com/entry/2019/12/01/133740#PostgreSQLコンテナの停止)
- [PostgreSQLコンテナの再開](https://www.kimullaa.com/entry/2019/12/01/133740#PostgreSQLコンテナの再開)
- [初期データを設定する](https://www.kimullaa.com/entry/2019/12/01/133740#初期データを設定する)
- PostgreSQLコンテナの削除
  - [PostgreSQLコンテナを消すとDBデータはどうなる?](https://www.kimullaa.com/entry/2019/12/01/133740#PostgreSQLコンテナを消すとDBデータはどうなる)
- [DBデータを永続化する](https://www.kimullaa.com/entry/2019/12/01/133740#DBデータを永続化する)
- ロケール設定
  - [日本語化する](https://www.kimullaa.com/entry/2019/12/01/133740#日本語化する)
- [なぜユーザ名やパスワードが起動引数で変えられるのか？](https://www.kimullaa.com/entry/2019/12/01/133740#なぜユーザ名やパスワードが起動引数で変えられるのか)
- [PostgreSQLコンテナの構成を知る](https://www.kimullaa.com/entry/2019/12/01/133740#PostgreSQLコンテナの構成を知る)
- PostgreSQLの設定の変更
  - [設定ファイルの編集](https://www.kimullaa.com/entry/2019/12/01/133740#設定ファイルの編集)
  - [編集した設定ファイルの配置](https://www.kimullaa.com/entry/2019/12/01/133740#編集した設定ファイルの配置)
- [設定をまとめる](https://www.kimullaa.com/entry/2019/12/01/133740#設定をまとめる)
- [関連記事](https://www.kimullaa.com/entry/2019/12/01/133740#関連記事)

## 目的

この記事はDockerで開発環境用のPostgreSQLを用意することを目的にしています。運用レベルの考慮はしていません。

## 検証環境

```
$ cat /etc/redhat-release
CentOS Linux release 7.1.1503 (Core)
$ docker -v
Docker version 18.03.1-ce, build 9ee9f40
$ docker-compose -v
docker-compose version 1.21.2, build a133471
```

## PostgreSQLを起動する

```
$ docker run --name my-db -p 5432:5432 -e POSTGRES_USER=dev -e POSTGRES_PASSWORD=secret -d postgres:9.6
```

|      オプション      | 役割                                                         |
| :------------------: | :----------------------------------------------------------- |
|     --name my-db     | コンテナ名                                                   |
|     -p 5432:5432     | ホストのポート:コンテナのポート                              |
|   -e POSTGRES_USER   | スーパユーザ名(省略時は"postgres")                           |
| -e POSTGRES_PASSWORD | スーパユーザのパスワード(省略時はパスワードなしでログイン可) |
|    -e POSTGRES_DB    | データベース名(省略時はPOSTGRES_USERと同じ)デフォルトだとユーザ名と同じデータベースが存在しないとエラーになるため、変える機会は少なそう |
|      -e PGDATA       | PostgreSQLのデータの格納先ディレクトリ(省略時は/var/lib/postgresql/data) |

起動しているコンテナの一覧は`docker ps`で確認できる。

```
# --name my-dbで指定した値がNAMESに表示される
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                    NAMES
ebb9cf2ff57f        postgres:9.6        "docker-entrypoint.s…"   14 seconds ago      Up 10 seconds       0.0.0.0:5432->5432/tcp   my-db
```

## PostgreSQLに接続する

### ホストからアクセスする

`-p xxxx:5432`で指定したホストのポートにアクセスすると、コンテナにパケットが転送される。PostgreSQLがコンテナで動いていることを意識する必要はない。ローカルで動くアプリケーションから接続するときは、この方法になる。

```
$ psql -h localhost -U dev
ユーザ dev のパスワード:
psql (9.6.9)
"help" でヘルプを表示します.

dev=# 
```

### コンテナを利用してアクセスする パターン1

psqlクライアント用にコンテナを立ち上げ、`--link`を利用して接続する。

```
$ docker run -it --rm --link my-db:db postgres:9.6 psql -h db -U dev
Password for user dev:
psql (9.6.9)
Type "help" for help.

dev=#
```

|    オプション     | 役割                                                         |
| :---------------: | :----------------------------------------------------------- |
|       --rm        | 実行が終わったらpsqlのコンテナを破棄する                     |
|  --link my-db:db  | コンテナの/etc/hostsに、my-dbにアクセス可能なIPを持つホスト名(db)が設定される |
| psql -h db -U dev | コンテナ内で実行するコマンド。 接続先には `--link`で設定したホスト名を指定する |

ただし `--link`は古い機能のため扱いに注意。
参考 [Legacy container links](https://docs.docker.com/engine/userguide/networking/default_network/dockerlinks/)

### コンテナを利用してアクセスする パターン2

現在立ち上がっているコンテナで別プロセスを立ち上げる。

```
$ docker exec -it my-db psql -U dev
psql (9.6.9)
Type "help" for help.

dev=#
```

## PostgreSQLコンテナの停止

`docker stop`するだけ。

```
$ docker stop my-db
```

## PostgreSQLコンテナの再開

`docker rm`してなければ、`docker start`できる。 DBデータは消えずに利用できる。

```
$ docker start my-db
```

## 初期データを設定する

コンテナの`/docker-entrypoint-initdb.d`ディレクトリに`*.sql, *.sql.gz, or *.sh`のファイルを配置すると、コンテナ起動時に実行される。アプリケーション用のスキーマを毎回実行したいときに利用すると便利。

```
$ cat sql/schema.sql
CREATE TABLE SAMPLE();
$ docker run -it --name my-db -v $(pwd)/sql:/docker-entrypoint-initdb.d  -e POSTGRES_PASSWORD=secret -d postgres:9.6
56fdad8741465fd86b618b62ed2ca7bf613b715b73e3ccf90e3618155d58d5d3
[kimura@localhost docker]$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS               NAMES
56fdad874146        postgres:9.6        "docker-entrypoint.s…"   5 seconds ago       Up 2 seconds        5432/tcp            my-db
[kimura@localhost docker]$ docker exec -it my-db psql -U postgres
psql (9.6.9)
Type "help" for help.

postgres=# \d
         List of relations
 Schema |  Name  | Type  |  Owner
--------+--------+-------+----------
 public | sample | table | postgres
(1 row)
```

## PostgreSQLコンテナの削除

`docker rm`するだけ。

```
$ docker rm my-db
```

`docker run`すると、DBデータはまっさらな状態になる。

### PostgreSQLコンテナを消すとDBデータはどうなる?

PostgreSQLのDockerfileに、以下のような記述がある。
参考 [Dockerfile](https://github.com/docker-library/postgres/blob/818958f40191c614b1e17dc9e5249ddbb7406a60/9.6/Dockerfile#L169)

```
VOLUME /var/lib/postgresql/data
```

`VOLUME`の指定があると、コンテナ起動時にボリュームが自動で作成される。 作成されたボリューム名は、`docker inspect`で確認できる。

```
$ docker inspect my-db
...
        "Mounts": [
            {
                "Type": "volume",
                "Name": "813a6def1e8a1e07e6e5cb317bd4577a83ecce8d6d2846449014ea3e32e56fd4",
                "Source": "/var/lib/docker/volumes/813a6def1e8a1e07e6e5cb317bd4577a83ecce8d6d2846449014ea3e32e56fd4/_data",
                "Destination": "/var/lib/postgresql/data",
                "Driver": "local",
                "Mode": "",
                "RW": true,
                "Propagation": ""
            }
        ],
...
```

**コンテナを削除するだけではボリュームは削除されない**。

> Removing the service does not remove any volumes created by the service. Volume removal is a separate step.

参考 [Manage data in Docker](https://docs.docker.com/storage/)

そのため、ボリュームを指定すれば昔のDBデータにアクセスできる。

```
$ docker run -it -v 813a6def1e8a1e07e6e5cb317bd4577a83ecce8d6d2846449014ea3e32e56fd4:/var/lib/postgresql/data --name my-db -e POSTGRES_PA
SSWORD=secret -d postgres:9.6
```

逆に、明示的にボリュームを消さないとデータは残り続ける。 コンテナを削除するときにボリュームも消す場合は、`-v`オプションを使う。

```
$ docker rm -v my-db
```

また、以下のコマンドで不要なボリュームを一括で削除できる。

```
$ docker volume prune
```

## DBデータを永続化する

ボリュームに名前を付けておけば、手軽にコンテナにアタッチできる。
参考 [Use volumes](https://docs.docker.com/storage/volumes/#start-a-container-with-a-volume)

```
$ docker volume create --name pgdata
pgdata

$ docker volume inspect pgdata
[
    {
        "CreatedAt": "2018-06-14T03:55:37+09:00",
        "Driver": "local",
        "Labels": {},
        "Mountpoint": "/var/lib/docker/volumes/pgdata/_data",
        "Name": "pgdata",
        "Options": {},
        "Scope": "local"
    }
]
```

起動時にData Volumeを指定してテーブルを作成する。

```
$ docker run -it --name my-db -v pgdata:/var/lib/postgresql/data -e POSTGRES_PASSWORD=secret -d postgres:9.6

$ docker exec -it my-db -U postgres
psql (9.6.9)
Type "help" for help.

postgres=#  create table book();
CREATE TABLE
```

コンテナを停止して削除する。

```
$ docker stop my-db
$ docker rm my-db
```

起動時に同じData Volumeを指定すれば、永続化したデータが参照できる。

```
$ docker run -it --name my-db -v pgdata:/var/lib/postgresql/data -e POSTGRES_PASSWORD=secret -d postgres:9.6
$ docker exec -it my-db -U postgres
psql (9.6.9)
Type "help" for help.

postgres=# \d
        List of relations
 Schema | Name | Type  |  Owner
--------+------+-------+----------
 public | book | table | postgres
(1 row)
```

ただしボリュームにデータを永続化した場合、`/docker-entrypoint-initdb.d`のスクリプトはコンテナ起動ごとに実行されない点に注意。

## ロケール設定

デフォルトではサーバロケールがen_US.utf8になっている。 ja_JP.utf8にしたい人は以下のように、Dockerfileを作成する必要がある。

> You can also extend the image with a simple Dockerfile to set a different locale. The following example will set the default locale to de_DE.utf8:
>
> FROM postgres:9.4 RUN localedef -i de_DE -c -f UTF-8 -A /usr/share/locale/locale.alias de_DE.UTF-8 ENV LANG de_DE.utf8
>
> Since database initialization only happens on container startup, this allows us to set the language before it is created.

参考 [OFFICIAL REPOSITORY](https://hub.docker.com/_/postgres/)

### 日本語化する

ロケールを追加し、LANG環境変数に設定する。

```
$ cat Dockerfile
FROM postgres:9.6
RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
ENV LANG ja_JP.utf8
```

Dockerfileをビルドし、起動すると日本語になる。

```
$ docker build -t dev-postgres -f Dockerfile .
$ docker run -it --name my-db -e POSTGRES_PASSWORD=secret -d dev-postgres 

$ docker exec -it my-db -U postgres
psql (9.6.9)
Type "help" for help.

postgres=# \l
                                 List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges
-----------+----------+----------+------------+------------+-----------------------
 postgres  | postgres | UTF8     | ja_JP.utf8 | ja_JP.utf8 |
 template0 | postgres | UTF8     | ja_JP.utf8 | ja_JP.utf8 | =c/postgres          +
           |          |          |            |            | postgres=CTc/postgres
```

## なぜユーザ名やパスワードが起動引数で変えられるのか？

-eオプションでコンテナの環境変数を設定できる仕組みがDockerにある。
参考 [Docker run リファレンス](http://docs.docker.jp/engine/reference/run.html#env)

以下のようなコマンドを実行すると `-e SAMPLE=sample`で指定した値が環境変数に設定されることがわかる。

```
$ docker run -it -e SAPMLE=sample --rm centos:7 env
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=c650fe32fc3b
TERM=xterm
SAPMLE=sample
HOME=/root
```

また、既に設定された環境変数を上書きすることもできる。 例えば、`HOME=/root`を`HOME=/tmp`で上書きしてみる。

```
$ docker run -it -e HOME=/tmp --rm  centos:7 env
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=4c5e7303b0fa
TERM=xterm
HOME=/tmp
```

上記の仕組みを利用し、起動スクリプトの動きを動的に変えている様子。

起動スクリプトの詳細を確認するときは`docker inspect`を利用する。

```
$ docker inspect postgres:9.6
...
  "WorkingDir": "",
  "Entrypoint": [  // 起動時に呼ばれるコマンド
    "docker-entrypoint.sh"
  ],
...
```

起動スクリプト(docker-entrypoint.sh)は以下のような処理になっている。
参考 [docker-entrypoint.sh](https://github.com/docker-library/postgres/blob/eff90effc6b5578be90bef93d96b3fceb1082a7c/9.6/docker-entrypoint.sh)

```
   file_env 'POSTGRES_USER' 'postgres'
   file_env 'POSTGRES_DB' "$POSTGRES_USER"

   psql=( psql -v ON_ERROR_STOP=1 )

   if [ "$POSTGRES_DB" != 'postgres' ]; then
           "${psql[@]}" --username postgres <<-EOSQL
                   CREATE DATABASE "$POSTGRES_DB" ;
           EOSQL
           echo
   fi

   if [ "$POSTGRES_USER" = 'postgres' ]; then
           op='ALTER'
```

## PostgreSQLコンテナの構成を知る

起動しているコンテナにbashで入ることができる。

スクリプトやディレクトリ構成を確認できる。

```
$ docker exec -it my-db /bin/bash
root@f43f50193fa6:/# ls
bin   dev                         docker-entrypoint.sh  home  lib64  mnt  proc  run   srv  tmp  var
boot  docker-entrypoint-initdb.d  etc                   lib   media  opt  root  sbin  sys  usr
```

## PostgreSQLの設定の変更

### 設定ファイルの編集

まず、コンテナからpostgresql.confをコピーする。

```
$ docker cp my-db:/var/lib/postgresql/data/postgresql.conf .
```

次に、ホスト側で設定ファイルを編集する。 （今回はmax_connectionsを100から5にする）

```
$ cat postgresql.conf
...
max_connections = 5
...
```

### 編集した設定ファイルの配置

コンテナ側にファイルを配置するだけだと、以下のようにエラーになる。
参考 [Unable to replace postgresql.conf](https://github.com/docker-library/postgres/issues/105)

```
$ docker run -it --name my-db -v $(pwd)/postgresql.conf:/var/lib/postgresql/data/postgresql.conf -e POSTGRES_PASSWORD=secret postgres:9.6
The files belonging to this database system will be owned by user "postgres".
This user must also own the server process.

The database cluster will be initialized with locale "en_US.utf8".
The default database encoding has accordingly been set to "UTF8".
The default text search configuration will be set to "english".

Data page checksums are disabled.

initdb: directory "/var/lib/postgresql/data" exists but is not empty
If you want to create a new database system, either remove or empty
the directory "/var/lib/postgresql/data" or run initdb
with an argument other than "/var/lib/postgresql/data".
```

以下のように、`$PGDATA`に設定ファイルを配置しないでPostgreSQLの起動時に`config_file`で設定ファイルのパスを指定すればよい。

```
$ docker run -it --name my-db -v $(pwd)/postgresql.conf:/etc/postgresql.conf -e POSTGRES_PASSWORD=secret -d postgres:9.6 -c config_file=/etc/postgresql.conf
bf55391afe2ae709e84efd81025b4d7f542f3db4bb42d13c348807cc9342b47f
$ $ docker exec -it my-db psql -U postgres
psql (9.6.9)
Type "help" for help.

postgres=# SHOW max_connections;
 max_connections
-----------------
 5
(1 row)
```

## 設定をまとめる

docker-composeを使ってコマンドをまとめる。

```
$ ll
合計 8
-rw-rw-r--. 1 kimura kimura 119  6月 14 07:02 Dockerfile
drwxrwxr-x. 2 kimura kimura  28  6月 14 06:58 conf
-rw-rw-r--. 1 kimura kimura 438  6月 14 13:11 docker-compose.yml
drwxrwxr-x. 2 kimura kimura  23  6月 14 13:41 sql
FROM postgres:9.6
RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
ENV LANG ja_JP.utf8
version: "3"
volumes:
   pgdata:
     driver: "local"
services:
  my-db:
    build: .
    container_name: "my-db"
    ports:
      - "5432:5432"
    environment:
      POSTGRES_USER: "dev"
      POSTGRES_PASSWORD: "secret"
    command: 'postgres -c config_file="/etc/postgresql.conf"'
    volumes:
      - "pgdata:/var/lib/postgresql/data"
      - "./conf/postgresql.conf:/etc/postgresql.conf"
      - "./sql:/docker-entrypoint-initdb.d"
```

これでDockerfileのビルドから起動までコマンド一発。

```
$ docker-compose up -d
Creating network "postgres_default" with the default driver
Creating volume "postgres_pgdata" with local driver
Building my-db
Step 1/3 : FROM postgres:9.6
 ---> d92dad241eff
Step 2/3 : RUN localedef -i ja_JP -c -f UTF-8 -A /usr/share/locale/locale.alias ja_JP.UTF-8
 ---> Running in 4ae5282dc582
Removing intermediate container 4ae5282dc582
 ---> ae2a5b8d31fe
Step 3/3 : ENV LANG ja_JP.utf8
 ---> Running in aa9c01cdc114
Removing intermediate container aa9c01cdc114
 ---> 117a670283ca
Successfully built 117a670283ca
Successfully tagged postgres_my-db:latest
WARNING: Image for service my-db was built because it did not already exist. To rebuild this image you must use `docker-compose build` or `docker-compose up --build`.
Creating my-db ... done
```

PostgreSQLに接続したい場合は以下のようにする。

```
$ docker-compose exec my-db psql -U dev
psql (9.6.9)
"help" でヘルプを表示します.

dev=# \d
          リレーションの一覧
 スキーマ |  名前  |    型    | 所有者
----------+--------+----------+--------
 public   | sample | テーブル | dev
(1 行)

dev=# show max_connections;
 max_connections
-----------------
 5
(1 行)
```

`docker exec`でもアクセスできる。

```
$ docker exec -it my-db psql -U dev
psql (9.6.9)
"help" でヘルプを表示します.

dev=# show max_connections ;
 max_connections
-----------------
 5
(1 行)
```

コンテナとデータのどちらも削除したい場合は、`-v`オプションを付ける。

```
$ docker-compose down -v
Removing network postgres_default
Removing volume postgres_pgdata
```

参考 [docker-compose down](https://docs.docker.com/compose/reference/down/)

## 関連記事

[www.kimullaa.com](https://www.kimullaa.com/entry/2020/03/14/134232)