# GitHub User Activity CLI

A simple command-line interface (CLI) tool to fetch and display recent GitHub user activity. This project is part of the [roadmap.sh projects](https://roadmap.sh/projects/github-user-activity).

## Description

This CLI tool fetches the recent activity of a GitHub user and displays it in a clean, readable format in the terminal. It uses the GitHub API to fetch user events and presents them in a simple list format.

## Features

- Fetch recent GitHub user activity
- Display events in a clean, readable format
- Handle various event types (Push, Issues, Stars)
- Error handling for API failures and invalid usernames
- Rate limiting awareness with proper User-Agent headers

## Requirements

- Go 1.16 or higher
- Internet connection to access GitHub API

## Installation

```bash
# Clone the repository
git clone https://github.com/hasanm95/github-activity.git

# Navigate to the project directory
cd github-activity

# Build the project
go build
```

## Usage

```bash
# Run the program with a GitHub username
./github-activity <username>

# Example
./github-activity hasanm95
```

### Example Output

```
- Pushed 3 commits to hasanm95/github-activity
- Opened a new issue in hasanm95/github-activity
- Starred hasanm95/github-activity
```

## Project Structure

```
.
├── main.go          # Main entry point
├── events.go        # GitHub events fetching logic
├── cmd.go           # Command-line flag handling
└── README.md        # This file
```

## Error Handling

The tool handles various error cases:
- Invalid GitHub usernames
- API rate limiting
- Network connectivity issues
- Invalid API responses

## Contributing

Feel free to submit issues and enhancement requests!

## License

This project is open source and available under the MIT License.
