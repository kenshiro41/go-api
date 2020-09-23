## URL
https://www.ken41.com


## 使用技術
言語:　Go
フレームワーク・ライブラリ: GraphQL, gqlgen, Gorm
DevTool: Docker

## 開発時
Dockerが必要です。

### docker-composeを起動&シェルに入る

```
make
```

### マイグレーション実行(初回起動時)

```
make migration
```

### GraphQL自動生成(gqlgen)

```
make gen
```

