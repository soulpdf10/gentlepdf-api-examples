using System;
using System.IO;
using System.Net.Http;
using System.Threading.Tasks;

namespace GentlePdfExamples
{
    class Program
    {
        private static readonly string ApiKey = Environment.GetEnvironmentVariable("GENTLEPDF_API_KEY") ?? "YOUR_API_KEY";
        private static readonly string BaseUrl = Environment.GetEnvironmentVariable("GENTLEPDF_BASE_URL") ?? "https://api.gentlepdf.com";

        static async Task Main(string[] args)
        {
            Console.WriteLine("GentlePDF C# .NET API Examples");
            Console.WriteLine("Ensure your GENTLEPDF_API_KEY environment variable is configured.");
        }

        public static async Task MergePdfsAsync(string[] filePaths, string outputPath = "merged.pdf")
        {
            Console.WriteLine($"[*] Merging {filePaths.Length} files in C#...");

            using var httpClient = new HttpClient();
            httpClient.DefaultRequestHeaders.Add("x-api-key", ApiKey);

            using var form = new MultipartFormDataContent();
            foreach (var path in filePaths)
            {
                var fileBytes = await File.ReadAllBytesAsync(path);
                form.Add(new ByteArrayContent(fileBytes), "files", Path.GetFileName(path));
            }

            var response = await httpClient.PostAsync($"{BaseUrl}/merge-pdf", form);
            if (response.IsSuccessStatusCode)
            {
                var content = await response.Content.ReadAsByteArrayAsync();
                await File.WriteAllBytesAsync(outputPath, content);
                Console.WriteLine($"[+] Successfully merged into '{outputPath}' ({content.Length} bytes)");
            }
            else
            {
                var errorText = await response.Content.ReadAsStringAsync();
                Console.WriteLine($"[-] Error {(int)response.StatusCode}: {errorText}");
            }
        }

        public static async Task CompressPdfAsync(string inputPath, string quality = "ebook", string outputPath = "compressed.pdf")
        {
            Console.WriteLine($"[*] Compressing '{inputPath}' in C#...");

            using var httpClient = new HttpClient();
            httpClient.DefaultRequestHeaders.Add("x-api-key", ApiKey);

            using var form = new MultipartFormDataContent();
            var fileBytes = await File.ReadAllBytesAsync(inputPath);
            form.Add(new ByteArrayContent(fileBytes), "file", Path.GetFileName(inputPath));
            form.Add(new StringContent(quality), "quality");

            var response = await httpClient.PostAsync($"{BaseUrl}/compress-pdf", form);
            if (response.IsSuccessStatusCode)
            {
                var content = await response.Content.ReadAsByteArrayAsync();
                await File.WriteAllBytesAsync(outputPath, content);
                Console.WriteLine($"[+] Compressed PDF saved to '{outputPath}' ({content.Length} bytes)");
            }
            else
            {
                var errorText = await response.Content.ReadAsStringAsync();
                Console.WriteLine($"[-] Error {(int)response.StatusCode}: {errorText}");
            }
        }
    }
}
