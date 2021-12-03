# GoCrudBook

<p align="center">
<img src="https://github.com/3n0ugh/GoCrudBook/blob/main/gopher-lib.jpeg" alt="drawing" width="900" height="570" align="center"/>
</p>

CRUD API example is written in go language using [net/http](https://pkg.go.dev/net/http) 
package and [MySQL](https://www.mysql.com/) database.

## Requirements
- [Go](go.dev)
- [MySQL](https://www.mysql.com/)
- Code Editor
    
## Project Structure
```bash
GoCrudBook  
├── LICENSE.md   
├── cmd 
│   └── web
│       ├── config
│       │   └── config.go
│       ├── database
│       │   └── connection.go
│       ├── handler
│       │   └── handler.go
│       ├── main.go
│       └── router
│           └── router.go
├── go.mod
├── go.sum
├── gopher-lib.jpeg
├── pkg
│   └── models
│       ├── book.go
│       └── mysql
│           ├── book.sql
│           └── mysql.go
└── to-do.md
```

## Instructions

- Clone the repository and move to the project directory:
  ```bash
    git clone https://github.com/3n0ugh/GoCrudBook.git
    cd GoCrudBook
  ```
- To load the database
  - To create a database into MySQL:
    ```sql
      CREATE DATABASE library;
    ```
  - Then, create a table and insert sample values into the database:
    ```bash
      mysql -u root library < pkg/models/book.sql
    ```
 - Tidy the modules
    ```bash
      go mod tidy
    ```
 - Run the main.go
    ```bash
      go run ./cmd/web/main.go
    ```

## Usage Examples
-  Get all book
  ```bash
    curl -X GET localhost:5000/book
  ```
-  Get book by id
  ```bash
    curl -X GET "localhost:5000/book/id?id=1933988673"
  ```
-  Get book by name (enough to contain the word)
  ```bash
    curl -X GET "localhost:5000/book/name?name=android"
  ```
-  Add book
  ```bash
    curl -X POST localhost:5000/book/create -d
      '{ "isbn" : 1932394161,
          "book_name" : "Time To Go Forward",
          "author": "3n0ugh",
          "page_count" : 14,
          "book_count": 1,
          "borrow_times": 3,
          "barrow_date": null,
          "last_recieved_dat": null
         }'
  ```
-  Delete Book (by id)
  ```bash
    curl -X DELETE "localhost:5000/book/delete?id=1932394161"
  ```
-  Update Book
  ```bash
    curl -X PUT localhost:5000/book/update -d 
       '{ "isbn" : 1932394161,
          "book_name" : "Time To Go Forward",
          "author": "3n0ugh",
        	"page_count" : 14,
          "book_count": 1,
          "borrow_times": 3,
          "barrow_date": null,
          "last_recieved_dat": null
         }'
  ```
