FROM gitpod/workspace-full

USER gitpod # This is so all files created here will still be accessible

# Install custom tools, runtime, etc.
RUN sudo apt-get update \
    && sudo apt-get install -y \
        hugo \
    && sudo rm -rf /var/lib/apt/lists/* # This is needed because the root user owns these
