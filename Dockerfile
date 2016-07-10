FROM busybox
COPY ca-certificates.crt /etc/ssl/certs/
COPY app /
COPY creds.json creds.json
COPY public /public
EXPOSE 5000
CMD ["/app"]
