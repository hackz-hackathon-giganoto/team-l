# team-l

## client
- http://localhost:3000
#### ドメイン

- 書籍一覧ページ : /home 
- 書籍詳細ページ : /books/{book-id}
- サインインページ : /sign-in
- サインアップページ : /sign-up
- プロフィール(本棚)ページ : /{user-name}


## api-server
- http://localhost:8080

## DB(MySQL)

```mermaid
erDiagram

books_article{
  article_id string
  user_id string
  isbn string
  date int
  article string
  lend bool
}

user{
  user_id string
  name string
  introduction string
  Twitter string
  github string
  image string
  auth_uid string
}

tag{
  tag_id string
  tag_name string
}

article_tag{
  tag_map_id string
  tag_id string
  article_id string
}

```
