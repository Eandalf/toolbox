/**
 * To test an api right inside a browser's dev tool console.
 * 
 * Some assuptions: Request (JSON) -> Response (JSON)
 * 
 * API_URL: the endpoint of the API for testing.
 * 
 * Param method could be "", use "POST" by default.
 * 
 * e.g.,
 * let payload = { key1: value1, key2: value2 };
 * await testApi("GET", "/to/uri", payload);
 * @param {string} method 
 * @param {string} path 
 * @param {Object} payload 
 * @returns Promise<Object>
 */
async function testApi(method, path, payload) {
  const API_URL = "http://localhost:3000";
  
  try {
    const response = await fetch(API_URL + path, {
      method: method ? method : "POST",
      headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/json',
      },
      body: payload ? JSON.stringify(payload) : undefined,
    });
    
    if (!response.ok) {
      throw new Error(`Response status: ${response.status}`);
    }

    const json = await response.json();
    return json;
  } catch (error) {
    console.error(error.message);
    return null;
  }
}
