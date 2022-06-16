FROM python:3.9.1-slim-buster

WORKDIR /usr/src

COPY ./requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt

COPY . /usr/src

CMD [ "python3", "./app/hello-world.py" ]