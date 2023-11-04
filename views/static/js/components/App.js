import { h } from "https://esm.sh/preact@10.18.1";
import {
  useEffect,
  useState,
  useRef,
} from "https://esm.sh/preact@10.18.1/hooks";
import htm from "https://esm.sh/htm@3.1.1";

import CommandView from "./CommandView.js";
import InlineView from "./InlineView.js";
import { BASE_URL, COMMAND_COLORS } from "../constants.js";
import {
  useFilteredCommands,
  useFilteredHistory,
  useMessageStatusCount,
} from "../hooks.js";

const html = htm.bind(h);

/**
 * @typedef {"command" | "inline"} ViewMode
 * @typedef {"error" | "information"} MessageType
 * @typedef {{commandName: string, messageBody: string, messageType: MessageType}} Message
 * @typedef {Object.<string, {history: Message[], color: string}>} CommandHash
 */

function App() {
  /**
   * @type {[Message[], (history) => Message[]]}
   */
  const [history, setHistory] = useState([]);
  /**
   * @type {[CommandHash, (commands) => CommandHash]}
   */
  const [commands, setCommands] = useState({});
  /**
   * @type {[ViewMode, () => ViewMode]}
   */
  const [viewMode, setViewMode] = useState("command");

  /**
   * @type {[string, () => string]}
   */
  const [searchTerm, setSearchTerm] = useState();

  const filteredCommands = useFilteredCommands(searchTerm, commands);
  const filteredHistory = useFilteredHistory(searchTerm, history);
  const statusCount = useMessageStatusCount(filteredHistory);

  /**
   * @param {Message} message
   */
  const addMessageToCommands = (message) => {
    setCommands((prevCommands) => {
      const { commandName } = message;
      const updatedCommands = { ...prevCommands };

      if (commandName in updatedCommands) {
        updatedCommands[commandName].history.push(message);
      } else {
        const index = Object.keys(updatedCommands).length;
        const color = COMMAND_COLORS[index % COMMAND_COLORS.length];
        updatedCommands[commandName] = { history: [message], color };
      }

      return updatedCommands;
    });
  };

  const getHistory = async () => {
    const res = await fetch(`http://${BASE_URL}/api/history`);
    /**
     * @type {Message[]}
     */
    const json = await res.json();
    const messages = json.filter((m) => {
      m.messageBody = m.messageBody.trim();
      if (!m.messageBody) return;

      addMessageToCommands(m);
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
    addMessageToCommands(message);
  };

  useEffect(() => {
    getHistory();
    const socket = new WebSocket(`ws://${BASE_URL}/ws`);
    socket.onmessage = handleSocketResponse;
  }, []);

  const contentRef = useRef(null);

  return html`<div className="app">
    <aside className="app__sidebar">
      <img
        className="app__logo-img"
        alt="Commando logo"
        src="/static/img/commando.png"
      />
      <span className="app__logo">Commando</span>
      <hr />
      <span className="app__field-header">Status</span>
      <section className="app__status">
        <div className="app__status-count">
          <span className="app__status-value"
            >${statusCount.information ?? 0}</span
          >
          stdin
        </div>
        <div className="app__status-count app__status--error">
          <span className="app__status-value">${statusCount.error ?? 0}</span>
          stderr
        </div>
      </section>
      <hr />
      <section>
        <div className="app__field">
          <label className="app__field-header" for="view-mode">View mode</label>
          <select
            id="view-mode"
            value=${viewMode}
            onChange=${(e) => setViewMode(e.target.value)}
          >
            <option value="command">Grid</option>
            <option value="inline">Unified</option>
          </select>
        </div>
        <div className="app__field">
          <label className="app__field-header" for="search-all-logs"
            >Search</label
          >
          <input
            id="search-all-logs"
            placeholder="search logs"
            onChange=${(e) => setSearchTerm(e.target.value)}
          />
        </div>
      </section>
    </aside>
    <main ref=${contentRef} className="app__content">
      ${viewMode === "command"
        ? html`<${CommandView} commands=${filteredCommands} />`
        : html`<${InlineView}
            history=${filteredHistory}
            commands=${commands}
            contentRef=${contentRef}
          />`}
    </main>
  </div> `;
}

export default App;
