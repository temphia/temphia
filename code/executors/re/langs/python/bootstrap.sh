#!/bin/bash

# Activate virtual environment
source ./bin/activate

# Install required pip packages if requirement.txt exists
if [ -f "requirement.txt" ]; then
    pip install -r requirement.txt
fi

# Run setup.sh if it exists
if [ -f "setup.sh" ]; then
    chmod +x setup.sh
    ./setup.sh
fi
