
# ShaleApps Full Stack Engineering Challenge (Todo) - Amirul

## Description

This project is a Todo List app built as part of the ShaleApps Full Stack Engineering Challenge. The app allows users to manage their tasks efficiently without the hassle of authentication. It is implemented using Go for the backend and Angular for the web client.

## Basic Criteria

- **Storage**: Items are stored in a relational/SQL database
- **Version Control**: The project is version controlled using Git and hosted on GitHub: https://github.com/mirul22/shaleapps-challenge
- **Deployment**: The app is deployed online for accessibility on Heroku

## Technologies Used

- **Backend**: Go
- **Frontend**: Angular
- **Docker**: For containerization
- **Heroku**: For deployment


## Features

- **Task Management**: Easily add, remove, and update tasks.
- **User-friendly Interface**: Intuitive design for seamless user experience.
- **Efficient Storage**: Utilizes a relational database for efficient data management.

## Deployment

- **Frontend**: https://shaleapps-challenge-frontend-2d7ea34c0d13.herokuapp.com/
- **Backend**: https://blooming-everglades-24065-95422c9c99ea.herokuapp.com/api/todos

## Installation and Usage

To run this project locally using Docker, follow these steps (Run in your terminal):

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/mirul22/shaleapps-challenge.git
   ```

### Frontend

2. Navigate to the project directory (frontend):

    ```bash
    cd frontend
    ```

3. Install all dependencies

    ```bash
    npm install
    ```
4. Serve the project

    ```bash
    ng serve
    ```

5. Access the app via your web browser at http://localhost:4200.

### Backend

2. Navigate to the project directory (backend):

    ```bash
    cd backend
    ```

3. Build and start the Docker containers:

    ```bash
    docker-compose up --build
    ```
5. Access the app via your web browser at http://localhost:8000/api/todos.

## Getting Started

To get started with the project, ensure you have Docker installed on your machine. Docker makes it easy to set up and run the project in any environment without worrying about dependencies.

## Contributors

- **Amirul** - Lead Developer

## License

This project is licensed under the [MIT License](LICENSE).

## Acknowledgments

- Special thanks to ShaleApps for providing the opportunity to work on this challenge.
- Thanks to the developers of Go and Angular for creating such powerful tools.

## Resources

### Technologies

- [Go Documentation](https://go.dev/)
- [Angular Documentation](https://angular.io/)
- [Docker Documentation](https://docs.docker.com/)
- [Heroku Documentation](https://devcenter.heroku.com/)

### Go Learning Resources

- [ByteSizeGo: Learning Golang 2024](https://www.bytesizego.com/blog/learning-golang-2024)
- [Go Tour](https://go.dev/tour/welcome/1)
- [Go By Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Avoid ORM in Go: Use Pure SQL Instead](https://betterprogramming.pub/avoid-orm-in-go-use-pure-sql-instead-3ae7f0485b37)
