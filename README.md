# GoLang Reatful API

## Run Command
```bash
# start
docker-compose up
# restart
docker-compose up --build
# stop
docker-compose stop
# remove
docker-compose rm
# has no memory
docker system prune
```

## Restful Api

```bash
http://localhost:8000/api/book/:id            // PATCH     # update books
http://localhost:8000/api/book/:id            // GET       # get book
http://localhost:8000/api/book                // GET       # get books
http://localhost:8000/api/book/:id            // DELETE    # delete books
http://localhost:8000/api/book                // POST      # create books
```
