# Step 1: Use official Node.js image as the base image
FROM node:18 AS build

# Step 2: Set the working directory in the container
WORKDIR /app

# Step 3: Copy package.json and package-lock.json (or yarn.lock) into the container
COPY package*.json ./

# Step 4: Install dependencies
RUN npm install

# Step 5: Copy the rest of the application code
COPY . .

# Step 6: Build the project using Vite
RUN npm run build

# Step 7: Use a lightweight web server to serve the app
FROM nginx:alpine

# Step 8: Copy the build output from the build stage to Nginx's default HTML directory
COPY --from=build /app/dist /usr/share/nginx/html

# Step 9: Expose port 80 for the web server
EXPOSE 80

# Step 10: Start Nginx server
CMD ["nginx", "-g", "daemon off;"]

