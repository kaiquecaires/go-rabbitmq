# Go RabbitMQ
This project is designed for practicing the concepts of RabbitMQ. You can get the full tutorial from: https://www.rabbitmq.com/tutorials.

# Prerequisites
- Docker: Install Docker on your system by following the official installation guide: [Install Docker](https://docs.docker.com/get-docker/)

# Setup
1. Run RabbitMQ instance <br />
Open your terminal and execute the following command to start a RabbitMQ instance using Docker:
```bash
docker run -it --rm --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3.12-management
```

2. Accessing the RabbitMQ Management Dashboard <br />
Open a web browser and navigate to http://localhost:15672/ to access the RabbitMQ management dashboard. Use the default login credentials: `guest` (username) and `guest` (password).

# Additional Notes
- <strong>Detached Mode</strong>: If you prefer to run the container in the background and continue using the terminal, add the -d flag to the docker run command.

- <strong>Data Persistence</strong>: To persist data and configurations between container restarts, consider using Docker volumes. See the example in the "Run RabbitMQ Instance" section.

- <strong>Network Isolation</strong>: If you plan to have multiple containers communicating with each other, consider creating a Docker network to isolate them.

# Cleanup
To stop and remove the RabbitMQ container, run the following command:

```bash
docker stop rabbitmq
```
