# Forum
## Overview

This project is a web forum application built with Go, JS, and SQLite. It enables users to communicate by creating posts and comments, associating categories with posts, liking/disliking content, and filtering posts. Docker is used for containerization to ensure easy setup and deployment.

## Authors
- Yusuf (yjawad)
- Marwa (malkhuza)
- Ali (aabdulhu)
- Sayed Ali (sayedalawi)
- Yusuf (yrahma)

## Features
- **User Authentication**
  - User registration with email, username, and password.
  - Password encryption.
  - Unique sessions using cookies with expiration dates.
  - UUID for session management.

- **Communication**
  - Registered users can create posts and comments.
  - Posts can have one or more categories.
  - All users can view posts and comments, but only registered users can create and interact with them.

- **Likes and Dislikes**
  - Registered users can like or dislike posts and comments.
  - Like and dislike counts are visible to all users.

- **Post Filtering**
  - Filter posts by categories.
  - Filter by posts created by the user.
  - Filter by posts liked by the user.


## Setup and Installation

### Prerequisites
- Go
- Docker
- SQLite

### Installation

1. **Clone the repository:**
    git clone https://learn.reboot01.com/git/malkhuza/forum.git
    cd forum

2. **Build the Docker image and Run the Docker container:**
    bash run.sh

3. **Access the application:**
    Open your browser and go to `http://localhost:5050`.

## Database

The forum uses SQLite for data storage.
- **Creating Tables**

- **Inserting Data**

- **Querying Data**

## Authentication

Users can register and log in to create and interact with posts and comments. Session management is handled via cookies, ensuring secure and single-session logins.

### Registration

- Requires email, username, and password.
- Checks for unique email.

### Login

- Validates email and password.
- Creates a session with a cookie that has an expiration date.

## Communication

Registered users can:

- Create posts with categories.
- Comment on posts.
- View all posts and comments.

## Likes and Dislikes

- Only registered users can like or dislike content.
- Like and dislike counts are visible to all users.

## Error Handling

The application handles various errors such as:

- HTTP status errors.
- Technical errors.
- User-friendly error messages.
- If either the key or the pem weren't found, we will terminate the server as we cannot establish an https server.

### HTTPS Setup
Enabling HTTPS

1. **Generate an SSL Certificate**

2. **Configure the Server**

## Best Practices

- Follows Go coding conventions.
- Structured and maintainable codebase.

## Packages

- Go standard packages.
- `sqlite3`
- `bcrypt`
- `UUID`


## Learning Outcomes

- Basics of web development:
  - HTML
  - HTTP
  - Sessions and cookies
- Using and setting up Docker:
  - Containerizing an application
  - Managing dependencies
  - Creating Docker images
- SQL and database management:
  - Creating and manipulating databases
  - Writing SQL queries
- Basics of encryption and security.

## Acknowledgments

- SQLite for the database management.
- Docker for containerization.
- Go standard packages for building the application.

