# WorDetect
### WordPress URL Detector

### Overview

This Go script is designed to efficiently check a list of URLs to determine if they are WordPress sites. It specifically checks for redirects to `/wp-content` and `/blog/wp-content` endpoints. Utilizing goroutines for concurrent processing, the script offers a fast and reliable way to identify WordPress URLs.

### Requirements

Go (Version 1.21 or later recommended)

### Installation

No additional libraries are required. Simply compile the Go script with the following command:

```bash
go install github.com/phor3nsic/wordetect@latest
```

### Usage

The script can be executed in two ways:

From a File: To check multiple URLs, you can pass a file with a list of URLs:

```bash
cat urls.txt | wordetect
```
Ensure urls.txt contains one URL per line.

Single URL Check: To check a single URL, use:
```bash
echo "https://example.com" | wordetect
```

### How It Works

The script reads URLs from the standard input, then uses goroutines to concurrently make HTTP requests to each URL appended with /wp-content and /blog/wp-content. It checks if the HTTP response indicates a redirection and if the redirection URL contains the specific WordPress endpoint. If both conditions are met, the script considers the URL as a WordPress site and prints it to the standard output.

> Note
>
>The script is a basic implementation and might need adjustments based on specific use cases or different WordPress configurations.
>Network reliability and the specific configuration of the target WordPress sites can affect the script's accuracy.
>Contributing

>Contributions to improve the script or suggestions for additional features are welcome.

### License

MIT

This version of the README provides the same comprehensive guide as before but without any markdown formatting.