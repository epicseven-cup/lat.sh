# lat.sh
just a fancy url shortener


# Setup
Using the docker setup you can run it directly, just make sure to update the database password


on the project folder start the backend and database
```bash
docker compose up --build
```

compile the frontend and place the files into a static host application like nginx
```bash
npm run build
```