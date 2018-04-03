# Setup

I tried beforehand to setup a docker environment with two containers (one for the Go application, one for the MySQL database), but for some reason the go_app container does not seem to work properly. You will still find the related configuration files if you want to review them. So, I will describe the steps to set it up on your local.

- Requirements:
    * Git
    * Go
    * MySQL

## Git

First of all, you need to have Git installed on your marchine. It will be needed to retrieve the dependencies for the Go application.

## Go

Then, please install the Go language. On MacOS, it is fairly straight-forward as it is. You just have to download a .pkg file and follow the installation prompt. On Linux, it may be a bit trickier because you will need to choose the appropriate archive file and setup the $GOPATH. Please find the instructions [here](https://golang.org/doc/install#tarball).

Alternatively, a MacOS static build of the project is provided, you just need to run it ("./user-api-endpoint") in a MacOS environment so you would not need to install Go to run the project.

If you want to build the project (in case you are on Linux), please run these commands at the root of the project:
- "go get"
- "go build"

This will generate a new binary for the current environment. You can then execute this binary.

## MySQL

To setup the MySQL database, you just need to import the tables.sql into your MySQL system. You can either do it with PHPMyAdmin or run this command: "mysql -u root -p < tables.sql".

# Testing

- To run the unit tests, please go to the users and locations folders and run this command: "go test -v" (needs to have Go installed)
- To test manually the REST API endpoints, please run curl on the POST endpoints:
	* Insert user: "curl -d '{"username":"scatter", "phone_number":"04593412782", "email":"r@gmail.com"}' -H "Content-Type: application/json" -X POST http://localhost:8888/user"
	* Insert/Update location: "curl -d '{"longitude":153.35, "latitude":64.265, "user_id":1}' -H "Content-Type: application/json" -X POST http://localhost:8888/user/location"
- To test manually the GET endpoint, please hit this endpoint: http://localhost:8888/user/:id (with :id being the related user id)

# RFI

- Why did you choose these specific technologies?
I decided to use Go and MySQL because it was suggested by Tom to use Go and also it is well-suited for the creation of a REST API. I used also MySQL because that was the most straight-forward system to setup and it would be useful if we need to extend this project. Also, MySQL is really useful for the uniqueness constraints on the user details.

- What were the major challenges you faced (at least one, even if you think it is all easy, the least easiest bit)?
The major challenge I faced was when I tried to create the environment setup with Docker because I was not used to it. The Dockerfiles and docker-compose.yml files are still in the project for review. The setup of the MySQL container seems fine but I had some trouble with the Golang container.

I also had some difficulty to setup the MySQL connection within the Go project, but there is some good documentation on internet so I succeeded to do it.

- What performance issues do you think may arise in your solution as the number of users increases?
One of the performance issue that may arise is while inserting an user and inserting or updating his location.

- What could you do to avoid these issues?
To avoid these issues, I could have created a connection pool for the connection with the MySQL database and also use the concurrency patterns of Go (channels or mutexes).

- What did you like/dislike about the test (please provide at least one of each)?
I really appreciated working with Go and setting up the REST endpoints. I did not like trying to setup the Docker environment because I am not used to it.

- What changes could we make to improve the test (make it more interesting, harder, easier, etc...)?
To make the test easier, it would be nice to allow the use of regular expressions for validation and not require to setup docker files and a docker-compose.yml file.

To make the test harder and more interesting, it would have been interesting to rely on a third party api to retrieve some informations. For instance, we could have setup an authentication system relying on the Google OAuth2 authentication system.
