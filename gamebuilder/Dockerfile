FROM gcr.io/cloud-builders/gcloud:latest


RUN apt-get update && \
  apt-get install -y ca-certificates curl wget build-essential nodejs 
RUN curl -sL https://deb.nodesource.com/setup_12.x | bash -
RUN apt-get install -y nodejs gconf-service libasound2 libatk1.0-0 \
  libatk-bridge2.0-0 libc6 libcairo2 \
  libcups2 libdbus-1-3 libexpat1 libfontconfig1 libgcc1 libgconf-2-4 \
  libgdk-pixbuf2.0-0 libglib2.0-0 libgtk-3-0 libnspr4 libpango-1.0-0 \
  libpangocairo-1.0-0 libstdc++6 libx11-6 libx11-xcb1 libxcb1 libxcomposite1 \
  libxcursor1 libxdamage1 libxext6 libxfixes3 libxi6 libxrandr2 libxrender1 \
  libxss1 libxtst6 fonts-liberation libappindicator1 libnss3 lsb-release \
  xdg-utils chromium-browser   
RUN node -v
# Install pngquant
RUN apt-get install -y libpng-dev
RUN wget http://pngquant.org/pngquant-2.9.0-src.tar.gz
RUN tar xvf pngquant-2.9.0-src.tar.gz
RUN cd pngquant-2.9.0 && ./configure
RUN cd pngquant-2.9.0 && make all install

ENTRYPOINT ["make"]
