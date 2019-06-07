FROM gcr.io/cloud-builders/gcloud:latest


RUN apt-get update && \
  apt-get install -f -y ca-certificates curl wget build-essential util-linux \
  gconf-service libasound2 libatk1.0-0 libatk-bridge2.0-0 libc6 libcairo2 \
  libcups2 libdbus-1-3 libexpat1 libfontconfig1 libgcc1 libgconf-2-4 \
  libgdk-pixbuf2.0-0 libglib2.0-0 libgtk-3-0 libnspr4 libpango-1.0-0 \
  libpangocairo-1.0-0 libstdc++6 libx11-6 libx11-xcb1 libxcb1 libxcomposite1 \
  libxcursor1 libxdamage1 libxext6 libxfixes3 libxi6 libxrandr2 libxrender1 \
  libxss1 libxtst6 fonts-liberation libappindicator1 libnss3 lsb-release \
  xdg-utils wget chromium-browser   
RUN curl -sL https://deb.nodesource.com/setup_12.x | bash -  
RUN apt-get install -y nodejs
RUN npm install http-server -g
RUN groupadd -r node && useradd -r -g node node
RUN sysctl -w kernel.unprivileged_userns_clone=1
# RUN chown -R node /builder/google-cloud-sdk


# RUN mkdir -p /builder/home/.config/gcloud
# RUN chown -R node:node /builder/home/.config/gcloud
# RUN chmod -R 777 /builder/home/.config/gcloud

RUN mkdir -p /workspace/
RUN chown -R node /workspace/
RUN chmod -R g+rwxs /workspace/

USER node

ENTRYPOINT ["make"]
