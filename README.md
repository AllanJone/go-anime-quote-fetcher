# Anime Quote Fetcher

This is a simple serverless Go application deployed on Vercel that fetches a random anime quote from `waifu.it`.

## Structure

- `api/quote.go`: This is a serverless function that fetches and responds with a random anime quote.
- `index.html`: The frontend HTML file that displays the fetched quote and provides an option to regenerate and fetch a new one.

## Setup

1. **Clone the Repository**:

   ```bash
   git clone <repository-url>
   cd <repository-directory>
   ```

2. **Environment Variable**:
   You need to set the `AUTH_TOKEN` environment variable for the authentication token required by the `waifu.it` API. This can be done in the Vercel dashboard or by using a `.env` file during local development.

   > **Note**: The `AUTH_TOKEN` is crucial for the application to communicate with the `waifu.it` API.

3. **Deploying on Vercel**:
   ```bash
   vercel
   ```

## Access

After deploying, you can directly access the frontend at:

```
https://your-project-name.vercel.app/
```

The API endpoint (serverless function) will be available at:

```
https://your-project-name.vercel.app/api/quote
```

## Contributions

Contributions, issues, and feature requests are welcome!
