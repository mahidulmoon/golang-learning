# UPSKILL APIs 

## API endpoints 

```
    /api/login
    /api/register
	/api/newsletter
	/api/workshops
    /api/workshops/{id}

 //restricted endpoints will require Authorization token. 
	/api/users/{id}
	/api/workshops
	/api/workshops/{id}  METHOD POST  -> will create or join workshops
	/api/workshops/{id} METHOD DELETE -> will DELETE/Cancel workshop

```
## Test

```
$ go test ./... -v 
```

## Install 

To run the server first create a .env file in the root with the following variable.

```
DB_USER=
DB_NAME= 
DB_HOST= 
DB_PORT=
DB_PASSWORD=
TEST_DB_NAME= 
APP_SECRET_KEY=
TOKEN_EXPIRE_TIME=
```

Then run this

```
$ go build .
$./upskill 
```
