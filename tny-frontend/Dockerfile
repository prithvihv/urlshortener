FROM node:lts
WORKDIR /app
COPY . .
RUN yarn install
RUN yarn build
RUN yarn global add serve
CMD ["serve" , "-l", "8080", "-s","./dist"]