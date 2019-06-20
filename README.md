# gorestapi

This aplication is to be used as a template for creating a REST API in Go. It is a WORK IN PROGRESS. Please review the code if you plan on using it. Use at your own risk. Comments, Suggestions, and Pull Requests are welcome.   

## Configuration

The app uses the AWS SDK to fetch the config from secretsmanager so you will need to export your AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY as environment variables inside your running container. Ideally these credentials should be better protected, but for this example it will do. You will need to setup a secret in your AWS account called gorestapiconfig. Your configruation should be modeled after the 'sampleconfig.json' in this repo and stored as JSON in your secret as its value. THE API is TLS enabled and it has been tested using a self-signed certificate, but you can import a certificate of your choosing. 

The app was tested using a backend AWS RDS database instance as a data store but you could use your own. Modify the config as needed to achieve this. There is a sampleDBSchema.sql provided in the repo that will create the scheama required to test the sample data. You will need to setup your own user and grants to the database after importing the schema 

## Building

git pull

go get -d ./...

go build .

Put it in a container and run. The Docker file provided can be used to build your own image to run the API on Kubernetes or your container orchestrator of choice. 

## Usage

The sample application allows you to list / update and delete fictitious employees from the database. 

To use https://<serverName>:<port>/<path> using curl. 

Use JSON formated data for POST/PATCH methods.

```curl https://host/path -d '{"foo":"bar"}```  

### Paths:
GET /v1/employees - Provides a JSON list of all employees

POST /v1/employeesinsert - Inserts a new employee 

PATCH /v1/employeesupdate - Updates an existing employee by given EmpID. 

POST /v1/employeesdelete - Removes an employee


## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
[GPL3]

## TODO:
Add authentication for POST/PATCH URI's
APISpec for the API

