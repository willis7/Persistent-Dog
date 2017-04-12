# Persistent Dog

RESTful API for tracking software releases.

### Build Status

[![Build Status](https://travis-ci.org/willis7/Persistent-Dog.svg?branch=master)](https://travis-ci.org/willis7/Persistent-Dog) 

[![Go Report Card](https://goreportcard.com/badge/github.com/willis7/Persistent-Dog)](https://goreportcard.com/report/github.com/willis7/Persistent-Dog)


### Credit
 
Project name generator:
http://online-generator.com/name-generator/project-name-generator.php

### Architectue

Here's how the application is structured.

* **common**: Implements some utility functions and provides initialization logic for the application
* **handlers**: Implements the application’s application handlers
* **data**: Implements the persistence logic with the MongoDB database
* **models**: Describes the data model of the application
* **routers**: Implements the HTTP request routers for the RESTful API
