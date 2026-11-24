#!/usr/bin/env bash
# GentlePDF Merge PDF Example (cURL)

API_KEY="${GENTLEPDF_API_KEY:-YOUR_API_KEY}"
BASE_URL="${GENTLEPDF_BASE_URL:-https://api.gentlepdf.com}"

if [ -z "$1" ] || [ -z "$2" ]; then
  echo "Usage: ./merge.sh <file1.pdf> <file2.pdf> [output.pdf]"
  exit 1
fi

FILE1="$1"
FILE2="$2"
OUTPUT="${3:-merged_output.pdf}"

echo "[*] Merging '$FILE1' and '$FILE2'..."

curl -s -X POST "$BASE_URL/merge-pdf" \
  -H "x-api-key: $API_KEY" \
  -F "files=@$FILE1" \
  -F "files=@$FILE2" \
  --output "$OUTPUT"

if [ -f "$OUTPUT" ] && [ -s "$OUTPUT" ]; then
  echo "[+] Success! Output saved to $OUTPUT"
else
  echo "[-] Merge failed. Check your API key and input files."
fi
