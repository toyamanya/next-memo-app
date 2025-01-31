# GolangとNext.jsとDockerを使ってメモアプリを作る

## コマンド
コンテナ起動
docker-compose up 
docker-compose up --build

キャッシュのクリア後にコンテナ再起動
docker-compose down --volumes --remove-orphans
docker-compose build --no-cache
docker-compose up

フロントエンドコンテナに入る
docker exec -it memo-app-frontend-1 sh


バックエンド
http://localhost:8080

フロントエンド
http://localhost:3000


## リンク
Github
https://github.com/toyamanya/next-memo-app

chatGPT
https://chatgpt.com/c/6764353d-4ed8-8013-ad48-3d974a34c751

hackmd
https://hackmd.io/20U-s4ymRk2CtNEcKfTp7w


## 参考文献
docker上のnodejs環境でのnode_moduleの扱いについて
https://gist.github.com/ychikazawa/d4d5548de08287b3fe1c9ae9453c3d7f
https://zenn.dev/yumemi_inc/articles/3d327557af3554




