import { h } from "https://esm.sh/preact@10.18.1";
import {
  useEffect,
  useState,
  useRef,
} from "https://esm.sh/preact@10.18.1/hooks";
import htm from "https://esm.sh/htm@3.1.1";

import CommandView from "./CommandView.js";
import InlineView from "./InlineView.js";
import { BASE_API_URL, BASE_WS_URL, COMMAND_COLORS } from "../constants.js";
import {
  useFilteredCommands,
  useFilteredHistory,
  formatMessageBody,
} from "../helpers.js";

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

  /**
   * @type {[boolean, () => boolean]}
   */
  const [darkMode, setDarkMode] = useState(true);

  const filteredCommands = useFilteredCommands(searchTerm, commands);
  const filteredHistory = useFilteredHistory(searchTerm, history);

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
    const res = await fetch(`${BASE_API_URL}/history`);
    /**
     * @type {Message[]}
     */
    const json = await res.json();
    const messages = json.filter((m) => {
      m.messageBody = formatMessageBody(m.messageBody);
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
    message.messageBody = formatMessageBody(message.messageBody);
    if (!message.messageBody) return;

    setHistory((prevHistory) => [...prevHistory, message]);
    addMessageToCommands(message);
  };

  useEffect(() => {
    getHistory();
    const socket = new WebSocket(BASE_WS_URL);
    socket.onmessage = handleSocketResponse;
  }, []);

  useEffect(() => {
    document.documentElement.dataset.theme = darkMode ? "dark" : "light";
  }, [darkMode]);

  const contentRef = useRef(null);

  return html`<div className="app">
    <aside className="app__sidebar">
      <span className="app__logo--image"></span>
      <span className="app__logo">cmdo</span>
      <hr />
      <section className="app__fields">
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
        <div className="app__field app__field--inline">
          <label className="app__field-header" for="dark-mode">Dark mode</label>
          <input
            type="checkbox"
            id="dark-mode"
            checked=${darkMode}
            onChange=${(e) => setDarkMode(e.target.checked)}
          />
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
