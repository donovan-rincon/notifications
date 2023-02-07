# Notifications application

This application is a notification system that has the ability to receive a message and depending on the
category and subscribers, notify these users in the channels they are registered.
It will be 3 message categories:
- Sports
- Finance
- Movies

And there will be 3 types of notifications, each type should have its own class to manage the logic of
sending the message independently.

- SMS
- E-Mail
- Push Notification
---
## Pre requisites


- Docker installed (if wanted to run dockerized) otherwise it can be run if go and vue are installed, running from each backend and frontend folders

---
## How to build it


This command will build the frontend app and create a docker image called frontend
```
docker build -t frontend .
```

Once the image is built, we run it with the following command
```
docker run -it -p 8000:80 --rm --name frontend frontend
```

This command will build the backend and create a docker image called backend
```
docker build -t backend .
```

Once the image is built, we run it with the following command
```
docker run -it -p 8001:8080 --rm --name backend backend
```

### Notes 
The front end app already has in its `.env` file the proper url that the backend docker container will expose

---

## Missing functionality

- No unit tests were added (due to late start on the challenge, no test were added in order to achieve desired functionality of the challenge)
- No users addition or edit
  