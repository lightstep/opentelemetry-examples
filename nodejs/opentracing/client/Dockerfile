FROM node:12

# Create app directory
WORKDIR /app/ls-client

# Install app dependencies
COPY package*.json ./
RUN npm install

# Bundle app source
COPY . .

CMD [ "node", "index.js" ]
