/**
 * GentlePDF Node.js API Examples
 * Demonstrates merging, compressing, converting, and OCR processing.
 */

import fs from 'fs';
import FormData from 'form-data';
import fetch from 'node-fetch';

const API_KEY = process.env.GENTLEPDF_API_KEY || 'YOUR_API_KEY';
const BASE_URL = process.env.GENTLEPDF_BASE_URL || 'https://api.gentlepdf.com';

/**
 * Merge 2 or more PDF documents into one
 */
export async function mergePDFs(filePaths, outputPath = 'merged.pdf') {
  console.log(`[*] Merging ${filePaths.length} PDFs...`);
  const form = new FormData();

  for (const fp of filePaths) {
    form.append('files', fs.createReadStream(fp));
  }

  const response = await fetch(`${BASE_URL}/merge-pdf`, {
    method: 'POST',
    headers: {
      'x-api-key': API_KEY,
      ...form.getHeaders(),
    },
    body: form,
  });

  if (response.ok) {
    const buffer = await response.buffer();
    fs.writeFileSync(outputPath, buffer);
    console.log(`[+] Merged PDF saved to ${outputPath} (${buffer.length} bytes)`);
  } else {
    console.error(`[-] Error ${response.status}:`, await response.text());
  }
}

/**
 * Compress a PDF document
 */
export async function compressPDF(filePath, quality = 'ebook', outputPath = 'compressed.pdf') {
  console.log(`[*] Compressing ${filePath} with quality '${quality}'...`);
  const form = new FormData();
  form.append('file', fs.createReadStream(filePath));
  form.append('quality', quality);

  const response = await fetch(`${BASE_URL}/compress-pdf`, {
    method: 'POST',
    headers: {
      'x-api-key': API_KEY,
      ...form.getHeaders(),
    },
    body: form,
  });

  if (response.ok) {
    const buffer = await response.buffer();
    fs.writeFileSync(outputPath, buffer);
    console.log(`[+] Compressed PDF saved to ${outputPath} (${buffer.length} bytes)`);
  } else {
    console.error(`[-] Error ${response.status}:`, await response.text());
  }
}

/**
 * Convert PDF to editable Microsoft Word document (.docx)
 */
export async function pdfToWord(filePath, outputPath = 'converted.docx') {
  console.log(`[*] Converting ${filePath} to Word DOCX...`);
  const form = new FormData();
  form.append('file', fs.createReadStream(filePath));

  const response = await fetch(`${BASE_URL}/pdf-to-word`, {
    method: 'POST',
    headers: {
      'x-api-key': API_KEY,
      ...form.getHeaders(),
    },
    body: form,
  });

  if (response.ok) {
    const buffer = await response.buffer();
    fs.writeFileSync(outputPath, buffer);
    console.log(`[+] Converted DOCX saved to ${outputPath}`);
  } else {
    console.error(`[-] Error ${response.status}:`, await response.text());
  }
}

console.log('GentlePDF Node.js SDK Examples Loaded.');
