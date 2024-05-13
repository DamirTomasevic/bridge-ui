#!/bin/bash

# Define the list of contracts to inspect
contracts=(
    "TokenUnlocking"
)

# Empty the output file initially
output_file="contract_layout.md"
> $output_file

# Loop over each contract
for contract in "${contracts[@]}"; do
    # Run forge inspect and append to the file
    # Ensure correct concatenation of the command without commas
    echo "forge inspect ${contract} storagelayout --pretty >> $output_file"

    echo "## ${contract}" >> $output_file
    forge inspect ${contract} storagelayout --pretty >> $output_file
    echo "" >> $output_file
done
