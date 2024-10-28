# BakingUp-Backend

Welcome to the BakingUp Backend Repository!

## Technology Stacks
<b>Programming Languages:</b>
- Go

<b>Development Tools:</b>
- <b>Backend API:</b> Go Fiber
- <b>Database:</b> PostgreSQL
- <b>Database ORM:</b> Prisma
- <b>3rd Party APIs:</b> Gemini AI
- <b>Container Management:</b> Docker
- <b>Hosting:</b> Azure Virtual Machine
- <b>DNS:</b> Cloudflare
- <b>CI/CD:</b> GitHub Actions

## Installation Guide
### 1. Clone the Repository
- For viewers on GitHub:
  
  ```bash
  git clone https://github.com/BakingUp/BakingUp-Backend.git
  ```
- For viewers in CS GitLab:
  
  ```bash
  git clone https://csgit.sit.kmutt.ac.th/csc498-499-bakingup/bakingup-backend.git
  ```
### 2. Generate Prisma Client
- Navigate to the root directory and run the following command:
  
  ```bash
  go run github.com/steebchen/prisma-client-go generate
  ```
### 3. Setup Environment Variables
 - Create a .env file in the root directory of the project and add the following environment variables:
   
   ```env
   HTTP_PORT=
   DATABASE_URL=postgresql://username:password@localhost:5432/yourdatabase
   HTTP_ALLOWED_ORIGINS=
   ```
 - **Note:** Replace `username`, `password`, and `yourdatabase` with your actual PostgreSQL credentials.
### 4. Start the Server
 - Run the following command to start the server:
   
   ```bash
   go run cmd/http/main.go
   ```

### 5. Access API Documentation (Swagger)
- Open your browser and go to the following URL to access the Swagger documentation for API endpoints:

    ```env
    http://localhost:8000/swagger/index.html
    ```
- **Note:** Ensure that the backend server is running on port 8000 for the documentation to load properly.

## More Information
For more information, please refer to the "Wiki" section at 
- For viewers in GitHub: https://github.com/BakingUp/BakingUp-Backend/wiki.
- For viewers in CS GitLab: https://csgit.sit.kmutt.ac.th/csc498-499-bakingup/bakingup-backend/-/wikis/.
