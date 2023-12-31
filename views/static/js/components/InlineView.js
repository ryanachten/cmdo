import { h } from "https://esm.sh/preact@10.18.1";
import htm from "https://esm.sh/htm@3.1.1";
import { useEffect } from "https://esm.sh/preact@10.18.1/hooks";

const html = htm.bind(h);

/**
 * @param {{ history: import('./App.js').Message[], commands: import('./App.js').CommandHash[] }}
 */
function InlineView({ history, commands, contentRef }) {
  useEffect(() => {
    contentRef.current.scroll({
      behavior: "smooth",
      top: contentRef.current.scrollHeight,
    });
  }, [history]);

  return html`<div className="inline-view">
    <div className="terminal__container">
      <span className="terminal__tab command__heading">cmdo</span>
      <ul>
        ${history.map(({ commandName, messageBody, messageType }) => {
          const className = `inline-view__item command--${
            commands[commandName].color
          } ${messageType === "error" ? "error" : ""}`;

          return html`<li className="${className}">
            <span className="command__heading inline-view__heading"
              >${commandName}</span
            >
            ${messageBody}
          </li>`;
        })}
      </ul>
      ${history.length === 0
        ? html`<div className="terminal__empty-state">No logs found</div>`
        : ""}
    </div>
  </div>`;
}

export default InlineView;
