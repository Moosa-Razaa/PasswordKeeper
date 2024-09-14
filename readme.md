# PasswordKeeper

PasswordKeeper is a secure and efficient password management application. This project uses Vue with TypeScript for the frontend and Go for the backend.

## Features

- Password generation
- Perform CRUD operation to store, read, update and delete passwords
- Easy authentication

## Tasks

- [x] Set up Vue project with TypeScript
- [x] Set up Go project for backend
- [x] Create password generation feature
- [ ] Create API to save passwords(create)
- [ ] Create API to read passwords(read)
- [ ] Create API to modify existing passwords(update)
- [ ] Creaet API to remove passwords(delete)
- [ ] Build UI for generating passwords
- [ ] Build UI for saving passwords
- [ ] Build UI for reading passwords
- [ ] Build UI for modifying passwords
- [ ] Build UI for removing passwords
- [ ] Integrate frontend with backend

## Getting Started

### Prerequisites

- Node.js
- Go

### Installation

1. Clone the repository:
    ```sh
    git clone https://github.com/yourusername/PasswordKeeper.git
    ```
2. Navigate to the project directory:
    ```sh
    cd PasswordKeeper
    ```
3. Install frontend dependencies:
    ```sh
    cd frontend
    npm install
    ```
4. Install backend dependencies:
    ```sh
    cd ../backend
    go mod tidy
    ```

### Running the Application

1. Start the backend server:
    ```sh
    cd backend
    go run main.go
    ```
2. Start the frontend development server:
    ```sh
    cd ../frontend
    npm run serve
    ```

## License

This project is licensed under the MIT License.