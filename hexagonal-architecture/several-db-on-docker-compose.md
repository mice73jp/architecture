# もはや DB は Docker でインストールする時代！初心者のための DB インストール on Docker

[MySQL](https://qiita.com/tags/mysql)[MongoDB](https://qiita.com/tags/mongodb)[SQLServer](https://qiita.com/tags/sqlserver)[PostgreSQL](https://qiita.com/tags/postgresql)[Docker](https://qiita.com/tags/docker)

# この記事について

**データベース**、これは 2020 年になっても、エンタープライズなどのアプリケーション開発において、必ずと言っていいほど使用されるものだと思います。
ただ、従来同様、本番環境や検証環境、開発環境に至るまで、未だに人力で頑張って`ホスト OS に直接インストール`している方が多いというのが現状です。

(少し誇張が過ぎるかもですが)
**開発環境のデータベースのインストールに 5 分以上時間をかけるのは時間の無駄です！**

今、このタイミングで、Docker でサクッと DB をインストールできることを覚えましょう。
ただし、本番環境など、冗長化構成や可用性を担保する必要がある場合などは、きちんと要求に合致したインストール方法を実施をしましょう。

# Oracle Database on Docker

※ちょっとインストール方法が面倒なので、後日追記します。早く情報を知りたい方は、[参考情報](https://qiita.com/ymasaoka/items/ca2bb2cb19ebeafe1ccc#参考情報)に記載の内容をご参照ください。

# SQL Server on Docker

- [Docker Hub - Microsoft SQL Server](https://hub.docker.com/_/microsoft-mssql-server) (Ubuntu ベース)

GitHub にて公開しているこちらの[リポジトリ](https://github.com/ymasaoka/docker-mssql)を参考に、docker-compose.yaml を作成してください。
※下記の yaml は RHEL ベースの SQL Server (on Linux) on Docker です。
※GitHub のリポジトリでは、ブランチで RHEL ベースと Ubuntu ベースの 2 つを用意しています。(2017 は Ubuntu のみ)

docker-compose.yaml

```
version: '3'

services:
  mssql:
    image: mcr.microsoft.com/mssql/rhel/server:2019-latest
    container_name: 'mssql2019-latest-rhel'
    environment:
      - MSSQL_SA_PASSWORD=<your_strong_password>
      - ACCEPT_EULA=Y
      # - MSSQL_PID=<your_product_id> # default: Developer
      # - MSSQL_PID=Express
      # - MSSQL_PID=Standard
      # - MSSQL_PID=Enterprise
      # - MSSQL_PID=EnterpriseCore
    ports:
      - 1433:1433
    # volumes: # Mounting a volume does not work on Docker for Mac
    #   - ./mssql/log:/var/opt/mssql/log
    #   - ./mssql/data:/var/opt/mssql/data
```

`docker-compose up -d`で、コンテナーを起動すれば、SQL Server のインストールは完了です。

```
docker-compose up -d
```

# MySQL on Docker

- [Docker Hub - mysql](https://hub.docker.com/_/mysql)
- [Docker Hub - adminer](https://hub.docker.com/_/adminer)

GitHub にて公開しているこちらの[リポジトリ](https://github.com/ymasaoka/docker-mysql)を参考に、docker-compose.yaml を作成してください。

docker-compose.yaml

```
version: '3'

services:
  db:
    image: mysql:8
    container_name: mysql
    restart: always
    environment:
      MYSQL_ROOT_PASSWORD: P@ssw0rd #required
      # MYSQL_DATABASE: employees #optional
      # MYSQL_USER: user #optional
      # MYSQL_PASSWORD: P@ssw0rd #optional
      # MYSQL_ALLOW_EMPTY_PASSWORD: "yes" #optional
      # MYSQL_RANDOM_ROOT_PASSWORD: "yes" #optional
      # MYSQL_ONETIME_PASSWORD: "yes" #optional (MySQL 5.6 or above)
      # MYSQL_INITDB_SKIP_TZINFO: "" #optional
    ports:
        - 3306:3306
    volumes:
        - ./data/mysql:/var/lib/mysql
        - ./conf:/etc/mysql/conf.d
```

`docker-compose up -d`で、コンテナーを起動すれば、MySQL のインストールは完了です。

```
docker-compose up -d
```

# PostgreSQL on Docker

- [Docker Hub - postgres](https://hub.docker.com/_/postgres)
- [Docker Hub - adminer](https://hub.docker.com/_/adminer)

GitHub にて公開しているこちらの[リポジトリ](https://github.com/ymasaoka/docker-postgres)を参考に、docker-compose.yaml を作成してください。

docker-compose.yaml

```
version: '3'

services:
  db:
    image: postgres:13
    container_name: postgres
    restart: always
    environment:
      POSTGRES_PASSWORD: P@ssw0rd #required
      # POSTGRES_USER: postgres #optional
      # POSTGRES_DB: postgres #optional
      # POSTGRES_INITDB_ARGS: "--data-checksums" #optional
      # POSTGRES_INITDB_WALDIR: "" #optional (PostgreSQL 10+ or above)
      # POSTGRES_INITDB_XLOGDIR: "" #optional (PostgreSQL 9.x only)
      # POSTGRES_HOST_AUTH_METHOD: trust #optional
      # PGDATA: /var/lib/postgresql/data/pgdata #optional
    ports:
        - 5432:5432
    volumes:
      - ./data:/var/lib/postgresql/data
```

`docker-compose up -d`で、コンテナーを起動すれば、PostgreSQL のインストールは完了です。

```
docker-compose up -d
```

# MongoDB on Docker

- [Docker Hub - mongo](https://hub.docker.com/_/mongo/)

GitHub にて公開しているこちらの[リポジトリ](https://github.com/ymasaoka/docker-mongodb)を参考に、docker-compose.yaml を作成してください。

docker-compose.yaml

```
version: '3'

services:
  mongo:
    image: mongo:latest
    container_name: mongodb
    restart: always
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: P@ssw0rd
    ports:
        - 27017:27017
    volumes:
        - ./data/db:/data/db
        - ./data/configdb:/data/configdb
    # Command 1: Customize configuration without configuration file
    # Command 2: Setting WiredTiger cache size limits
    # command: >
    #   --serviceExecutor adaptive
    #   --wiredTigerCacheSizeGB 1.5

  mongo-express:
    image: mongo-express:latest
    container_name: mongo-express
    restart: always
    ports:
      - 8081:8081
    environment:
      ME_CONFIG_MONGODB_ADMINUSERNAME: root
      ME_CONFIG_MONGODB_ADMINPASSWORD: P@ssw0rd
```

`docker-compose up -d`で、コンテナーを起動すれば、MongoDB のインストールは完了です。

```
docker-compose up -d
```

# 管理ツール

管理ツールについては、各自好きなものを使用すれば良いと思いますが、おすすめは **Visual Studio Code** を使う方法です。
Docker Compose も DB 操作も Visual Studio Code 内で行えるため、とても便利です。
詳細は以下の記事を参照してください。

- [Visual Studio Code 上で SQL Database を操作するための便利な拡張機能](https://qiita.com/ymasaoka/items/b31049b284597b91742b)
- [Visual Studio Code 上で MySQL を操作するための便利な拡張機能](https://qiita.com/ymasaoka/items/aa03323bbac7e7c5f1be)
- [Visual Studio Code 上で PostgreSQL と Cosmos DB を操作するための便利な拡張機能](https://qiita.com/ymasaoka/items/48d4a9be65856c41dafd)
- [Visual Studio Code 上で MongoDB を操作するための便利な拡張機能](https://qiita.com/ymasaoka/items/690f9a49c2e5dcbd9e5a)

また、MySQL と PostgreSQL のところで記載している adminer のコンテナーを一緒に作るのもアリかと思います。

# 参考情報

## Oracle Database

- [GitHub - oracle/docker-images](https://github.com/oracle/docker-images)
- [公式 Oracle Database の Docker イメージを構築](https://b.chiroito.dev/entry/2016/12/28/235627)
- [Oracle Database 19c available on GitHub](https://blogs.oracle.com/database/oracle-database-19c-available-on-github-v2)
- [[Oracle Database\] 公式Docker Imageを利用してOracle Database 19c環境を構築してみた](https://itedge.stars.ne.jp/docker_image_oracle_database_19c/)

## SQL Server

- [Docker Hub](https://hub.docker.com/_/microsoft-mssql-server)
- [クイック スタート:Docker を使用して SQL Server コンテナー イメージを実行する](https://docs.microsoft.com/ja-jp/sql/linux/quickstart-install-connect-docker?view=sql-server-ver15&pivots=cs1-bash)
- [Windows上でSQL Serverを使用してC#アプリを作成する](https://qiita.com/ymasaoka/items/944e8a5f1987cc9e0d37)

## MySQL

- [Docker Hub](https://hub.docker.com/_/mysql)

## PostgresSQL

- [Docker Hub](https://hub.docker.com/_/postgres)

## MongoDB

- [Docker Hub](https://hub.docker.com/_/mongo/)