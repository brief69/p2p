# P2P Microservice in Go

This project is a generic and high-speed P2P communication microservice built in Go.

## Project Structure

The project is structured as follows:

- `main.go`: The entry point of the application. It loads the configuration, creates a new P2P service, starts the service, and waits for a termination signal to stop the service.
- `config.go`: Contains the logic to load the configuration from a JSON file.
- `p2p.go`: Defines the P2P service, including the creation of a new service, starting and stopping the service, and handling incoming streams.
- `service.go`: Defines the P2P service, including the creation of a new service, starting and stopping the service, and handling incoming streams.
- `utils.go`: Contains utility functions for error handling and converting strings to multiaddresses.
- `Dockerfile`: Defines the Docker image for the application.
- `.gitignore`: Specifies intentionally untracked files to ignore.
- `README.md`: This file, which provides an overview of the project.

## Getting Started

To get started, clone the repository and navigate into the directory:

```bash
git clone https://github.com/yourusername/yourrepository.git
cd yourrepository
```

Next, build the Docker image:

```bash
docker build -t p2p-microservice .
```

Then, run the Docker container:

```bash
docker run -p 8080:8080 p2p-microservice
```

The application will now be running at `localhost:8080`.

## Configuration

The application's configuration is stored in a `config.json` file. The configuration includes the host and port for the P2P service. Here is an example configuration:

```json
{
  "p2p": {
    "host": "localhost",
    "port": 8080
  }
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the terms of the MIT license.
