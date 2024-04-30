# Groupie Tracker

Groupie Tracker is a web application that provides information about various music artists, their locations, concert dates, and relations.

## Authors

- Hasan Dhaif (hadhaif)
- Ahmed Alhamed (ahalhamed)

## Project Structure

The project is structured into several packages:

- `models`: Contains the data structures used in the project, such as `Artist`, `Locations`, `Dates`, and `Relation`.
- `handlers`: Contains the HTTP handlers that handle requests to various endpoints.
- `middleware`: Contains the middleware functions used to add functionality to the HTTP handlers.
- `templates`: Contains the HTML templates used to render the web pages.

## Installation and Setup

1. Clone the repository: `git clone https://learn.reboot01.com/git/hadhaif/groupie-tracker`
2. Navigate to the project directory: `cd groupie-tracker`
3. Run the project: `go run cmd/groupie-tracker/main.go`

## Usage

Once the server is running, navigate to `http://localhost:8080` in your web browser. You will see a list of artists. Click on an artist to see more information about them, including their locations, concert dates, and relations.
