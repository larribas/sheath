FROM scratch
EXPOSE 1827
COPY sheathe /
ENTRYPOINT ["/sheathe"]