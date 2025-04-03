import http from "k6/http";
import { parseHTML } from "k6/html";
import { sleep, check } from "k6";

export const options = {
  vus: 3,
  duration: "15s",
};

const baseUrl = "http://localhost:8080";

export default function () {
  let res = http.get(baseUrl);
  check(res, {
    "get home page": (r) => r.status == 200,
  });
}
