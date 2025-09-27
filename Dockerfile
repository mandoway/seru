FROM eclipse-temurin:25 AS jre-build

RUN $JAVA_HOME/bin/jlink \
         --add-modules java.base,java.logging \
         --strip-debug \
         --no-man-pages \
         --no-header-files \
         --compress=2 \
         --output /javaruntime


FROM golang:1.23.2-bookworm AS build
ENV CUE_VERSION=v0.12.1

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go generate ./...
RUN CGO_ENABLED=1 GOOS=linux go build -o seru
RUN go install cuelang.org/go/cmd/cue@${CUE_VERSION}

FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends \
      ca-certificates \
      git \
    && rm -rf /var/lib/apt/lists/*

ENV JAVA_HOME=/opt/java/openjdk
ENV PATH="${JAVA_HOME}/bin:${PATH}"
COPY --from=jre-build /javaruntime $JAVA_HOME

COPY --from=build /usr/local/go /usr/local/go
COPY --from=build /go/bin/cue /usr/local/bin/cue
ENV PATH="/usr/local/go/bin:${PATH}"

WORKDIR /app
COPY --from=build /app/seru .
COPY --from=build /app/cue.so .
COPY --from=build /app/cue.jar .
COPY --from=build /app/perses_deploy.jar .

ENTRYPOINT ["./seru"]