# Docs


## Database

Database layer is sits on top of [upperdb](https://github.com/upper/db).

| Vendor       | Support   | Notes                                         
|--------------|-----------|-----------------------------------------------
| Postgres     | ✔️         |  using upper.                                 
| SQLite       | ✔️         |  using upper.                                 
| MySQL        | ❌         |  suppored by upper, should be "trival" to add.
| MSSQL        | ❌         |  suppored by upper, should be "trival" to add.

## Executors

| Vendor       |type    | Notes                                         |
|--------------|--------|-----------------------
| javascript   |`lang`  |  using [goja](https://github.com/dop251/goja)
| webassembly  |`lang`  |  using [wazero](github.com/tetratelabs/wazero)
| pagedash     |`loader`|  dashbaord with js hooks  
| pageform     |`loader`|  form (wizard) with js hooks


- [FAQ](./faq.md)
- [Architecture](./arch.md)
- [Glossary](./glossary.md)