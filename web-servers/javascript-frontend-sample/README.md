# Getting Started with Create React App

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Deploy and View
Build this app into an OCI image using the Paketo Web Servers buildpack.  There
are two [project
descriptor](https://buildpacks.io/docs/app-developer-guide/using-project-descriptor/)
files in this directory, `httpd.toml` and `nginx.toml`. Each contain the
configuration necessary to build the React source code into static assets and
serve those assets with the selected server.

The HTTPD and NGINX buildpacks will automatically generate a default
configuration file for the server.

### Build with `pack` | HTTPD
```bash
pack build react-sample --descriptor httpd.toml
```

### Build with `pack` | NGINX
```bash
pack build react-sample --descriptor nginx.toml
```

### Run the app
```
docker run  -it -p 8080:8080 --env PORT=8080  react-sample
```

### View the app
Visit `localhost:8080` in your browser. Since push-state routing is enabled in
this build with `BP_WEB_SERVER_ENABLE_PUSH_STATE`, visiting
`localhost:8080/foo` – or any other endpoint – will serve the same page.

## Available Scripts

In the project directory, you can run:

### `npm start`

Runs the app in the development mode.\
Open [http://localhost:3000](http://localhost:3000) to view it in your browser.

The page will reload when you make changes.\
You may also see any lint errors in the console.

### `npm run build`

Builds the app for production to the `build` folder.\
It correctly bundles React in production mode and optimizes the build for the best performance.

The build is minified and the filenames include the hashes.\
Your app is ready to be deployed!

See the section about
[deployment](https://facebook.github.io/create-react-app/docs/deployment) for
more information.

