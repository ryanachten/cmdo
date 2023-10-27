import { h } from "https://esm.sh/preact@10.18.1";
import htm from "https://esm.sh/htm@3.1.1";

const html = htm.bind(h);
const commandColors = ["cyan", "magenta", "green", "blue"];

/**
 * @param {{ commands: import('./App.js').CommandHash }}
 */
function CommandView({ commands }) {
  return html`<div className="content">
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

  return html`<section className="container container--${color}">
    <span className="container__heading">${commandName}</span>
    <ul>
      ${messages.map((message) => html`<li>${message.messageBody}</li>`)}
    </ul>
  </section>`;
}

export default CommandView;
