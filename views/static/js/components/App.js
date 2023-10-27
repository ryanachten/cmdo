import { h } from "https://esm.sh/preact@10.18.1";
import {
  useEffect,
  useState,
  useMemo,
} from "https://esm.sh/preact@10.18.1/hooks";
import htm from "https://esm.sh/htm@3.1.1";

import CommandView from "./CommandView.js";
import InlineView from "./InlineView.js";
import { BASE_URL, COMMAND_COLORS } from "../constants.js";

const html = htm.bind(h);

/**
 * @typedef {"command" | "inline"} ViewMode
 * @typedef {{commandName: string, messageBody: string, messageType: "error" | "information"}} Message
 * @typedef {Object.<string, {history: Message[], color: string}>} CommandHash
 */

function App() {
  /**
   * @type {[Message[], (Message[]) => Message[]]}
   */
  const [history, setHistory] = useState([]);
  /**
   * @type {[ViewMode, () => ViewMode]}
   */
  const [viewMode, setViewMode] = useState("command");

  const getHistory = async () => {
    const res = await fetch(`http://${BASE_URL}/api/history`);
    /**
     * @type {Message[]}
     */
    const json = await res.json();
    setHistory((prevHistory) => [...prevHistory, ...json]);
  };

  const handleSocketResponse = (event) => {
    /**
     * @type {Message}
     */
    const message = JSON.parse(event.data);
    setHistory((prevHistory) => [...prevHistory, message]);
  };

  useEffect(() => {
    getHistory();
    const socket = new WebSocket(`ws://${BASE_URL}/ws`);
    socket.onmessage = handleSocketResponse;
  }, []);

  /**
   * @type {CommandHash}
   */
  const commands = useMemo(
    () =>
      history.reduce((aggregate, message) => {
        const { commandName } = message;
        if (commandName in aggregate) {
          aggregate[commandName].history.push(message);
        } else {
          const index = Object.keys(aggregate).length;
          const color = COMMAND_COLORS[index % COMMAND_COLORS.length];
          aggregate[commandName] = { history: [message], color };
        }
        return aggregate;
      }, {}),
    [history]
  );

  return html`<div class="app">
    <label>Use inline view</label>
    <input
      type="checkbox"
      onChange=${(e) => setViewMode(e.target.checked ? "inline" : "command")}
    />
    ${viewMode === "command"
      ? html`<${CommandView} commands=${commands} />`
      : html`<${InlineView} history=${history} commands=${commands} />}`}
  </div> `;
}

export default App;
