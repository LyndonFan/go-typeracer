#!/bin/bash

page=1
totalPages="0"

baseURL="https://api.quotable.io/quotes?minLength=100"

# Looping checks with != as it's hard to do math in bash
while [ "$page" != $totalPages ]
do
    curl -s "$baseURL&page=$page" -o "response.json"
    totalPages=$(jq -r '.totalPages' "response.json")
    jq -r '.results' "response.json" > "page_$page.json"

    echo "Page $page of $totalPages"
    ((page++))
    
    # rate limit is 180/min = 3 requests per second
    sleep 0.4
done

page=1
touch quotes.txt
while [ "$page" != $totalPages ]
do
    jq -r '.[] | .content' "page_$page.json" >> "quotes.txt"
    ((page++))
done

# Cleanup
rm response.json page_*.json