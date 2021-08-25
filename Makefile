SHELL := $(shell which bash)

export PROJECT = k8s-helloworld

devbox:
	@$(SHELL) .dev/devbox.sh

