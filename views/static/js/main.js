import { getHistory, createWebSocketEvent } from "./api.js";

const main = async () => {
  await getHistory();
  createWebSocketEvent();
};
main();
