# lat.sh
just a fancy url shortener

# Deployment
[lat.sh](https://lat.sh)

redirect:
[lat.sh/google](https://lat.sh/google)

The main difference between this url shortener and others are that it provides a set url query parameters to help thoughs using this application to know more information about the redirect link before click on it.

examples like:

redirect link: `lat.sh/hello`

by using the `o` flag
`lat.sh/hello?o=true` you are able to tell what is the destination of the link and how many has already clicked on the link before ever getting redirected.

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
