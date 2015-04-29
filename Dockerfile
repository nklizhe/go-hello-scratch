FROM scratch
ADD bin/linux/hello /usr/bin/
EXPOSE 3000
ENTRYPOINT [ "/usr/bin/hello" ]

