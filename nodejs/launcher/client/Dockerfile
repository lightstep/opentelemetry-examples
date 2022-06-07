FROM node:14-alpine3.15

# Create app directory
WORKDIR /usr/src/client

# Install app dependencies
COPY package*.json ./
RUN npm install

# Bundle app source
COPY . .

CMD [ "node", "client.js" ]
