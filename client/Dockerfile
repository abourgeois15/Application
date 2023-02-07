FROM node:13.12.0-alpine

WORKDIR /app

COPY node_modules ./
COPY package.json ./
COPY package-lock.json ./

RUN npm config set proxy http://localhost:3129
RUN npm config set https-proxy http://localhost:3129
RUN npm config set registry http://registry.npmjs.org/

# RUN npm install

COPY . ./

CMD ["npm", "start"]