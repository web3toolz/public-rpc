FROM node:18-alpine AS builder

WORKDIR /home/app

COPY package.json yarn.lock* package-lock.json* pnpm-lock.yaml* /home/app/

RUN yarn install --frozen-lockfile --non-interactive

COPY . /home/app

ARG GATSBY_API_URL
ARG GATSBY_GTAG_ID
ENV GATSBY_API_URL $GATSBY_API_URL
ENV GATSBY_GTAG_ID $GATSBY_GTAG_ID
ENV NODE_ENV production

RUN yarn build

FROM nginx:alpine

COPY --from=builder /home/app/public /usr/share/nginx/html

RUN rm /etc/nginx/conf.d/default.conf

COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]