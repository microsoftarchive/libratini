### How to use
```
config := libratini.Config{
  Collate: 100,
  Prefix:  "my-awesome-service",
  Source:  "this-machine's hostname",
  Token:   "librato access token",
  User:    "librato user email"
}
handler := martini.Classic()
handler.Use(libratini.Middleware(config))
```