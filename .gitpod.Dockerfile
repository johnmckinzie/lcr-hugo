FROM gitpod/workspace-full
                    
USER gitpod

RUN sudo apt-get -q update \
    && sudo apt-get install -yq hugo \
    && sudo rm -rf /var/lib/apt/lists/*
