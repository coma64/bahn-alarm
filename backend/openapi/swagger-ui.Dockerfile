FROM swaggerapi/swagger-ui:v5.7.2

RUN apk add yq

COPY openapi.yml /app/openapi.yml

RUN cat /app/openapi.yml | yq -p yaml -o json > /app/swagger.json

