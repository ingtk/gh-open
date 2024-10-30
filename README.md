# gh-open

Opens the remote URL of the specified Git repository in a browser.

[日本語](README.ja.md)

## Features

- Accepts the path of the Git repository as a command-line argument (not required if in the current directory).
- Retrieves the URL of `origin` from the remote settings of the repository.
- Opens the GitHub repository page in a browser based on the retrieved URL.

## How to Run

Navigate to the directory under the git repository and enter the following command:

```
gh-open
```

## Installation Instructions

1. Clone or download this repository.
2. Install the program.
```bash
go install .
```

## Notes

- An error message will be displayed if the `origin` remote is not found.
- Supported platforms are Linux, Windows, and macOS, but only macOS has been verified.
