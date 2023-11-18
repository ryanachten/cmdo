import { h } from "https://esm.sh/preact@10.18.1";
import htm from "https://esm.sh/htm@3.1.1";
import {
  useEffect,
  useRef,
  useState,
} from "https://esm.sh/preact@10.18.1/hooks";

import { useFilteredHistory } from "../helpers.js";
import { BASE_API_URL } from "../constants.js";

const html = htm.bind(h);

/**
 * @param {{ commands: import('./App.js').CommandHash }}
 */
function CommandView({ commands }) {
  return html`<div className="command-view">
    ${Object.entries(commands).map(
      ([commandName, command]) =>
        html`<${CommandList}
          commandName=${commandName}
          messages=${command.history}
          color=${command.color}
        />`
    )}
  </div>`;
}

/**
 * @param {{ commandName: string, messages: import('./App.js').Message[], color: string }}
 */
function CommandList({ commandName, messages, color }) {
  const listRef = useRef(null);
  /**
   * @type {[string, () => string]}
   */
  const [searchTerm, setSearchTerm] = useState();
  const filteredMessages = useFilteredHistory(searchTerm, messages);

  useEffect(() => {
    listRef.current.scroll({
      behavior: "smooth",
      top: listRef.current.scrollHeight,
    });
  }, [filteredMessages.length]);

  const stopCommand = async () =>
    await fetch(`${BASE_API_URL}/command`, {
      method: "POST",
      body: JSON.stringify({
        commandName,
        requestedState: "stop",
      }),
    });

  return html`<section
    className="terminal__container command-view__container command--${color}"
  >
    <div className="terminal__header">
      <span className="terminal__tab">${commandName}</span>
      <div className="command-view__heading-inputs">
        <input
          className="command-view__container"
          placeholder="Search logs"
          onChange=${(e) => setSearchTerm(e.target.value)}
        />
        <button className="command-view__stop-button" onClick=${stopCommand}>
          Stop
        </button>
      </div>
    </div>
    <ul ref=${listRef}>
      ${filteredMessages.map(
        (message) =>
          html`<li
            className="${message.messageType === "error" ? "error" : ""}"
          >
            ${message.messageBody}
          </li>`
      )}
    </ul>
    ${filteredMessages.length === 0
      ? html`<div className="terminal__empty-state">No logs found</div>`
      : ""}
  </section>`;
}

export default CommandView;
