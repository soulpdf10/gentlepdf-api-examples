package com.gentlepdf.examples;

import java.io.File;
import java.io.FileOutputStream;
import java.io.InputStream;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.file.Files;
import java.nio.file.Path;

/**
 * GentlePDF Java API Examples
 * Demonstrates HTTP multipart requests using modern Java 17+ HttpClient.
 */
public class GentlePdfClient {

    private static final String API_KEY = System.getenv().getOrDefault("GENTLEPDF_API_KEY", "YOUR_API_KEY");
    private static final String BASE_URL = System.getenv().getOrDefault("GENTLEPDF_BASE_URL", "https://api.gentlepdf.com");

    public static void main(String[] args) {
        System.out.println("GentlePDF Java 17+ API Examples Loaded.");
        System.out.println("Base URL: " + BASE_URL);
    }
}
