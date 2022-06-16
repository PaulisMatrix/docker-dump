# Simple flask application using docker
Running my simple python flask application inside a docker container

### Please run below commands to run the application.
* Preq : Ofc you need to have docker installed for Mac/Windows/Linux. Please travel to [this link](https://docs.docker.com/get-docker/) for installation.

1. Build the image whose specifications you have mentioned in the Dockerfile. The Dockerfile should be present in the main directory.

    `docker image build -t flask-docker .`

2. Now run the image inside a container:

    `docker run -p 3000:3000 flask-docker`

   In detach mode i.e run container in the background:

    `docker run -p 3000:3000 -d flask-docker`

   Mounting volumes: This is necessary when you have changes in your files and you don't have to build the image everytime so that those changes will get reflected. Changes will be automatically loaded inside the container.

   Add -v flag : -v {source_path}:{dest_path}. Consider absolute paths. So in this case it becomes:

    `docker run -p 3000:3000 -d -v ./docker-dump/flask-docker/:./usr/src/ flask-docker`

3. Head on to ```localhost:3000/hello``` to see magic.

4. Stop the container

    `docker stop <container_id>`

5. You can find more docker commands in my Notes folder.



