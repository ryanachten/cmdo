import { h } from "https://esm.sh/preact@10.18.1";
import htm from "https://esm.sh/htm@3.1.1";
import { useEffect, useRef } from "https://esm.sh/preact@10.18.1/hooks";

import MessageBody from "./MessageBody.js";

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

  useEffect(() => {
    listRef.current.scroll({
      behavior: "smooth",
      top: listRef.current.scrollHeight,
    });
  }, [messages.length]);

  return html`<section
    className="terminal__container command-view__container command--${color}"
  >
    <span className="command__heading terminal__tab">${commandName}</span>
    <ul ref=${listRef}>
      ${messages.map(
        (message) =>
          html`<li
            className="${message.messageType === "error" ? "error" : ""}"
          >
            <${MessageBody} messageBody=${message.messageBody} />
          </li>`
      )}
    </ul>
  </section>`;
}

export default CommandView;
