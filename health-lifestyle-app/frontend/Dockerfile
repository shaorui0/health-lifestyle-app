# Use Node.js Alpine image
FROM node:16-alpine

WORKDIR /app

# Copy package.json and package-lock.json
COPY package.json package-lock.json ./

# Install dependencies
RUN npm install

# Copy the rest of the frontend code
COPY . .

# Set environment variables
ENV VITE_API_URL=http://localhost:8000

# Expose port 3000
EXPOSE 3000

# Start the development server
CMD ["npm", "run", "dev"]
