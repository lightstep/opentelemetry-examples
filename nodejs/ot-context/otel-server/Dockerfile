FROM node:12

# Create app directory
WORKDIR /usr/src/otel-server

# Install app dependencies
COPY package*.json ./
RUN npm install

# Bundle app source
COPY . .

CMD [ "npm run start" ]
