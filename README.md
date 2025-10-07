# golang-atlas-practice

## go mod 初期化
```
go mod init my-app
```

```
$ go mod init my-app
go: creating new go.mod: module my-app
go: to add module requirements and sums:
	go mod tidy
```

## Atlas のセットアップ
```
curl -sSf https://atlasgo.sh | sh
chmod +x atlas
mv atlas /usr/local/bin/
```

### セットアップ時に　homebrew　を使う場合
```
brew install ariga/tap/atlas
```

### Atlas のバージョン確認
```
atlas version
```

```
$ atlas version
atlas version v0.35.1-46f5e79-canary
https://github.com/ariga/atlas/releases/latest
```

## DB migration 実行手順

### migration ファイルの作成
```
atlas migrate new add_users_table
```
ファイルの中身は後追いで作成しても良い

### golang から Atlas を使用して migration を実行する
golang では os/exec を用いて `atlas migrate apply` を呼び出すのが一般的

### MySQL の起動
```
docker compose up -d
```

### マイグレーションファイル作成済みの状態で golang を実行
```
go run main.go
```

成功した場合、以下のようなメッセージが表示される
```
❯ go run main.go
Running DB migrations with Atls...
Migrating to version 20250920120000 (1 migrations in total):

  -- migrating version 20250920120000
    -> CREATE TABLE `users` (
         `id` BIGINT AUTO_INCREMENT PRIMARY KEY,
         `name` VARCHAR(255) NOT NULL,
         `email` VARCHAR(255) NOT NULL UNIQUE,
         `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
         `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
       ) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
  -- ok (7.444583ms)

  -------------------------
  -- 16.361167ms
  -- 1 migration
  -- 1 sql statement
Migration applied successfully!
```

### MySQL への反映を手動で実行したいとき
```
atlas migrate apply --env dev
```

## マイグレーションを追加するとき
`atlas migrate new <name>` で作成したファイルを編集した場合は、`atlas migrate hash` を再実行して更新しなければならない
また、checksum ファイル(atlas.sum) は、チーム開発で整合性を保つ必要があるため、バージョン管理(Git) にコミットしておくこと

### checksum ファイル生成
```
atlas migrate new add_users_table
atlas migrate hash
```

実行すると atlas.sum というファイルが migrations/ 配下に作成される<br>
h1:abcdef1234567890...

### 再度 migration を実行する　
```
go run main.go
```

## migration 結果確認

### docker 内にある MySQL を確認する
docker コンテナが起動していることを確認する
```
docker ps
```

```
docker exec -it mysql8 bash
```

```
mysql -u appuser -p appdb
※ パスワードを入力する
```
