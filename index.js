const axios = require("axios");

async function fetchCats() {
  try {
    const response = await axios.get("http://localhost:8080/cats");
    console.log(response.data);
  } catch (error) {
    console.error("Error fetching cats:", error);
  }
}

fetchCats();
