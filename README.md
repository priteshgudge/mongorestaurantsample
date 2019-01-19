#Mongo Sample for Golang read/write

## Data Retreival Code sample:


## Sample data source:
https://github.com/ozlerhakan/mongodb-json-files/blob/master/datasets/restaurant.json

## Tasks:

1. Write a script - python/golang to import the documents separated by newline into a mongo collection named "restaurant"

2. Write an implementation of the given interface which reads/writes data & write a main file which demonstrates this.

3. Write a command line interface which accepts command line arguments into the terminal which retreives data in golang(extra credit).

Following commands should work:
./runprogram find --type_of_food=Thai
 --> returns name:postcode of the all the matching restaurants
./runprogram find --postcode=8FY 
 --> returns name:type_of_food of all matching restaurants

./runprogram count --type_of_food=Thai
 --> returns count of all restarants matching criteria as above
 

