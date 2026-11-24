"""
GentlePDF Python Examples
Merge, Compress, Convert, and OCR PDFs using the GentlePDF REST API.
"""

import os
import requests

API_KEY = os.getenv("GENTLEPDF_API_KEY", "YOUR_API_KEY")
BASE_URL = os.getenv("GENTLEPDF_BASE_URL", "https://api.gentlepdf.com")

HEADERS = {
    "x-api-key": API_KEY
}


def merge_pdfs(file_paths: list[str], output_filename: str = "merged.pdf"):
    """Merge 2 or more PDF documents into one."""
    print(f"[*] Merging {len(file_paths)} files...")
    url = f"{BASE_URL}/merge-pdf"

    # Prepare multipart files list
    files = [("files", open(fp, "rb")) for fp in file_paths]

    try:
        response = requests.post(url, headers=HEADERS, files=files)
        if response.status_code == 200:
            with open(output_filename, "wb") as f:
                f.write(response.content)
            print(f"[+] Successfully merged into '{output_filename}' ({len(response.content)} bytes)")
        else:
            print(f"[-] Error {response.status_code}: {response.text}")
    finally:
        for _, f in files:
            f.close()


def compress_pdf(input_path: str, quality: str = "ebook", output_filename: str = "compressed.pdf"):
    """
    Compress a PDF document.
    Quality presets: 'screen' (72 dpi), 'ebook' (150 dpi), 'printer' (300 dpi)
    """
    print(f"[*] Compressing '{input_path}' with quality '{quality}'...")
    url = f"{BASE_URL}/compress-pdf"

    with open(input_path, "rb") as f:
        files = {"file": f}
        data = {"quality": quality}
        response = requests.post(url, headers=HEADERS, files=files, data=data)

    if response.status_code == 200:
        with open(output_filename, "wb") as out:
            out.write(response.content)
        print(f"[+] Compressed PDF saved to '{output_filename}' ({len(response.content)} bytes)")
    else:
        print(f"[-] Error {response.status_code}: {response.text}")


def pdf_to_word(input_path: str, output_filename: str = "output.docx"):
    """Convert a PDF into an editable Word (.docx) document."""
    print(f"[*] Converting '{input_path}' to Word DOCX...")
    url = f"{BASE_URL}/pdf-to-word"

    with open(input_path, "rb") as f:
        files = {"file": f}
        response = requests.post(url, headers=HEADERS, files=files)

    if response.status_code == 200:
        with open(output_filename, "wb") as out:
            out.write(response.content)
        print(f"[+] Converted DOCX saved to '{output_filename}'")
    else:
        print(f"[-] Error {response.status_code}: {response.text}")


def ocr_pdf(input_path: str, language: str = "eng"):
    """Perform OCR optical character recognition to extract text."""
    print(f"[*] Running OCR on '{input_path}' (lang: {language})...")
    url = f"{BASE_URL}/ocr-pdf"

    with open(input_path, "rb") as f:
        files = {"file": f}
        data = {"language": language}
        response = requests.post(url, headers=HEADERS, files=files, data=data)

    if response.status_code == 200:
        result = response.json()
        print("[+] OCR Result:")
        print(result.get("text", result))
    else:
        print(f"[-] Error {response.status_code}: {response.text}")


if __name__ == "__main__":
    print("GentlePDF API Python Examples Loaded.")
    print("Set your GENTLEPDF_API_KEY environment variable to test live endpoints.")
