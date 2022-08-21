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