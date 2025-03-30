FROM alpine:3.21 AS system

RUN apk add --update --no-cache nodejs npm && \
    npm install -g serve

WORKDIR /app

FROM system AS builder

COPY package.json package-lock.json ./
COPY ./src ./src
COPY ./public ./public

ARG REACT_APP_API_HOST

ENV REACT_APP_API_HOST=${REACT_APP_API_HOST}

RUN npm install --force

RUN REACT_APP_API_HOST=${REACT_APP_API_HOST} npm run build

FROM system AS production

COPY --from=builder /app/build ./build

RUN adduser -D -g 'nonroot' nonroot && chown -R nonroot:nonroot /app/

USER nonroot

CMD ["sh", "-c", "serve -s build -l $FRONT_PORT"]