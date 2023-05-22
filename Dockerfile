FROM node:alpine

ENV NODE_ENV=production
ENV HOST 0.0.0.0

WORKDIR /sharecvweb

COPY ./.nuxt ./.nuxt
COPY ./.output ./output
COPY ./package.json ./package.json

EXPOSE 3000

CMD ["node", "output/server/index.mjs"]
