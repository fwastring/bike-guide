import axios from "axios";

const api = axios.create({
  baseURL: "/api",
  timeout: 10000,
  headers: {
    "Content-Type": "application/json",
	// "Authorization": localStorage.getItem("jwt")
  },
});

export default api;
