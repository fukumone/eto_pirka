# EtoPirka

### Create Database

```
$ mysql -u root
>> create database eto_pirka_development;
```

### Create .env.development file

FB_CLIENT_IDとFB_SECRET_KEYは個人で設定

```
$ echo 'BasicAuthUSER=root\nBasicAuthPASSWORD=password\nFB_CLIENT_ID=xxx\nFB_SECRET_KEY=xxx\nFB_HOST=http://localhost:3000/auth/callback/facebook\nDB_USER_NAME=root\nDATABASE_NAME=eto_pirka_development' > .env.development
```

### migrate

```
$ goose up
```

### seed

```
$ go run seed/seed.go
```

### Server Run

```
$ go run main.go
```
