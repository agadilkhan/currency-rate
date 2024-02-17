# Currency-Rate
### Service for getting the currency rate
***
## Built With
* [![Golang][Golang-badge]][Golang-url]
* [![Gin][Gin-badge]][Gin-url]
* [![PostgreSQL][PostgreSQL-badge]][PostgreSQL-url]


[Golang-badge]: https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Golang-url]: https://golang.org/
[Gin-badge]: https://img.shields.io/badge/Gin-00ADD8?style=for-the-badge&logo=go&logoColor=white
[Gin-url]: https://gin-gonic.com/
[PostgreSQL-badge]: https://img.shields.io/badge/PostgreSQL-336791?style=for-the-badge&logo=postgresql&logoColor=white
[PostgreSQL-url]: https://www.postgresql.org/

## Getting Started
### Installation
1. Clone the repository
```sh
git clone https://github.com/agadilkhan/currency-rate.git
```
2. Install dependencies
```sh
go mod tidy
   ```
3. Create config.yaml file in config folder, copy config.yaml.example and fill missing values

### Migrate
To create the currency table
```sh
migrate -path migrations/ -database "postgresql://username:password@localhost:5432/database_name?sslmode=disable" -verbose up
```

### Launch
```
go run cmd/main.go
```

### Endpoints
- ```http://localhost:8080/api/currency``` - list of all currencies
- ```http://localhost:8080/api/currency/:code``` - get currency by code
- ```http://localhost:8080/api/currency/force-update``` - force update

## Contact

[Alfarabi Agadilkhan](https://t.me/agadilkhan) (aagadilkhan@gmail.com)