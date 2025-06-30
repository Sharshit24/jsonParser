# jsonParser

A simple Go program to parse JSON data from a file.

## Project Structure

- `main.go`: Main application code that reads and parses `data.json`.
- `data.json`: Sample JSON file with user data.
- `go.mod`: Go module definition.

## Usage

1. Place your JSON data in `data.json`. Example:
    ```json
    {
        "name": "Alice",
        "age": 30,
        "skills": ["Go", "Python", "JavaScript"]
    }
    ```

2. Run the program:
    ```sh
    go run main.go
    ```

3. Output:
    ```
    Name: Alice, Age: 30
    ```

## Requirements

- Go 1.24.3 or later# jsonParser

A simple Go program to parse JSON data from a file.

## Project Structure

- `main.go`: Main application code that reads and parses `data.json`.
- `data.json`: Sample JSON file with user data.
- `go.mod`: Go module definition.

## Usage

1. Place your JSON data in `data.json`. Example:
    ```json
    {
        "name": "Alice",
        "age": 30
    }
    ```

2. Run the program:
    ```sh
    go run main.go
    ```

3. Output:
    ```
    Name: Alice, Age: 30
    ```

## Requirements

- Go 1.24.3 or later

## License