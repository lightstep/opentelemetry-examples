##########################################
FROM golang:1.18-alpine as builder

ARG REPO
ARG BRANCH
ENV REPO_URL=https://github.com/${REPO}
ENV REPO_PATH=/go/src/github.com/${REPO}
ENV REPO_BRANCH=${BRANCH}

RUN mkdir -p ${REPO_PATH}
RUN apk add --no-cache git

RUN git clone -b ${REPO_BRANCH} --single-branch --depth 1 ${REPO_URL} ${REPO_PATH}

WORKDIR ${REPO_PATH}
RUN go mod tidy
# must pass flags for static or will not work on scratch
RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux go build -ldflags '-w -s' -ldflags '-extldflags "-lm -stdc++ -static"' -o /otelcontribcol ./cmd/otelcontribcol

FROM scratch 
# copy ca certs
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# copy bin
COPY --from=builder /otelcontribcol /otelcontribcol

COPY collector.yml /collector.yml
