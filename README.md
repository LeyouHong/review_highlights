# About project
Used [shield](github.com/eaigner/shield) to traning data and save the data into 2 redis.<br>
One is collect emotion data, another is collect noun data. For every review, we can use emotion data and noun data<br>
to score it.

# Install

## Install docker-compose
Please see the [docker-compose](https://docs.docker.com/compose/install/#prerequisites)

## Install Golang in Linux.
Download Golang last version: [download](https://golang.org/dl)<br>
[install go](https://golang.org/doc/install)<br> 
Set the env in ~/.profile(GOPATH and GOROOT based on your env)

## Build rh
````
go get github.com/LeyouHong/review_highlights

make

./rh reviews.txt 2
````