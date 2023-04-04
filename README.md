# Snippet Box
Snippet Box is a simple web application that allows users to save, view, and edit snippets of code. It is written in the Go programming language and uses the Gorilla web toolkit for routing and handling HTTP requests.

# Features
* Create, view, and edit snippets of code
* Syntax highlighting for a variety of programming languages
* User authentication and session management
* Flash messages for user feedback
* Secure password storage with bcrypt

# Installation
* Clone the repository:
bash
`
git clone https://github.com/your-username/snippet-box.git`


# Navigate to the project directory:
bash

`cd snippet-box`

# Install dependencies:
go
`go mod download`

# Set environment variables:

`
export SNIPPETBOX_DB_USER=<database-username>
export SNIPPETBOX_DB_PASS=<database-password>
export SNIPPETBOX_DB_NAME=<database-name>`

# Build and run the application:


`go build ./snippet-box`

# Usage
Once the application is running, you can access it by navigating to http://localhost:4000 in your web browser. You can create an account by clicking the "Sign up" link on the homepage, or log in with an existing account by clicking the "Log in" link.

Once you are logged in, you can create new snippets by clicking the "New snippet" button on the homepage. You can view and edit existing snippets by clicking on their titles in the snippet list.

