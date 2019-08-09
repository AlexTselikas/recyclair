# Recyclair.eu.org

Recyclair.eu.org is an online interactive map that displays the garbage and recycling bins that are near you.All bin data has been provided by the community

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Prerequisites
- Pure Go Postgres driver for database/sql https://github.com/lib/pq
- Go net/http configurable handler to handle CORS requests https://github.com/rs/cors
- PostgreSQL https://www.postgresql.org/
- Any recent Go version  https://golang.org/


### Installing
- Clone the project to a location on your computer.
- Open the folder with your favorite code editor.
- Create a database with
- - 1 Table called ```bin_data```
- - 5 Columns like the picture below

<img  src="https://i.imgur.com/AEQQZOp.png">

## Deployment

Compile and run the ```server.go``` 
```
go build server.go
```

## Built With
* [Pure Go Postgres driver for database/sql](https://github.com/lib/pq) - PostgreSQL driver for Go
* [Go net/http configurable handler to handle CORS requests](https://github.com/rs/cors) - CORS reqests library
* [PostgreSQL](https://www.postgresql.org/) - PostgreSQL database

## Contributing
You are free to either contribute on adding bin data from your area or maintaining the code for the front and back end, or both! ;)
 
## Authors

* **Alexandros Tselikas** - *Initial work* - [AlexTselikas](https://github.com/AlexTselikas)

See also the list of [contributors](https://github.com/recyclerapp/recyclair/contributors) who participated in this project.

## Acknowledgments

* Big thanks for GRNET S.A for free server space for hosting the backend.
* Also big thanks to GitHub for the front end hosting
* Thanks to OpenStreetMap.org and their contributors
