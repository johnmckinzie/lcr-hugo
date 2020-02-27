FROM gitpod/workspace-full

# Install custom tools, runtime, etc.
RUN sudo apt-get update \
    && sudo apt-get install -y \
        hugo \
        mysql-server \
        mysql-client \
    && sudo rm -rf /var/lib/apt/lists/* # This is needed because the root user owns these
