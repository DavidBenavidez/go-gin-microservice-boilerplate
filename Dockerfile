FROM scratch
WORKDIR /root/
COPY app ./
CMD ["./app"]
EXPOSE 8080