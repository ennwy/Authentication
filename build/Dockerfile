FROM golang:1.19 as build

ENV BIN_FILE /opt/auth/tests
ENV CODE_DIR /go/src/

WORKDIR ${CODE_DIR}

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY . ${CODE_DIR}

ARG LDFLAGS

RUN CGO_ENABLED=0 go build \
        -o ${BIN_FILE} ./cmd/...

FROM alpine:3.9

ENV BIN_FILE /opt/auth/tests
ENV CODE_DIR /go/src/

COPY --from=build ${BIN_FILE} ${BIN_FILE}
#COPY --from=build ${CODE_DIR}/templates templates
 
CMD ${BIN_FILE}