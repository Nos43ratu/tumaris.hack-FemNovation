FROM node:18-alpine AS deps

FROM deps as dependencies
COPY . .
RUN yarn --frozen-lockfile
RUN yarn build

FROM nginx:alpine
COPY --from=dependencies /dist /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY --from=dependencies /nginx/nginx.conf /etc/nginx/conf.d

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
