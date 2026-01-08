FROM node:alpine AS builder
WORKDIR /app
COPY . .
RUN npm install
EXPOSE 3000

#CMD ["npm", "start"]
RUN npm run build
FROM nginx:alpine
RUN rm /etc/nginx/conf.d/default.conf
COPY --from=builder /app/build /usr/share/nginx/html
COPY conf/nginx.conf  /etc/nginx/conf.d/default.conf

EXPOSE 3000
CMD ["nginx", "-g", "daemon off;"]
