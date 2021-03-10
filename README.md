# Audio CRUD for Nooble
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![PR's Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](http://makeapullrequest.com) 
[![GitHub followers](https://img.shields.io/github/followers/saivittalb.svg?style=social&label=Follow)](https://github.com/saivittalb?tab=followers) 
[![Twitter Follow](https://img.shields.io/twitter/follow/saivittalb.svg?style=social)](https://twitter.com/saivittalb) 

A simple backend repo for creating and maintaining an audio shorts directory. APIs were built to Create, Read, Update and Delete (CRUD) Audio Short and List Audio Shorts. Done as a part of my intro assignment for interning with Nooble.

The tech stack used was Go and PostgreSQL. PostgreSQL was utilized using the cloud-based setup with the help of ElephantSQL.
 
###### Note 
Developed with Go 1.16.1.
Editor used was Visual Studio Code 1.54.1.
 
## License
This project is licensed under the MIT License, a short and simple permissive license with conditions only requiring preservation of copyright and license notices. Licensed works, modifications and larger works may be distributed under different terms and without source code.

<p align="center"> Copyright (c) 2021 Sai Vittal B. All rights reserved.</p>

## Setup

### Prerequisites
- Install Go.
- Install the dependencies using the following commands in terminal/console window:
```bash
$ go get -u github.com/gorilla/mux
$ go get github.com/lib/pq
$ go get github.com/joho/godotenv
```

### ElephantSQL/Custom Setup
- Setup a method to use PostgreSQL; or create an ElephantSQL account for easy setup and follow the below steps.
- Create an instance using their free plan (which doesn't require any credit card).
- Once your instance is created, go to the ```Browser``` tab in the ElephantSQL and paste the code from the file ```table.sql``` and execute it.
- Open the ```Details``` tab and copy the URL and paste it to the ```POSTGRES_URL``` in the ```.env``` file.

### Testing the APIs
- Run the following command in terminal/console window:
```bash
$ go run main.go
```
- Test the APIs with Postman by importing the Postman Collection file which is in the JSON format in this repo.
- If you don't have Postman installed, test the APIs with curl requests using the routes defined in the file ```router/router.go```.

## Contributing
- Fork this project by clicking the ```Fork``` button on top right corner of this page.
- Open terminal/console window. 
- Clone the repository by running following command in git:
 ```bash
$ git clone https://github.com/[YOUR-USERNAME]/audio-crud-nooble.git
```
- Add all changes by running this command.
```bash
$ git add .
```
- Or to add specific files only, run this command.
```bash
$ git add path/to/your/file
```
- Commit changes by running these commands.
```bash
$ git commit -m "DESCRIBE YOUR CHANGES HERE"

$ git push origin
```
- Create a Pull Request by clicking the ```New pull request``` button on your repository page.

[![ForTheBadge built-with-love](http://ForTheBadge.com/images/badges/built-with-love.svg)](https://GitHub.com/saivittalb/) 
[![ForTheBadge powered-by-electricity](http://ForTheBadge.com/images/badges/powered-by-electricity.svg)](http://ForTheBadge.com)

<p align="center"> Copyright (c) 2021 Sai Vittal B. All rights reserved.</p>
<p align="center"> Made with ❤ by <a href="https://github.com/saivittalb">Sai Vittal B</a></p>
