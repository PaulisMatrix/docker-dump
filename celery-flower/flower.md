# Flower monitoring setup for celery

* Followed this guide: https://flower.readthedocs.io/en/latest/install.html

* Note: 

    * To send broker stats to flower, we need to enable rabbitmq management plugin and give access that to the current rabbitmq user.

    * Sending broker(rabbitmq) stats to celery flower dashboard

        `rabbitmq-plugins enable rabbitmq_management_agent`

        `rabbitmqctl set_user_tags <USERNAME> management`

    * Flower will persist the state of each task executed if we pass flag --persistent=True in a sqlite db in the root directory.

    * If we want to check past tasks statues then can enable this otherwise only those tasks info will be visible on the dashboard which are in the current session.
