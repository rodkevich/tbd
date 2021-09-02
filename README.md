tbd

#### Вообще эта задача для NoSQL базы...

***


Locally:   
RUN: `go mod init` - if not exists  
RUN: `go mod vendor`  
RUN: `go run ./cmd/task1/task1.go `

For tasks using db run on Docker navigate ./cmd/task5/:  
`docker-compose build --parallel`  
`docker-compose up`  
RUN: `go run ./cmd/task5/server/server.go `

WARNING:  
If you have any docker volumes attached to a postgres instance - init.sql won't work.  
Before running a server detach them using :  
`docker-compose down --volumes`
Or : create migrations from ./cmd/task5/config/migrations/init.sql manually

Structure:  
`pkg`  
|-`app`   
|-`datasource`  
| |-`postgress`  
|-`tickets`  - **will be also used for last task**  
| |-interface.go  *used for taskN 3*  
| |-functions.go  *used for taskN 1*  
| |-methods.go  *used for taskN 1*  
| |-structures.go  *used for taskN 1*  
| |-`types`  - **additional types used for development**  
| | |-currency.go  
| | |-description.go  
| | |-phone.go  
| | |-link.go

`cmd`  *runnable files and requirements for each task*  
|-**task1 - structures, methods & functions**  
| |-task1.go  
| |-req.yaml  
|-**task2 - for loop**  
| |-task2.go  
| |-main.go  
| |-req.yaml  
|-**task3 - interfaces**  
| |-req.yaml  
| |-task3.go  
|-**task4 - algorithm** `?q=is_this_task_ok`   
| |-task4.go  
| |-req.yaml  
|-**task5 - last task app**  
| |-validations.go - **example**
| |-req.yaml | |-`server` - **application**
| |-`docker-compose.yml

`internal` *different small helpers*  
|-msg  
| |-messages.go

|-**`test`**  *test files*

README.md
***  
