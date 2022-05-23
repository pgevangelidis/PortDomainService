# PortDomainService service

The PortDomainService is a service that accepts a Port struct using gRPC and stores the port data to database. 

## Memory store as database 
The database is a memory implementation using the `sync.Map` type from the `sync` library. 
The service's datalayer benefits from the abstraction, using interfaces, giving the option in the future to replace the database with another data store provider like PostgreSQL or NoSQL like store.

## gRPC as communication protocol
The service uses gRPC as communication protocol. The proto file is very simple and only satisfies the exercise requirement, which is insert or update a record to the database. Obviously, this is not the appropriate way to store things. Ideally, we would like to split the logic into separate functions Insert and Update. 

## docker to contain the build
The pds service has its own `Dockerfile` that creates a docker image. 
The service also has a config package that returns to required variables, the `HOST` and the `PORT`. 
The service has a Makefile to streamline some of the commands to build the service. The Makefile is basically naked for the moment, and requires further work to run appropriately.

## Client for PortDomainService 
The client of our service has been developed independently, with its own go module and Dockerfile. 

## Client communication
The client uses http protocol to allow communication between the user and the client, having only one endpoint that read the file. 
The communication between the client and ther server is with gRPC. 

## Json streaming
The system requirements on memory, prevents us from loading big json files in memory. Instead, we use the standard library `encoding/json`  and the `json.Decoder` type to stream json objects from big files. 

## docker for the client
Similarly to the server, we use a another Dockerfile to build the client and potentially deploy it on a docker network using a script and `docker-compose` command.

## Future development
- Think of alternative ways to read big json files. Maybe FTP or something else. 
- Expand REST API.
- Expand gRPC protocol.
- Add a docker-compose.yml file
- Improve Makefiles