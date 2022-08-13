# Simple golang server running inside a docker container.

Running a simple golang server using gorilla mux inside docker container. 

## Steps:
    1. Using gorilla mux library to create a router and a handler and passing it to the http.Server

    2. Same basics steps to run the container. 
        a. docker build -t go-docker .
        b. docker run -p 3000:3000 go-docker.
        c. docker start <container_id> -> to start the container
        d. docker stop <container_id> -> to stop the container.
        e. docker remove <container_id> -> to remove the container. 
        (You know the drill.)
    
    3. Goto localhost:3000/ping or curl localhost:3000/ping

    4. Have used multi-stage build wherein the builder stage builds the image and creates a binary executable of your program.
    And deploy stage to run the bin exe. 

    5. Multi-stage builds are used to reduce the deployment image. With the help of this the file image size can be reduced to as much  as 12MBs. 


## Different api endpoints:

    1. GET: localhost:3000/albums -> return the list of albums you have currently.
    2. GET: localhost:3000/albums/{id} -> return a specific album.
    3. POST: localhost:3000/albums -> add a new album.
    4. PATCH: localhost:3000/albums/{id} -> update an album given an id.
    5. DELETE: localhost:3000/albums/{id} -> delete an album given an id. 
