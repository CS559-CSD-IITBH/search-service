# Search Service

The Search Service is a Go-based service that provides search functionality using Elasticsearch for a collection of stores and their items.

## Features

- **Elasticsearch Integration:** Utilizes the Elasticsearch database to perform efficient searches on store data.
- **Gin Web Framework:** Uses the Gin web framework for handling HTTP requests and responses.
- **Structured Query:** Supports structured queries, including nested queries.

## Prerequisites

Before running the service, make sure you have the following dependencies installed:

- Go (version 1.13 or higher)
- Elasticsearch (version 6.x)

## Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/search-service.git
   cd search-service

## Notes

1. Remember to add the **.env** file in the root directory.
2. Add the Mongo URL hosted on cloud in config file for monstache.
