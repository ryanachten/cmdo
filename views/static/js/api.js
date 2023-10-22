/**
 * @typedef {{commandName: string, messageBody: string, messageType: "error" | "information"}} Message
 */
import { createMessageItem } from "./dom.js";

const baseUrl = "localhost:1111";

export const getHistory = async () => {
  const res = await fetch(`http://${baseUrl}/api/history`);
  /**
   * @type {Message[]}
   */
  const history = await res.json();
  history.forEach((message) => {
    createMessageItem(message);
  });
};

const handleSocketResponse = (event) => {
  /**
   * @type {Message}
   */
  const message = JSON.parse(event.data);
  createMessageItem(message);
};

export const createWebSocketEvent = () => {
  const socket = new WebSocket(`ws://${baseUrl}/ws`);
  socket.onmessage = handleSocketResponse;
};
