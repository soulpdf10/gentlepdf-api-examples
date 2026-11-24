# 🚀 GentlePDF API — Multi-Language Integration Examples

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
[![API Status](https://img.shields.io/badge/API-Online%20(99.95%25)-emerald)](https://gentlepdf.com/api-docs)
[![Python](https://img.shields.io/badge/Python-3.8+-3776AB?logo=python&logoColor=white)](python/)
[![Node.js](https://img.shields.io/badge/Node.js-18+-339933?logo=node.js&logoColor=white)](nodejs/)
[![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go&logoColor=white)](go/)
[![PHP](https://img.shields.io/badge/PHP-8.0+-777BB4?logo=php&logoColor=white)](php/)
[![.NET](https://img.shields.io/badge/.NET-8.0+-512BD4?logo=dotnet&logoColor=white)](csharp/)
[![Java](https://img.shields.io/badge/Java-17+-007396?logo=openjdk&logoColor=white)](java/)

Official runnable examples demonstrating how to use the [GentlePDF REST API](https://gentlepdf.com/api-docs) to merge, compress, convert, OCR, split, and protect PDFs in 7 programming languages.

---

## ⚡ 60-Second Quickstart

### 1. Get your API Key
1. Sign up at [gentlepdf.com/register](https://gentlepdf.com/register) to get **100 free API credits**.
2. Copy your secret API key from the Developer Dashboard.
3. Export it in your environment:
   ```bash
   export GENTLEPDF_API_KEY="gpdf_live_your_actual_key_here"
   ```

### 2. Run an Example with cURL
```bash
curl -X POST "https://api.gentlepdf.com/merge-pdf" \
  -H "x-api-key: $GENTLEPDF_API_KEY" \
  -F "files=@document1.pdf" \
  -F "files=@document2.pdf" \
  --output merged.pdf
```

---

## 📂 Repository Structure

| Language | Directory | Description |
| :--- | :--- | :--- |
| **cURL / Shell** | [`curl/`](curl/) | Ready-to-run `.sh` Bash scripts for quick terminal automation |
| **Python** | [`python/`](python/) | Python examples using `requests` with streaming downloads |
| **Node.js / TS** | [`nodejs/`](nodejs/) | Modern ESM & TypeScript examples using native `fetch` & `FormData` |
| **Go** | [`go/`](go/) | Idiomatic Go code using standard library `net/http` & `mime/multipart` |
| **PHP** | [`php/`](php/) | Modern PHP 8+ using native `CURLFile` |
| **C# (.NET)** | [`csharp/`](csharp/) | .NET 8 console application with `HttpClient` & async streams |
| **Java** | [`java/`](java/) | Maven project using Java 17+ `java.net.http.HttpClient` |

---

## 🛠️ API Operations & Credit Costs

| Endpoint | Method | Credits | Description |
| :--- | :---: | :---: | :--- |
| `/merge-pdf` | `POST` | **1** | Merges 2+ PDFs into a single document |
| `/compress-pdf` | `POST` | **1** | Compresses PDF using screen, ebook, or printer presets |
| `/split-pdf` | `POST` | **1** | Extracts custom page ranges or returns a ZIP |
| `/pdf-to-word` | `POST` | **2** | Converts PDF into editable Microsoft Word (`.docx`) |
| `/word-to-pdf` | `POST` | **2** | Converts Word documents (`.docx`/`.doc`) into PDF |
| `/ocr-pdf` | `POST` | **3** | Scans image/scanned PDF into text or searchable PDF |
| `/watermark-pdf`| `POST` | **1** | Stamps customized watermark text or stamp |
| `/protect-pdf` | `POST` | **1** | 256-bit AES encryption with password protection |
| `/unlock-pdf`  | `POST` | **1** | Removes password and decryption restrictions |

> 💡 **Credits never expire!** Review our pay-as-you-go packs and usage calculator at [gentlepdf.com/api-pricing](https://gentlepdf.com/api-pricing).

---

## ⚠️ HTTP Status Codes & Error Handling

```json
{
  "detail": "File 'large_scan.pdf' (65 MB) exceeds the 50 MB limit for your plan."
}
```

* `200 OK` — Success. Processed binary file or JSON response returned.
* `400 Bad Request` — Invalid file format, missing parameters, or corrupted document.
* `401 Unauthorized` — Invalid or missing `x-api-key` header.
* `402 Payment Required` — Zero API credits remaining. [Refill credits here](https://gentlepdf.com/api-pricing).
* `413 Payload Too Large` — File size exceeds your current credit pack threshold.
* `429 Too Many Requests` — Rate limit exceeded. Backoff and retry.

---

## 📄 License
This repository is released under the [MIT License](LICENSE).
Feel free to copy, modify, and distribute these code snippets in your own commercial or open-source projects.
