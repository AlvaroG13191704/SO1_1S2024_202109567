FROM node:lts-alpine as builder

WORKDIR /app

COPY . .

RUN yarn install \
  --prefer-offline \
  --pure-lockfile \
  --non-interactive \
  && yarn build \
  && rm -rf node_modules \
  && NODE_ENV=production yarn install \
  --prefer-offline \
  --pure-lockfile \
  --non-interactive \
  --production=true

FROM node:lts-alpine

WORKDIR /app

COPY --from=builder /app  .

ENV HOST 0.0.0.0
EXPOSE 3000

CMD [ "yarn", "start" ]

