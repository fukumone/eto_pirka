# EtoPirka
・Go言語で簡潔なSNSアプリケーション

### Install

```
go get github.com/fukumone/eto_pirka
```

### Create Database

```
$ mysql -u root
>> create database eto_pirka_development;
```

### Create .env.development file

FB_CLIENT_IDとFB_SECRET_KEYは個人で設定

```
$ echo 'BasicAuthUSER=root\nBasicAuthPASSWORD=password\nFB_CLIENT_ID=xxx\nFB_SECRET_KEY=xxx\nFB_HOST=http://localhost:3000/auth/callback/facebook\nDATABASE_USER_PASSWORD=root\nDATABASE_NAME=eto_pirka_development\nDATABASE_USER_PASSWORD=password' > .env.development
```

### Migrate

```
$ goose up
```

### Server Run

```
$ go build -o ./eto_pirka && ./eto_pirka
```
