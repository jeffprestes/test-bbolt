#env GOOS=linux GOARCH=amd64 go build
#docker build -t mercurius:test-bbolt .
#docker run -p 8080:8080 -d mercurius:test-bbolt

FROM scratch

ADD test-bbolt /
ADD conf/ /conf
ADD public/ /public
ADD locale/ /locale

CMD [ "/test-bbolt" ]