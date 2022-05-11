# Stage 1: Compile and Build angular codebase

# Use official node image as the base image
FROM node:latest as build

# Set the working directory
WORKDIR /app

# Add the source code to app
COPY ./angular.json .
COPY ./jest.config.js .
COPY ./package.json .
COPY ./package-lock.json .
COPY ./tsconfig.app.json .
COPY ./tsconfig.json .

# Install all the dependencies
RUN npm install --save-dev

# Copy over the rest of the source files
COPY ./conduit-web ./conduit-web

# Generate the build of the application
RUN npm run build --prod

# Stage 2: Serve app with nginx server

# Use official nginx image as the base image
FROM nginx:latest

# Copy the build output to replace the default nginx contents.
COPY --from=build /app/dist/conduit-web /usr/share/nginx/html

EXPOSE 80
