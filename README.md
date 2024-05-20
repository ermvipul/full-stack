# full-stack assignment
This repo consist of a demo app with go, next.js and wordpress app code.

I have tried to divide this tasks into two parts: -
1. In the first implementation, I have used a simple go app which gives a "hello, docker" message & then initialised a next-js app which used the go-app api & renders the message on browser. I have used a docker-compose config to deploy the wordpress app.
2. In the second part, I have created a post in wordpress app & then golang app to get the posts data from the wordpress api. Then the api has been used by the Nextjs app & the data has been displayed to the user.

In order to run the config in local with wordpress post api. Please run the docker-compose file at the root dir using command : -> docker compse up --build or docker compse up -d . Note: After running the compose command, please create post in wordpress & enable permalinks in settings & click on save changes. This is solve the issue related to wordpress API's

In order to run the application on the server please use the workflow file.

If you want to run only golang app & nextjs app using docker. Please use the below commands
1. Please create a docker network using docker network create <name> command
2. Please run the golang docker app using docker run --network assignment -p 1190:1190 go-app(you can use bind mounts if you want to see live changes)
3. Please run the nextjs docker app using docker run --network assignment -p 1190:1190 next-js app

Notes:
1. To build any app using dockerfile, please use command docker build -t <image_name> .
2. There any two files which are responsible for defining cross-application communications: 1. main.go file in golang app in docker-ps-ping folder. 2. Index file in docker-nextjs-app/pages folder

To-do:
1. Test cases
2. Linting