# Technology

To build this project, we used the following software technologies: 
- Go for backend logic and API handling. 
- Vue.js as the frontend framework, along with PrimeVue and Tailwind CSS for styling and UI components. 
- PostgreSQL as the relational database system and a local development setup using tools like pgAdmin and Vite for database management and testing.

# Installation 
- Install Go: https://webinstall.dev/golang/
- Install Node: https://nodejs.org/en
- Install PostgreSQL: https://www.postgresql.org/download/

Create a local user and database with the following information:
- User name “tuser”
- User password “password123”
- Database owner “tuser”
- Database name “mydatabase”
- Port (if needed): 5432
From the project root, run:
```sh
psql -U tuser -d mydatabase -h localhost -a -f ./backend/config/config.sql`
```

You’ll be prompted to type in a password, it will be “password123”
If you see some output saying tables have been dropped, the DBMS has been configured correctly.

In a terminal, navigate to `./backend` and do the following:
```
go mod download
go run main.go
```

In a separate terminal, navigate to `./frontend` and do the following:
```
npm install
npm run dev
```

Navigate to localhost:3001

If you see the login page like in the video, you've successfully installed the project. You can now login and interact with the DBMS with the following names (case sensitive):
- Customer: Jonathan Johnson
- Employee: Carlos Smith
- Manager: Bryan Smith

# Additional DDL Information
All SQL queries used by the UI to interact with the DBMS can be found in the go files located in `./backend/queries/*`

All DDLs can be found in `./backend/config/config.sql`
