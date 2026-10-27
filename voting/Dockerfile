FROM maven:3.9-eclipse-temurin-17-alpine AS builder

WORKDIR /app

COPY pom.xml ./
RUN mvn dependency:go-offline

COPY src ./src

RUN mvn clean package

FROM eclipse-temurin:17-jre AS runner

RUN groupadd --system app && useradd --system --gid app app

WORKDIR /app

COPY --from=builder --chown=app:app /app/target/*.jar ./voting.jar

EXPOSE 8080

USER app

CMD ["java", "-jar", "voting.jar"]