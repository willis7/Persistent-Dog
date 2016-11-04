# VERSION:		  0.1
# DESCRIPTION:	Runs a simple docker microservice which publishes metrics
# USAGE:
#
# # Build image
# docker build -t heimdall .
#
# # Run container
# docker run -it --rm --name my-heimdall -p 80:80 heimdall
#
FROM golang:1.7.0-onbuild

EXPOSE 80
