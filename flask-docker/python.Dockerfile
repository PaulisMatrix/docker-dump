# using a python docker base image
# using a multi-stage build to save space
FROM python:3.10-slim AS builder-image

# create a venv to make copying packages folder easier b/w the different stages
# declaring the path to avoid path issues w/ packages
RUN python3.10 -m venv /home/myuser/venv 
ENV PATH="/home/myuser/venv/bin:$PATH"

# copying and installing requirements first will avoid a re-install of all dependencies
# when re-building the image
COPY requirements.txt .
RUN pip3 install --no-cache-dir wheel 
RUN pip3 install --no-cache-dir -r requirements.txt 


# now starting the next layer, the one that runs the program
# FROM ubuntu:22.04 AS runner-image 
FROM python:3.10-slim AS runner-image

# creating a separate user to run the docker commands, helps with security as root 
# access is denied to this user
RUN useradd --create-home myuser

# copy over the virtualenv created in the previous layer to the current user's dir
COPY --from=builder-image /home/myuser/venv /home/myuser/venv

# create and set workdir, copy over all data to the folder 
USER myuser
RUN mkdir /home/myuser/code 
WORKDIR /home/myuser/code
COPY . . 

EXPOSE 7777 

# using this ensures that all messages are printed out on the terminal
ENV PYTHONUNBUFFERED=1

# activate the virtualenv
ENV VIRTUAL_ENV=/home/myuser/venv
ENV PATH="/home/myuser/venv/bin:$PATH"

# using /dev/shm as the worker temp dir will help prevent random locks and freezes 
# used for gunicorn heartbeat
CMD [ "gunicorn", "-b", "0.0.0.0:7777", "-w", "4", "-k", "uvicorn.workers.UvicornWorker", "--worker-tmp-dir", "/dev/shm", "main:app" ]

# alternatively: 
# CMD [ "uvicorn", "--host", "0.0.0.0","--port", "7777", "main:app" ]