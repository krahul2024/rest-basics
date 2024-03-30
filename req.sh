#!/bin/bash


created_at=$(date -u +"%Y-%m-%dT%H:%M:%SZ")
updated_at=$(date -u +"%Y-%m-%dT%H:%M:%SZ")


curl -v -X PUT \
    -H "Content-Type: application/json" \
    -d '{
        "Name":        "Software",
        "Description": "Highly efficient software!",
        "Price":       8999.99,
        "InternalId":  "sw-001",
        "CreatedAt":   "'"$created_at"'",
        "UpdatedAt":   "'"$updated_at"'"
    }' \
    localhost:3300/products/34
