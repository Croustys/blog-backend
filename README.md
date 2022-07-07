# Blog backend

## Developed by Ákos Barabás & written in Go

For development purposes place `.env` file in `root` folder

### API routes

`POST`

- `/login`
- `/register`
- `/create` - post creation

`GET`

- `/user/{id}`
- `/posts?offset={1}&limit={1}` --optional parameters for lazy loading
- `/posts` -- all posts
- `/post/{id}` -- specific post

## Technologies used

- [JWT](https://github.com/golang-jwt/jwt) - token based authentication
- [Bcrypt](https://pkg.go.dev/golang.org/x/crypto/bcrypt) - password hashing
- [MongoDB](https://www.mongodb.com/) - Document based Database

## Usage

- register

```js
fetch("api.domain.com/register", {
  method: "POST",
  headers: {
    "Content-type": "application/json",
  },
  body: JSON.stringify({ username, email, password }),
});
```

- login

```js
fetch("api.domain.com/login", {
  method: "POST",
  headers: {
    "Content-type": "application/json",
  },
  body: JSON.stringify({ email, password }),
});
```

- create post

```js
fetch("api.domain.com/create", {
  method: "POST",
  headers: {
    "Content-type": "application/json",
  },
  body: JSON.stringify({ title, content }),
});
```
