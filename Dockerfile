# VERSION:		  0.1
# DESCRIPTION:	Runs a simple docker microservice
# USAGE:
#
# # Build image
# docker build -t dog .
#
# # Run container
# docker run -it --rm --name persistent-dog -p 80:80 dog
#
FROM golang:1.7.0-onbuild

EXPOSE 80
