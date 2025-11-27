<h1 align="center">Welcome to <span style="color:mediumseagreen">GO Echo DDD boilerplate</span></h1>

It's an API Skeleton project based on Echo framework.
Our aim is reducing development time on default features that you can meet very often when your work on API.
There is a useful set of tools that described below. Feel free to contribute!

## What's inside:

- Registration
- Authentication with JWT
- CRUD API for posts
- Migrations
- Request validation
- Environment configuration
- Docker development environment

## Usage

1. Copy .env.dist to .env and set the environment variables. There are examples for all the environment variables except
   COMPOSE_USER_ID, COMPOSE_GROUP_ID which are used by the linter. To get the current user ID, run in terminal:

   `echo $UID`

   In the .env file set these variables:

   `COMPOSE_USER_ID="username in current system"` - your username in system

   `COMPOSE_GROUP_ID="user uid"` - the user ID which you got in the terminal

2. Run your application using the command in the terminal:

   `docker-compose up`
3. Make requests to register a user (if necessary) and login.
4. After the successful login, copy a token from the response, use as "apiKey" in a form:
   "Bearer {token}". For example:
   Bearer
   eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODk0NDA5NjYsIm9yaWdfaWF0IjoxNTg5NDM5OTY2LCJ1c2VyX2lkIjo1fQ.f8dSG3NxFLHwyA5-XIYALT5GtXm4eiH-motqtqAUBOI

Now, you are able to make requests which require authentication.

## Directories

1. **/cmd** entry points.

2. **/config** has structures which contains service config.

3. **/db** has seeders and method for connecting to the database.

4. **/deploy** contains the container (Docker) package configuration and template(docker-compose) for project
   deployment.

5. **/development** includes Docker and docker-compose files for setup linter.

6. **/migrations** has files for run migrations.

7. **/models** includes structures describing data models.

8. **/repositories** contains methods for selecting entities from the database.

9. **/requests** has structures describing the parameters of incoming requests, and validator.

10. **/responses** includes structures describing the parameters of outgoing response.

11. **/server** is the main project folder. This folder contains the executable server.go.

12. **/server/builders** contains builders for initializing entities.

13. **/server/handlers** contains request handlers.

14. **/server/routes** has a file for configuring routes.

15. **/services** contains methods for creating entities.

16. **/tests**  includes tests and test data.
