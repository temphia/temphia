#!/bin/bash

# Create Python virtual environment
python -m venv myenv

# Activate virtual environment
source myenv/bin/activate

# Install required pip packages if requirement.txt exists
if [ -f "requirement.txt" ]; then
    pip install -r requirement.txt
fi

# Run setup.sh if it exists
if [ -f "setup.sh" ]; then
    chmod +x setup.sh
    ./setup.sh

    # Check if setup.sh ran successfully
    if [ $? -eq 0 ]; then
        # Execute main.py
        python main.py
    else
        echo "setup.sh failed to run successfully."
    fi
else
    # Execute main.py if setup.sh does not exist
    python main.py
fi
