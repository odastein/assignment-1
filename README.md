# Assignment 1
###### By Oda Steinsholt

## What the service does
This is a service that gets information, about universities 
and countries from two different web services. Further it 
treats and combines the information, and gives it to this 
web service. 

## How to use the program
The main function is placed in its own file called main.go, 
which is located right under the assignment-1 folder. To use 
the program, click run, and write in the API's URL in the web 
browser. 

## Structure of the program
The program is made up of three go-files, one for each task, 
alongside the main file, a file with structs and a file with
constants. 

The first go-file is called universityInfoHandler.go and 
contains the functionality for the first task. The first task is 
about finding information about a university, from a list of 
universities, and a list of countries. In this file we have the 
UniInfoHandler function, and this uses the getRequest1Handler 
function. This function again uses getCountry, and getUniversity.

The second go-file is neighbourUniversityHandler, and its 
responsibility is to do the second task. The second task is 
about finding info about universities, based on a search word 
for their neighbour country, and a search word for the 
university name. It can also contain a number, which decides the 
maximum universities returned, per neighbour country. We have the 
NeighbourUnisHandler function that uses the getRequest2Handler
function. This function uses getCountryByName, 
getUniByCountryAndName and getCountry.

The third go-file handles the third task, and is called 
diagnostics.go. The third task is where we check the status 
codes for country and university, the version, and the time 
that the service has been running. This file contains the 
DiagHandler function that uses the getRequest3Handler function. 
This again uses the upTime function.

## Example 
#### One request the user can send us this: 
* http://localhost:8080/unisearcher/v1/uniinfo/uruguay. 

That will 
return a list of universities with "uruguay" in the name. 

#### Another example is
* http://localhost:8080/unisearcher/v1/neighbouruni/turkey/middle/8.

This will return the first 8universities in each of Turkeys 
neighbouring countries, that contains the word "middle". This 
could be many, since this area is a part of the middle east.

## Dependencies
This service is dependent on two different web services, in 
order to get the information needed. The two services are:
* http://universities.hipolabs.com
* https://restcountries.com
