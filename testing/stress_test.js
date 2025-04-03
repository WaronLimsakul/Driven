import http from "k6/http";
import { parseHTML } from "k6/html";
import { sleep, check } from "k6";

export const options = {
  stages: [
    { duration: "3m", target: 100 },
    { duration: "3m", target: 100 },
    { duration: "2m", target: 0 },
  ],
  thresholds: {
    http_req_duration: ["p(95)<200"], // 95% of requests should be  < 200ms
    http_req_failed: ["rate<0.01"], // fail < 1% of requests
  },
};

const baseUrl = "http://localhost:8080";
const params = {
  headers: {
    "HX-Request": "true",
  },
}; // some route will check HTMX header

// 1. go to landing page
// 2. get sign in page
// 3. sign in
// 4. get home page
// 5. create a task
// 6. done the task
// 7. undone the task
// 8. delete the task
// 9. sign out
export default function () {
  let res = http.get(baseUrl);
  check(res, {
    "get home page": (r) => r.status == 200,
  });

  res = http.get(`${baseUrl}/signin`);
  check(res, {
    "get signin page": (r) => r.status == 200,
  });

  // just pass an object to second parameter, it's default x-www-form-urlencoded
  res = http.post(
    `${baseUrl}/signin`,
    {
      email: "tester@gmail.com",
      password: "123456",
    },
    params,
  );

  check(res, {
    "is signed in": (r) =>
      r.status == 201 && r.body.includes("You're signed in!"),
  });

  res = http.get(`${baseUrl}/home`);
  check(res, {
    "can get home": (r) => r.status == 200,
  });

  res = http.post(
    `${baseUrl}/tasks/week`,
    {
      "task-name": "test-task",
      "task-priority": "0",
      "task-date": "2026-01-01",
    },
    params,
  );
  check(res, {
    "create task success": (r) => r.status == 201,
  });

  const createTaskRes = parseHTML(res.body);
  const smallTaskID = createTaskRes.find("div").attr("id");

  res = http.put(`${baseUrl}/tasks/week/${smallTaskID}/done`, {}, params);
  check(res, {
    "task done": (r) => r.status == 201,
  });

  res = http.put(`${baseUrl}/tasks/week/${smallTaskID}/undone`, {}, params);
  check(res, {
    "task undone": (r) => r.status == 201,
  });

  res = http.del(`${baseUrl}/tasks/${smallTaskID}`);
  check(res, {
    "task deleted": (r) => r.status == 204,
  });

  res = http.post(`${baseUrl}/signout`, {}, params);
  check(res, {
    "is signed out": (r) => r.status == 201,
  });
}
