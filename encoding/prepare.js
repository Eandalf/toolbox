// Load binary streams and turn them into blobs.

// test -> gzip
const data = "H4sIAAAAAAAA/ypJLS4BBAAA//8Mfn/YBAAAAA==";

// option 1
const byteCharacters = atob(data);

const byteNumbers = new Array(byteCharacters.length);
for (let i = 0; i < byteCharacters.length; i++) {
  byteNumbers[i] = byteCharacters.charCodeAt(i);
}

let payload = new Blob(new Uint8Array(byteNumbers));

// option 2
payload = Uint8Array.from(atob(data), (c) => c.charCodeAt(0));

// option 3 (This one actually worked!!!)
// Ref: https://dev.to/ternentdotdev/json-compression-in-the-browser-with-gzip-and-the-compression-streams-api-4135
// create a stream
const stream = new Blob(["test"], { type: "text/plain" }).stream();

// gzip stream
const gStream = stream.pipeThrough(
  new CompressionStream("gzip")
);

// create a response
const res = new Response(gStream);

payload = await res.blob();
