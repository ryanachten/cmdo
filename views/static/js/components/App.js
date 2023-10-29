import { h } from "https://esm.sh/preact@10.18.1";
import {
  useEffect,
  useState,
  useMemo,
} from "https://esm.sh/preact@10.18.1/hooks";
import htm from "https://esm.sh/htm@3.1.1";

import CommandView from "./CommandView.js";
import InlineView from "./InlineView.js";
import Logo from "./Logo.js";
import { BASE_URL, COMMAND_COLORS } from "../constants.js";

const html = htm.bind(h);

/**
 * @typedef {"command" | "inline"} ViewMode
 * @typedef {{commandName: string, messageBody: string, messageType: "error" | "information"}} Message
 * @typedef {Object.<string, {history: Message[], color: string}>} CommandHash
 */

function App() {
  /**
   * @type {[Message[], (history) => Message[]]}
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
    const messages = json.filter((m) => {
      m.messageBody = m.messageBody.trim();
      if (!m.messageBody) return;

      return m;
    });

    setHistory((prevHistory) => [...prevHistory, ...messages]);
  };

  const handleSocketResponse = (event) => {
    /**
     * @type {Message}
     */
    const message = JSON.parse(event.data);
    message.messageBody = message.messageBody.trim();
    if (!message.messageBody) return;

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
  // TODO: this is really inefficient - add command as own state updated only as needed
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

  return html`<div className="app">
    <aside className="app__sidebar">
      ${html`<${Logo} />`}
      <div>
        <label for="view-mode">View mode</label>
        <select
          id="view-mode"
          value=${viewMode}
          onChange=${(e) => setViewMode(e.target.value)}
        >
          <option value="command">Grid</option>
          <option value="inline">Unified</option>
        </select>
      </div>
    </aside>
    <main className="app__content">
      ${viewMode === "command"
        ? html`<${CommandView} commands=${commands} />`
        : html`<${InlineView} history=${history} commands=${commands} />}`}
    </main>
  </div> `;
}

export default App;
