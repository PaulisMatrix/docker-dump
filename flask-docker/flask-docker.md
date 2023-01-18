# Simple flask application using docker

Running my simple python flask application inside a docker container

### **Please refer below commands to run the application.**

- Preq : Ofc you need to have docker installed for Mac/Windows/Linux. Please travel to [this link](https://docs.docker.com/get-docker/) for installation.

1. Build the image whose specifications you have mentioned in the Dockerfile. The Dockerfile should be present in the main directory.

   `docker image build -t flask-docker .`

2. Now run the image inside a container:

   `docker run -p 3000:3000 flask-docker`

   In detach mode i.e run container in the background:

   `docker run -p 3000:3000 -d flask-docker`

   Mounting volumes: This is necessary when you have changes in your files and you don't have to build the image everytime so that those changes will get reflected. Changes will be automatically loaded inside the container.

   Add -v flag : -v {source_path}:{dest_path}. Consider absolute paths. So in this case it becomes:

   `docker run -p 3000:3000 -d -v ./docker-dump/flask-docker/:./usr/src/ flask-docker`

3. Head on to `localhost:3000/hello` to see magic.

4. Stop the container

   `docker stop <container_id>`

5. You can find more docker commands in my Notes folder.

6. "Best" dockerfile for python applications:

   https://luis-sena.medium.com/creating-the-perfect-python-dockerfile-51bdec41f1c8

### **Compose file to specify different services/dependencies.**

- When you want to run different services which are dependent on eachother, you can specify its settings in a docker-compose.yml file.

Suppose you have a flask app running with celery configured to consume tasks that are being pushed to rabbitmq and redis being used for caching.

Celery acts as both, a producer and a consumer. The producer is called client/publisher and the consumer as workers.

**Setup required**:

1.  Celery setup:

    a. Configuring rabbitmq which we use as a broker to push our tasks to and celery :

        * Add a user -> rabbitmqctl add_user rabbit_mq_user rabbit_mq_user_pass

        * Add a virtual host which is basically a namespace provided by rmq which allows to separate our queues depending on use cases-> rabbitmqctl add_vhost default_tasks

        * Add user to this virtual host -> rabbitmqctl set_permissions -p default_tasks rabbit_mq_user “.” “.” “.*”

        * To view queues assigned to your vhost -> rabbitmqctl list_queues -p default_tasks name messages consumers

        * Your celery url to be added in the settings become -> CELERY_BROKER_URL = "amqp://{}:{}@{}/{}".format(RABBIT_MQ_USER, RABBIT_MQ_PASSWORD, RABBIT_MQ_HOST, RABBIT_MQ_VHOST)
        where:
            1. RABBIT_MQ_USER  = os.getenv('RABBIT_MQ_USER', "")

            2. RABBIT_MQ_PASSWORD = os.getenv('RABBIT_MQ_PASSWORD', "")

            3. RABBIT_MQ_HOST = os.getenv('RABBIT_MQ_HOST',"")

            4. RABBIT_MQ_VHOST = os.getenv('RABBIT_MQ_VHOST',"")

    b. Configure redis :

        * Specify port -> REDIS_PORT = os.getenv("REDIS_PORT", 6379)

        * Specify host -> REDIS_HOST = os.getenv("REDIS_HOST", "localhost")

        * Specify password -> REDIS_PASSWORD = os.getenv("REDIS_PASSWORD", "")

        * Specify redis namespace/context db -> REDIS_DB = os.getenv("REDIS_CONTEXT_DB", 10)

    c. Refer to the compose file wherein we are settings these variables.

    d. Few commands :

        * Start your flask app -> gunicorn -c gunicorn_conf.py app:app

        * Start your celery worker -> celery -A app.celery worker --loglevel=INFO

        * Use rabbitmq and redis CLI to see your queues/tasks, keys being set resp.

        * After all is set do :

            1. docker-compose up -> to start all your services.

            2. docker-compose down -> to stop and remove all your services.

            3. docker-compose up <service_name> -> to start a specific service.(Pass -d flag to run in the background)

            4. You can start and stop without removing containers by using -> docker-compose start/stop
