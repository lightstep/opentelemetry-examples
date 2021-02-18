FROM node:12

# Create app directory
WORKDIR /app/otel-client-ottrace

# Install app dependencies
COPY package*.json ./
RUN npm install

# Bundle app source
COPY . .

CMD [ "node", "index.js" ]
