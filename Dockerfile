FROM balenalib/raspberrypi3-debian
RUN install_packages netcat bash
COPY dispatcher /
COPY run.sh /
RUN chmod +x run.sh
ENTRYPOINT ["./run.sh"]
