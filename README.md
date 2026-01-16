# Celeste Classic Net

A simple web application with a Go backend server and HTML/CSS/JS frontend.

## Project Structure

```
├── server/          # Go backend server
│   ├── main.go     # Server implementation
│   └── go.mod      # Go module file
└── website/        # Frontend website
    ├── index.html  # Main HTML page
    ├── style.css   # Styling
    └── script.js   # JavaScript for button interaction
```

## Features

- Go HTTP server that serves static files and handles API requests
- Simple website with a button that sends requests to the server
- Button clicks are logged to the Go console
- Clean, responsive UI with gradient styling

## How to Run

1. Navigate to the server directory:
   ```bash
   cd server
   ```

2. Run the Go server:
   ```bash
   go run main.go
   ```

3. Open your browser and navigate to:
   ```
   http://localhost:8080
   ```

4. Click the "Click Me!" button and watch the Go console for log messages!

## How it Works

- The Go server (`server/main.go`) serves static files from the `website/` directory
- When the button is clicked, JavaScript (`website/script.js`) sends a POST request to `/api/button-click`
- The Go server receives the request and logs "Button was clicked!" to the console
- The server responds with a JSON message confirming receipt

## To Do

- modify p8 cart for celeste
- ping go server from p8 on js + html
- more...
