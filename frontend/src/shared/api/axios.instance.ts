import axios from "axios";

const isProd = process.env.NODE_ENV === "production";

export const instance = axios.create({
  baseURL: isProd ? "" : "https://kustoma.shop",
  timeout: 1000,
  headers: { "X-Custom-Header": "foobar" },
});
