# go-technical-test-bankina
backend developer technical test at bank ina

# Technical Test Description:
Test Overview:

The test focuses on building a component of a rest api in golang.
The test assesses skills in designing and implementing Create a RESTful API for a user and tasks feature using a Clean architecture golang

# Project
The project is using clean architecture. there are inside the /src folder. In this repo have a entity, repository, service and handler.

the authenticate is using jwt to make it secure.

# Usage
Inside this project will running on default port localhost:8080 with the default database in port 3306 in your local

user
 "email": "agus12342gmail.com",
 "password": "agus12345"

# Endpoint
API Documentation are attached in 
**postman folder** and you could import to other device to see the detailed body and response message.

# Migrate
Import to local database using mysql. inside the /db folder there is a database file, you can import that file into your local database. (mysql)

Use command `make migration_up` to migrate the tables

# Run Project
Use command `make run`