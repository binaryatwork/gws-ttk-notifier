FROM alpine
RUN  apk update && apk upgrade && apk add ca-certificates
RUN  mkdir -p /usr/local/ttk
ADD  notifier /usr/local/ttk/
RUN  chmod +x /usr/local/ttk/notifier
CMD  /usr/local/ttk/notifier