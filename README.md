# Basic Web API in Go

A simple web API built in Go using the `net/http` package. This project demonstrates how to create endpoints, handle HTTP requests, and return responses in Go.

## How to Execute

1. Make sure you have Go installed on your system.

2. Clone the repository:

```
git clone https://github.com/seu-usuario/api-web-basica-go
```

3. Navigate to the project directory:
```
cd api-web-basica-go
```

4. Run the server:
```
go run main.go
```

5. The server will be running at http://localhost:8080.

- Endpoints
```
/welcome: Returns a simple welcome message.
```

```
/form: Presents an HTML form where you can enter a name. When the form is submitted, it returns a personalized message with the name.
```

## Example of use
- Visit http://localhost:8080/welcome to see the welcome message.
- Visit http://localhost:8080/form to access the form.

## License
This project is licensed under the MIT License.