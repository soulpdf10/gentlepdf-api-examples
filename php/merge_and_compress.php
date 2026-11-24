<?php
/**
 * GentlePDF PHP API Client Examples
 * Merge and Compress PDF files with native PHP cURL.
 */

$apiKey = getenv('GENTLEPDF_API_KEY') ?: 'YOUR_API_KEY';
$baseURL = getenv('GENTLEPDF_BASE_URL') ?: 'https://api.gentlepdf.com';

/**
 * Merge two or more PDF files
 */
function mergePDFs(array $files, string $outputPath = 'merged.pdf'): bool {
    global $apiKey, $baseURL;
    echo "[*] Merging " . count($files) . " files with PHP...\n";

    $postData = [];
    foreach ($files as $i => $filePath) {
        if (!file_exists($filePath)) {
            echo "[-] Error: file not found: $filePath\n";
            return false;
        }
        $postData["files[$i]"] = new CURLFile($filePath, 'application/pdf');
    }

    $ch = curl_init("$baseURL/merge-pdf");
    curl_setopt_array($ch, [
        CURLOPT_POST => true,
        CURLOPT_POSTFIELDS => $postData,
        CURLOPT_HTTPHEADER => [
            "x-api-key: $apiKey"
        ],
        CURLOPT_RETURNTRANSFER => true,
        CURLOPT_SSL_VERIFYPEER => true,
    ]);

    $response = curl_exec($ch);
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    curl_close($ch);

    if ($httpCode === 200) {
        file_put_contents($outputPath, $response);
        echo "[+] Success! Merged PDF saved to $outputPath\n";
        return true;
    } else {
        echo "[-] Error $httpCode: $response\n";
        return false;
    }
}

/**
 * Compress a PDF file
 */
function compressPDF(string $filePath, string $quality = 'ebook', string $outputPath = 'compressed.pdf'): bool {
    global $apiKey, $baseURL;
    echo "[*] Compressing $filePath (quality: $quality)...\n";

    if (!file_exists($filePath)) {
        echo "[-] Error: file not found: $filePath\n";
        return false;
    }

    $postData = [
        'file' => new CURLFile($filePath, 'application/pdf'),
        'quality' => $quality,
    ];

    $ch = curl_init("$baseURL/compress-pdf");
    curl_setopt_array($ch, [
        CURLOPT_POST => true,
        CURLOPT_POSTFIELDS => $postData,
        CURLOPT_HTTPHEADER => [
            "x-api-key: $apiKey"
        ],
        CURLOPT_RETURNTRANSFER => true,
    ]);

    $response = curl_exec($ch);
    $httpCode = curl_getinfo($ch, CURLINFO_HTTP_CODE);
    curl_close($ch);

    if ($httpCode === 200) {
        file_put_contents($outputPath, $response);
        echo "[+] Success! Compressed PDF saved to $outputPath\n";
        return true;
    } else {
        echo "[-] Error $httpCode: $response\n";
        return false;
    }
}

echo "GentlePDF PHP Examples Loaded.\n";
