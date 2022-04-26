# Getting Started with Create React App

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## Deploy and View
Build this app into an OCI image using the Paketo Node.js and NGINX buildpacks.
The
[`project.toml`](https://buildpacks.io/docs/app-developer-guide/using-project-descriptor/)
in this directory contains all of the required configuration to bulid the React
source into static assets, then serve those assets with NGINX.

The NGINX buildpack will automatically generate a default `nginx.conf` to
configure NGINX to serve the app.

### Build with `pack`
```bash
pack build react-nginx
```

### Run the app
```
docker run  -it -p 8080:8080 --env PORT=8080  react-nginx
```

### View the app
Visit `localhost:8080` in your browser. Since push-state routing is enabled in
this build with `BP_WEB_SERVER_ENABLE_PUSH_STATE`, visiting
`localhost:8080/foo`, or any other endpoint, will serve the same page.

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

