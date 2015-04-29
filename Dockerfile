FROM scratch
EXPOSE 1827
COPY sheath /
ENTRYPOINT ["/sheath"]