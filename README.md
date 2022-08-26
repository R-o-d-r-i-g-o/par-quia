# Scheduler


## routers design

```
Scheduler.com/Welcome
	/Login
	/Register

Scheduler.com/Home
	/Profile/:name
		/Configurations
		/Calendar
		
	/Community
		/:name
	
	/Events
		/:date
		
	/Pastors
		/:name
			/Members
			/Configurations
			/Schedules-list
				/:name
					/edit
					/view

```


## Project Architeture

```

├── api
│   ├── router.go
│   └── server.go
├── common
│   ├── configs
│   │   ├── catch.go
│   │   └── env.go
│   ├── library
│   │   ├── const_errors.go
│   │   └── const_names.go
│   └── util
│       ├── util.go
│       └── util_test.go
├── core
│   ├── event
│   │   ├── controller.go
│   │   ├── routes.go
│   │   └── service.go
│   ├── pastor
│   │   ├── controller.go
│   │   ├── routes.go
│   │   └── service.go
│   ├── schedule
│   │   ├── controller.go
│   │   ├── routes.go
│   │   └── service.go
│   ├── section
│   │   ├── controller.go
│   │   └── routes.go
│   └── user
│       ├── controller.go
│       ├── repository.go
│       ├── routes.go
│       └── service.go
├── database
│   ├── database.go
│   ├── entity
│   │   ├── seed.go
│   │   └── tables_gorm.go
│   ├── model
│   │   └── tables_Json.go
│   └── runner.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── Makefile
└── README.md


```