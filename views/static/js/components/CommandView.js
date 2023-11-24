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
          state=${command.state}
        />`
    )}
  </div>`;
}

/**
 * @param {{ commandName: string, messages: import('./App.js').Message[], color: string, state: import('./App.js').CommandState }}
 */
function CommandList({ commandName, messages, color, state }) {
  const listRef = useRef(null);
  /**
   * @type {[string, () => string]}
   */
  const [searchTerm, setSearchTerm] = useState();
  /**
   * @type {[bool, () => bool]}
   */
  const [isActive, setActive] = useState(true); // TODO: really this should be determined by the backend
  const filteredMessages = useFilteredHistory(searchTerm, messages);

  useEffect(() => {
    listRef.current.scroll({
      behavior: "smooth",
      top: listRef.current.scrollHeight,
    });
  }, [filteredMessages.length]);

  useEffect(() => {
    if (state === "started") {
      setActive(true);
    } else {
      setActive(false);
    }
  }, [state]);

  /**
   * @param {"stop" | "start"} requestedState
   */
  const updateCommandState = async (requestedState) => {
    await fetch(`${BASE_API_URL}/command`, {
      method: "POST",
      body: JSON.stringify({
        commandName,
        requestedState,
      }),
    });
    setActive(requestedState === "start");
  };

  return html`<section
    className="terminal__container command-view__container command--${color} command-view--${state}"
  >
    <div className="terminal__header">
      <span className="terminal__tab">${commandName}</span>
      <div className="command-view__heading-inputs">
        <input
          className="command-view__container"
          placeholder="Search logs"
          onChange=${(e) => setSearchTerm(e.target.value)}
        />
        ${isActive
          ? html`<button
              className="command-view__button command-view__button--stop"
              onClick=${() => updateCommandState("stop")}
            >
              Stop
            </button>`
          : html`<button
              className="command-view__button command-view__button--start"
              onClick=${() => updateCommandState("start")}
            >
              Start
            </button>`}
      </div>
    </div>
    ${state === "stopped" &&
    html`<div className="command-view__status command-view__status--stopped">
      ${commandName} has stopped
    </div>`}
    ${state === "failed" &&
    html`<div className="command-view__status command-view__status--failed">
      ${commandName} has failed
    </div>`}
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
