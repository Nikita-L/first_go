# Test task  

3 microservices: `parser`, `centre`, `database`.  

## Parser  
Parser gets *data.tar* file from local file system, untars it, parses it line by line (just to not use much RAM) and sends clean data to centre, using grpc through `api.proto`.    

## Centre  
Centre is grpc server implementing `api.proto` and writing data to database. Connection to database starts in each goroutine on requests.  

## Database
Postgres