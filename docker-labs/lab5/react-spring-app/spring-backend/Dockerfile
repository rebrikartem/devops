FROM alpine:3.21 AS system

RUN apk add --update --no-cache openjdk17

WORKDIR /app

FROM system AS build

COPY .mvn ./.mvn
COPY src ./src
COPY mvnw pom.xml ./

RUN sed -i 's/\r$//' mvnw && chmod +x mvnw && \
    ./mvnw compile && \
    ./mvnw package -DskipTests

FROM system AS production

COPY --from=build /app/target/*.jar ./spring-backend.jar
COPY --from=build /app/target/classes/application.properties ./classes/application.properties

RUN adduser -D -g 'nonroot' nonroot && chown -R nonroot:nonroot /usr/lib/jvm/java-17-openjdk /app/

USER nonroot

CMD ["java", "-jar", "/app/spring-backend.jar"]