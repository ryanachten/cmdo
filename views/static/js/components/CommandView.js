import { h } from "https://esm.sh/preact@10.18.1";
import htm from "https://esm.sh/htm@3.1.1";
import { useEffect, useRef } from "https://esm.sh/preact@10.18.1/hooks";

const html = htm.bind(h);
const commandColors = ["cyan", "magenta", "green", "blue"];

/**
 * @param {{ commands: import('./App.js').CommandHash }}
 */
function CommandView({ commands }) {
  return html`<div className="command-view">
    ${Object.entries(commands).map(
      ([commandName, command], index) =>
        html`<${CommandList}
          commandName=${commandName}
          messages=${command.history}
          index=${index}
        />`
    )}
  </div>`;
}

/**
 * @param {{ commandName: string, messages: import('./App.js').Message[], index: number }}
 */
function CommandList({ commandName, messages, index }) {
  const color = commandColors[index % commandColors.length];
  const listRef = useRef(null);

  useEffect(() => {
    listRef.current.scroll({
      behavior: "smooth",
      top: listRef.current.scrollHeight,
    });
  }, [messages]);

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
            ${message.messageBody}
          </li>`
      )}
    </ul>
  </section>`;
}

export default CommandView;
