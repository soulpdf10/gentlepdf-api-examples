#!/usr/bin/env bash
# GentlePDF Compress PDF Example (cURL)

API_KEY="${GENTLEPDF_API_KEY:-YOUR_API_KEY}"
BASE_URL="${GENTLEPDF_BASE_URL:-https://api.gentlepdf.com}"

if [ -z "$1" ]; then
  echo "Usage: ./compress.sh <input.pdf> [quality] [output.pdf]"
  echo "Quality options: screen (max compression), ebook (recommended), printer"
  exit 1
fi

INPUT="$1"
QUALITY="${2:-ebook}"
OUTPUT="${3:-compressed_output.pdf}"

echo "[*] Compressing '$INPUT' with quality '$QUALITY'..."

curl -s -X POST "$BASE_URL/compress-pdf" \
  -H "x-api-key: $API_KEY" \
  -F "file=@$INPUT" \
  -F "quality=$QUALITY" \
  --output "$OUTPUT"

if [ -f "$OUTPUT" ] && [ -s "$OUTPUT" ]; then
  echo "[+] Success! Compressed PDF saved to $OUTPUT"
else
  echo "[-] Compression failed. Verify API key and document integrity."
fi
