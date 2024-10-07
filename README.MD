# SIMPLE API WITH VALIDATION AND AUTHENTICATION JWT

This is a simple REST API WITH JWT AUTHENTICATION build in golang

## API Reference

#### Get all users

```http
  GET /user?query=&page=1&limit=10
```

| Parameter | Type     | Description              |
| :-------- | :------- | :----------------------- |
| `query`   | `string` | For filter name or email |
| `page`    | `int`    | For default page is 1    |
| `limit`   | `int`    | For default limit is 10  |

#### Get users

```http
  GET /user/${id}
```

| Parameter | Type  | Description                       |
| :-------- | :---- | :-------------------------------- |
| `id`      | `int` | **Required**. Id of user to fetch |

#### Create users

```http
  POST /user
```

| Parameter | Type     | Description        |
| :-------- | :------- | :----------------- |
| `name`    | `string` | **Required**. Body |
| `email`   | `string` | **Required**. Body |
| `phone`   | `string` | Body               |
| `address` | `string` | Body               |

#### Update Users

```http
  PATCH /user/{id}
```

| Parameter | Type     | Description                        |
| :-------- | :------- | :--------------------------------- |
| `id`      | `int`    | **Required**. Id of user to update |
| `name`    | `string` | **Required**. Body                 |
| `email`   | `string` | **Required**. Body                 |
| `phone`   | `string` | Body                               |
| `address` | `string` | Body                               |

#### delete users

```http
  DELETE /user/${id}
```

| Parameter | Type  | Description                        |
| :-------- | :---- | :--------------------------------- |
| `id`      | `int` | **Required**. Id of user to delete |
