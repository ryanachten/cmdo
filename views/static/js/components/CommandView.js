import { h } from "https://esm.sh/preact@10.18.1";
import htm from "https://esm.sh/htm@3.1.1";
import {
  useEffect,
  useRef,
  useState,
} from "https://esm.sh/preact@10.18.1/hooks";

import { useFilteredHistory } from "../helpers.js";

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

  return html`<section
    className="terminal__container command-view__container command--${color}"
  >
    <div className="terminal__header">
      <span className="command__heading terminal__tab">${commandName}</span>
      <input
        placeholder="Search logs"
        onChange=${(e) => setSearchTerm(e.target.value)}
      />
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
