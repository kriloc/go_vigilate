# Go - Vigilate

A dead simple monitoring service, intended to replace things like Nagios.

## Requirements

Vigilate requires:
- Postgres 11 or later (db is set up as a repository, so other databases are possible)
- An account with [Pusher](https://pusher.com/), or a Pusher alternative
  (like [ipê](https://github.com/dimiro1/ipe))

## Run

First, make sure ipê is running (if you're using ipê):

On Mac/Linux
~~~
cd ipe
./ipe 
~~~


Run with flags:

~~~
./vigilate \
-dbuser='tcs' \
-pusherHost='localhost' \
-pusherPort='4001' \
-pusherKey='123abc' \
-pusherSecret='abc123' \
-pusherApp="1" \
-pusherSecure=false
~~~~

## notes
https://github.com/tsawler/vigilate

`cp database.yml.example database.yml`

使用 soda 來migrate:

在執行soda命令根目錄下有database.yml或者config/database.yml時，並且資料庫伺服器正在執行，Soda可以在database.yml使用一個簡單的命令在檔案中建立所有資料庫，如下： -e development 代表建立開發環境下的資料庫，當然也可以是 test 和 production。

https://learnku.com/docs/buffalo-doc-cn/buffalo-doc-cn-migrations/5092

執行 `./soda migrate`

會讀取models/models.go

`GO111MODULE=on go build -o vigilate cmd/web/*.go`
`GO111MODULE=on ./run.sh`

login default:
admin@example.com / passowrd

https://github.com/go-chi/chi

chi is a lightweight, idiomatic and composable router for building Go HTTP services.

https://github.com/dimiro1/ipe

An open source Pusher server implementation compatible with Pusher client libraries written in Go.

`./ipe`

## Test

### SQL for Test

* The Dashboard shows status count

```sql
SELECT
(select count(id) from host_services where active =1 and status='pending') as pending,
(select count(id) from host_services where active =1 and status='healthy') as healthy,
(select count(id) from host_services where active =1 and status='warning') as warning,
(select count(id) from host_services where active =1 and status='problem') as problem
```

## libs
[scs](https://github.com/alexedwards/scs) : HTTP Session Management for Go

[Pusher Channels HTTP Go Library](https://github.com/pusher/pusher-http-go)
：The Golang library for interacting with the Pusher Channels HTTP API.