<h1 align="center">
  URL-Shortner
</h1>

## Technologies

- ### Back end

  - [Echo](https://echo.labstack.com/)- Go web frameworkfor building the REST Apis
  - [Mongodb](http://mongodb.com/)- Document oriented NoSQL database.
  - [Mongo Express](https://github.com/mongo-express/mongo-express) - Web-based MongoDB admin.

- ### Front end

  - [React](https://reactjs.org/) - JavaScript library for building user interfaces.
  - [Bootstrap css](https://reactstrap.github.io/)- Responsive front-end framework.

## Getting Started

#### Clone the project

```sh
# clone it
git clone https://github.com/Mersock/react-golang-URL-shortener.git
cd react-golang-URL-shortener
```

#### Run back end && front end

#### Make sure your machine has install Docker [Docker](https://www.docker.com/)

```sh
# start
docker-compose up -d --build

# stop
docker-compose down
```

#### Url (only local machine)

Front End Landing Page - http://localhost:3000

Front End Admin Page - http://localhost:3000/adminPage

Back End - http://localhost:8080

Mongo Express - http://localhost:8081

## ☑ TODO

- [x] User page to input a url and return a shortened URL
- [x] Input URL should be validated and respond with error if not a valid URL
- [x] Visiting the Shortened URLs must redirect to the original URL with a HTTP 302 redirect
- [x] Hit counter for shortened URLs (increment with every hit)
- [x] Admin page with URL List View
- [x] Show list of URL records with:
- [x] Short Code
- [x] Full Url
- [x] Number of hits
- [x] Filter List by Short Code
- [ ] Regex based blacklist for URLs, urls that match the blacklist respond with an error
- [ ] Expiry (if any)
