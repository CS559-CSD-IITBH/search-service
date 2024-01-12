# Search Service

It is a Go-based service that provides search functionality using Elasticsearch for a collection of stores and their items.

## Features

- **Elasticsearch with MongoDB:** Utilizes [Elasticsearch](https://github.com/elastic/elasticsearch) to perform efficient searches on store data. Using [go-elasticsearch](https://github.com/elastic/go-elasticsearch) and [Monstache](https://github.com/rwynn/monstache) for quick and easy integration with [MongoDB](https://github.com/mongodb/mongo). 
- **Gin Web Framework:** Uses the [Gin](https://github.com/gin-gonic/gin) web framework for handling HTTP requests and responses.
- **Structured Query:** Supports structured queries, including nested queries. Integrated dashboard with [Kibana](https://github.com/elastic/kibana).

## Prerequisites

Before running the service, make sure you have the following dependencies installed:

- Go (version 1.20 or higher)
- Elasticsearch (version 6.x)
- Docker 

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/search-service.git
   cd search-service

2. Add your personal **MongoDB** URL to the `config.toml` file in the `monstache` directory.

   ```
   mongo-url = <url for your personal mongo instance>
   ```
   
4. Add a `.env` file in the root directory. It should contain the following fields.

   ```
   PORT=<port for the api-service>
   ELASTICSEARCH_URL=<url for your personal elasticsearch instance>
   MONGODB_URL=<url for your personal mongo instance>
   MONGODB_COLLECTION=<database-name.collection-name in your mongo instance> 
   ```

5. Build the docker image for the **api-service**. Run the following command in the root directory.

   ```
   docker build -t search-service-api:latest .
   ```

6. Start the services using **docker-compose**. Note that it is expected your *Mongo* is hosted on cloud.

   ```
   sudo docker-compose compose build && sudo docker-compose up
   ```
   When you start the services for the first time, your Elastic/Kibana credentials will be in the application      logs, on your terminal screen.  

7. After the services are up and running, you can visit the **Kibana** dashboard at http://localhost:5601/.        Enter the credentials required. Now you can run queries!
