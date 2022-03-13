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
  article_id char
  user_id char
  isbn char
  date datetime
  article text
  lend char
}

user{
  user_id char
  name char
  introduction char
  Twitter char
  github char
  image char
  auth_uid char
}

tag{
  tag_id char
  tag_name char
}

article_tag{
  tag_map_id char
  tag_id char
  article_id char
}

```
