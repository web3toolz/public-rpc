FROM node:18-alpine AS builder

WORKDIR /home/app

COPY package.json yarn.lock* package-lock.json* pnpm-lock.yaml* /home/app/

RUN yarn install --frozen-lockfile --non-interactive

COPY . /home/app

ARG NEXT_PUBLIC_API_URL
ARG NEXT_PUBLIC_GTAG_ID
ENV NEXT_PUBLIC_API_URL $NEXT_PUBLIC_API_URL
ENV NEXT_PUBLIC_GTAG_ID $NEXT_PUBLIC_GTAG_ID

ENV NODE_ENV production

RUN yarn build

FROM nginx:alpine

COPY --from=builder /home/app/out /usr/share/nginx/html

RUN rm /etc/nginx/conf.d/default.conf

COPY nginx/nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]